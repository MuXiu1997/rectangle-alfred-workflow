package main

import (
	"fmt"
	"regexp"
	"strings"

	aw "github.com/deanishe/awgo"
)

var (
	wf      *aw.Workflow
	icon    *aw.Icon
	actions []*Action
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

	for _, action := range actions {
		if compileExpr.MatchString(action.PinYin) || compileExpr.MatchString(action.NameZh) {
			item := wf.NewItem(fmt.Sprintf("Rectangle - %s", action.NameZh)).
				Subtitle(action.Name)
			item.Arg(action.Name).Icon(icon).Valid(true)
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
