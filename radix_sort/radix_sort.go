package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	input := []int{16, 26, 36, 46, 16, 26, 36, 46, 56}

	RadixSort(input)

	fmt.Println(input)
}

// RadixSort sorts a slice of ints, provided each int is positive and has max 4 digits
func RadixSort(input []int) {
	// convert to byte slices
	convertedInput := convertToByteSlices(input)

	for digit := 0; digit < 2; digit++ {
		partial := make([][][]byte, 10, 10)
		for p := range partial {
			partial[p] = make([][]byte, 0, 0)
		}

		for _, v := range convertedInput {
			// fmt.Println(v, "-", digit)
			// fmt.Println("v[digit]", (int)(v[digit]), v)
			partial[(int)(v[digit])] = append(partial[(int)(v[digit])], v)
		}

		convertedInput := make([][]byte, 0, 0)
		for _, v := range partial {
			for _, number := range v {
				convertedInput = append(convertedInput, number)
			}
		}

		fmt.Println("convertedInput", convertedInput)
	}

	for i, c := range convertedInput {
		fmt.Println("----", c)

		input[i] = (int)(c[0])
	}
}

func convertToByteSlices(in []int) [][]byte {
	r := make([][]byte, len(in), len(in))

	for i, v := range in {
		buf := new(bytes.Buffer)

		err := binary.Write(buf, binary.LittleEndian, (uint16)(v))
		if err != nil {
			fmt.Println(v, "binary.Write failed:", err)
		}

		r[i] = buf.Bytes()
	}

	return r
}
