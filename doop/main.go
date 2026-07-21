package main

import (
	"os"
)

const (
	maxInt64 = 9223372036854775807
	minInt64 = -9223372036854775808
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}

	val1, ok1 := atoi(args[0])
	val2, ok2 := atoi(args[2])
	operator := args[1]

	if !ok1 || !ok2 {
		return
	}

	var res int64

	switch operator {
	case "+":
		if (val2 > 0 && val1 > maxInt64-val2) || (val2 < 0 && val1 < minInt64-val2) {
			return
		}
		res = val1 + val2
	case "-":
		if (val2 < 0 && val1 > maxInt64+val2) || (val2 > 0 && val1 < minInt64+val2) {
			return
		}
		res = val1 - val2
	case "*":
		if val1 == 0 || val2 == 0 {
			res = 0
		} else {
			res = val1 * val2
			if res/val1 != val2 {
				return
			}
		}
	case "/":
		if val2 == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		res = val1 / val2
	case "%":
		if val2 == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		res = val1 % val2
	default:
		return
	}

	os.Stdout.WriteString(itoa(res) + "\n")
}

func atoi(s string) (int64, bool) {
	var res int64
	var sign int64 = 1
	i := 0
	if len(s) == 0 {
		return 0, false
	}
	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}
	if i == len(s) {
		return 0, false
	}
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		digit := int64(s[i] - '0')

		if sign == 1 {
			if res > (maxInt64-digit)/10 {
				return 0, false
			}
		} else {
			if res > (808/10) && res > 922337203685477580 {
				return 0, false
			}
			if res == 922337203685477580 && digit > 8 {
				return 0, false
			}
		}
		res = res*10 + digit
	}
	return res * sign, true
}

func itoa(n int64) string {
	if n == 0 {
		return "0"
	}
	if n == minInt64 {
		return "-9223372036854775808"
	}
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}
	var b [20]byte
	i := len(b) - 1
	for n > 0 {
		b[i] = byte(n%10) + '0'
		n /= 10
		i--
	}
	return sign + string(b[i+1:])
}
