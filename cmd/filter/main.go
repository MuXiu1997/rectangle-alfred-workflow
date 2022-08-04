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

	options := map[string]string{
		"左半屏":  "leftHalf",
		"右半屏":  "rightHalf",
		"中间半屏": "centerHalf",
		"上半屏":  "topHalf",
		"下半屏":  "bottomHalf",

		"左上": "topLeft",
		"右上": "topRight",
		"左下": "bottomLeft",
		"右下": "bottomRight",

		"最大化":   "maximize",
		"接近最大化": "almostMaximize",
		"最大化高度": "maximizeHeight",
		"缩小":    "makeSmaller",
		"放大":    "makeLarger",
		"中央":    "center",
		"恢复":    "restore",

		"下一个显示器": "nextDisplay",
		"上ー个显示器": "previousDisplay",

		"前1/3":  "firstThird",
		"中间1/3": "centerThird",
		"后1/3":  "lastThird",
		"前2/3":  "firstTwoThirds",
		"后2/3":  "lastTwoThirds",

		"左首1/4": "firstFourth",
		"左二1/4": "secondFourth",
		"右二1/4": "thirdFourth",
		"右首1/4": "lastFourth",
		"左侧3/4": "firstThreeFourths",
		"右侧3/4": "lastThreeFourths",

		"向左移动": "moveLeft",
		"向右移动": "moveRight",
		"向上移动": "moveUp",
		"向下移动": "moveDown",

		"左上1/6": "topLeftSixth",
		"中上1/6": "topCenterSixth",
		"右上1/6": "topRightSixth",
		"左下1/6": "bottomLeftSixth",
		"中下1/6": "bottomCenterSixth",
		"右下1/6": "bottomRightSixth",
	}

	for optionZH, optionEN := range options {
		pinyinKey := toPinyinKey(optionZH)
		if compileExpr.MatchString(optionZH) || compileExpr.MatchString(pinyinKey) {
			item := wf.NewItem(fmt.Sprintf("Rectangle - %s", optionZH))
			item.Arg(optionEN).Icon(icon).Valid(true)
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
