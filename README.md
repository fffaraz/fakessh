# FakeSSH
A fake dockerized SSH server that logs login attempts

```
docker run -it --rm -p 22:22 fffaraz/fakessh
```

```
docker run -d --restart=always -p 22:22 --name fakessh fffaraz/fakessh
docker logs -f fakessh
```
