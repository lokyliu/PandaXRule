package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/message"
)

type messageTypeSwitchNode struct {
	bareNode
}
type messageTypeSwitchNodeFactory struct{}

func (f messageTypeSwitchNodeFactory) Name() string     { return "MessageTypeSwitchNode" }
func (f messageTypeSwitchNodeFactory) Category() string { return NODE_CATEGORY_FILTER }
func (f messageTypeSwitchNodeFactory) Labels() []string {
	return []string{
		message.RowMes,
		message.AttributesMes,
		message.TelemetryMes,
		message.RpcRequestMes,
		message.AlarmMes,
		message.UpEventMes,
		message.ConnectMes,
		message.ConnectMes,
		message.DisConnectMes,
	}
}
func (f messageTypeSwitchNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &messageTypeSwitchNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *messageTypeSwitchNode) Handle(msg *message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.MsgType)
	msg.Metadata = map[string]interface{}{"AA": "BB", "deviceName": "fff"}
	nodes := n.GetLinkedNodes()
	messageType := msg.MsgType
	for label, node := range nodes {
		if messageType == label {
			return node.Handle(msg)
		}
	}
	return nil
}
