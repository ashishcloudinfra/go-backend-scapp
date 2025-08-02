package helpers

import "fmt"

func FloatToString(input_num float64) string {
	str := fmt.Sprintf("%.2f", input_num)
	return str
}
