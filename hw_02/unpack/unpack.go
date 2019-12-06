package unpack

import (
	"strconv"
	"strings"
)

// Unpack unpacks string to string
func Unpack(s string) string {

	ans := ""
	length := len(s)
	repeatFlag := 0

	for counter, value := range s {

		// проверим текущий символ
		_, err := strconv.ParseInt(string(value), 10, 32)

		switch err == nil {
		// если текущий символ инт - увеличим флаг повторений и выйдем если он больше 1
		case true:
			repeatFlag++
			if repeatFlag > 1 {
				return ""
			}

		// если ткущий символ - буква
		case false:
			// обнулим флаг повторений
			repeatFlag = 0

			// если это последний символ - прибавим его и выйдем
			if counter >= length-1 {
				ans += string(value)
				return ans
			}

			// посмотрим следующий символ
			nextNumber, err2 := strconv.ParseInt(string(s[counter+1]), 10, 32)

			// если это строка, то добавим текущий и идем дальше
			if err2 != nil {
				ans += string(value)
				// если число - то добавим символ, умноженный на это число
			} else {
				ans += strings.Repeat(string(value), int(nextNumber))
			}

		}

	}

	return ans
}
