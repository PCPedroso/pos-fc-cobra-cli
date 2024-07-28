/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/PCPedroso/pos-fc-cobra-cli/internal/database"
	"github.com/spf13/cobra"
)

var _name, _description string = "name", "description"

// createCmd represents the create command
// para criar o pacote foi utilizado a seguinte linha de comando
// cobra-cli add create -p 'categoryCmd'
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func newCreateCmd(categoryDB database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a new category",
		Long:  "Create a new category",
		RunE:  runCreate(categoryDB),
	}
}

func runCreate(categoryDB database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString(_name)
		description, _ := cmd.Flags().GetString(_description)
		_, err := categoryDB.Create(name, description)
		if err != nil {
			return err
		}
		return nil
	}
}

func init() {
	createCmd := newCreateCmd(GetCategoryDB(GetDB()))

	categoryCmd.AddCommand(createCmd)

	createCmd.Flags().StringP(_name, "n", "", "Name of the category")
	createCmd.Flags().StringP(_description, "d", "", "Description of the category")

	createCmd.MarkFlagsRequiredTogether(_name, _description)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
