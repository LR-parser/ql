package util

import (
	"ql/token"
	"strconv"
)

func StringLiteralTokensToString(token *token.Token) (str string) {
	if token == nil {
		panic("Nil token passed to StringLiteralTokensToString")
	}

	astr, err := strconv.Unquote(string(token.Lit))
	if err != nil {
		return ""
	}

	return astr
}
