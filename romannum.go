package romannum

import (
	"errors"
	"fmt"
	"strings"
)

type symbol struct {
	value int
	digit string
}

func listOfSymbol() []symbol {
	return []symbol{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
}

func ToRoman(num int) (result string, err error) {
	if num < 1 || num > 3999 {
		err = errors.New("out of range")
		return
	}
	list := listOfSymbol()
	for _, l := range list {
		for num >= l.value {
			result = fmt.Sprintf("%s%s", result, l.digit)
			num -= l.value
		}
	}
	return
}

func ToInt(roman string) (result int, err error) {
	list := listOfSymbol()
	for len(roman) > 0 {
		found := false
		for _, l := range list {
			if strings.HasPrefix(roman, l.digit) {
				result += l.value
				roman = strings.Replace(roman, l.digit, "", 1)
				found = true
			}
		}
		if !found {
			err = errors.New("invalid source")
			return
		}
	}
	return
}
