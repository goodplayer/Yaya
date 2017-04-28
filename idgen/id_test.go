package idgen_test

import (
	"testing"

	"github.com/goodplayer/yaya/idgen"
)

func BenchmarkIdgen_Next(b *testing.B) {
	id := idgen.NewIdGen(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		idgen.ConvertToString(id.Next())
	}
}
