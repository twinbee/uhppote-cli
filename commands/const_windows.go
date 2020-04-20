package commands

import (
	"golang.org/x/sys/windows"
	"path/filepath"
)

// DefaultConfig is the default file path for the uhppoted configuration file
var DefaultConfig = filepath.Join(workdir(), "uhppoted.conf")

func workdir() string {
	programData, err := windows.KnownFolderPath(windows.FOLDERID_ProgramData, windows.KF_FLAG_DEFAULT)
	if err != nil {
		return `C:\uhppoted`
	}

	return filepath.Join(programData, "uhppoted")
}
