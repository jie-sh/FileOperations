package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

//判断文件、文件夹是否存在
func CheckFileExists(filename string) bool {
	//os.Stat(name) 返回描述文件的FileInfo类型值，用来获取文件属性
	if _, err := os.Stat(filename); err != nil {
		//若文件不存在则 err != nil，故 IsNotExist 为 false
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func WriteFile(filename string, strdata string) {
	if CheckFileExists(filename) {
		fmt.Printf("文件已存在！\n")
	} else {
		//创建文件
		_, err := os.Create(filename)
		check(err)
		fmt.Printf("文件不存在！创建新文件\n")
	}
	//打开文件
	f, err := os.OpenFile(filename, os.O_APPEND, 0666)
	check(err)
	//写入字符串
	data, err := f.WriteString(strdata)
	check(err)
	fmt.Printf("写入了 %d 个字符串！\n", data)
	f.Close()
}

func WriteFile2(filename string, strdata string) {
	bytes := []byte(strdata)
	//将字节写入文件，若不存在则创建文件
	//使用 ioutil.WriteFile 写入文件时，如果文件存在，则首先会清空文件后再写入
	err := ioutil.WriteFile(filename, bytes, 0666)
	check(err)

	//读取文件写入的内容
	fileContent, err := ioutil.ReadFile(filename)
	check(err)
	fmt.Printf("数据写入成功！\n")
	fmt.Println(string(fileContent))
}

func WriteFile3(filename string, strdata string) {
	if CheckFileExists(filename) {
		fmt.Printf("文件已存在！\n")
	} else {
		//创建文件
		_, err := os.Create(filename)
		check(err)
		fmt.Printf("文件不存在！创建新文件\n")
	}
	//打开文件
	f, err := os.OpenFile(filename, os.O_APPEND, 0666)
	check(err)
	bytes := []byte(strdata)
	//写字节数组到文件
	data, err := f.Write(bytes)
	check(err)
	fmt.Printf("写入了 %d 个字节到文件！\n", data)
	f.Close()
}

func WriteFile4(filename string, strdata string) {
	if CheckFileExists(filename) {
		fmt.Printf("文件已存在！\n")
	} else {
		//创建文件
		_, err := os.Create(filename)
		check(err)
		fmt.Printf("文件不存在！创建新文件\n")
	}
	//打开文件
	f, err := os.OpenFile(filename, os.O_APPEND, 0666)
	check(err)
	//创建新的 Writer 对象
	w := bufio.NewWriter(f)
	data, err := w.WriteString(strdata)
	fmt.Printf("写入了 %d 个字节到文件！\n", data)
	//flush() 刷新缓冲区，即将缓冲区中的数据立刻写入文件，同时清空缓冲区，不需要被动的等待输出缓冲区写入
	//一般情况下，文件关闭后会自动刷新缓冲区，但有时需要在关闭前刷新它，这时就可以使用 flush() 方法
	w.Flush()
	f.Close()
}

func main() {
	filename := "D:\\Netch\\xz\\test.txt"
	strdata := "1234567890qwertyuiop"

	//io.WriteString
	//WriteFile(filename, strdata)

	//ioutil.WriteFile
	//WriteFile2(filename, strdata)

	//file.Write
	//WriteFile3(filename, strdata)

	//bufio.NewWriter
	WriteFile4(filename, strdata)
}
