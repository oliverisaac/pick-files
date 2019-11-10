package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func Test_isFile(t *testing.T) {
	type args struct {
		fname string
	}

	existingFile, err := ioutil.TempFile(os.TempDir(), "test")
	if err != nil {
		log.Fatal("Failed to create temporary file")
	}
	existingFile.Close()
	defer os.Remove(existingFile.Name())

	nonExistingFile, err := ioutil.TempFile(os.TempDir(), "test")
	if err != nil {
		log.Fatal("Failed to create temporary file for deletion")
	}
	nonExistingFile.Close()
	os.Remove(nonExistingFile.Name())

	existingDirectory := os.TempDir()

	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Test real-file",
			args: args{fname: existingFile.Name()},
			want: true,
		},
		{
			name: "Test non-file",
			args: args{fname: nonExistingFile.Name()},
			want: false,
		},
		{
			name: "Test against directory",
			args: args{fname: existingDirectory},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFile(tt.args.fname); got != tt.want {
				t.Errorf("isFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isDir(t *testing.T) {
	type args struct {
		fname string
	}

	existingFile, err := ioutil.TempFile(os.TempDir(), "test")
	if err != nil {
		log.Fatal("Failed to create temporary file")
	}
	existingFile.Close()
	defer os.Remove(existingFile.Name())

	nonExistingFile, err := ioutil.TempFile(os.TempDir(), "test")
	if err != nil {
		log.Fatal("Failed to create temporary file for deletion")
	}
	nonExistingFile.Close()
	os.Remove(nonExistingFile.Name())

	existingDirectory := os.TempDir()

	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Test real-file",
			args: args{fname: existingFile.Name()},
			want: false,
		},
		{
			name: "Test non-file",
			args: args{fname: nonExistingFile.Name()},
			want: false,
		},
		{
			name: "Test against directory",
			args: args{fname: existingDirectory},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDir(tt.args.fname); got != tt.want {
				t.Errorf("isDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
