# jrctl
> Command line tool to help interact with our API

![](https://img.shields.io/badge/License-JetRails_License-green.svg?style=for-the-badge&labelColor=89BA40&color=282F38)
![](https://img.shields.io/badge/Version-2.1.4-green.svg?style=for-the-badge&labelColor=89BA40&color=282F38)
![](https://img.shields.io/badge/OS-MacOS/Linux-green.svg?style=for-the-badge&labelColor=89BA40&color=282F38)

## About

Our command-line tool will help you interact with the JetRails API. Services are currently limited because we are actively working on exposing our API to the public. Future versions of our CLI tool will extend functionality. Stay tuned!

## Installation (MacOS)

```shell
$ brew tap jetrails/tap
$ brew install jrctl
```

## Installation (RedHat)

```shell
$ rpm -i https://github.com/jetrails/jrctl/releases/download/latest/jrctl_latest_linux_amd64.rpm
```

## Installation (Debian)

```shell
$ curl -sL -o /var/cache/apt/archives/jrctl_latest_linux_amd64.deb https://github.com/jetrails/jrctl/releases/download/latest/jrctl_latest_linux_amd64.deb
$ dpkg -i /var/cache/apt/archives/jrctl_latest_linux_amd64.deb
$ rm /var/cache/apt/archives/jrctl_latest_linux_amd64.deb
```

## Building & Running

This project uses a simple [Makefile](./Makefile) as a build system.

|     Command        | Description                               |
|:------------------:|-------------------------------------------|
|  `make build`      | Build binaries and output to `bin` folder |
|  `make build-all`  | Build for all platforms                   |
|  `make clean`      | Delete built binaries                     |
|  `make docs`       | Generate documentation                    |
| `make package`     | Package binary for many distributions     |
| `make format`      | Format code with goimports                |
|  `make help`       | Display available commands                |

## Configuration File

A configuration file is used to load user settings. It can be found in `~/.jrctl/config.yaml`. Some features can be configured via environmental variables.
