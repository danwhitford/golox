package runlengthencoder

import "fmt"


type RunLengthEncoder struct{
	arr []int
}

func (rle *RunLengthEncoder) Append(i int) {
	if rle.arr == nil {
		rle.arr = []int{i, 1}
		return
	}

	if rle.arr[len(rle.arr) - 2] == i {
		rle.arr[len(rle.arr) - 1] += 1
		return
	}

	rle.arr = append(rle.arr, i)
	rle.arr = append(rle.arr, 1)
}

func (rle *RunLengthEncoder) Get(idx int) int {
	winmin, winmax := 0, 0
	for i := 0; i < len(rle.arr); i += 2 {
		winmax += rle.arr[i+1]

		fmt.Println(winmin, winmax, idx)
		if winmin <= idx && idx < winmax {
			return rle.arr[i]
		}
		winmin = winmax
	}
	panic("not found or index out-of-range")
}

func (rle *RunLengthEncoder) Size() int {
	return len(rle.arr)
}
