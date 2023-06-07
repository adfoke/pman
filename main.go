package main

import (
	"fmt"
	"strings"
	"os/exec"
	"strconv"
	"syscall"
	)

type Process struct {
	cmd    *exec.Cmd
	pid    int
	status string
}

func startProcess(command string, args []string) *Process {
	cmd := exec.Command(command, args...)
	err := cmd.Start()
	if err != nil {
	fmt.Printf("Error starting process: %v\n", err)
	return nil
	}
process := &Process{
	cmd:    cmd,
	pid:    cmd.Process.Pid,
	status: "running",
	}
	
	return process
}

func stopProcess(process *Process) {
	err := process.cmd.Process.Signal(syscall.SIGTERM)
	if err != nil {
	fmt.Printf("Error stopping process: %v\n", err)
	return
	}
	
	process.status = "stopped"
}

func printProcessStatus(process *Process) {
	fmt.Printf("PID: %d, Status: %s\n", process.pid, process.status)
}

func main() {
	var processes []*Process
	
	for {
	fmt.Print("Enter command (start/stop/status/exit): ")
	var command string
	fmt.Scanln(&command)
	
	switch command {
	case "start":
	fmt.Print("Enter process command: ")
	var processCommand string
	fmt.Scanln(&processCommand)
	
	fmt.Print("Enter process arguments (separated by space): ")
	var processArgs string
	fmt.Scanln(&processArgs)
	
	args := strings.Split(processArgs, " ")
	process := startProcess(processCommand, args)
	if process != nil {
	processes = append(processes, process)
	}
	
	case "stop":
	fmt.Print("Enter process PID: ")
	var pidStr string
	fmt.Scanln(&pidStr)
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
	fmt.Println("Invalid PID")
	continue
	}
	
	var processToStop *Process
	for _, process := range processes {
	if process.pid == pid {
	processToStop = process
	break
	}
	}
	
	if processToStop != nil {
	stopProcess(processToStop)
	} else {
	fmt.Println("Process not found")
	}
	
	case "status":
	for _, process := range processes {
	printProcessStatus(process)
	}
	
	case "exit":
	return
	
	default:
	fmt.Println("Invalid command")
	}
	}
}
	