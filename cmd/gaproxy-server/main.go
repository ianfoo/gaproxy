// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: feec04999c
// Version Date: Sat Mar 16 17:27:47 UTC 2019

package main

import (
	"flag"

	// This Service
	"github.com/ianfoo/gaproxy/svc/server"
	"github.com/ianfoo/gaproxy/svc/server/cli"
)

func main() {
	// Update addresses if they have been overwritten by flags
	flag.Parse()

	server.Run(cli.Config)
}
