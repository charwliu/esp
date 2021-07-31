package main

import (
	"flag"
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var directory = flag.String("pwd", "", "Working directory of ESP eprotogen.")

func main() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage of eprotogen:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if !filepath.IsAbs(*directory) {
		pwd, wdErr := os.Getwd()
		if wdErr != nil {
			fmt.Println("Can not get current working directory.")
			os.Exit(1)
		}
		*directory = filepath.Join(pwd, *directory)
	}

	pwd := *directory
	GOBIN := GetGOBIN()
	binPath := os.Getenv("PATH")
	pathSlice := []string{binPath, GOBIN, pwd}
	binPath = strings.Join(pathSlice, string(os.PathListSeparator))
	os.Setenv("PATH", binPath)

	EXE := ""
	if runtime.GOOS == "windows" {
		EXE = ".exe"
	}
	protoc := "protoc" + EXE

	if path, err := exec.LookPath(protoc); err != nil {
		fmt.Println("Make sure that you have `" + protoc + "` in your system path or current path. To download it, please visit https://github.com/protocolbuffers/protobuf/releases")
		os.Exit(1)
	} else {
		protoc = path
	}

	protoFilesMap := make(map[string][]string)
	walkErr := filepath.Walk(pwd, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		dir := filepath.Dir(path)
		filename := filepath.Base(path)
		if strings.HasSuffix(filename, ".proto") {
			path = path[len(pwd)+1:]
			protoFilesMap[dir] = append(protoFilesMap[dir], path)
		}

		return nil
	})
	if walkErr != nil {
		fmt.Println(walkErr)
		os.Exit(1)
	}

	for _, files := range protoFilesMap {
		for _, relProtoFile := range files {
			var args []string

			args = []string{"--go_out", pwd, "--go_opt", "paths=source_relative", "--go-grpc_out", pwd, "--go-grpc_opt", "paths=source_relative", "--plugin", "protoc-gen-go=" + GOBIN + "/protoc-gen-go" + EXE, "--plugin", "protoc-gen-go-grpc=" + GOBIN + "/protoc-gen-go-grpc" + EXE}

			args = append(args, relProtoFile)
			cmd := exec.Command(protoc, args...)
			cmd.Env = append(cmd.Env, os.Environ()...)
			cmd.Dir = pwd
			output, cmdErr := cmd.CombinedOutput()
			if len(output) > 0 {
				fmt.Println(string(output))
			}
			if cmdErr != nil {
				fmt.Println(cmdErr)
				os.Exit(1)
			}
		}
	}
}

func envFile() (string, error) {
	if file := os.Getenv("GOENV"); file != "" {
		if file == "off" {
			return "", fmt.Errorf("GOENV=off")
		}
		return file, nil
	}
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	if dir == "" {
		return "", fmt.Errorf("missing user-config dir")
	}
	return filepath.Join(dir, "go", "env"), nil
}

// GetRuntimeEnv returns the value of runtime environment variable,
// that is set by running following command: `go env -w key=value`.
func GetRuntimeEnv(key string) (string, error) {
	file, err := envFile()
	if err != nil {
		return "", err
	}
	if file == "" {
		return "", fmt.Errorf("missing runtime env file")
	}
	var data []byte
	var runtimeEnv string
	data, readErr := os.ReadFile(file)
	if readErr != nil {
		return "", readErr
	}
	envStrings := strings.Split(string(data), "\n")
	for _, envItem := range envStrings {
		envItem = strings.TrimSuffix(envItem, "\r")
		envKeyValue := strings.Split(envItem, "=")
		if strings.EqualFold(strings.TrimSpace(envKeyValue[0]), key) {
			runtimeEnv = strings.TrimSpace(envKeyValue[1])
		}
	}
	return runtimeEnv, nil
}

// GetGOBIN returns GOBIN environment variable as a string. It will NOT be empty.
func GetGOBIN() string {
	// The one set by user explicitly by `export GOBIN=/path` or `env GOBIN=/path command`
	GOBIN := os.Getenv("GOBIN")
	if GOBIN == "" {
		var err error
		// The one set by user by running `go env -w GOBIN=/path`
		GOBIN, err = GetRuntimeEnv("GOBIN")
		if err != nil {
			// The default one that Golang uses
			return filepath.Join(build.Default.GOPATH, "bin")
		}
		if GOBIN == "" {
			return filepath.Join(build.Default.GOPATH, "bin")
		}
		return GOBIN
	}
	return GOBIN
}
