// Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
//
// This file originated from: http://opensource.conformal.com/
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

// Go bindings for Pango.
package pango

// #cgo pkg-config: fontconfig gobject-2.0 glib-2.0 pango pangocairo
// #include <pango/pango.h>
// #include "pango.go.h"
import "C"
import "errors"

//	"github.com/andre-hub/gotk3/glib"
//	"unsafe"

func init() {

}

// Finalizer is a function that when called will finalize an object
type Finalizer func()

// FinalizerStrategy will be called by every runtime finalizer in gotk3
// The simple version will just call the finalizer given as an argument
// but in larger programs this might cause problems with the UI thread.
// The FinalizerStrategy function will always be called in the goroutine that
// `runtime.SetFinalizer` uses. It is a `var` to explicitly allow clients to
// change the strategy to something more advanced.
var FinalizerStrategy = func(f Finalizer) {
	f()
}

/*
 * Type conversions
 */

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}
func gobool(b C.gboolean) bool {
	if b != 0 {
		return true
	}
	return false
}

/*
 * Unexported vars
 */

var nilPtrErr = errors.New("cgo returned unexpected nil pointer")

/*
 * Constantes
 */

const (
	SCALE int = 1024
)
