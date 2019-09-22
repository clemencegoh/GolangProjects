package finding

import "fmt"

// given an array and a number, return indices of the 2 positions which sum up to number
// eg. [1,9,2,4,3], num == 10:
// return [1,2]


// custom type tuple
type customPair struct {
	number int
	position int
}

func SolveSumNumber( array []int, target int) (answer []int){

	overHead := make(map[int]customPair)
	for pos, i := range array{
		// find if key exists in map
		if _, ok := overHead[target - i]; ok {
			// found answer
			fmt.Println("Map at this time:",overHead)
			answer = []int{overHead[target - i].position,pos}
			return answer
		}else{
			// put into map
			overHead[i] = customPair{
				target - i,
				pos,
			}
		}
	}

	// if not found
	fmt.Println("No sum found")
	return nil
}
