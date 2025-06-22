package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// llmCmd represents the llm command
var llmCmd = &cobra.Command{
	Use:   "llm",
	Short: "Integrate with a local LLM for command suggestions",
	Long: `The llm command interacts with a locally running LLM (e.g., Ollama or Mistral)
to provide intelligent suggestions based on user input.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "Please provide input for the LLM.")
			return
		}

		// Example: Call a local LLM via a shell command
		input := args[0]
		output, err := exec.Command("ollama", "run", "gemma3:4b", input).Output()
		if err != nil {
			fmt.Fprintf(cmd.OutOrStdout(), "Error interacting with LLM: %v\n", err)
			return
		}

		fmt.Fprintf(cmd.OutOrStdout(), "LLM Response: %s\n", string(output))
	},
}

func init() {
	rootCmd.AddCommand(llmCmd)
}
