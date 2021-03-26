# screenToRemoteCapture
A simple tool to send screenshots of the screen to the clipboard of another computer.

# Platforms
- Linux
- Windows (only client)

# Requirements
Go 1.16 needed.<br>
`gcc` is needed.
`xclip` is needed on Linux.

# Usage
- Clone repository and `go build` it.
- Edit configuration file (or create a new one following the same schema) to configure screen region to capture and server address and port where to send screenshots. (<b>Client</b>)
- Edit configuration file (or create a new one following the same schema) to configure server listening address and port to receive incoming screenshots. (<b>Server</b>)
