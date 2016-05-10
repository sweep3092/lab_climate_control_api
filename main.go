package main

import (
	"github.com/tarm/serial"
	"log"
)

func main() {
	c := &serial.Config{Name: "PORT", Baud: 9600, ReadTimeout: time.Second * 5}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
}

func up(s) {
	s.Write([]byte("up"))
}

func down(s) {
	s.Write([]byte("down"))
}

func cur(s) {
	s.Write([]byte("cur"))
	buf := make([]byte, 128)
	n, _ = s.Read(buf)
	log.Printf("%q", buf[:n])
}

func set(s) {
	s.Write([]byte("set"))
	buf := make([]byte, 128)
	n, _ = s.Read(buf)
	log.Printf("%q", buf[:n])
}
