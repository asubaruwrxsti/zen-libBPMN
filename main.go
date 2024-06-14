package main

import (
	libbpmn "test/main/lib-bpmn"
	zenEngine "test/main/zen"

	zen "github.com/gorules/zen-go"
)

func main() {
	bpmnConf := libbpmn.BpmnEngineConfig{
		EngineName: "libBpmn",
		BpmnFile:   "simple_task.bpmn",
		TaskId:     "hello-world",
		Variables:  map[string]interface{}{"temp": "hot", "act": "outside"},
	}

	engine := libbpmn.BpmnEngine{}
	context, err := engine.Execute(bpmnConf)
	if err != nil {
		panic(err)
	}

	zenEngine.Execute(zen.EngineConfig{}, "./graph.json", map[string]interface{}{
		"temp": context.GetVariable("temp"), "act": context.GetVariable("act"),
	})
}
