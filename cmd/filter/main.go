package main

import (
	"fmt"
	"regexp"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/elliotchance/orderedmap/v2"
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

	options := orderedmap.NewOrderedMap[string, string]()

	options.Set("左半屏", "leftHalf")
	options.Set("右半屏", "rightHalf")
	options.Set("中间半屏", "centerHalf")
	options.Set("上半屏", "topHalf")
	options.Set("下半屏", "bottomHalf")

	options.Set("左上", "topLeft")
	options.Set("右上", "topRight")
	options.Set("左下", "bottomLeft")
	options.Set("右下", "bottomRight")

	options.Set("最大化", "maximize")
	options.Set("接近最大化", "almostMaximize")
	options.Set("最大化高度", "maximizeHeight")
	options.Set("缩小", "smaller")
	options.Set("放大", "larger")
	options.Set("中央", "center")
	options.Set("恢复", "restore")

	options.Set("下一个显示器", "nextDisplay")
	options.Set("上ー个显示器", "previousDisplay")

	options.Set("前1/3", "firstThird")
	options.Set("中间1/3", "centerThird")
	options.Set("后1/3", "lastThird")
	options.Set("前2/3", "firstTwoThirds")
	options.Set("后2/3", "lastTwoThirds")

	options.Set("左首1/4", "firstFourth")
	options.Set("左二1/4", "secondFourth")
	options.Set("右二1/4", "thirdFourth")
	options.Set("右首1/4", "lastFourth")
	options.Set("左侧3/4", "firstThreeFourths")
	options.Set("右侧3/4", "lastThreeFourths")

	options.Set("向左移动", "moveLeft")
	options.Set("向右移动", "moveRight")
	options.Set("向上移动", "moveUp")
	options.Set("向下移动", "moveDown")

	options.Set("左上1/6", "topLeftSixth")
	options.Set("中上1/6", "topCenterSixth")
	options.Set("右上1/6", "topRightSixth")
	options.Set("左下1/6", "bottomLeftSixth")
	options.Set("中下1/6", "bottomCenterSixth")
	options.Set("右下1/6", "bottomRightSixth")

	for option := options.Front(); option != nil; option = option.Next() {
		optionZH := option.Key
		optionEN := option.Value
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
