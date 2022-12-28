package main

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
)

var errPassword = errors.New("password authentication failed")

func main() {
	if len(os.Args) > 1 {
		logPath := os.Args[1]
		logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer logFile.Close()
		log.SetOutput(logFile)
	}

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	serverConfig := &ssh.ServerConfig{
		MaxAuthTries:     6,
		PasswordCallback: passwordCallback,
		ServerVersion:    "SSH-2.0-OpenSSH_6.6.1p1 Ubuntu-2ubuntu2.3",
	}

	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	signer, _ := ssh.NewSignerFromSigner(privateKey)
	serverConfig.AddHostKey(signer)

	listener, err := net.Listen("tcp", ":22")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept:", err)
			continue
		}
		go handleConn(conn, serverConfig)
	}
}

func passwordCallback(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
	log.Println(conn.RemoteAddr(), string(conn.ClientVersion()), conn.User(), string(password))
	return nil, errPassword
}

func handleConn(conn net.Conn, serverConfig *ssh.ServerConfig) {
	defer conn.Close()
	log.Println(conn.RemoteAddr())
	ssh.NewServerConn(conn, serverConfig)
}
