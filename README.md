# jrctl
> Command line tool to help interact with our API

![](https://img.shields.io/badge/License-JetRails_License-green.svg?style=for-the-badge&labelColor=89BA40&color=282F38)
![](https://img.shields.io/badge/Version-1.0.4-green.svg?style=for-the-badge&labelColor=89BA40&color=282F38)
![](https://img.shields.io/badge/OS-MacOS/Linux-green.svg?style=for-the-badge&labelColor=89BA40&color=282F38)

## About

Our command-line tool will help you interact with the JetRails API. Services are currently limited because we are actively working on exposing our API to the public. Future versions of our CLI tool will extend functionality. Stay tuned!

## Installation (MacOS)

```shell
$ brew tap jetrails/tap
$ brew install jrctl
```

## Installation (Linux)

```shell
$ curl -Ls https://github.com/jetrails/jrctl/releases/download/1.0.4/jrctl_linux_amd64 -o jrctl
$ chmod +x jrctl
$ sudo mv jrctl /usr/local/bin/jrctl
```

## Currently Exposed Services

|   Command   | Description                               |
|:-----------:|-------------------------------------------|
| `secret`    | Interact with our one-time secret service |
| `whitelist` | Whitelist IP addresses & ports            |

## Building & Running

This project uses a simple [Makefile](./Makefile) to build the final binary.

|    Command   | Description                             |
|:------------:|-----------------------------------------|
| `make build` | Build binary and output to `bin` folder |
| `make clean` | Delete built binaries                   |
| `make help`  | Display available commands              |

## Configuration File

A configuration file is used to load user settings. It can be found in `~/.jrctl.yaml`. Environmental variables can also be used to override the configured settings.
