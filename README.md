# An Implementation of FakeTLS

<p align="center">
  <img src="./_images/server-after-handshake.PNG" width="738">
</p>

<p align="center">
  <img src="./_images/wireshark-handshake.PNG" width="738">
</p>

## Compiling and Running

### Server (Go)

- If necessary, change the `SERVER_PORT` constant, if necessary (default is `443`).
- Compile the server binary by running `make server`.
- Run the server with `sudo ./server`.
- Make note of the server IP address that is printed when starting the program.

### Client (C)

- Change the `SERVER_IP` definition to the IP address printed by the server program.
- If necessary, change the `SERVER_PORT` constant to the same used by the server code (default is `443`).
- Compile the client binary by running `make client`.
- Run the client with `./client`.

## Usage

<p align="center">
  <img src="./_images/reverse-shell-demo.PNG" width="738">
</p>

<p align="center">
  <img src="./_images/ls-command.PNG" width="738">
</p>

<p align="center">
  <img src="./_images/ls-response.PNG" width="738">
</p>
