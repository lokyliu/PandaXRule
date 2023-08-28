package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/message"
)

type transformRenameKeyNode struct {
	bareNode
	FormType string    `json:"formType" yaml:"formType"` //msg metadata
	Keys     []KeyName `json:"keys" yaml:"keys"`
}
type KeyName struct {
	OldName string `json:"oldName" yaml:"oldName"`
	NewName string `json:"newName" yaml:"newName"`
}
type transformRenameKeyNodeFactory struct{}

func (f transformRenameKeyNodeFactory) Name() string     { return "RenameKeyNode" }
func (f transformRenameKeyNodeFactory) Category() string { return NODE_CATEGORY_TRANSFORM }
func (f transformRenameKeyNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f transformRenameKeyNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &transformRenameKeyNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *transformRenameKeyNode) Handle(msg *message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.MsgType)

	successLabelNode := n.GetLinkedNode("Success")
	failureLabelNode := n.GetLinkedNode("Failure")
	if n.FormType == "msg" {
		data := msg.Msg
		for _, key := range n.Keys {
			if _, found := data[key.OldName]; found {
				data[key.NewName] = data[key.OldName]
				delete(data, key.OldName)
				msg.Msg = data
			}
		}
	} else if n.FormType == "metadata" {
		data := msg.Metadata
		for _, key := range n.Keys {
			if data.GetValue(key.OldName) != nil {
				data[key.NewName] = data[key.OldName]
				delete(data, key.OldName)
				msg.Metadata = data
			}
		}
	} else {
		if failureLabelNode != nil {
			return failureLabelNode.Handle(msg)
		}
	}
	if successLabelNode != nil {
		return successLabelNode.Handle(msg)
	}
	return nil
}
