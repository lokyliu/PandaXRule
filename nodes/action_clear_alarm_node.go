package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/message"
)

const ClearAlarmNodeName = "ClearAlarmNode"

type clearAlarmNodeFactory struct{}

type clearAlarmNode struct {
	bareNode
	Script    string `json:"script" yaml:"script"`
	AlarmType string `json:"alarmType" yaml:"alarmType"`
}

func (f clearAlarmNodeFactory) Name() string     { return ClearAlarmNodeName }
func (f clearAlarmNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f clearAlarmNodeFactory) Labels() []string { return []string{"Cleared", "Failure"} }
func (f clearAlarmNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &clearAlarmNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *clearAlarmNode) Handle(msg *message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.MsgType)
	//cleared := n.GetLinkedNode("Cleared")
	//failure := n.GetLinkedNode("Failure")

	return nil
}
