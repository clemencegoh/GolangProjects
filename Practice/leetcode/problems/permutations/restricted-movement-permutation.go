package permutations

// given a M x N matrix, starting at the bottom left, restricted to moving right and up
// find all possible ways to move to top right of matrix

// SolveMatrix solves for number of possible paths that can be taken
// take note that m and n denote steps
func SolveMatrix(m int, n int) int {
	// given m steps up and n steps right, find total number of permutations
	// use combination formula: C(m,n) => m!/m!(m-n)!
	return FactorialIterative(m)/(FactorialIterative(n)*FactorialIterative(m-n))
}

// Factorial returns factorial of a number
func FactorialRecursive (number int) int {
	// recursive approach

	if number==1{
		return number
	}

	return number * FactorialRecursive(number -1)
}

func FactorialIterative (number int) int {
	if number <= 1{
		return number
	}

	ans := 1
	for i:=1;i<=number;i++{
		ans *= i
	}

	return ans
}
