// Copyright 2017 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tcpopt_test

import (
	"log"
	"net"
	"time"

	"github.com/liuxp0827/tcp-option"
	"github.com/mikioh/tcp"
)

func ExampleOption() {
	c, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	tc, err := tcp.NewConn(c)
	if err != nil {
		log.Fatal(err)
	}
	if err := tc.SetOption(tcpopt.KeepAlive(true)); err != nil {
		log.Fatal(err)
	}
	if err := tc.SetOption(tcpopt.MSS(1100)); err != nil {
		log.Fatal(err)
	}
	if err := tc.SetOption(tcpopt.TOA{
		Port: 9000,
		Ip:   net.ParseIP("127.0.0.5"),
	}); err != nil {
		log.Fatal(err)
	}
	if err := tc.SetOption(tcpopt.KeepAliveIdleInterval(3 * time.Minute)); err != nil {
		log.Fatal(err)
	}
	if err := tc.SetOption(tcpopt.KeepAliveProbeInterval(30 * time.Second)); err != nil {
		log.Fatal(err)
	}
	if err := tc.SetOption(tcpopt.KeepAliveProbeCount(3)); err != nil {
		log.Fatal(err)
	}

	if err := c.SetReadDeadline(time.Now().Add(3 * time.Second)); err != nil {
		log.Fatal(err)
	}

	log.Println("hello world!")

}
