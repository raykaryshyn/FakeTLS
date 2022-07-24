# An Implementation of FakeTLS

## Usage

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
