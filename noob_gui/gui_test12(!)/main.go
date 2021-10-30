package main
 
import (
	"log"
	"strconv"
	"fmt"
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

//不能改变
const FORMATWIDGET int = 80		//框框宽
const FORMATHEIGHT int = 55		//框框高

//可以改变
const MAXCOUNT float64 = 10		//X轴坐标最大值
const MAXSerial int = 6000		//Y轴坐标最大值
const MINSerial int = -6000		//Y轴坐标最小值

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
				Paint:               mw.CoordinateSystemInit,
			},
		},
	}).Run(); err != nil {
		log.Fatal(err)
	}
	//按关闭后才会执行（bug）
	a := 0
	for a < 100 {
		fmt.Printf("a=%d\n", a)
		a++
	}

}


func (mw *MyMainWindow) CoordinateSystemInit(canvas *walk.Canvas, updateBounds walk.Rectangle) error {
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
	
	rectPen, err := walk.NewCosmeticPen(walk.PenSolid, walk.RGB(139, 69, 19))
	if err != nil {
		return err
	}
	defer rectPen.Dispose()
	linePen, err := walk.NewCosmeticPen(walk.PenSolid, walk.RGB(205, 201, 201))
	if err != nil {
		return err
	}
	defer linePen.Dispose()
	if err := canvas.DrawRectanglePixels(rectPen, bounds); err != nil {
		return err
	}
	font, err := walk.NewFont("Arial", 8, walk.FontBold)
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
			var nowcount_str string
			if int(MAXCOUNT) < bounds.Width / FORMATWIDGET {
				nowcount_str = strconv.FormatFloat(MAXCOUNT / float64(lines_Width_count) * float64(count), 'g', 5, 64)
			} else {
				nowcount_str = strconv.Itoa(int(MAXCOUNT) / lines_Width_count * count)
			}
			bounds_font := walk.Rectangle{X:lines_Width[key].X - len(nowcount_str)*2, Y:bounds.Y + bounds.Height, Width:35, Height:17}
			
			if err := canvas.DrawTextPixels(nowcount_str, font, walk.RGB(10, 10, 10), bounds_font, walk.TextWordbreak); err != nil {
				return err
			}
			count++
		}
	}

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
