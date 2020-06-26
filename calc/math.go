package calc

func Add(operands ...int) int {
	sum := 0
	for _, operand := range operands {
		sum += operand
	}
	return sum
}
