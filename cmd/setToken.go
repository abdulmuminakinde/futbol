/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/abdulmuminakinde/futbol/internal/token"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// setTokenCmd represents the setToken command
var setTokenCmd = &cobra.Command{
	Use:   "setToken",
	Short: "setToken collects and stores the APIKey in an interactive session",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("setToken called")
		m := token.InitialModel()
		if _, err := tea.NewProgram(m).Run(); err != nil {
			fmt.Printf("Oh oh: there was an error spinning up: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(setTokenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setTokenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setTokenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
