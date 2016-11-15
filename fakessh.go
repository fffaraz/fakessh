package main

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

var tf = "01-02 15:04:05"

func main() {
	serverConfig := &ssh.ServerConfig{
		PasswordCallback: func(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
			fmt.Printf("%s\t%s\t%s\t%s\t%s\n", time.Now().Format(tf), conn.RemoteAddr(), string(conn.ClientVersion()), conn.User(), string(password))
			return nil, errors.New("Password authentication failed")
		},
		ServerVersion: "SSH-2.0-OpenSSH_6.6.1p1 Ubuntu-2ubuntu2.3",
	}
	keyBytes, _ := rsa.GenerateKey(rand.Reader, 2048)
	key, _ := ssh.NewSignerFromSigner(keyBytes)
	serverConfig.AddHostKey(key)
	listener, err := net.Listen("tcp", "0.0.0.0:22")
	if err != nil {
		log.Fatal("Failed to listen:", err.Error())
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Printf("%s\t%s\tConnected\n", time.Now().Format(tf), conn.RemoteAddr())
		go handleConn(serverConfig, conn)
	}
}

func handleConn(serverConfig *ssh.ServerConfig, conn net.Conn) {
	defer conn.Close()
	_, _, _, err := ssh.NewServerConn(conn, serverConfig)
	if err != nil {
		fmt.Printf("%s\t%s\t%s\n", time.Now().Format(tf), conn.RemoteAddr(), err.Error())
	}
	fmt.Printf("%s\t%s\tDisconnected\n", time.Now().Format(tf), conn.RemoteAddr())
}
