package main
 
import (
	"log"
	"strconv"
	// "fmt"
	// "math"
)
 
import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)
 

type MyMainWindow struct {
	*walk.MainWindow
	paintWidget *walk.CustomWidget
}

const LINEWIDTHNUM int = 11
const LINEHEIGTNUM int = 8
const MAXCOUNT int = 800

const FORMATWIDGET int = 80
const FORMATHEIGHT int = 55

const MAXSerial int = 6000
const MINSerial int = -6000

func main() {
	mw := new(MyMainWindow)

	if _, err := (MainWindow{
		AssignTo: 	&mw.MainWindow,
		Title: 		"坐标图-2",
		MinSize:	Size{400, 400},
		Size:		Size{800, 600},
		Layout: 	VBox{MarginsZero: true},
		Children: []Widget{
			CustomWidget{
				AssignTo:            &mw.paintWidget,
				ClearsBackground:    true,
				InvalidatesOnResize: true,
				Paint:               mw.drawStuff,
			},
		},
	}).Run(); err != nil {
		log.Fatal(err)
	}
}


func (mw *MyMainWindow) drawStuff(canvas *walk.Canvas, updateBounds walk.Rectangle) error {
	bmp, err := createBitmap()
	if err != nil {
		return err
	}
	defer bmp.Dispose()
	//ClientBounds获取边界，可以随着拖动变化
	bounds := mw.paintWidget.ClientBoundsPixels()
	bounds.X += 40
	bounds.Y += 20
	bounds.Width -= 60
	bounds.Height -= 40

	//Snow3 RGB(205, 201, 201)
	//NewCosmeticPen()只能生成任意PEN类型和任意颜色的画笔
	//不能设置宽度，默认为1
	//PenSolid			实线
	//PenDash			长虚线
	//PenDot			短虚线
	//PenDashDot		点划线
	//PenDashDotDot		短线划线
	//PenNull			无
	//PenInsideFrame	暂无
	//PenUserStyle		暂无
	//PenAlternate		密集点
	
	rectPen, err := walk.NewCosmeticPen(walk.PenSolid, walk.RGB(139, 69, 19))
	if err != nil {
		return err
	}
	defer rectPen.Dispose()
	// //打印画笔宽度
	// fmt.Println(rectPen.Width())

	linePen, err := walk.NewCosmeticPen(walk.PenSolid, walk.RGB(205, 201, 201))
	if err != nil {
		return err
	}
	defer linePen.Dispose()
	//画矩形
	if err := canvas.DrawRectanglePixels(rectPen, bounds); err != nil {
		return err
	}

	//step2:建立坐标系

	//新建字体
	//FontBold			粗体
	//FontItalic		斜体
	//FontUnderline		下划线
	//FontStrikeOut		删除线
	font, err := walk.NewFont("Times New Roman", 8, walk.FontBold)
	if err != nil {
		return err
	}
	defer font.Dispose()
	
	var lines_Width_count int
	var lines_Height_count int

	lines_Width_count = bounds.Width / FORMATWIDGET
	lines_Height_count = bounds.Height / FORMATHEIGHT

	var lines_Width []walk.Point = make([]walk.Point, lines_Width_count*2)
	var lines_Height []walk.Point = make([]walk.Point, lines_Height_count*2)
	//放置0
	if err := canvas.DrawTextPixels(
		"0", 
		font, 
		walk.RGB(10, 10, 10), 
		walk.Rectangle{X:bounds.X, Y:bounds.Y + bounds.Height, Width:25, Height:20}, 
		walk.TextWordbreak); 
		err != nil {
		return err
	}

	count := 1
	for key, _ := range lines_Width {
		if key % 2 == 0 {
			lines_Width[key] = walk.Point{X: bounds.X + FORMATWIDGET*count, Y: bounds.Y }
		} else {
			lines_Width[key] = walk.Point{X: bounds.X + FORMATWIDGET*count, Y: bounds.Y + bounds.Height}
			if err := canvas.DrawLinePixels(linePen, lines_Width[key-1], lines_Width[key]); err != nil {
				return err
			}

			nowcount_str := strconv.Itoa(MAXCOUNT / lines_Width_count  * count)
			bounds_font := walk.Rectangle{X:lines_Width[key].X - len(nowcount_str) - 5, Y:bounds.Y + bounds.Height, Width:25, Height:20}
			
			if err := canvas.DrawTextPixels(nowcount_str, font, walk.RGB(10, 10, 10), bounds_font, walk.TextWordbreak); err != nil {
				return err
			}
			count++
		}
	}

	//放置MAXSerial
	if err := canvas.DrawTextPixels(
		strconv.Itoa(MAXSerial), 
		font, 
		walk.RGB(10, 10, 10), 
		walk.Rectangle{X:bounds.X - 23, Y:bounds.Y - 5, Width:25, Height:20}, 
		walk.TextWordbreak); 
		err != nil {
		return err
	}

	count = 0
	for key, _ := range lines_Height {
		if key % 2 == 0 {
			count++
			lines_Height[key] = walk.Point{X: bounds.X, Y: bounds.Y + FORMATHEIGHT*count}
		} else {
			lines_Height[key] = walk.Point{X: bounds.X + bounds.Width, Y: bounds.Y + FORMATHEIGHT*count}
			if err := canvas.DrawLinePixels(linePen, lines_Height[key-1], lines_Height[key]); err != nil {
				return err
			}
			nowserial_str := strconv.Itoa(MAXSerial - ((MAXSerial - MINSerial) / lines_Height_count) * count)
			
			bounds_font := walk.Rectangle{X:bounds.X - len(nowserial_str)*4 - 8, Y:lines_Height[key].Y - 6, Width:25, Height:20}
			
			if err := canvas.DrawTextPixels(nowserial_str, font, walk.RGB(10, 10, 10), bounds_font, walk.TextWordbreak); err != nil {
				return err
			}
		}
	}


	// //画圆的画笔
	// ellipseBrush, err := walk.NewHatchBrush(walk.RGB(0, 255, 0), walk.HatchDiagonalCross)
	// if err != nil {
	// 	return err
	// }
	// defer ellipseBrush.Dispose()
 
	// if err := canvas.FillEllipsePixels(ellipseBrush, bounds); err != nil {
	// 	return err
	// }
 
	return nil
}

func createBitmap() (*walk.Bitmap, error) {
	bounds := walk.Rectangle{Width: 600, Height: 600}
 
	bmp, err := walk.NewBitmapForDPI(bounds.Size(), 192)
	if err != nil {
		return nil, err
	}
 
	succeeded := false
	defer func() {
		if !succeeded {
			bmp.Dispose()
		}
	}()
 
	// canvas, err := walk.NewCanvasFromImage(bmp)
	// if err != nil {
	// 	return nil, err
	// }
	// defer canvas.Dispose()
 
	succeeded = true
 
	return bmp, nil
}
