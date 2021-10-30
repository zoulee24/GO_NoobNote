package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	// "github.com/lxn/win"
	
)

import (
	// "fmt"
	"strings"
)

const Menu1_str 	string =  "设置"
const Test1_str 	string =  "测试"
//可变
var Button1_str 	string =  "打开端口"

type Serial_Windows struct {
	*walk.MainWindow
	//菜单
	action1 		*walk.Action
	//文本框
    textedit1 		*walk.TextEdit
	//可能未来的文本框
	scrollview1		*walk.ScrollView
	//打开关闭端口按钮
    OnOffButton1 		*walk.PushButton
	//打开端口
	Serial_state	bool
}

func GUI_INIT() (MainWindow, *Serial_Windows) {
	sw := new(Serial_Windows)

	serial_windows := (MainWindow{
		AssignTo:		&sw.MainWindow,
		Title:			"串口示波器-v0.0.1",
		MinSize:		Size{400, 300},
		Size:			Size{800, 600},
		Layout: 	HBox{MarginsZero: true},
		//菜单们
		MenuItems: []MenuItem{
			//菜单1
			Menu{
				Text:	Menu1_str,
				//菜单1子项
				Items: []MenuItem{
					Action{
						AssignTo:	&sw.action1,
						Text:	Test1_str,
						//回调函数
						OnTriggered: func() {
							walk.MsgBox(sw, "点击", "菜单项1", walk.MsgBoxOKCancel)
						},
					},
					//分割线
					Separator{},
					Action{
						Text:	Test1_str,
						//回调函数
						OnTriggered: func() {
							walk.MsgBox(sw, "点击", "菜单项2", walk.MsgBoxOKCancel)
						},
					},
				},
			},
		},
		//控件们
		Children: []Widget{
			PushButton{
				AssignTo:      &sw.OnOffButton1,
				Text: Button1_str,
				StretchFactor:8,
				//回调函数
				OnClicked: func() {
					sw.Serial_state = !sw.Serial_state
					if sw.Serial_state {
						sw.OnOffButton1.SetText("关闭端口")
						sw.textedit1.SetText(strings.ToUpper("已经打开端口\r\n"))
						walk.MsgBox(sw, "提示", "已经打开端口", walk.MsgBoxOK)
					} else {
						sw.OnOffButton1.SetText("打开端口")
						sw.textedit1.SetText(strings.ToUpper("已经关闭端口\r\n"))
						walk.MsgBox(sw, "提示", "已经关闭端口", walk.MsgBoxOK)
					}
				},
			},
			TextEdit{
				// StretchFactor: 1,
				AssignTo:      &sw.textedit1,
				ReadOnly:      true,
			},
			
			ScrollView{
				AssignTo:		&sw.scrollview1,
				Name:			"test",
			},

		},


	})

	return serial_windows, sw
}

func (this *Serial_Windows) Show_Msg_Box(title, text string) {
	walk.MsgBox(this, title, text, walk.MsgBoxOKCancel)
}

func (this *Serial_Windows) Show_Text_On_Gui(str string) error {
	if this.textedit1 != nil{
		if err := this.textedit1.SetText(strings.ToUpper(this.textedit1.Text())+str); err != nil {
			return err
		}
		// if err := this.textedit1.SetText(str + "\r\n"); err != nil {
		// 	return err
		// }
	}
	return nil
}