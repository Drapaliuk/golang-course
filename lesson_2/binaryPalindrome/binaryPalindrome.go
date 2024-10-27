package binaryPalindrome // це окей писати назву пакету через camelCase

func decimalToBinary(n int) string {
	if n == 0 {
		return "0"
	}

	binary := ""

	for n > 0 {
		remainder := n % 2
		binary = string(rune(remainder+'0')) + binary
		n = n / 2
	}

	return binary
}

func reverseString(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}

	return
}

func IsBinaryPalindrome(n int) bool {
	binaryStr := decimalToBinary(n)
	revertedBinaryStr := reverseString(binaryStr)

	return binaryStr == revertedBinaryStr

}
