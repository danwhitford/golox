package runlengthencoder

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRunLengthEncoder(t *testing.T) {
	table := []struct{
		inputs []int
	}{
		{ []int{1} },
		{ []int{1, 2} },
		{ []int{15, 25, 35, 45, 55} },
		{ []int{7, 7, 7, 7, 7, 7, 7} },
		{ []int{3, 3, 77, 77, 77, 77, 8, 14, 14} },
	}

	for tsti, tst := range table {
		t.Run(fmt.Sprintf("RLE Test %d", tsti), func(t *testing.T) {
			rle := RunLengthEncoder{}
			for _, i := range tst.inputs {
				rle.Append(i)
			}
	
			var got []int
			for i := range tst.inputs {
				j := rle.Get(i)
				got = append(got, j)
			}
	
			if diff := cmp.Diff(tst.inputs, got); diff != "" {
				t.Errorf("Mismatch test %d (-want +got):\n%s", tsti, diff)
			}
		})
	}
}

func TestRunLengthEncoderSize(t *testing.T) {
	table := []struct{
		inputs []int
		want int
	}{
		{ []int{1}, 2 },
		{ []int{1, 2}, 4 },
		{ []int{15, 25, 35, 45, 55}, 10 },
		{ []int{7, 7, 7, 7, 7, 7, 7}, 2 },
		{ []int{3, 3, 77, 77, 77, 77, 8, 14, 14} , 8 },
	}

	for tsti, tst := range table {
		t.Run(fmt.Sprintf("RLE Test %d", tsti), func(t *testing.T) {
			rle := RunLengthEncoder{}
			for _, i := range tst.inputs {
				rle.Append(i)
			}
	
			got := rle.Size()
	
			if diff := cmp.Diff(tst.want, got); diff != "" {
				t.Errorf("Mismatch test %d (-want +got):\n%s", tsti, diff)
			}
		})
	}
}
