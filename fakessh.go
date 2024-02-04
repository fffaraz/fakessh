package main

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

var (
	errBadPassword = errors.New("permission denied")
	serverVersions = []string{
		"SSH-2.0-OpenSSH_6.6.1p1 Ubuntu-2ubuntu2.3",
		"SSH-2.0-OpenSSH_6.7p1 Debian-5+deb8u3",
		"SSH-2.0-OpenSSH_7.2p2 Ubuntu-4ubuntu2.10",
		"SSH-2.0-OpenSSH_7.4",
		"SSH-2.0-OpenSSH_8.0",
		"SSH-2.0-OpenSSH_8.4p1 Debian-2~bpo10+1",
		"SSH-2.0-OpenSSH_8.4p1 Debian-5+deb11u1",
		"SSH-2.0-OpenSSH_8.9p1 Ubuntu-3ubuntu0.6",
	}
)

func main() {
	if len(os.Args) > 1 {
		logPath := fmt.Sprintf("%s/fakessh-%s.log", os.Args[1], time.Now().Format("2006-01-02-15-04-05-000"))
		logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			log.Println("Failed to open log file:", logPath, err)
			return
		}
		defer logFile.Close()
		log.SetOutput(logFile)
	}

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	serverConfig := &ssh.ServerConfig{
		MaxAuthTries:     6,
		PasswordCallback: passwordCallback,
		ServerVersion:    serverVersions[0],
	}

	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	signer, _ := ssh.NewSignerFromSigner(privateKey)
	serverConfig.AddHostKey(signer)

	listener, err := net.Listen("tcp", ":22")
	if err != nil {
		log.Println("Failed to listen:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept:", err)
			break
		}
		go handleConn(conn, serverConfig)
	}
}

func passwordCallback(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
	log.Println(conn.RemoteAddr(), string(conn.ClientVersion()), conn.User(), string(password))
	time.Sleep(100 * time.Millisecond)
	return nil, errBadPassword
}

func handleConn(conn net.Conn, serverConfig *ssh.ServerConfig) {
	defer conn.Close()
	log.Println(conn.RemoteAddr())
	ssh.NewServerConn(conn, serverConfig)
}
