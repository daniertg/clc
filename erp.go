import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("tidak bisa dibagi 0")
	}
	return a / b, nil
}
