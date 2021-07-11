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
// key is file name and value is file content
func ReadDir(path string) (map[string]string, error) {
	err := IsDirExists(path)
	if err != nil {
		return nil, fmt.Errorf("[utils.ReadDir] invalid path (%s) provided: %w", path, err)
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("[utils.ReadDir] cannot scan directory: %w", err)
	}

	if len(files) == 0 {
		return nil, fmt.Errorf("[utils.ReadDir] provided directory %s is empty", path)
	}

	env := make(map[string]string)
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		content, err := ioutil.ReadFile(path + "/" + file.Name())
		if err != nil {
			log.Printf("[utils.ReadDir] cannot read file content %s: %s", path + "/" + file.Name(), err)

			env[file.Name()] = ""
		} else {
			env[file.Name()] = string(content)
		}
	}

	return env, nil
}

// RunCommand run command with arguments and environment variables
func RunCommand(cmd []string, env map[string]string) int {
	var envStrings []string
	for key,val := range env {
		envStrings = append(envStrings, fmt.Sprintf("%s=%s", key, val))
	}

	command := exec.Command(cmd[0], cmd[1:]...)
	command.Env = envStrings
	command.Stdout = os.Stdout

	err := command.Run()
	if err != nil {
		return 1
	}

	return 0
}

// IsDirExists check if provided path exists and it's directory
func IsDirExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("[utils.IsDirExists] directory %s is not exists", path)
	}

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("[utils.IsDirExists] cannot open file: %w", err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("[utils.IsDirExists] cannot read file info: %w", err)
	}

	if !fileInfo.IsDir() {
		return fmt.Errorf("[utils.IsDirExists] provided path is not a directory")
	}

	return nil
}