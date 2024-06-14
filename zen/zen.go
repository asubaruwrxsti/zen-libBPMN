package zenEngine

import (
	"fmt"
	"os"

	"github.com/gorules/zen-go"
)

func Execute(config zen.EngineConfig, graphName string, variables map[string]any) ([]byte, error) {
	fmt.Printf("\n\n////////////////////////////////////////\n")

	engine := zen.NewEngine(config)
	fmt.Printf("Starting Zen Engine with config %+v\n", config)

	graph, err := os.ReadFile(graphName)
	if err != nil {
		return nil, fmt.Errorf("error reading graph file: %w", err)
	}
	fmt.Printf("Loaded graph file: %s\n", graphName)

	decision, err := engine.CreateDecision(graph)
	if err != nil {
		return nil, fmt.Errorf("error creating decision: %w", err)
	}
	fmt.Printf("Created decision from graph\n")

	response, err := decision.Evaluate(variables)
	if err != nil {
		return nil, fmt.Errorf("error evaluating decision: %w", err)
	}
	fmt.Printf("Evaluated decision with variables: %v\n", variables)

	responseJSON, err := response.Result.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("error converting response to JSON: %w", err)
	}
	fmt.Printf("Response: %s\n", responseJSON)
	fmt.Printf("////////////////////////////////////////\n\n")
	return responseJSON, nil
}
