package help

func Sum(input []int) int {
	sum := 0
	for _, v := range input {
		sum += v
	}
	return sum
}

func Max(input []int) int {
	max := 0
	for _, v := range input {
		if v > max {
			max = v
		}
	}
	return max
}

func Min(input []int) int {
	min := input[0]
	for _, v := range input {
		if v < min {
			min = v
		}
	}
	return min
}

func Contains[T comparable](input []T, item T) bool {
	for _, v := range input {
		if v == item {
			return true
		}
	}
	return false
}
