/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/metapox/logical-streaming-app/ingester"
	"github.com/spf13/cobra"
)

// ingestCmd represents the ingest command
var ingestCmd = &cobra.Command{
	Use:   "ingest",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ingester, err := ingester.CreateIngester("postgres")
		if err != nil {
			panic(err)
		}
		ingester.Start()
		fmt.Println("ingest called")
	},
}

func init() {
	rootCmd.AddCommand(ingestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ingestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ingestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
