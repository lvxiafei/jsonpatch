package jsonpatch_test

import (
	"testing"

	"gomodules.xyz/jsonpatch/v2"
)

func FuzzCreatePatch(f *testing.F) {
	add := func(a, b string) {
		f.Add([]byte(a), []byte(b))
	}
	add(simpleA, simpleB)
	add(superComplexBase, superComplexA)
	add(hyperComplexBase, hyperComplexA)
	add(arraySrc, arrayDst)
	add(empty, simpleA)
	add(point, lineString)
	f.Fuzz(func(t *testing.T, a, b []byte) {
		_, _ = jsonpatch.CreatePatch(a, b)
	})
}
