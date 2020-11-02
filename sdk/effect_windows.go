package sdk

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

const (
	// Extension of Vst2 files
	Extension = ".dll"
)

var (
	// ScanPaths of Vst2 files
	scanPaths = []string{
		"C:\\Program Files (x86)\\Steinberg\\VSTPlugins",
		"C:\\Program Files\\Steinberg\\VSTPlugins ",
	}
)

func init() {
	envVstPath := os.Getenv("VST_PATH")
	if len(envVstPath) > 0 {
		scanPaths = append(scanPaths, envVstPath)
	}
}

// Open loads the plugin entry point into memory. It's DLL in windows.
func Open(path string) (*EntryPoint, error) {
	//Load plugin by path
	dll, err := syscall.LoadDLL(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load VST from '%s': %v\n", path, err)
	}

	//Get pointer to plugin's Main function
	m, err := syscall.GetProcAddress(dll.Handle, main)
	if err == nil {
		return &EntryPoint{
			main:   effectMain(unsafe.Pointer(m)),
			handle: uintptr(dll.Handle),
			Name:   dll.Name,
		}, nil
	}

	err = fmt.Errorf("failed to get entry point for plugin '%s': %w\n", path, err)
	if errRelease := dll.Release(); errRelease != nil {
		return nil, fmt.Errorf("failed to release DLL '%s': %w after: %v", path, errRelease, err)
	}
	return nil, err
}

// Close frees plugin handle.
func (m *EntryPoint) Close() error {
	if err := syscall.FreeLibrary(syscall.Handle(m.handle)); err != nil {
		return fmt.Errorf("failed to release VST handle: %w", err)
	}
	return nil
}
