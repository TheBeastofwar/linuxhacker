package plugins

import (
	"fmt"
	"os/exec"
	"strings"
)

func VM_Check() {
	fmt.Println("virtual machine checking...")
	cmd := exec.Command("dmidecode", "-s", "system-product-name")
	result, err := cmd.CombinedOutput()
	if err != nil {
		return
	}
	fmt.Println("result:")
	if strings.Contains(string(result), "VM") {
		fmt.Println("it may be in the virtual machine")
	} else {
		fmt.Println("it may not be in the virtual machine")
	}
}
