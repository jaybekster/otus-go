package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
)

func init() {
	flag.Parse()
}

func main() {
	command := flag.Arg(1)
	path := flag.Arg(0)

	fmt.Println(path)

	var err error

	commandWithArgs := append([]string{command}, flag.Args()[2:]...)

	envs, err := ReadDir(path)

	if err != nil {
		log.Fatalf("Failed read directory %s\n", err)
	}

	err = RunCmd(commandWithArgs, envs)

	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func ReadDir(dir string) (map[string]string, error) {
	envs := make(map[string]string)

	files, err := ioutil.ReadDir(dir)

	if err != nil {
		return envs, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()

		content, err := ioutil.ReadFile(path.Join(dir, fileName))

		if err != nil {
			return envs, err
		}

		envs[fileName] = string(content)
	}

	return envs, nil
}

func RunCmd(commandWithArgs []string, env map[string]string) error {
	var envsArray []string

	for key, value := range env {
		if len(value) != 0 {
			envsArray = append(envsArray, key+"="+value)
		} else {
			os.Unsetenv(key)
		}
	}

	cmd := exec.Command(commandWithArgs[0], commandWithArgs[1:]...)
	cmd.Env = append(os.Environ(), envsArray...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
