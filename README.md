# gpctl

NB: I threw this together in one and a half hours. It is likely that something will blow up if you use it.

## Introduction

A simple CLI interface for managing your GoPro Hero 12.

## Usage

```shell
# List all available devices
gpctl devices

# List all files and directories in the root folder
gpctl ls /

# Download a file from the device
gpctl cp /100GOPRO/GX010008.MP4 /home/user/Downloads/awesome-vid.mp4

# Delete a file from the device
gpctl rm /100GOPRO/GX010008.MP4

# Start webcam mode
gpctl webcam start

# Stop webcam mode
gpctl webcam stop
```

## Installation

To build the code, you will need to have `Go 1.21` or higher installed.

```shell
# Install in default location (~/.local/bin/gpctl)
make build && make install

# Use PREFIX to change the install location
make build && make install PREFIX=/usr/local/bin

# Uninstall from default location
make uninstall
```

## Sources

- [KonradIT's awesome code](https://github.com/KonradIT/mmt)
- [Open GoPro's HTTP specification](https://gopro.github.io/OpenGoPro/)
