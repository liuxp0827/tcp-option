// Copyright 2016 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin dragonfly freebsd linux netbsd openbsd solaris

package tcpopt

import (
	"time"
	"unsafe"
)

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
	v := int32(time.Duration(ka) / options[soKeepidle].uot)
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (ka KeepAliveProbeInterval) Marshal() ([]byte, error) {
	ka += KeepAliveProbeInterval(options[soKeepintvl].uot - time.Nanosecond)
	v := int32(time.Duration(ka) / options[soKeepintvl].uot)
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (ka KeepAliveProbeCount) Marshal() ([]byte, error) {
	v := int32(ka)
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (ck Cork) Marshal() ([]byte, error) {
	v := boolint32(bool(ck))
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (ns NotSentLowWMK) Marshal() ([]byte, error) {
	v := int32(ns)
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

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

func parseKeepAliveIdleInterval(b []byte) (Option, error) {
	if len(b) < 4 {
		return nil, errBufferTooShort
	}
	v := time.Duration(nativeEndian.Uint32(b)) * options[soKeepidle].uot
	return KeepAliveIdleInterval(v), nil
}

func parseKeepAliveProbeInterval(b []byte) (Option, error) {
	if len(b) < 4 {
		return nil, errBufferTooShort
	}
	v := time.Duration(nativeEndian.Uint32(b)) * options[soKeepintvl].uot
	return KeepAliveProbeInterval(v), nil
}

func parseKeepAliveProbeCount(b []byte) (Option, error) {
	if len(b) < 4 {
		return nil, errBufferTooShort
	}
	return KeepAliveProbeCount(nativeEndian.Uint32(b)), nil
}

func parseCork(b []byte) (Option, error) {
	if len(b) < 4 {
		return nil, errBufferTooShort
	}
	return Cork(uint32bool(nativeEndian.Uint32(b))), nil
}

func parseNotSentLowWMK(b []byte) (Option, error) {
	if len(b) < 4 {
		return nil, errBufferTooShort
	}
	return NotSentLowWMK(nativeEndian.Uint32(b)), nil
}
