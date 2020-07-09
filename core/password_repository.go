package core

//This file generate the password

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

//Functions type for the generate functions
type Functions func() string

//GenerateMin generates one rand char
func GenerateMin() string {
	min := map[int64]rune{0: 'a', 1: 'b', 2: 'c', 3: 'd', 4: 'e', 5: 'f', 6: 'g', 7: 'h', 8: 'i', 9: 'j', 10: 'k', 11: 'l', 12: 'm', 13: 'n', 14: 'ñ',
		15: 'o', 16: 'p', 17: 'q', 18: 'r', 19: 's', 20: 't', 21: 'u', 22: 'v', 23: 'w', 24: 'x', 25: 'y', 26: 'z'}

	num, _ := rand.Int(rand.Reader, big.NewInt(27))
	aux := num.Int64()
	return string(min[aux])
}

//GenerateMayus generates one rand mayus char
func GenerateMayus() string {
	max := map[int64]rune{0: 'A', 1: 'B', 2: 'C', 3: 'D', 4: 'E', 5: 'F', 6: 'G', 7: 'H', 8: 'I', 9: 'J', 10: 'K', 11: 'L', 12: 'M', 13: 'N', 14: 'Ñ',
		15: 'O', 16: 'P', 17: 'Q', 18: 'R', 19: 'S', 20: 'T', 21: 'U', 22: 'V', 23: 'W', 24: 'X', 25: 'Y', 26: 'Z'}

	num, _ := rand.Int(rand.Reader, big.NewInt(27))
	aux := num.Int64()
	return string(max[aux])
}

//GenerateSpecial generates one special character from the list
func GenerateSpecial() string {
	special := map[int64]rune{0: '~', 1: '!', 2: '@', 3: '#', 4: '#', 5: '$', 6: '%', 7: '^', 8: '&', 9: '*', 10: '(', 11: ')', 12: '-', 13: '+', 14: '=',
		15: '{', 16: '}', 17: '[', 18: ']', 19: '\\', 20: '|', 21: '<', 22: '>', 23: '?', 24: '¿', 25: '¡', 26: ':'}

	num, _ := rand.Int(rand.Reader, big.NewInt(27))
	aux := num.Int64()
	return string(special[aux])
}

//GenerateNumber the same method to generate one rand number as a char
func GenerateNumber() string {
	num, _ := rand.Int(rand.Reader, big.NewInt(10))
	aux := num.Int64()
	return strconv.Itoa(int(aux))
}

//ProcessPassword makes the password base on the functions in the arguments
func ProcessPassword(count int, min Functions, mayus Functions, special Functions, numbers Functions) string {
	pass := ""
	var array []Functions

	if min != nil {
		array = append(array, min)
	}

	if mayus != nil {
		array = append(array, mayus)
	}

	if special != nil {
		array = append(array, special)
	}

	if numbers != nil {
		array = append(array, numbers)
	}

	for i := 0; i < count; i++ {
		nBig, _ := rand.Int(rand.Reader, big.NewInt(int64(len(array)))) //New big int between 0 and the size of the array
		n := nBig.Int64()
		pass += array[n]() //We invoke one of the functions that we have on the array
	}

	// We return one string with the password
	return pass
}

//GeneratePassRand determines how the password must be done
func GeneratePassRand(long int, mayus bool, special bool, numbers bool) string {
	switch long >= 8 {
	case mayus && special && numbers:
		return ProcessPassword(long, GenerateMin, GenerateMayus, GenerateSpecial, GenerateNumber)

	case mayus && special && !numbers:
		return ProcessPassword(long, GenerateMin, GenerateMayus, GenerateSpecial, nil)

	case mayus && !special && numbers:
		return ProcessPassword(long, GenerateMin, GenerateMayus, nil, GenerateNumber)

	case !mayus && special && numbers:
		return ProcessPassword(long, GenerateMin, nil, GenerateSpecial, GenerateNumber)

	case mayus && !special && !numbers:
		return ProcessPassword(long, GenerateMin, GenerateMayus, nil, nil)

	case !mayus && special && !numbers:
		return ProcessPassword(long, GenerateMin, nil, GenerateSpecial, nil)

	case !mayus && !special && numbers:
		return ProcessPassword(long, GenerateMin, nil, nil, GenerateNumber)

	case !mayus && !special && !numbers:
		return ProcessPassword(long, GenerateMin, nil, nil, nil)

	default:
		return ""
	}
}
