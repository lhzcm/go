// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && s390x
// +build linux
// +build s390x

package runtime

const (
	// vdsoArrayMax is the byte-size of a maximally sized array on this architecture.
	// See cmd/compile/internal/s390x/galign.go arch.MAXWIDTH initialization.
	vdsoArrayMax = 1<<50 - 1
)

var vdsoLinuxVersion = vdsoVersionKey{"LINUX_2.6.29", 0x75fcbb9}

var vdsoSymbolKeys = []vdsoSymbolKey{
	{"__kernel_clock_gettime", 0xb0cd725, 0xdfa941fd, &vdsoClockgettimeSym},
}

// initialize with vsyscall fallbacks
var (
	vdsoClockgettimeSym uintptr = 0
)
