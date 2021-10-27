package _test

import (
	"time"
	"fmt"
)
//停止位
const (
	//一位(常用)
	Stop1     byte = 1
	Stop1Half byte = 15
	Stop2     byte = 2
)
//校验位
const (
	//无(常用)
	ParityNone  byte = 'N'
	ParityOdd   byte = 'O'
	ParityEven  byte = 'E'
	ParityMark  byte = 'M' // parity bit is always 1
	ParitySpace byte = 'S' // parity bit is always 0
)

type Config struct {
	//端口名字
	Name        string
	//速率
	Baud        int//默认	115200
	//超时时间
	ReadTimeout time.Duration

	// 数据位
	Size 		byte//默认	8位

	// 校验位
	Parity 		byte//默认	无

	// 停止位
	StopBits 	byte//默认	1位
}

// ErrBadSize 错误的数据位
var ErrBadSize error = errors.New("unsupported serial data size")

// ErrBadStopBits 错误的停止位
var ErrBadStopBits error = errors.New("unsupported stop bit setting")

// ErrBadParity 错误的校验位
var ErrBadParity error = errors.New("unsupported parity setting")

const (
	//无(常用)
	DataSize_5  	byte = 5
	DataSize_6   	byte = 6
	DataSize_7  	byte = 7
	DataSize_8  	byte = 8
)

func OpenPort(c *Config) (*Port, error) {
	size, par, stop := c.Size, c.Parity, c.StopBits
	if size == 0 {
		size = DataSize_8
	}
	if par == 0 {
		par = ParityNone
	}
	if stop == 0 {
		stop = Stop1
	}
	return openPort(c.Name, c.Baud, size, par, stop, c.ReadTimeout)
}

func main() {
	now := time.Now()
	fmt.Println(now)
}