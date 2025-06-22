package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Display help information for commands",
	Long:  `The help command provides detailed information about available commands and their usage in the Riqo CLI tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && (args[0] == "ja") {
			fmt.Println("Riqo CLI ヘルプ:")
			fmt.Println("利用可能なコマンド:")
			fmt.Println("- suggest: 自然言語入力に基づいてCLIコマンドを提案します。")
			fmt.Println("- history: コマンド履歴を管理します（表示、クリア）。")
			fmt.Println("- docs: ドキュメントを自動的に整理および更新します。")
			fmt.Println("- realtime: リアルタイムでコマンド提案を提供します。")
			fmt.Println("- llm: ローカルLLMと統合してコマンド提案を行います。")
			fmt.Println("詳細情報については 'riqo [command] --help' を使用してください。")
		} else {
			fmt.Println("Riqo CLI Help:")
			fmt.Println("Available Commands:")
			fmt.Println("- suggest: Suggest a CLI command based on natural language input.")
			fmt.Println("- history: Manage command history (view, clear).")
			fmt.Println("- docs: Automatically organize and update documentation.")
			fmt.Println("- realtime: Provide real-time command suggestions.")
			fmt.Println("- llm: Integrate with a local LLM for command suggestions.")
			fmt.Println("Use 'riqo [command] --help' for detailed information about a specific command.")
		}
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)
}
