package validParentheses

import "math"

const (
	roundOpenBrackerCodePoint   rune = 40
	roundCloseBrackerCodePoint  rune = 41
	curlyOpenBracketCodePoint   rune = 123
	curlyCloseBracketCodePoint  rune = 125
	squareOpenBracketCodePoint  rune = 91
	squareCloseBracketCodePoint rune = 93
)

func isBracket(n rune) bool {
	return n == roundOpenBrackerCodePoint || n == roundCloseBrackerCodePoint || n == curlyOpenBracketCodePoint || n == curlyCloseBracketCodePoint || n == squareOpenBracketCodePoint || n == squareCloseBracketCodePoint
}

func isClosedBracketsPair(openBracket, closeBracket rune) bool {
	switch openBracket {
	case roundOpenBrackerCodePoint:
		return closeBracket == roundCloseBrackerCodePoint
	case curlyOpenBracketCodePoint:
		return closeBracket == curlyCloseBracketCodePoint
	case squareOpenBracketCodePoint:
		return closeBracket == squareCloseBracketCodePoint
	default:
		return false
	}
}

func ValidParentheses(s string) bool {
	strInRunes := []rune(s)
	middleIdx := int(math.Round(float64(len(strInRunes) / 2)))

	isBracketInTheMiddle := len(strInRunes)%2 != 0 && isBracket(strInRunes[middleIdx])
	if isBracketInTheMiddle {
		return false
	}

	for i := 0; i < middleIdx; i++ {
		curChar := strInRunes[i]
		oppositeChar := strInRunes[len(strInRunes)-1-i]

		if !isClosedBracketsPair(curChar, oppositeChar) {
			return false
		}
	}

	return true
}
