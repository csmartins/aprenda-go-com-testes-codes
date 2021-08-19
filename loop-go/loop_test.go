package iteractions

import (
	"math/rand"
	"testing"
	"strings"
)

func TestLoop(t *testing.T){

	verifyCorrectness := func(t *testing.T, result, expected string){
		t.Helper()
		if result != expected {
			t.Errorf("got '%s', but expected '%s'", result, expected)
		}
	}
	t.Run("repeat 5 times", func(t *testing.T){
		result := Loop("a", 5)
		expected := "aaaaa"
	
		verifyCorrectness(t, result, expected)
	})
	t.Run("no character", func (t *testing.T) {
		result := Loop("a", 0)
		expected := ""

		verifyCorrectness(t, result, expected)
	})
	t.Run("repeat rand times", func (t *testing.T) {
		x := rand.Intn(100)
		result := Loop("a", x)
		expected := strings.Repeat("a", x)

		verifyCorrectness(t, result, expected)
	})
}

func BenchmarkLoop(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Loop("a", i)
    }
}
