package checkdocker

import (
	"errors"
	"log"
	"os/exec"
	"strings"
)

func Check(dockerNames []string) error {
	out, err := exec.Command("docker", "ps", "--format", "{{.Names}}").Output()
	if err != nil {
		log.Fatal(err)
	}

	activeDocker := strings.Split(string(out), "\n")

	for _, v1 := range dockerNames {
		found := false
		for _, v2 := range activeDocker {
			if v1 == v2 {
				found = true
			}
		}

		if found == false {
			return errors.New("docker " + v1 + " not found ")
		}
	}

	return nil
}
