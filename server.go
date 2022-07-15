package main

import (
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
        fmt.Println("Server Running...")
        server, err := net.Listen(SERVER_TYPE, ":" + SERVER_PORT)
        if err != nil {
                fmt.Println("Error listening:", err.Error())
                os.Exit(1)
        }
        defer server.Close()
        fmt.Println("Listening on port " + SERVER_PORT)
        fmt.Println("Waiting for client...")
        for {
                connection, err := server.Accept()
                if err != nil {
                        fmt.Println("Error accepting: ", err.Error())
                        os.Exit(1)
                }
                fmt.Println("client connected")
                go processClient(connection)
        }
}

func processClient(connection net.Conn) {
        stdoutDumper := hex.Dumper(os.Stdout)
        defer stdoutDumper.Close()

        buffer := make([]byte, 50000)
        mLen, err := connection.Read(buffer)
        fmt.Println("Length: ", mLen)
        if err != nil {
                fmt.Println("Error reading:", err.Error())
        }
        //fmt.Println("Received: ", string(buffer[:mLen]))
        stdoutDumper.Write([]byte(buffer[:mLen]))
        _, err = connection.Write([]byte("ls -al"))
        connection.Close()
}
