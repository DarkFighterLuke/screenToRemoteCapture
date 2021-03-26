# screenToRemoteCapture
A simple tool to send screenshots of the screen to the clipboard of another computer.

# Platforms
- Linux
- Windows (only client)

# Dependencies
No dependencies are needed on any platforms if you are running the client.
`xclip` is needed on Linux for running the server.

# Usage
Clone repository and `go build` it.
Edit configuration file (or create a new one following the same schema) to configure screen region to capture and server address and port where to send screenshots. (Client)
Edit configuration file (or create a new one following the same schema) to configure server listening address and port to receive incoming screenshots. (Server)
