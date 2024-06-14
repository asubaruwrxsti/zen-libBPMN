package libbpmn

import (
	"fmt"
	"log"
	"time"

	"github.com/nitram509/lib-bpmn-engine/pkg/bpmn_engine"
)

type BpmnEngine struct{}

type BpmnEngineConfig struct {
	EngineName string
	BpmnFile   string
	TaskId     string
	Variables  map[string]interface{}
}

func (BpmnEngine) Execute(config BpmnEngineConfig) (*bpmn_engine.ProcessInstanceInfo, error) {
	engine := bpmn_engine.New(config.EngineName)
	process, err := engine.LoadFromFile(config.BpmnFile)
	if err != nil {
		log.Fatalf("file \"%s\" can't be read: %v", config.BpmnFile, err)
	}
	fmt.Printf("Loaded BPMN file: %s\n", config.BpmnFile)

	engine.AddTaskHandler(config.TaskId, userTaskHandler())
	fmt.Printf("Added task handler for task: %s\n", config.TaskId)

	variables := config.Variables
	if variables == nil {
		variables = map[string]interface{}{}
	}
	fmt.Printf("Starting process instance with variables: %v\n", variables)

	engine.CreateAndRunInstance(process.ProcessKey, variables)
	return engine.RunOrContinueInstance(process.ProcessKey)
}

func printContextHandler(job bpmn_engine.ActivatedJob) {
	fmt.Printf("\n\n////////////////////////////////////////\n")
	fmt.Printf("ElementId:                = %s\n", job.GetElementId())
	fmt.Printf("BpmnProcessId:            = %s\n", job.GetBpmnProcessId())
	fmt.Printf("ProcessDefinitionKey:     = %d\n", job.GetProcessDefinitionKey())
	fmt.Printf("ProcessDefinitionVersion: = %d\n", job.GetProcessDefinitionVersion())
	fmt.Printf("CreatedAt:                = %s\n", job.GetCreatedAt())
	fmt.Printf("Temprature:           	  = %s\n", job.GetVariable("temp"))
	fmt.Printf("Activity:                 = %s\n", job.GetVariable("act"))
	fmt.Printf("////////////////////////////////////////\n\n")

	job.Complete()
}

var externalEvent = "none"

func userTaskHandler() func(job bpmn_engine.ActivatedJob) {
	return func(job bpmn_engine.ActivatedJob) {
		if externalEvent == "none" {
			// send notification to user
			fmt.Printf("Sending notification to user ...\n")
			time.Sleep(5 * time.Second)

			// wait for user response
			fmt.Printf("Waiting for user response ...\n")
			time.Sleep(5 * time.Second)

			// simulate user response
			externalEvent = "user is done"
		}
		if externalEvent == "user is done" {
			fmt.Printf("User is done!\n")
			job.Complete()
		}
		if externalEvent == "user is done but wrong response" {
			job.Fail("error in user task")
		}
		// just return and so 'pause' the process instance
	}
}
