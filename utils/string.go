package utils

import "strconv"

func IsNum(n string) bool {
	_, err := strconv.Atoi(n)
	return err == nil
}

func ConvertToInt(n string) int {
	num, _ := strconv.Atoi(n)
	return num
}
