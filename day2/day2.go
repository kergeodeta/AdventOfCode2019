package day2

import (
	"fmt"
	"log"
)

var puzzleInput = []int{
	1, 12, 2, 3,
	1, 1, 2, 3,
	1, 3, 4, 3,
	1, 5, 0, 3,
	2, 10, 1, 19,
	1, 5, 19, 23,
	1, 23, 5, 27,
	2, 27, 10, 31,
	1, 5, 31, 35,
	2, 35, 6, 39,
	1, 6, 39, 43,
	2, 13, 43, 47,
	2, 9, 47, 51,
	1, 6, 51, 55,
	1, 55, 9, 59,
	2, 6, 59, 63,
	1, 5, 63, 67,
	2, 67, 13, 71,
	1, 9, 71, 75,
	1, 75, 9, 79,
	2, 79, 10, 83,
	1, 6, 83, 87,
	1, 5, 87, 91,
	1, 6, 91, 95,
	1, 95, 13, 99,
	1, 10, 99, 103,
	2, 6, 103, 107,
	1, 107, 5, 111,
	1, 111, 13, 115,
	1, 115, 13, 119,
	1, 13, 119, 123,
	2, 123, 13, 127,
	1, 127, 6, 131,
	1, 131, 9, 135,
	1, 5, 135, 139,
	2, 139, 6, 143,
	2, 6, 143, 147,
	1, 5, 147, 151,
	1, 151, 2, 155,
	1, 9, 155, 0,
	99,
	2, 14, 0, 0,
}

func calculate(in []int) (out []int) {
	i := 0
	out = append(out, in...)
	for {
		code := out[i+0]
		if code == 99 {
			break
		}

		arg1 := out[i+1] // Position of the first parameter
		arg2 := out[i+2] // Position of the second parameter
		rPos := out[i+3] // Position of the result

		var result int
		switch code {
		case 1:
			result = out[arg1] + out[arg2]
		case 2:
			result = out[arg1] * out[arg2]
		default:
			log.Fatalf("Incorrect operation code! Program exiting.")
		}

		out[rPos] = result
		i += 4
	}

	return
}

func Run() {
	fmt.Println(calculate(puzzleInput))
}
