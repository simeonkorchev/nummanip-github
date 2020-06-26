package calc
import "errors"
func Add(operands ...int) (int, error) {
	if len(operands) < 2 {
		return 0, errors.New("provide more than  1 numbers")
	}
	sum := 0
	for _, operand := range operands {
		sum += operand
	}
	return sum, nil
}
