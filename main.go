package main

import (
	"github.com/docker-machine-driver-aliyunecs/aliyunecs"
	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(aliyunecs.NewDriver("", ""))
}
