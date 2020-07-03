package main

import (
	"errors"
	"fmt"
)

type commandLineSetter func (arguments *commandLineArguments, vale string)

var commandLineSetters = map[string]commandLineSetter {
	"-addr": func(arguments *commandLineArguments, value string) {arguments.addr = value},
	"-token": func(arguments *commandLineArguments, value string) { arguments.githubToken = value },
	"-role": func(arguments *commandLineArguments, value string) { arguments.role = value },
	"-path": func(arguments *commandLineArguments, value string){ arguments.path = value },
	"-key": func(arguments *commandLineArguments, value string) { arguments.key = value	},
}
func getCommandLineArguments(arr []string) (*commandLineArguments, error) {
	if len(arr) == 0 || len(arr) == 1 {
		return nil, errors.New("expected arguments, got none")
	}
	a := arr[1:]
	if len(a) % 2 != 0 {
		return nil, errors.New("expected even number of arguments")
	}
	args := &commandLineArguments{}
	for i := 0; i < len(a); i = i + 2 {
		var setter commandLineSetter
		var ok bool
		if setter, ok = commandLineSetters[a[i]]; !ok {
			return nil, errors.New(fmt.Sprintf("unknown argument flag: %s", a[i]))
		}
		setter(args, a[i+1])
	}
	return args, nil
}

type commandLineArguments struct {
	addr        string
	githubToken string
	role        string
	path        string
	key         string
}
