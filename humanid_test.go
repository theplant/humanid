package humanid_test

import (
	"testing"

	"fmt"

	"github.com/theplant/humanid"
)

func TestNew(t *testing.T) {
	for i := 0; i < 20; i++ {
		id := humanid.New()
		fmt.Println(id)
		if len(id) != 9 {
			t.Error(id)
		}
	}
}

func TestNewNumber(t *testing.T) {
	for i := 0; i < 20; i++ {
		id := humanid.NewNumber()
		fmt.Println(id)
		if len(id) != 8 {
			t.Error(id)
		}
	}
}

func BenchmarkNew(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		humanid.New()
	}
}
