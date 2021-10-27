package _11

import "fmt"
// import "unsafe"
//结构体
//格式：
/* 
type 变量名字 struct {
	变量名字1 数据类型
	变量名字2 数据类型
}
*/
type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {

	var book1 Books

	book1.title = "Go语言新手入门"
	book1.author = "github.com/zoulee24/GO_NoobNote"
	book1.subject = "Go语言教程"
	book1.book_id = 876582827

	fmt.Println(book1)
	fmt.Println("")
	print_Books(book1)
}
//结构体可以作为函数参数
//和c语言一样
func print_Books( book Books) {
	fmt.Printf( "Book title : %s\n", book.title)
	fmt.Printf( "Book author : %s\n", book.author)
	fmt.Printf( "Book subject : %s\n", book.subject)
	fmt.Printf( "Book book_id : %d\n", book.book_id)
}
