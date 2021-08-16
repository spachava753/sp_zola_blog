package main

import (
	"os"
	"testing"
)

func Test_runCommand(t *testing.T) {
	type args struct {
		command string
		args    []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "echo",
			args: args{
				command: "echo",
				args:    []string{"hello"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := runCommand(tt.args.command, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("runCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getDeployments(t *testing.T) {
	type args struct {
		accountId string
		email     string
		apiToken  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "get deployments",
			args: args{
				accountId: os.Getenv("ACCOUNT_ID"),
				email:     os.Getenv("EMAIL"),
				apiToken:  os.Getenv("API_TOKEN"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if url, err := getLatestDeployment(tt.args.accountId, tt.args.email, tt.args.apiToken); (err != nil) != tt.wantErr {
				t.Errorf("getLatestDeployment() error = %v, wantErr %v", err, tt.wantErr)
				t.Logf("recieved url: %s\n", url)
			}
		})
	}
}
