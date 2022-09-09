package singleton

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetLazySingleton(t *testing.T) {
	assert.Equal(t, GetLazySingleton(), GetLazySingleton())
}

func BenchmarkGetLazySingleton(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetLazySingleton() != GetLazySingleton() {
				b.Error("test fail")
			}
		}
	})
}
