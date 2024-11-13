package utils

import (
	"math/rand"
	"strconv"
)

func GenerateId(name string) string {
	result := ""
	result = result + name + strconv.Itoa(rand.Intn(1000)+1)
	return result
}
