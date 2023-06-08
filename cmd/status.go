/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	"os"
	"github.com/spf13/cobra"
	"github.com/mitchellh/go-ps"
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
		pid, err := strconv.Atoi(args[0])
		if err != nil {
            fmt.Println("Invalid pid:", args[0])
            os.Exit(1)
        }
		processes, err := ps.Processes()
        if err != nil {
            fmt.Println("Failed to get processes:", err)
            os.Exit(1)
        }
		var process ps.Process
        for _, p := range processes {
            if p.Pid() == pid {
                process = p
                break
            }
        }
		if process.Pid() == 0 {
            fmt.Println("Process not found:", pid)
            os.Exit(1)
        }
		//print the detail of the pid
		fmt.Printf("Process ID: %d\n", process.Pid())
		fmt.Printf("Parent Process ID: %d\n", process.PPid())
		fmt.Printf("Executable: %s\n", process.Executable())
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
