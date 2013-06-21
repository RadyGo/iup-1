/*
	Copyright (C) 2011 by Jeremy Cowgar <jeremy@cowgar.com>

	This file is part of go-iup.

	go-iup is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as
	published by the Free Software Foundation, either version 3 of
	the License, or (at your option) any later version.

	go-iup is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU Lesser General Public
	License along with go-iup.  If not, see <http://www.gnu.org/licenses/>.
*/

package glcanvas

/*
#cgo LDFLAGS: -liup -liupcontrols -liupgl
#cgo windows LDFLAGS: -lgdi32 -lole32 -lcomdlg32 -lcomctl32

#include <stdlib.h>
#include <iup.h>
#include <iupgl.h>
*/
import "C"
import . "github.com/grd/iup"

func GLCanvas(opts ...interface{}) *Ihandle {
	OpenControlLib()

	ih := (*Ihandle)(C.IupGLCanvas(nil))

	for _, o := range opts {
		switch o.(type) {
		default:
			Decorate(ih, o)
		}
	}

	return ih
}

func GLMakeCurrent(ih *Ihandle) {
	C.IupGLMakeCurrent(ih.C())
}

func GLIsCurrent(ih *Ihandle) int {
	return int(C.IupGLIsCurrent(ih.C()))
}

func GLSwapBuffers(ih *Ihandle) {
	C.IupGLSwapBuffers(ih.C())
}

func GLPalette(ih *Ihandle, index int, r, g, b float64) {
	C.IupGLPalette(ih.C(), C.int(index), C.float(r), C.float(g), C.float(b))
}

func GLUseFont(ih *Ihandle, first, count, list_base int) {
	C.IupGLUseFont(ih.C(), C.int(first), C.int(count), C.int(list_base))
}

func GLWait(gl int) {
	C.IupGLWait(C.int(gl))
}
