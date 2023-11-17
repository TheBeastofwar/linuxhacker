package privilege_check

import (
	"fmt"
	"os/exec"
)

func Kernel_Check() {
	fmt.Println("kernel vulnerability checking...")
	//"uname","-r","|cut -d'-' -f1"
	cmd := exec.Command("uname -r |cut -d'-' -f1")
	result, err := cmd.CombinedOutput()
	if err != nil {
		return
	}
	fmt.Println(string(result))
}
