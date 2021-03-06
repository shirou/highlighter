package lexer

import (
	"testing"
)

func TestStack(t *testing.T) {
	ti := 10
	ti2 := 20
	t.Run("push", func(t *testing.T) {
		s := NewStack(1)
		s.Push(ti)
		if s.Pop() != ti {
			t.Errorf("wrong pop")
		}
		s.Push(ti)
		if s.Pop() != ti {
			t.Errorf("wrong pop")
		}
		s.Push(ti)
		s.Push(ti2)
		if s.Pop() != ti2 {
			t.Errorf("wrong pop")
		}
		if s.Pop() != ti {
			t.Errorf("wrong pop")
		}
	})
	t.Run("empty", func(t *testing.T) {
		s := NewStack(1)
		if s.Pop() != StackEmpty {
			t.Errorf("not empty")
		}
	})

}

func BenchmarkStackPush(b *testing.B) {
	s := NewStack(b.N + 1)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}
func BenchmarkStackPop(b *testing.B) {
	s := NewStack(b.N + 1)
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}
