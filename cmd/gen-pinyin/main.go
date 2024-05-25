package main

import (
	"encoding/json"
	pinyin2 "github.com/mozillazg/go-pinyin"
	"os"
	"strings"
)

type Action struct {
	Name   string `json:"name"`
	NameZh string `json:"name-zh"`
}

func main() {
	actionNameJsonFile := os.Args[1]
	outputGoFile := os.Args[2]

	jsonBytes, err := os.ReadFile(actionNameJsonFile)
	if err != nil {
		panic(err)
	}
	var data []Action
	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		panic(err)
	}

	sb := strings.Builder{}
	sb.WriteString("// GENERATED CODE, DO NOT EDIT\n\n")
	sb.WriteString("package main\n\n")
	sb.WriteString("func init() {\n")
	for _, action := range data {
		//	actions = append(actions, &Action{
		//		Name: "leftHalf",
		//		NameZh: "左半屏",
		//		PinYin: "zuo-ban-ping-fang",
		//	})
		pinyinKey := toPinyinKey(action.NameZh)
		sb.WriteString("\tactions = append(actions, &Action{\n")
		sb.WriteString("\t\tName: \"" + action.Name + "\",\n")
		sb.WriteString("\t\tNameZh: \"" + action.NameZh + "\",\n")
		sb.WriteString("\t\tPinYin: \"" + pinyinKey + "\",\n")
		sb.WriteString("\t})\n")
	}
	sb.WriteString("}\n")
	err = os.WriteFile(outputGoFile, []byte(sb.String()), 0644)
	if err != nil {
		panic(err)
	}
}

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
