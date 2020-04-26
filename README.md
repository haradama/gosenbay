# AutomotiveSenbay

## Environment

- Ubuntu

- gocv.io/x/gocv
- github.com/spf13/cobra
- github.com/kbinani/screenshot

## Install
### Install dependance libraries
```
sudo apt-get update
sudo apt-get upgrade
sudo apt-get install can-utils
```

### Add Virtual CAN
```
sudo modprobe vcan
sudo ip link add dev vcan0 type vcan
sudo ip link set vcan0 up
```

## Usage

```
docker run -it --rm --privileged --name autobay autobay /bin/bash
```

## License
AutomotiveSenbay is available under the Apache License, Version 2.0 license. See the LICENSE file for more info.