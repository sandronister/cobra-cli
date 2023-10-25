/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/sandronister/cobra-cli/internal/database"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
func newCreateCmd(categoryDB database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a new category",
		Long:  `Create a new category`,
		RunE:  runCreate(categoryDB),
	}
}

func runCreate(categoryDB database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		_, err := categoryDB.Create(name, description)

		if err != nil {
			return err
		}

		return nil
	}
}

func init() {
	createCmd := newCreateCmd(GetCategoryDB(GetDb()))
	categoryCmd.AddCommand(createCmd)
	createCmd.PersistentFlags().StringP("name", "n", "", "Name of the category")
	createCmd.PersistentFlags().StringP("description", "d", "", "Description of category")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
