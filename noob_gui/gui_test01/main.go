package main
 
import (
	"strings"
 
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)
 
func main() {
	var inTE, outTE *walk.TextEdit
 
	MainWindow{
		Title:   "xiaochuan测试",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE, MaxLength: 10},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()
}