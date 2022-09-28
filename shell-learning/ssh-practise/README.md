## SSH-Expect

### 功能
**把用户名密码写入hosts，通过 myssh 连接，这样就不需要输入密码了**

### 需要安装 expect
`yum install -y expect`


**可以通过 myssh 查看当前哪些主机可以登录**
```
./myssh
当前可连接主机:
docker
连接远程主机:
myssh [主机名]
```

**可以通过 myssh 登录主机**
```
./myssh docker
连接主机:10.249.4.155
spawn ssh -p 22 root@10.249.4.155
root@10.249.4.155's password:
Activate the web console with: systemctl enable --now cockpit.socket

Last login: Wed Sep 28 18:19:46 2022 from 10.249.4.171
```