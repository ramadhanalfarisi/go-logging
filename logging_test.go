package go_logging

import (
	"fmt"
	"testing"
)

func TestLogging(t *testing.T) {
	get_config := ConfigLogging(true,"logs/")
	get_config.Debug("This is debug")
	get_config.Info("This is info")
	get_config.Warning("This is warning")
	get_config.Error(fmt.Errorf("This is error"))
}