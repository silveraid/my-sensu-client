package sensu

import (
	"errors"

	stdCheck "sensu-common/check"
	"sensu/check"
)

type CheckResponse struct {
	Check  stdCheck.CheckOutput `json:"check"`
	Client string               `json:"client"`
}

var commandKeyError = errors.New("Command key not filled")

func executeCheck(input *stdCheck.CheckRequest) (*stdCheck.CheckOutput, error) {
	var output stdCheck.CheckOutput

	if ch, ok := check.Store[input.Extension]; input.Extension != "" && ok {
		output = ch.Execute()
	} else if ch, ok := check.Store[input.Name]; ok {
		output = ch.Execute()
	} else if input.Command == "" {
		return nil, commandKeyError
	} else {
		output = (&check.ExternalCheck{Request: input}).Execute()
	}

	output.CheckRequest = input

	return &output, nil
}
