/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Category called with name: " + category)
		exists, _ := cmd.Flags().GetBool("exists")
		fmt.Println(exists)
		id, _ := cmd.Flags().GetInt16("id")
		fmt.Printf("ID: %d", id)
	},
}

var category string

func init() {
	categoryCmd.AddCommand(createCmd)
	createCmd.PersistentFlags().StringVarP(&category, "name", "n", "categoria", "Name of the category")
	createCmd.PersistentFlags().BoolP("exists", "e", false, "Verify category exists")
	createCmd.PersistentFlags().Int16P("id", "i", 0, "Id of category")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
