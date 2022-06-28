package commands

import (
	"server/store"
)

const (
	helpCommand        = "HELP"
	dumpCommand        = "DUMP"
	setCommand         = "SET"
	incrCommand        = "INCR"
	renameCommand      = "RENAME"
	success            = "SUCCESS"
	unknownCommand     = "UNKNOWN COMMAND"
	notFound           = "NOT FOUND"
	error              = "ERROR"
	invalidSetInput    = "INVALID SET INPUT, example : SET <key> <value>"
	invalidRenameInput = "INVALID RENAME INPUT, example : RENAME <key> <key>"
	invalidDumpInput   = "INVALID DUMP INPUT, example : DUMP <key>"
	invalidIncrInput   = "INVALID INCR INPUT, example : INCR <key>"
	helpCommandOutput  = "Commands supported : \nSET <key> <value> \nRENAME <key> <key>\nDUMP <key>\nINCR <key>"
)

var gm = store.NewStorageMap()

// ExecuteCommand takes a command and inspects it to check there is
// a matching executable command. If no command can be found, it returns unknown command
func ExecuteCommand(commandWithParameters []string) string {

	key := commandWithParameters[0]

	switch key {
	case helpCommand:
		return helpCommandOutput
	case dumpCommand:
		return executeDumpCommand(commandWithParameters)
	case setCommand:
		return executeSetCommand(commandWithParameters)
	case incrCommand:
		return executeIncrCommand(commandWithParameters)
	case renameCommand:
		return executeRenameCommand(commandWithParameters)
	default:
		break
	}

	return unknownCommand
}

func executeDumpCommand(commandWithParameters []string) string {
	if len(commandWithParameters) < 2 {
		return invalidDumpInput
	}
	value, ok := gm.Dump(commandWithParameters[1])
	if ok {
		return value
	}
	return notFound
}

func executeSetCommand(commandWithParameters []string) string {
	if len(commandWithParameters) < 3 {
		return invalidSetInput
	}
	ok := gm.Set(commandWithParameters[1], commandWithParameters[2])
	if ok {
		return success
	}
	return error
}

func executeIncrCommand(commandWithParameters []string) string {
	if len(commandWithParameters) < 2 {
		return invalidIncrInput
	}
	value, ok := gm.Incr(commandWithParameters[1])
	if ok {
		return value
	}
	return error
}

func executeRenameCommand(commandWithParameters []string) string {
	if len(commandWithParameters) < 3 {
		return invalidRenameInput
	}
	ok := gm.Rename(commandWithParameters[1], commandWithParameters[2])
	if ok {
		return success
	}
	return notFound
}
