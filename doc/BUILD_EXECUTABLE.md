# How to build executable for your OS

Instruction to build an executable for a custom Operation System.

## Pre requisites

- Golang >= 1.22.3
- Clone this repository

## How to

With the `shell/terminal/cmd` enter the cloned folder and run one of the bellow commands.
After that, the executable file will be created in this same folder and it'll be possible to run it like: `./gitclean` or `.gitclean` or `.\gitclean`.

### Build for my current OS

```shell
go build .
```

### Build for Linux 64bit

```shell
GOOS=linux GOARCH=amd64 go build .
```

### Build for Windows 64bit

```shell
GOOS=windows GOARCH=amd64 go build .
```

## List all available OS

```shell
go tool dist list
```
