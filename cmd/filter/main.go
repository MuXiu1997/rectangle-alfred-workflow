package main

import (
	"fmt"
	"regexp"
	"strings"

	aw "github.com/deanishe/awgo"
)

var (
	wf   *aw.Workflow
	icon *aw.Icon
)

func run() {
	defer func() {
		wf.SendFeedback()
	}()

	arg := ""
	if len(wf.Args()) != 0 {
		arg = wf.Args()[0]
	}
	expr := strings.Join(strings.Split(arg, ""), ".*")
	expr = ".*" + expr + ".*"
	compileExpr, _ := regexp.Compile(expr)

	options := []string{
		"左半屏",
		"右半屏",
		"中间半屏",
		"上半屏",
		"下半屏",

		"左上",
		"右上",
		"左下",
		"右下",

		"左侧1/3",
		"中间1/3",
		"右侧1/3",
		"左侧2/3",
		"右侧2/3",
		"最大化",
		"接近最大化",
		"最大化高度",
		"缩小",
		"放大",
		"中央",
		"恢复",
		"下一个显示器",
		"上ー个显示器",
	}

	for _, option := range options {
		pinyinKey := toPinyinKey(option)
		if compileExpr.MatchString(option) || compileExpr.MatchString(pinyinKey) {
			item := wf.NewItem(fmt.Sprintf("Rectangle - %s", option))
			item.Arg(option).Icon(icon).Valid(true)
		}

	}
}

func main() {
	wf = aw.New()
	icon = &aw.Icon{
		Value: "./icon.png",
		Type:  aw.IconTypeImage,
	}
	wf.Run(run)
}
