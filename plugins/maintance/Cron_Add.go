package maintance

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Cron_Add(dir_cron string) {
	fmt.Println("add something in the crontab to maintance the privilege")
	dircron := strings.Split(dir_cron, "#")
	dir := dircron[0]
	cron := dircron[1]
	cmd := exec.Command("sh", "-c", fmt.Sprintf("echo '%s' >> %s", dir, cron))
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("create the %s in the directory of %s", cron, dir))
	fmt.Println("success")
}
