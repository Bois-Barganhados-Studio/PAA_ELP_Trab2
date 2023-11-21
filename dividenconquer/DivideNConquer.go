package dividenconquer

var Qtd_comp int

func MaximumSubarray(nums []int) []int {
	Qtd_comp = 0
	return MaximumSubarrayRecur(nums)
}

func MaximumSubarrayRecur(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	mid := len(nums) / 2
	left := MaximumSubarray(nums[:mid])
	right := MaximumSubarray(nums[mid:])
	cross := NewMaximumSubarray(left, right)
	return cross
}

func NewMaximumSubarray(left, right []int) []int {
	Qtd_comp++
	if len(left) == 0 {

		return right
	}
	Qtd_comp++
	if len(right) == 0 {

		return left
	}
	Qtd_comp++
	if left[len(left)-1] < right[0] {
		return append(left, right...)
	}

	newArray := []int{}
	Qtd_comp++
	if len(left) < len(right) {
		// Percorre o array da esquerda até enquanto o valor seja menor que o da direta[0]
		for i := 0; i < len(left); i++ {
			Qtd_comp++
			if left[i] < right[0] {
				newArray = append(newArray, left[i])
			}
		}
		newArray = append(newArray, right...)
	} else {
		// Percorre o array da direita até enquanto o valor seja menor que o da esquerda[len-1]
		newArray = append(newArray, left...)
		for i := 0; i < len(right); i++ {
			Qtd_comp++
			if right[i] > left[len(left)-1] {
				newArray = append(newArray, right[i])
			}
		}
	}
	return newArray
}
