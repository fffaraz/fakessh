# FakeSSH
A fake dockerized SSH server that logs login attempts

```
docker run -it --rm -p 22:22 fffaraz/fakessh
```

```
docker run -d --restart=always -p 22:22 --name fakessh fffaraz/fakessh
docker logs -f fakessh
```

###See also

* [sshesame](https://github.com/jaksi/sshesame) - A fake SSH server that lets everyone in and logs their activity.
* [ssh-chat](https://github.com/shazow/ssh-chat) - Custom SSH server written in Go. Instead of a shell, you get a chat prompt.
