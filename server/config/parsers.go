package config

import (
	"errors"
	"fmt"
	"strings"
)

// Assert a non-empty byte stream or panic otherwise
func assertNonEmptyStream(bytes []byte) error {
	if len(bytes) == 0 {
		return errors.New("EMPTY STREAM")
	}
	return nil
}

//ParseRequest convert bytes into string for further processing
func ParseRequest(bytes []byte) ([]string, error) {

	err := assertNonEmptyStream(bytes)
	if err != nil {
		fmt.Errorf(err.Error())
		return nil, err
	}
	clientInput := string(bytes[:])
	clientInputFields := strings.Fields(clientInput)
	return clientInputFields, nil

}
