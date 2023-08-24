package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/message"
)

type createAlarmNode struct {
	bareNode
	AlarmType     string `json:"alarmType" yaml:"alarmType"`
	AlarmSeverity string `json:"alarmSeverity" yaml:"alarmSeverity"`
}

type createAlarmNodeFactory struct{}

func (f createAlarmNodeFactory) Name() string     { return "CreateAlarmNode" }
func (f createAlarmNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f createAlarmNodeFactory) Labels() []string { return []string{"Created", "Updated", "Failure"} }
func (f createAlarmNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &createAlarmNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *createAlarmNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	//created := n.GetLinkedNode("Created")
	//updated := n.GetLinkedNode("Updated")
	//failure := n.GetLinkedNode("Failure")

	return nil
}
