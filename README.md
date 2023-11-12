# linuxhacker
## 使用方法
```
  _ _                  _                _             
 | (_)_ __  _   ___  _| |__   __ _  ___| | _____ _ __ 
 | | | '_ \| | | \ \/ / '_ \ / _` |/ __| |/ / _ \ '__|
 | | | | | | |_| |>  <| | | | (_| | (__|   <  __/ |   
 |_|_|_| |_|\__,_/_/\_\_| |_|\__,_|\___|_|\_\___|_|
    				version: 1.0.0

Usage: golang -p [plugin] -c [choice]
Options:
-p	choose the plugin
-c	choose the choice for the corresponding plugin
```
插件:
- env_check  选项: docker(检查是否处于docker环境),vm(检查是否处于虚拟机中)
- privilege_check 选项: suid,sudo,cap,kernel
- maintance 选项: ssh_soft_link/ssh1,ssh_public_key/ssh2,ssh_key_logger/ssh3,ssh_stealth_login/ssh4,cron,file
- ssh_soft_link/ssh1 ssh软链接权限维持
- ssh_public_key/ssh2 ssh写公钥权限维持
- ssh_key_logger/ssh3 ssh key logger可以记录下别人ssh登录的密码
- ssh_stealth_login/ssh4 ssh隐身登录,w,last等命令查找不到
- cron 计划任务提权
- file 更改文件的atime,mtiem,实现webshell和一些恶意代码的权限维持
![XF8~C6G2(KEBC8(`~K)E7MG](https://github.com/TheBeastofwar/linuxhacker/assets/117450378/7e2f3549-90d3-410e-a4c6-d8e3a74d5c68)
![U~NJTH%OHJ5K IXLHKK9G6](https://github.com/TheBeastofwar/linuxhacker/assets/117450378/a43b5abe-5016-4b70-87fc-3f8358875fa9)
## 开发日志
### 2023 11.12 v1.0.0
- 可探测当前linux系统所处环境(是否是docker,vm等)
- 可探测当前linux系统的各种提权方法(suid,sudo等)
- 自带一些权限维持的模块
- 遗留问题:1.需要将perl写的内核漏洞探测工具用go重构 2.suid,sudo,capilities的if语句有中屎山代码的感觉 3.需要更改让部分模块的输入从fmt.Scanln改为直接Options传参,这样在使用webshell管理工具的时候能方便些
