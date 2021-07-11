// Package utils package provide functions that implement all basic features required for envdir cli tool
package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// ReadDir function parse all files in provided path and
// convert it to map of environment variables where
// key is file name and value is file content.
func ReadDir(path string) (map[string]string, error) {
	err := IsDirExists(path)
	if err != nil {
		return nil, &InvalidPathError{err, path}
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, &CannotScanDirError{err}
	}

	if len(files) == 0 {
		return nil, &EnvDirIsEmptyError{path}
	}

	env := make(map[string]string)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		content, err := ioutil.ReadFile(path + "/" + file.Name())
		if err != nil {
			log.Printf("[utils.ReadDir] cannot read file content %s: %s", path+"/"+file.Name(), err)

			env[file.Name()] = ""
		} else {
			env[file.Name()] = string(content)
		}
	}

	return env, nil
}

// RunCommand run command with arguments and environment variables.
func RunCommand(cmd []string, env map[string]string) int {
	envStrings := make([]string, 0, len(env))
	for key, val := range env {
		envStrings = append(envStrings, fmt.Sprintf("%s=%s", key, val))
	}

	command := exec.Command(cmd[0], cmd[1:]...)
	command.Env = envStrings
	command.Stdout = os.Stdout

	if command.Run() != nil {
		return 1
	}

	return 0
}

// IsDirExists check if provided path exists and it's directory.
func IsDirExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return &EnvDirIsNotExistsError{path}
	}

	file, err := os.Open(path)
	if err != nil {
		return &CannotOpenEnvDirError{err, path}
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return &CannotReadFileInfoError{err, path}
	}

	if !fileInfo.IsDir() {
		return &ProvidedPathIsNotDirError{path}
	}

	return nil
}
