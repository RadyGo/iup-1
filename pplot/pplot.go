/*
	Copyright (C) 2011 by Jeremy Cowgar <jeremy@cowgar.com>

	This file is part of go-

	go-iup is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as
	published by the Free Software Foundation, either version 3 of
	the License, or (at your option) any later version.

	go-iup is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU Lesser General Public
	License along with go-  If not, see <http://www.gnu.org/licenses/>.
*/

package pplot

/*
#cgo LDFLAGS: -liup -liupcontrols -liupcd -liup_pplot -lcd
#cgo windows LDFLAGS: -lgdi32 -lole32 -lcomdlg32 -lcomctl32

#include <stdlib.h>
#include <iup.h>
#include <iupcontrols.h>
#include <iup_pplot.h>
*/
import "C"

import (
	"unsafe"
	. "github.com/grd/iup"
)

/*
 Duplicate utility functions until C types can be exported in Go
*/

func stringArrayToC(strs []string) []*C.char {
	max := len(strs)
	result := make([]*C.char, max+1)

	for k, v := range strs {
		result[k] = C.CString(v)
	}
	result[max] = nil

	return result
}

func freeCStringArray(strs []*C.char) {
	for _, v := range strs {
		if v != nil {
			C.free(unsafe.Pointer(v))
		}
	}
}

func float64ArrayToC(nums []float64) []C.float {
	result := make([]C.float, len(nums))

	for k, v := range nums {
		result[k] = C.float(v)
	}

	return result
}

/*
 Actual library code
*/

var pplotLibOpened = false

func PPlot(opts ...interface{}) *Ihandle {
	OpenControlLib()

	if pplotLibOpened == false {
		C.IupPPlotOpen()
		pplotLibOpened = true
	}

	ih := (*Ihandle)(C.IupPPlot())

	for _, o := range opts {
		switch o.(type) {
		default:
			Decorate(ih, o)
		}
	}

	return ih
}

func PlotBegin(ih *Ihandle, strXdata int) {
	C.IupPPlotBegin(ih.C(), C.int(strXdata))
}

func PlotAdd(ih *Ihandle, x, y float64) {
	C.IupPPlotAdd(ih.C(), C.float(x), C.float(y))
}

func PlotAddStr(ih *Ihandle, x string, y float64) {
	cX := C.CString(x)
	defer C.free(unsafe.Pointer(cX))

	C.IupPPlotAddStr(ih.C(), cX, C.float(y))
}

func PlotEnd(ih *Ihandle) {
	C.IupPPlotEnd(ih.C())
}

func PlotInsert(ih *Ihandle, index, sample_index int, x, y float64) {
	C.IupPPlotInsert(ih.C(), C.int(index), C.int(sample_index), C.float(x), C.float(y))
}

func PlotInsertStr(ih *Ihandle, index, sample_index int, x string, y float64) {
	cX := C.CString(x)
	defer C.free(unsafe.Pointer(cX))

	C.IupPPlotInsertStr(ih.C(), C.int(index), C.int(sample_index), cX, C.float(y))
}

// Differs from IupPPlotInsertPoints as `count' is determined automatically in this case
func PlotInsertPoints(ih *Ihandle, index, sample_index int, x, y []float64) {
	count := len(x)
	cX := float64ArrayToC(x)
	cY := float64ArrayToC(y)

	C.IupPPlotInsertPoints(ih.C(), C.int(index), C.int(sample_index), &cX[0], &cY[0],
		C.int(count))
}

// Differs from IupPPlotInsertPoints as `count' is determined automatically in this case
func PlotInsertStrPoints(ih *Ihandle, index, sample_index int, x []string, y []float64) {
	count := len(x)

	cX := stringArrayToC(x)
	defer freeCStringArray(cX)

	cY := float64ArrayToC(y)

	C.IupPPlotInsertStrPoints(ih.C(), C.int(index), C.int(sample_index), &cX[0], &cY[0], C.int(count))
}

// Differs from IupPPlotInsertPoints as `count' is determined automatically in this case
func PlotAddPoints(ih *Ihandle, index int, x, y []float64) {
	count := len(x)
	cX := float64ArrayToC(x)
	cY := float64ArrayToC(y)

	C.IupPPlotAddPoints(ih.C(), C.int(index), &cX[0], &cY[0], C.int(count))
}

// Differs from IupPPlotInsertPoints as `count' is determined automatically in this case
func PlotAddStrPoints(ih *Ihandle, index int, x []string, y []float64) {
	count := len(x)

	cX := stringArrayToC(x)
	defer freeCStringArray(cX)

	cY := float64ArrayToC(y)

	C.IupPPlotAddStrPoints(ih.C(), C.int(index), &cX[0], &cY[0], C.int(count))
}

func PlotTransform(ih *Ihandle, x, y float64) (int, int) {
	cIX := new(C.int)
	cIY := new(C.int)

	C.IupPPlotTransform(ih.C(), C.float(x), C.float(y), cIX, cIY)

	return int(*cIX), int(*cIY)
}
