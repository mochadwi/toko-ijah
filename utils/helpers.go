package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// StrToInt64 func
func StrToInt64(word string) int64 {

	num, _ := strconv.ParseInt(word, 0, 64)
	return num
}

// StrToInt func
func StrToInt(word string) int {

	num, _ := strconv.Atoi(word)
	return num
}

// StrToUint func
func StrToUint(word string) uint {

	num, _ := strconv.Atoi(word)
	return uint(num)
}

// StrToUint64 func
func StrToUint64(word string) uint64 {

	num, _ := strconv.ParseUint(word, 0, 64)
	return num
}

// ShortenSizeStr func
func ShortenSizeStr(oldSize string) string {
	var newSize = ""

	switch oldSize {
	case "XXXL":
		newSize = "X3"
	case "XXL":
		newSize = "XX"
	case "XL":
		newSize = "XL"
	default:
		newSize = oldSize + oldSize
	}

	return newSize
}

// ShortenColourStr func
func ShortenColourStr(oldColour string) string {
	newColour := strings.ToUpper(oldColour[0:3])

	if strings.Contains(oldColour, " ") {
		slice := strings.Split(oldColour, " ")
		newColour = strings.ToUpper(slice[0][0:1] + slice[1][0:2])
	}

	return newColour
}

// RandomSKU func
func RandomSKU() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return strconv.Itoa(r1.Intn(99999999))
}

// RandomSKUByLength func
func RandomSKUByLength(digits int) string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	n := 9

	for i := 1; i <= digits; i++ {

		n *= 10
		fmt.Printf("digits: %d", n)
	}

	// fmt.Printf("digits: %d", n)

	return strconv.Itoa(r1.Intn(n))
}

// GenerateReceipt func 20171101-77541
func GenerateReceipt(date string) string {

	return date + "-" + RandomSKUByLength(5)
}

// GenerateSKU func
func GenerateSKU(size string, colour string) string {

	return "SSI-D" + RandomSKU() + "-" + ShortenSizeStr(size) + "-" + ShortenColourStr(colour)
}

// PrettifyName func
func PrettifyName(name string, size string, colour string) string {

	return name + " (" + size + ", " + colour + ")"
}
