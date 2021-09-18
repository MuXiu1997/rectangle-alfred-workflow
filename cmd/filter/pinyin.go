package main

import (
	pinyin2 "github.com/mozillazg/go-pinyin"
)

func toPinyinKey(s string) string {
	k := ""
	for _, c := range s {
		pinyin := pinyin2.LazyPinyin(string(c), pinyin2.NewArgs())
		if len(pinyin) == 0 {
			k += string(c)
		} else {
			k += pinyin[0][:1]
		}
	}
	return k
}
