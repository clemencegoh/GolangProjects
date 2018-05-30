package tests

import (
	"testing"
	"GolangProjects/leetcode/problems/finding"
	"github.com/magiconair/properties/assert"
)

func TestTwoSum(t *testing.T){
	mainTable := []int{
		12,2,5,8,7,9,3,4,2,1,8,13,
	}
	targetNums := []int{
		15,8,4,5,20,12,
	}


	for _, i := range targetNums{
		solution := finding.SolveSumNumber(mainTable,i)
		answer := solution[0] + solution[1]
		assert.Equal(t,answer,i,"Solutions do not match with target number %v",string(i))
	}
}
