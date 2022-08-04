package main

import (
	"strings"

	pinyin2 "github.com/mozillazg/go-pinyin"
)

func toPinyinKey(str string) string {
	strBuilder := strings.Builder{}
	for _, r := range str {
		char := string(r)
		pinyin := pinyin2.LazyPinyin(char, pinyin2.NewArgs())
		// Non-Chinese char
		if len(pinyin) == 0 {
			strBuilder.WriteString(char)
		} else {
			strBuilder.WriteString(pinyin[0][:1])
		}
	}
	return strBuilder.String()
}
