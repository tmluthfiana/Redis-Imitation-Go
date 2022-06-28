package commands

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInvalidSetParameters(t *testing.T) {
	commandWithParameters := []string{"SET"}
	result := executeSetCommand(commandWithParameters)
	require.Equal(t, result, "INVALID SET INPUT, example : SET <key> <value>")
}

func TestInvalidDumpParameters(t *testing.T) {
	commandWithParameters := []string{"DUMP"}
	result := executeDumpCommand(commandWithParameters)
	require.Equal(t, result, "INVALID DUMP INPUT, example : DUMP <key>")
}

func TestInvalidIncrParameters(t *testing.T) {
	commandWithParameters := []string{"INCR"}
	result := executeIncrCommand(commandWithParameters)
	require.Equal(t, result, "INVALID INCR INPUT, example : INCR <key>")
}

func TestInvalidRenameParameters(t *testing.T) {
	commandWithParameters := []string{"RENAME"}
	result := executeRenameCommand(commandWithParameters)
	require.Equal(t, result, "INVALID RENAME INPUT, example : RENAME <key> <key>")
}

func TestExecuteCommandError(t *testing.T) {
	commandWithParameters := []string{"123"}
	result := ExecuteCommand(commandWithParameters)
	require.Equal(t, result, "UNKNOWN COMMAND")
}

func TestHelpCommand(t *testing.T) {
	commandWithParameters := []string{"HELP"}
	result := ExecuteCommand(commandWithParameters)
	require.Equal(t, result, "Commands supported : \nSET <key> <value> \nRENAME <key> <key>\nDUMP <key>")
}
