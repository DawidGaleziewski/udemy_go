/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package net

import (
	"fmt"

	"github.com/spf13/cobra"
)

// netCmd represents the net command
var NetCmd = &cobra.Command{ // we will export net command
	Use:   "net",
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


func init() {
	pingCmd.Flags().StringVarP(&pingPath, "url", "n", "", "This is a url for the pinging target") // we can pass string variable pointer there to populate this referance
	if err := pingCmd.MarkFlagRequired("url"); err != nil { // shorthand to guard agains usage without flag
		fmt.Println(err)
	}
	NetCmd.AddCommand(pingCmd) // we can nest this way sub commands in init
}
