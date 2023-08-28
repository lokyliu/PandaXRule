package nodes

import (
	"pandax/message"
)

type logNode struct {
	bareNode
	Script string `json:"script"`
}

type logNodeFactory struct{}

func (f logNodeFactory) Name() string     { return "LogNode" }
func (f logNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f logNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f logNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &logNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *logNode) Handle(msg *message.Message) error {
	successLableNode := n.GetLinkedNode("Success")
	//failureLableNode := n.GetLinkedNode("Failure")

	/*scriptEngine := NewScriptEngine(*msg, "ToString", n.Script)
	logMessage, err := scriptEngine.ScriptToString()
	if err != nil {
		if failureLableNode != nil {
			return failureLableNode.Handle(msg)
		} else {
			return err
		}
	}*/

	if successLableNode != nil {
		return successLableNode.Handle(msg)
	}
	return nil
}
