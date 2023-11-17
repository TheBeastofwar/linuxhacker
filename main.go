package main

import (
	"flag"
	"fmt"
	plugins1 "linuxhacker/plugins/env_check"
	plugins3 "linuxhacker/plugins/maintance"
	plugins2 "linuxhacker/plugins/privilege_check"
)

func main() {
	banner()

	helpFlag := flag.Bool("h", false, "Show help information")
	pluginFlag := flag.String("p", "", "Choose the plugin")
	checkFlag := flag.String("c", "", "Choose the choice for the corresponding plugin")
	otherFlag := flag.String("q", "", "other parameter")

	flag.Parse()

	if *helpFlag || flag.NFlag() == 0 {
		fmt.Println("Usage: linuxhacker -p [plugin] -c [choice] -q [other]")
		fmt.Println("Options:")
		fmt.Println("-p\tchoose the plugin")
		fmt.Println("-c\tchoose the choice for the corresponding plugin")
		fmt.Println("-q\tother parameter")
		return
	}

	var plugin string
	var choice string
	var other string

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
	if *otherFlag != "" {
		other = *otherFlag
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
			fmt.Println("please input the dirctory to check(/usr/bin,/usr/sbin ... default to /usr/bin)")
			plugins2.SUID_Check(other)
		case "sudo":
			plugins2.SUDO_Check()
		case "cap":
			fmt.Println("please input the dirctory to check(/usr/bin,/usr/sbin ... default to /usr/bin)")
			plugins2.Capabilities_Check(other)
		case "kernel":
			plugins2.Kernel_Check()
		default:
			fmt.Println("Bad input privilege_check chocie(suid,sudo,cap,kernel)")
		}
	case "maintance", "3":
		switch choice {
		case "ssh_soft_link", "ssh1":
			plugins3.Soft_Link(other)
		case "ssh_public_key", "ssh2":
			plugins3.Public_Key(other)
		case "ssh_key_logger", "ssh3":
			plugins3.Key_Logger()
		case "ssh_stealth_login", "ssh4":
			plugins3.Stealth_Login()
		case "cron":
			plugins3.Cron_Add(other)
		case "file":
			plugins3.File_Change(other)
		case "command":
			fmt.Println("be careful if the command is overwritten!!!")
			plugins3.Command_Replace(other)
		default:
			fmt.Println("Bad input maintance choice(ssh_soft_link/ssh1,ssh_public_key/ssh2,ssh_key_logger/ssh3,ssh_stealth_login/ssh4,cron,file)")
		}
	default:
		fmt.Println("Bad input plugin(env_check/1,privilege_check/2,maintance/3)")
	}

}

func banner() {
	version := "1.1.0"
	fmt.Print(`
  _ _                  _                _             
 | (_)_ __  _   ___  _| |__   __ _  ___| | _____ _ __ 
 | | | '_ \| | | \ \/ / '_ \ / _` + "`" + ` |/ __| |/ / _ \ '__|
 | | | | | | |_| |>  <| | | | (_| | (__|   <  __/ |   
 |_|_|_| |_|\__,_/_/\_\_| |_|\__,_|\___|_|\_\___|_|
    `)
	fmt.Println("\t\t\t\tversion: " + version + "\n")
}
