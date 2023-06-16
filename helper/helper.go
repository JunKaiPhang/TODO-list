package helper

import (
	"fmt"
	"strings"
)

func RemovePasswordFromApiLogApiInput(apiInputX string) (apiInput string) {
	// make request body into array, split after every ","
	tInput := strings.Trim(apiInputX, "{}")
	arr := strings.SplitAfter(tInput, ",")

	// func pw return true if n string contains "password"
	pw := func(n string) bool { return strings.Contains(n, "password") }

	i := 0 // output index
	var subPwObj string
	var pwObj string

	// loop request body array
	for _, n := range arr {
		// if slices not contain "password", then keep slices
		if !pw(n) {
			arr[i] = n
			i++
		}

		if pw(n) {
			pwArr := strings.SplitAfter(n, ":")
			if strings.Contains(pwArr[1], ",") {
				pwArr[1] = `"",`
			} else {
				pwArr[1] = `""`
			}
			subPwObj = pwArr[0] + pwArr[1]
			pwObj += subPwObj
		}
	}
	// cut array length down to the length without password
	arr = arr[:i]

	// convert request body array into string, and trim string "[]"
	arr1 := strings.Trim(fmt.Sprint(arr), "[]")
	arr1 += pwObj

	apiInput = fmt.Sprint("{" + arr1 + "}")

	return apiInput
}

/* Use; shorten string to certain length*/ //
/* On Success; return shorted string */ //
/* On Error; no err */ //
func Short(s string, i int) string {
	runes := []rune(s)
	if len(runes) > i {
		return string(runes[:i])
	}
	return s
}
