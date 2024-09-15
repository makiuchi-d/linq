//go:build go1.23

package linq_test

import (
	"errors"
	"testing"

	"github.com/makiuchi-d/linq/v2"
)

var src = linq.Range(0, 100000)

func BenchmarkFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e := src()
		for {
			_, err := e.Next()
			if err != nil {
				if errors.Is(err, linq.EOC) {
					break
				}
			}
		}
	}
}

func BenchmarkForEach(b *testing.B) {
	for i := 0; i < b.N; i++ {
		linq.ForEach(src, func(n int) error {
			return nil
		})
	}
}

func BenchmarkRangeFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, _ = range src.All() {
		}
	}
}
