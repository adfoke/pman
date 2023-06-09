/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"github.com/spf13/cobra"
	"github.com/adfoke/pman/process"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args)!=1 {
			fmt.Println("Please specify one pid")
			os.Exit(1)
		}
		pid, err := strconv.ParseUint(args[0], 10, 32)
		if err != nil {
		fmt.Println("Invalid PID:", args[0])
		os.Exit(1)
		}

		process, err := process.GetProcessInfo(uint32(pid))
		if err != nil {
			fmt.Println("Process not found for PID:", pid)
			os.Exit(1)
		}
		//termintor the process
		err = process.TerminateProcess(uint32(pid))
		if err != nil {
			fmt.Println("Error terminating process:", err)
			os.Exit(1)
		}
		fmt.Println("Process terminated successfully")		

	},
}

func init() {
	rootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
