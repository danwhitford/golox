package stack

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDissembleBasicChunk(t *testing.T) {
	var st Stack[int]
	for i := 1; i < 6; i++ {
		st.Push(i)
	}

	want := []int{5, 4, 3, 2, 1}
	var got []int
	for !st.Empty() {
		g, err := st.Pop()
		if err != nil {
			t.Error(err)
		}
		got = append(got, g)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Mismatch (-want +got):\n%s", diff)
	}
}
