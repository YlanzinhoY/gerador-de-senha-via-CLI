/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"ylanzinhoy/services"

	"github.com/spf13/cobra"
)

// passwordGeneratorCmd represents the passwordGenerator command
var passwordGeneratorCmd = &cobra.Command{
	Use:   "passwordGenerator",
	Short: "Generated a new password",
	Run: func(cmd *cobra.Command, args []string) {

		lengthPassoword, err := cmd.Flags().GetInt("length")
		if err != nil {
			panic("error in parse length")
		}
		services.PasswordGenerator(lengthPassoword)
	},
}


func init() {
	rootCmd.AddCommand(passwordGeneratorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// passwordGeneratorCmd.PersistentFlags().String("foo", "", "A help for foo")

	passwordGeneratorCmd.Flags().Int("length", 12, "choice quantity of char")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// passwordGeneratorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
