package main

import (
	"fmt"
	libbpmn "test/main/lib-bpmn"
)

func main() {
	bpmnConf := libbpmn.BpmnEngineConfig{
		EngineName: "libBpmn",
		BpmnFile:   "simple_task_user.bpmn",
		TaskId:     "user-task",
		Variables:  map[string]interface{}{"temp": "hot", "act": "outside"},
	}

	engine := libbpmn.BpmnEngine{}
	context, err := engine.Execute(bpmnConf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Process instance completed: %+v\n", context)

	// zenEngine.Execute(zen.EngineConfig{}, "./graph.json", map[string]interface{}{
	// 	"temp": context.GetVariable("temp"), "act": context.GetVariable("act"),
	// })
}
