package main

import (
	"fmt"
	"math"
)

/*
 * Complete the 'decryptPassword' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 * decryptPassword("51Pa*0Lp*0e")
 */

func minimumStep(staircaseMap []int32) int32 {
	// Write your code here
	var stepMin, stepIndicator, countMax, defaultCount, lastNum int32
	// step = 0
	defaultCount = 3

	for i := 0; i < len(staircaseMap); i++ {
		if i == 0 {
			countMax++
			continue
		}

		if staircaseMap[i] == 0 {
			if lastNum != 1 && countMax > 0 {
				stepIndicator++
			}

			if i >= 2 && staircaseMap[i-2] == 0 {
				stepIndicator++
			}

			lastNum = 0
		} else {
			if lastNum == 1 {
				return 0
			}

			if countMax > 1 {
				defaultCount++
			}
			lastNum = 1
		}

		countMax++

		if (len(staircaseMap) - 1) == i {
			stepMin += stepIndicator
			break
		}

		if countMax >= defaultCount {
			if stepIndicator == 1 || stepIndicator == 2 {
				stepMin++
				stepIndicator = 0
				countMax = 0
				defaultCount = 3
				lastNum = 0
				i -= 1
			}
		}

		//// Start leader
		// if i + 1 <= len(staircaseMap) - 1 && staircaseMap[i + 1] == 1 {
		//     if i +
		//     return 0
		// }

		// if (i + 2) <= (len(staircaseMap) - 1) {
		//     if staircaseMap[i] == 0 && staircaseMap[i+2] == 0 {
		//         step++
		//         i++
		//         continue
		//     } else if {

		//     }
		// }
	}

	return stepMin
}

func main() {
	//fmt.Println(decryptPassword("51Pa*0Lp*0e"))
	//fmt.Println(minimumStep([]int32{0, 1, 0}))

	ratio := math.Pow(10, float64(2))
	fmt.Println("Ratio", ratio)

	round := math.Round(0.5 * ratio)
	fmt.Println(round / ratio)
}
