package main
 
import (
	"fmt"
	"io"
	"os"
	"strings"
 
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)
 
type MyMainWindow struct {
	*walk.MainWindow
	edit *walk.TextEdit
}
 
func main() {
	mw := &MyMainWindow{}
	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		MinSize:  Size{400, 300},
		Size:     Size{600, 400},
		MenuItems: []MenuItem{
			Menu{
				Text: "文件",
				Items: []MenuItem{
					Action{
						Text: "打开文件",
						Shortcut: Shortcut{ //定义快捷键后会有响应提示显示
							Modifiers: walk.ModControl,
							Key:       walk.KeyO,
						},
						OnTriggered: mw.openFileActionTriggered, //点击动作触发响应函数
					},
					Action{
						Text: "另存为",
						Shortcut: Shortcut{
							Modifiers: walk.ModControl | walk.ModShift,
							Key:       walk.KeyS,
						},
						OnTriggered: mw.saveFileActionTriggered,
					},
					Action{
						Text: "退出",
						OnTriggered: func() {
							mw.Close()
						},
					},
				},
			},
			Menu{
				Text: "帮助",
				Items: []MenuItem{
					Action{
						Text: "关于",
						OnTriggered: func() {
							walk.MsgBox(mw, "关于", "这是一个菜单和工具栏的实例",
								walk.MsgBoxIconInformation|walk.MsgBoxDefButton1)
						},
					},
				},
			},
		},
		ToolBar: ToolBar{ //工具栏
			ButtonStyle: ToolBarButtonTextOnly,
			Items: []MenuItem{
				Menu{
					Text: "New",
					Items: []MenuItem{
						Action{
							Text:        "A",
							OnTriggered: mw.newAction_Triggered,
						},
						Action{
							Text:        "B",
							OnTriggered: mw.newAction_Triggered,
						},
					},
					OnTriggered: mw.newAction_Triggered, //在菜单中不可如此定义，会无响应
				},
				Separator{}, //分隔符
				Action{
					Text:        "View",
					OnTriggered: mw.changeViewAction_Triggered,
				},
			},
		},
		Layout: VBox{},
		Children: []Widget{
			TextEdit{
				AssignTo: &mw.edit,
			},
		},
		OnDropFiles: mw.dropFiles, //放置文件事件响应函数
	}).Create(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
 
	mw.Run()
}
 
func (mw *MyMainWindow) openFileActionTriggered() {
	dlg := new(walk.FileDialog)
	dlg.Title = "打开文件"
	dlg.Filter = "文本文件 (*.txt)|*.txt|所有文件 (*.*)|*.*"
 
	if ok, err := dlg.ShowOpen(mw); err != nil {
		fmt.Fprintln(os.Stderr, "错误：打开文件时\r\n")
		return
	} else if !ok {
		fmt.Fprintln(os.Stderr, "用户取消\r\n")
		return
	}
 
	s := fmt.Sprintf("选择了：%s\r\n", dlg.FilePath)
	mw.edit.SetText(s)
}
 
func (mw *MyMainWindow) saveFileActionTriggered() {
	dlg := new(walk.FileDialog)
	dlg.Title = "另存为"
	dlg.Filter = "文本文件 (*.txt)|*.txt|所有文件 (*.*)|*.*"
 
	if ok, err := dlg.ShowSave(mw); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	} else if !ok {
		fmt.Fprintln(os.Stderr, "取消")
		return
	}
 
	data := mw.edit.Text()
	filename := dlg.FilePath
	f, err := os.Open(filename)
	if err != nil {
		f, _ = os.Create(filename)
	} else {
		f.Close()
		f, err = os.OpenFile(filename, os.O_WRONLY, 0x666)
	}
	if len(data) == 0 {
		f.Close()
		return
	}
	io.Copy(f, strings.NewReader(data))
	f.Close()
}
 
func (mw *MyMainWindow) newAction_Triggered() {
	walk.MsgBox(mw, "New", "Newing something up... or not.", walk.MsgBoxIconInformation)
}
 
func (mw *MyMainWindow) changeViewAction_Triggered() {
	walk.MsgBox(mw, "Change View", "By now you may have guessed it. Nothing changed.", walk.MsgBoxIconInformation)
}
 
func (mw *MyMainWindow) dropFiles(files []string) {
	mw.edit.SetText("")
	for _, v := range files {
		mw.edit.AppendText(v + "\r\n")
	}
}