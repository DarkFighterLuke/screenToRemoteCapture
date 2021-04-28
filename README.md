# screenToRemoteCapture
A simple tool to send screenshots to the clipboard of another computer.

# Platforms
- Linux
- Windows
- MacOS (not tested, waiting for someone)

# Requirements
`xclip` is needed on Linux.

# Usage
- Clone repository and `go build` it.
- Edit configuration file (or create a new one following the same schema) to configure screen region to capture and server address and port where to send screenshots. (<b>Client</b>)
- Edit configuration file (or create a new one following the same schema) to configure server listening address and port to receive incoming screenshots. (<b>Server</b>)

# Build
## Linux dependencies
- Go 1.16 needed
- `gcc` is needed
- `xcb` `libxcb-xkb-dev x11-xkb-utils` `libx11-dev` `libx11-xcb-dev` `libxkbcommon-x11-dev` are needed

## Windows dependencies
- Go 1.16 is needed
- `MINGW-W64` is needed (you can find the right 64 bit version [here](https://sourceforge.net/projects/mingw-w64/files/Toolchains%20targetting%20Win32/Personal%20Builds/mingw-builds/installer/mingw-w64-install.exe/download)

## Compile
```
git clone https://github.com/DarkFighterLuke/screenToRemoteCapture.git
cd screenToRemoteCapture
go build .
```
