package main

import (
	"fmt"
	"instance/model"
	"math"
)

func printArrayAsMatrix(l int, arr []byte) {
	for i := 0; i < l; i++ { //y
		for j := 0; j < l; j++ { //x
			fmt.Printf("% 4d", arr[i*l+j])
		}
		fmt.Println("")
	}
}

func recalc(x, y, l int) int {
	return y*l + x
}

func main() {

	lm := model.CreateEmptyLvlMap(25, 25)
	for i := 0; i < lm.Height*lm.Width; i++ {
		lm.Map[i] = byte(i)
	}
	lm.Map[8*lm.Width+7] = 9
	//sm := lm.GetSquare(11, 11, 5)
	//printArrayAsMatrix(sm.Length, sm.Square)

	Length := 21
	arr := make([]byte, Length*Length)

	for i := 0; i < Length*Length; i++ {
		arr[i] = 0 //byte(i)
	}

	// centerX := 5.
	// centerY := 5.

	// pointX := 0.
	// pointY := 4.

	// A := (centerX - pointX) / (centerY - pointY)
	// B := A*centerY - pointX
	// fmt.Printf("A=%.4f B=%.4f\n", A, B)
	// x := 0.
	// for i := 0; i < 4; i++ {
	// 	x = A*float64(i) - B
	// 	fmt.Printf("x=%.4f y=%.4f\n", x, float64(i))
	// 	arr[int((float64(i)+centerY)*float64(Length)+(x+centerX))] = 5
	// }

	offsetX := 10
	offsetY := 10
	//( x-x1)/(x2-x1) = (y-y1)/(y2-y1)
	//(x/5) = (y/4)
	//x = y * 5/4
	// pointX := 5
	// pointY := 4
	x := 0.
	// for i := -5; i < 6; i++ {
	// 	x = float64(i) * (-5. / 5.)
	// 	fmt.Printf("x=%.4f y=%.4f\n", x, float64(i))
	// 	arr[recalc(int(math.Round(x))+offsetX, i+offsetY, 11)]++
	// }
	for j := -10; j < 11; j++ {
		for i := -10; i < 11; i++ {
			x = float64(i) * (float64(-j) / 10.)
			//fmt.Printf("x=%.4f y=%.4f\n", x, float64(i))
			arr[recalc(int(math.Round(x))+offsetX, i+offsetY, Length)]++
		}
	}
	y:=0.
	for j := -9; j < 10; j++ {
		for i := -10; i < 11; i++ {
			y = float64(i) * (float64(-j) / 10.)
			//fmt.Printf("x=%.4f y=%.4f\n", x, float64(i))
			arr[recalc(i+offsetX, int(math.Round(y))+offsetY, Length)]++
		}
	}
	//	fmt.Printf("A=%.4f B=%.4f\n", A, B)

	//arr[recalc(5, 5, 11)] = 0
	printArrayAsMatrix(Length, arr)

	//

}
