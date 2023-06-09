package process

import(
	"fmt"
	"syscall"
	"unsafe"
	"golang.org/x/sys/windows"
)

type entry windows.ProcessEntry32

var (
	kernel32            = windows.NewLazySystemDLL("kernel32.dll")
	procTerminateProcess = kernel32.NewProc("TerminateProcess")
	procCloseHandle      = kernel32.NewProc("CloseHandle")
	)


func ListProcesses() ([]entry, error) {
	//get the list of windows processes
	procs, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return nil, fmt.Errorf("Failed to create processes snapshot: %v", err)
	}
	defer windows.CloseHandle(procs)
	//return the list of processes
	var procEntry windows.ProcessEntry32
	procEntry.Size = uint32(unsafe.Sizeof(procEntry))
	err = windows.Process32First(procs, &procEntry)
	if err != nil {
		return nil, fmt.Errorf("Failed to get first process: %v", err)
	}
	var entries []entry
	for {
		entries = append(entries, entry(procEntry))
		err = windows.Process32Next(procs, &procEntry)
		if err != nil {
			break
		}
	}
	return entries, nil	
}


func GetProcessInfo(pid uint32) (entry, error) {
	//find the process by pid
	procs, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, pid)
	if err != nil {
		return entry{}, fmt.Errorf("Failed to create processes snapshot: %v", err)
	}
	defer windows.CloseHandle(procs)
	//return the process info
	var procEntry windows.ProcessEntry32
	procEntry.Size = uint32(unsafe.Sizeof(procEntry))
	err = windows.Process32First(procs, &procEntry)
	if err != nil {
		return entry{}, fmt.Errorf("Failed to get first process: %v", err)
	}
	return entry(procEntry), nil	
}
func CreateProcess(appName string) (*entry, error) {
	//create the process by appname
	var startupInfo windows.StartupInfo
	var processInfo windows.ProcessInformation
	appNamePtr, err := syscall.UTF16PtrFromString(appName)
	if err != nil {
		return nil, fmt.Errorf("Failed to convert app name to UTF16: %v", err)
	}
	err = windows.CreateProcess(appNamePtr, nil, nil, nil, false, 0, nil, nil, &startupInfo, &processInfo)
	if err != nil {
		return nil, fmt.Errorf("Failed to create process: %v", err)
	}	

	//return the process info
	return &entry{ProcessID: processInfo.ProcessId}, nil

}

func TerminateProcess(pid uint32) error {
	handle, err := GetProcessInfo(pid)
	if err != nil {
		return fmt.Errorf("Failed to get process info: %v", err)
	}
	//terminate the process
	procTerminateProcess.Call(uintptr(handle.ProcessID), uintptr(0))
	//close the handle
	procCloseHandle.Call(uintptr(handle.ProcessID))
	fmt.Println("Process terminated")
	return nil	
}
