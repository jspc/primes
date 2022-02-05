//go:build eggos
// +build eggos

package main

import (
	"runtime"

	"github.com/icexin/eggos/console"
	"github.com/icexin/eggos/drivers/cga/fbcga"
	_ "github.com/icexin/eggos/drivers/e1000"
	"github.com/icexin/eggos/drivers/pci"
	"github.com/icexin/eggos/drivers/uart"
	"github.com/icexin/eggos/drivers/vbe"
	"github.com/icexin/eggos/inet"
	"github.com/icexin/eggos/kernel"
)

func kernelInit() {
	// trap and syscall threads use two Ps,
	// and the remainings are for other goroutines
	runtime.GOMAXPROCS(6)

	kernel.Init()
	uart.Init()
	console.Init()

	vbe.Init()
	fbcga.Init()
	pci.Init()
	inet.Init()
}

func init() {
	kernelInit()
}