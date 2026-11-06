# Jittermouse
Tired of:
- your computer getting blocked when you don't really want to?
- being forced to use Teams and deal with its erratic Absence detection?
- having to build cumbersome mechanisms in order to keep your PC as if it was busy?
- needing to purchase an external device?

Well, so do I, so I decided to build this.

A simple program that all it does is move your mouse in place a single pixel every ±60 seconds

## How to use
1. Download it from Releases tab or compile it from source
2. Run it via your command line of preference
3. ??
4. profit

### How to terminate it
You can terminate it anytime using `Ctrl+C`, but you can also use the key combo `Ctrl+Shift+L` and terminate it anywhere and do it gratiously (not that it matters much)

## Compile yourself
Is pretty straight forward, if you are doing it for your existing OS, `go build` will do it for you, but just in case you want to do it for other OS or architectures, you can also do it by using the proper tags.
> Note: There are flags for Windows that allow it to run on the background unnoticed, but that might flag the app as a virus, identifying it as a Keylogger (false-positive). Do what you will with that information.

### Build for MacOS

```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -tags "mac purego" -o build/jittermouse_macos_arm64 main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -tags "mac purego" -o build/jittermouse_macos_amd64 main.go
```

### Build for Windows

```bash
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -tags "win purego" -o build/jittermouse_win_x64.exe main.go
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -tags "win purego" -o build/jittermouse_win_x86.exe main.go
```

### Build for Linux

#### For X11

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags "x11 purego" -o build/jittermouse_x11_x64 main.go
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -tags "x11 purego" -o build/jittermouse_x11_x86 main.go
```

#### For Wayland

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags "wayland purego" -o build/jittermouse_wayland_x64 main.go
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -tags "wayland purego" -o build/jittermouse_wayland_x86 main.go
```

## Special thanks

- [Jitterbug](https://github.com/lthibault/jitterbug) for allowing some jitter on the mouse movement to make it not so mechanic.
- [RobotGo](https://github.com/go-vgo/robotgo) for all of the tools that allow reading and doing inputs.
- [GoHook](https://github.com/robotn/gohook) for allowing asynchronous and global key inputs.
