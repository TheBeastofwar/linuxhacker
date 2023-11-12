package privilege_check

import (
	"fmt"
	"os/exec"
)

func Kernel_Check() {
	fmt.Println("kernel vulnerability checking...")
	cmd := exec.Command("perl", "./plugins/poc/os_poc/linux-exploit-suggester-2.pl")
	result, err := cmd.CombinedOutput()
	if err != nil {
		return
	}
	fmt.Println(string(result))
}
