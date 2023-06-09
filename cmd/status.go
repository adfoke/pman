/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	"os"
	"github.com/spf13/cobra"
	"github.com/adfoke/pman/process"

)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show process status",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Please specify one pid")
			os.Exit(1)
		}
		//convert the pid to uint32
		pid, err := strconv.ParseUint(args[0], 10, 32)
		if err != nil {
			fmt.Println("Invalid PID:", args[0])
			os.Exit(1)
		}
		//get the process info
		appPtr, err := process.GetProcessInfo(uint32(pid))
		if err != nil {
			fmt.Println("Process not found for PID:", pid)
			os.Exit(1)
		}
		fmt.Println(appPtr)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
