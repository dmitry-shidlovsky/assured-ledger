// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package throw

import (
	"fmt"
	"runtime"
	"testing"
)

func TestCaptureStack(t *testing.T) {
	// TODO proper tests
	//fmt.Printf("%s\n==============\n", captureStackByDebug(0, false))
	//fmt.Printf("%s\n==============\n", captureStackByDebug(1, false))
	//fmt.Printf("%s\n==============\n", captureStackByDebug(0, true))
	//fmt.Printf("%s\n==============\n", captureStackByDebug(99, true))
	//
	//fmt.Println()
	//
	//fmt.Printf("%s\n==============\n", captureStackByCallers(0, false))
	//fmt.Printf("%s\n==============\n", captureStackByCallers(1, false))
	//fmt.Printf("%s\n==============\n", captureStackByCallers(0, true))

	fmt.Printf("%s\n==============\n", captureStack(0, false))
	fmt.Printf("%s\n==============\n", captureStack(0, true))
	fmt.Printf("%s\n==============\n", captureStack(99, true))
}

func BenchmarkCaptureStack(b *testing.B) {
	b.Run("captureStackByDebug-full", func(b *testing.B) {
		for i := b.N; i > 0; i-- {
			v := captureStackByDebug(0, false)
			runtime.KeepAlive(v)
		}
	})

	b.Run("captureStackByCallers-full", func(b *testing.B) {
		for i := b.N; i > 0; i-- {
			v := captureStackByCallers(0, false)
			runtime.KeepAlive(v)
		}
	})

	b.Run("captureStackByDebug-top", func(b *testing.B) {
		for i := b.N; i > 0; i-- {
			v := captureStackByDebug(0, true)
			runtime.KeepAlive(v)
		}
	})

	b.Run("captureStackByCallers-top", func(b *testing.B) {
		for i := b.N; i > 0; i-- {
			v := captureStackByCallers(0, true)
			runtime.KeepAlive(v)
		}
	})
}