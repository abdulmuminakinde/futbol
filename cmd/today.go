/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/abdulmuminakinde/futbol/internal/today"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := tea.NewProgram(today.InitialModel()).Run(); err != nil {
			fmt.Printf("Oh oh: there was an error spinning up: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// todayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// todayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
