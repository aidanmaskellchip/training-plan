package handlers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pontus-devoteam/agent-sdk-go/pkg/agent"
	"github.com/pontus-devoteam/agent-sdk-go/pkg/runner"
)

type ConversationHandler struct {
	agent  *agent.Agent
	runner *runner.Runner
}

func NewConversationHandler(agent *agent.Agent, runner *runner.Runner) *ConversationHandler {
	return &ConversationHandler{
		agent:  agent,
		runner: runner,
	}
}

func (h *ConversationHandler) StartInteractiveSession() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Training Plan Assistant")
	fmt.Println("Type 'quit' or 'exit' to stop")
	fmt.Println("------------------------")

	// Introduction message
	result, err := h.runner.RunSync(h.agent, &runner.RunOptions{
		Input: "Introduce yourself and explain your purpose.",
	})
	if err != nil {
		log.Fatalf("Error running agent: %v", err)
	}
	fmt.Println("Assistant:", result.FinalOutput)

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "quit" || input == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		if input == "" {
			continue
		}

		result, err := h.runner.RunSync(h.agent, &runner.RunOptions{
			Input: input,
		})

		if err != nil {
			log.Printf("Error running agent: %v", err)
			continue
		}

		fmt.Println("Assistant:", result.FinalOutput)
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
}
