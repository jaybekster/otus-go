package main

import (
	"flag"
	"os"
	"fmt"
	"io/ioutil"
	"path"
	"os/exec"
)

func init() {
	flag.Parse();
}

func main() {
	command := flag.Arg(1);
	// path := flag.Arg(0);

	// fmt.Println(command, path);

	// fmt.Println(ReadDir(path));

	cmdError := RunCmd(command, flag.Args()[2:]...)

	if cmdError != nil {
		fmt.Println(cmdError, "error")
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

func RunCmd(command string, args ...string) error {
	cmd := exec.Command(command, args...)

	cmd.Stdout = os.Stdout;

	if err := cmd.Run(); err != nil {
		return err;
	}

	return nil;
}
