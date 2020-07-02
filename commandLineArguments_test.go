package main

import (
	"reflect"
	"testing"
)

func Test_getCommandLineArguments(t *testing.T) {
	okArr := []string{"command", "-addr", "https://myvault", "-token", "myToken", "-role", "myRole"}
	expectedArgs := &commandLineArguments{
		addr:        "https://myvault",
		githubToken: "myToken",
		role:        "myRole",
	}
	type args struct {
		arr []string
	}
	tests := []struct {
		name    string
		args    args
		want    *commandLineArguments
		wantErr bool
	}{
		{"empty", args{[]string{"command"}}, nil, true},
		{"odd", args{[]string{"command", "-addr", "my-address", "ss"}}, nil, true},
		{"unknown", args{[]string{"command", "-addr", "my-address", "-unknown", "val"}}, nil, true},
		{"ok", args{okArr}, expectedArgs, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCommandLineArguments(tt.args.arr)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCommandLineArguments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCommandLineArguments() got = %v, want %v", got, tt.want)
			}
		})
	}
}