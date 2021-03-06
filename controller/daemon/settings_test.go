package daemon

import (
	"os"
	"runtime"
	"testing"

	"github.com/reef-pi/reef-pi/controller/storage"
)

func TestDevModeDetection(t *testing.T) {
	store, err := storage.TestDB()
	defer store.Close()

	if err != nil {
		t.Fatal(err)
	}
	if runtime.GOOS != "windows" {
		os.Unsetenv("DEV_MODE")
		s, err := initializeSettings(store)
		if err != nil {
			t.Error(err)
		}
		if s.Capabilities.DevMode {
			t.Error("Devmode is turned on, expected off")
		}
		os.Setenv("DEV_MODE", "1")
		s, err = initializeSettings(store)
		if err != nil {
			t.Error(err)
		}
		if !s.Capabilities.DevMode {
			t.Error("Devmode is turned off, expected on")
		}
		os.Unsetenv("DEV_MODE")
	}
}
