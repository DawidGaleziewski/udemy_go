/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package net

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

func ping(domain string)(int, error){
	url := "http://" + domain
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	resp.Body.Close()
	return resp.StatusCode, nil
}

var client = http.Client{
	Transport: &http.Transport {
		
	},
}
var pingPath string

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// logic 
		ping(pingPath)
	},
}

func init() {
	// rootCmd.AddCommand(pingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
		pingCmd.Flags().StringVarP(&pingPath, "url", "n", "", "This is a url for the pinging target") // we can pass string variable pointer there to populate this referance
	if err := pingCmd.MarkFlagRequired("url"); err != nil { // shorthand to guard agains usage without flag
		fmt.Println(err)
	}
}
