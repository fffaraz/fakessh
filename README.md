# FakeSSH
A dockerized fake SSH server that logs login attempts

```
docker run -it --rm -p 22:22 fffaraz/fakessh
```

OR

```
docker run -d --restart=always -p 22:22 --name fakessh fffaraz/fakessh
docker logs -f fakessh
```


### See also

* [sshesame](https://github.com/jaksi/sshesame) - A fake SSH server that lets everyone in and logs their activity.
* [ssh-chat](https://github.com/shazow/ssh-chat) - Custom SSH server written in Go. Instead of a shell, you get a chat prompt.
* [gliderlabs/ssh](https://github.com/gliderlabs/ssh) - Easy SSH servers in Golang.
* [gliderlabs/sshfront](https://github.com/gliderlabs/sshfront) - Programmable SSH frontend.
* [desaster/kippo](https://github.com/desaster/kippo) - Kippo - SSH Honeypot.
* [micheloosterhof/cowrie](https://github.com/micheloosterhof/cowrie) - Cowrie SSH/Telnet Honeypot.
