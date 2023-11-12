package plugins

import (
	"fmt"
	"os/exec"
)

func Docker_Check() {
	fmt.Println("docker environment checking...")
	cmd := exec.Command("fdisk", "-l")
	result, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("result:it may be in the docker volumn")
		return
	}
	if result != nil {
		fmt.Println("result:it may not be in the docker volumn")
	}
}
