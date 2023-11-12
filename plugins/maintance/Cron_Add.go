package maintance

import (
	"fmt"
	"os"
	"os/exec"
)

func Cron_Add(input string, input2 string) {
	fmt.Println("add something in the crontab to maintance the privilege")
	cmd := exec.Command("sh", "-c", fmt.Sprintf("echo '%s' >> %s", input, input2))
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	fmt.Println("success")
}
