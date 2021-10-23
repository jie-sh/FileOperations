package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//检查错误
func check(err error) {
	if err != nil {
		panic(err)
	}
}

//通过 ioutil 直接通过文件名来加载文件
//一次将整个文件加载进来， 粒度较大
func ReadFile(filepath string) {
	dat, err := ioutil.ReadFile(filepath)
	check(err)
	//输出字符的 ASCII 码
	fmt.Println(dat)
	//将文件内容转换为 string 输出
	fmt.Println(string(dat))

}

//通过 os.Open 的方式得到 *File 指针，通过这个指针 可以对文件进行更细粒度的操作
func ReadFile2(filepath string) {
	f, err := os.Open(filepath)
	check(err)

	//一
	//使用 make 函数来分配指定大小的缓冲区
	buffer := make([]byte, 15)
	//从文件中读取数据到缓冲区
	data, err := f.Read(buffer)
	check(err)
	fmt.Printf("%d bytes: \n%s\n", data, string(buffer))

	//二
	//通过 f.seek 进行更精细的操作，设置下一次读/写的位置
	//offset 设置相对偏移， whence 表示文件起始的相对位置
	//设置相对文件头偏移为5的位置开始往后读取数据
	s, err := f.Seek(5, 0)
	check(err)
	buffer2 := make([]byte, 5)
	data2, err := f.Read(buffer2)
	fmt.Printf("%d bytes after %d position: \n%s\n", data2, s, string(buffer2))

	//三
	//通过io包种的函数 也可以实现类似的功能
	o, err := f.Seek(5, 0)
	check(err)
	buffer3 := make([]byte, 5)
	//io.ReadAtLeast 从文件 f 中至少读取 len(buffer3) 个数据到 buffer3 中
	//在此代码中 len(buffer3) 为所能读取的最多数据，若超过则会报错
	data3, err := io.ReadAtLeast(f, buffer3, len(buffer3))
	check(err)
	fmt.Printf("%d bytes after %d position: \n%s\n", data3, o, string(buffer3))

	f.Close()
}

//通过bufio包来进行读取 bufio中又许多比较有用的函数 比如一次读入一整行的内容
func ReadFile3(filepath string) {
	f, err := os.Open(filepath)
	check(err)
	//调整文件指针的起始位置到最开始的地方
	_, err = f.Seek(0, 0)
	check(err)
	r := bufio.NewReader(f)

	//读取从头开始的5个字节
	//Peek() 返回输入流的下 n 个字节，而不会移动读取位置，返回的字节只在下一次调用读取操作前合法
	buffer, err := r.Peek(5)
	check(err)
	fmt.Printf("5 bytes: \n%s\n", string(buffer))

	//调整文件相对位置
	_, err = f.Seek(5, 0)
	check(err)
	r2 := bufio.NewReader(f)
	buffer2, err := r2.Peek(5)
	check(err)
	fmt.Printf("5 bytes: \n%s\n", string(buffer2))

	//使用 bufio 的其他函数
	for {
		//读出内容保存为string 每次读到以'\n'为标记的位置
		line, err := r.ReadString('\n')
		fmt.Print(line)
		//读取到文件结尾
		if err == io.EOF {
			break
		}
	}

	/* for {
		bytes, err := r.ReadByte()
		fmt.Printf("%c", bytes)
		if err == io.EOF {
			break
		}
	} */

	/* for {
		line, _, err := r.ReadLine()
		fmt.Printf("%s\n", line)
		if err == io.EOF {
			break
		}
	} */

	f.Close()
}

//ioutil.ReadAll
func ReadFile4(filepath string) {
	f, err := os.Open(filepath)
	check(err)
	bytes, err := ioutil.ReadAll(f)
	check(err)
	fmt.Println(bytes)
	fmt.Println(string(bytes))
}

func main() {
	file_path := "D:\\mbx3000d"

	//读取文件的四种方法，分别为：使用 ioutil.ReadFile 读取文件，使用 file.Read 读取文件，使用 bufio.NewReader 读取文件，使用 ioutil.ReadAll 读取文件

	//ioutil.ReadFile
	//ReadFile(file_path)

	//file.Read
	//ReadFile2(file_path)

	//bufio.NewReader
	//ReadFile3(file_path)

	//ioutil.ReadAll
	ReadFile4(file_path)
}
