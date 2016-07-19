// Copyright 2016 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tcpopt

import (
	"time"
	"unsafe"
)

const (
	sysSOL_SOCKET = 0xffff

	sysSO_SNDBUF    = 0x1001
	sysSO_RCVBUF    = 0x1002
	sysSO_KEEPALIVE = 0x8

	sysTCP_NODELAY = 0x1

	sysIOC_OUT            = 0x40000000
	sysIOC_IN             = 0x80000000
	sysIOC_VENDOR         = 0x18000000
	sysSIO_KEEPALIVE_VALS = sysIOC_IN | sysIOC_VENDOR | 4
)

var options = [soMax]option{
	soNodelay:      {ianaProtocolTCP, sysTCP_NODELAY, 0},
	soSndbuf:       {sysSOL_SOCKET, sysSO_SNDBUF, 0},
	soRcvbuf:       {sysSOL_SOCKET, sysSO_RCVBUF, 0},
	soKeepalive:    {sysSOL_SOCKET, sysSO_KEEPALIVE, 0},
	soKeepidle:     {ianaProtocolTCP, sysSIO_KEEPALIVE_VALS, time.Millisecond},
	soKeepintvl:    {ianaProtocolTCP, sysSIO_KEEPALIVE_VALS, time.Millisecond},
	soKeepcnt:      {ianaProtocolTCP, -1, 0},
	soCork:         {ianaProtocolTCP, -1, 0},
	soNotsentLOWAT: {ianaProtocolTCP, -1, 0},
	soError:        {sysSOL_SOCKET, -1, 0},
}

var parsers = map[int64]func([]byte) (Option, error){
	ianaProtocolTCP<<32 | sysTCP_NODELAY:        parseNoDelay,
	sysSOL_SOCKET<<32 | sysSO_SNDBUF:            parseSendBuffer,
	sysSOL_SOCKET<<32 | sysSO_RCVBUF:            parseReceiveBuffer,
	sysSOL_SOCKET<<32 | sysSO_KEEPALIVE:         parseKeepAlive,
	ianaProtocolTCP<<32 | sysSIO_KEEPALIVE_VALS: parseKeepAliveValues,
}

// Marshal implements the Marshal method of Option interface.
func (nd NoDelay) Marshal() ([]byte, error) {
	v := boolint32(bool(nd))
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (sb SendBuffer) Marshal() ([]byte, error) {
	v := int32(sb)
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (rb ReceiveBuffer) Marshal() ([]byte, error) {
	v := int32(rb)
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (ka KeepAlive) Marshal() ([]byte, error) {
	v := boolint32(bool(ka))
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (ka KeepAliveIdleInterval) Marshal() ([]byte, error) {
	ka += KeepAliveIdleInterval(options[soKeepidle].uot - time.Nanosecond)
	v := uint32(time.Duration(ka) / options[soKeepidle].uot)
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (ka KeepAliveProbeInterval) Marshal() ([]byte, error) {
	ka += KeepAliveProbeInterval(options[soKeepintvl].uot - time.Nanosecond)
	v := uint32(time.Duration(ka) / options[soKeepintvl].uot)
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (ka KeepAliveProbeCount) Marshal() ([]byte, error) { return nil, errOpNoSupport }

// Marshal implements the Marshal method of Option interface.
func (ck Cork) Marshal() ([]byte, error) { return nil, errOpNoSupport }

// Marshal implements the Marshal method of Option interface.
func (ns NotSentLowWMK) Marshal() ([]byte, error) { return nil, errOpNoSupport }

func parseNoDelay(b []byte) (Option, error) {
	if len(b) < 4 {
		return nil, errBufferTooShort
	}
	return NoDelay(uint32bool(nativeEndian.Uint32(b))), nil
}

func parseSendBuffer(b []byte) (Option, error) {
	if len(b) < 4 {
		return nil, errBufferTooShort
	}
	return SendBuffer(nativeEndian.Uint32(b)), nil
}

func parseReceiveBuffer(b []byte) (Option, error) {
	if len(b) < 4 {
		return nil, errBufferTooShort
	}
	return ReceiveBuffer(nativeEndian.Uint32(b)), nil
}

func parseKeepAlive(b []byte) (Option, error) {
	if len(b) < 4 {
		return nil, errBufferTooShort
	}
	return KeepAlive(uint32bool(nativeEndian.Uint32(b))), nil
}

func parseKeepAliveValues(b []byte) (Option, error) { return nil, errOpNoSupport }
