package main

import (
	"io/ioutil"
	"path"
	"os/exec"
)

func main() {

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

func RunCmd(command []string, env map[string]string) int {

	cmd := exec.Command(command)

	cmd.Run()
}
