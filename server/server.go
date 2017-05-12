package main

import (
	"log"
	"fmt"
	"net"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	//"bufio"
)

const (
	HOST = "localhost"
	PORT = "6969"
	TYPE = "tcp"
)

var file = os.Getenv("HOME")+"/todo.txt-test"
var todo []string
var txt string
var conn net.Conn
func main(){
	//list := []string

	dat, err := ioutil.ReadFile(file)
	txt = string(dat)
	todo = strings.Split(txt, "\n")

	l, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	print("Listening on " + HOST + ":" + PORT)
	for {
		// Listen for an incoming connection.
		var err error
		conn, err = l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		print("go handle")
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn){

	log.Println("Accepted new connection.")

	defer conn.Close()
	defer log.Println("Closed connection.")

	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			if err.Error() != "EOF" {
				print("error reading", err.Error())
			}
			return
		}

		if len(string(buf)) <3 {
			return
		}

		str := string(buf)
		cmd := str[0:3]
		dat := str[3:len(str)-1]

		print("cmd", cmd, "\n")
		print("dat", dat, "\n")

		switch cmd {
		case "get":
			send(txt)
		case "add":
			add(dat)
		case "rm-":
			rm(dat)
		default:
			print("\nUnkown response: ", str)
		}
	}
}

func add(s string) {
	print("add:", s)
	todo = append(todo, s)
	txt = strings.Join(todo, "\n")
	_ = ioutil.WriteFile(file, []byte(txt+"\n"), 0644)
}

func rm(s string){
	i, _ := strconv.Atoi(s)
	todo[i] = ""
	for !(todo[len(todo)-1] == "") {
		todo = todo[0:len(todo)-1]
	}
}

func send(s string){
	conn.Write([]byte(s))
}
