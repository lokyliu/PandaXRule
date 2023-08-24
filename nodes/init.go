package nodes

// init 注册所有节点
func init() {
	RegisterFactory(inputNodeFactory{})
	RegisterFactory(switchFilterNodeFactory{})
	RegisterFactory(scriptFilterNodeFactory{})
	RegisterFactory(messageTypeFilterNodeFactory{})
	RegisterFactory(messageTypeSwitchNodeFactory{})

	RegisterFactory(transformDeleteKeyNodeFactory{})
	RegisterFactory(transformRenameKeyNodeFactory{})
	RegisterFactory(transformScriptNodeFactory{})

	RegisterFactory(createAlarmNodeFactory{})
	RegisterFactory(clearAlarmNodeFactory{})
	RegisterFactory(logNodeFactory{})
	RegisterFactory(saveAttributesNodeFactory{})
	RegisterFactory(saveTimeSeriesNodeFactory{})
	RegisterFactory(delayNodeFactory{})

	RegisterFactory(externalDingNodeFactory{})
	RegisterFactory(externalWechatNodeFactory{})
	RegisterFactory(externalKafkaNodeFactory{})
	RegisterFactory(externalNatsNodeFactory{})
	RegisterFactory(externalMqttNodeFactory{})
	RegisterFactory(externalRestapiNodeFactory{})
	RegisterFactory(externalSendEmailNodeFactory{})
	RegisterFactory(externalSendSmsNodeFactory{})
	RegisterFactory(externalRuleChainNodeFactory{})

}
