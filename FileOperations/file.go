package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/kardianos/osext"
)

//获取目录中的所有文件
func GetDirectoryAllFile(dirname string) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		//panic(err)
		fmt.Println(err.Error())
	}
	for _, file := range files {
		if file.IsDir() {
			//递归 循环获取子目录中的文件
			GetDirectoryAllFile(dirname + "\\" + file.Name())
		} else {
			fmt.Println(dirname + "\\" + file.Name()) //输出文件名
			fmt.Println(file.Size())                  //输出文件字节大小
			fmt.Println(file.ModTime())               //输出文件修改时间
			fmt.Println("----------------")
		}
	}
}

//获取所有文件夹、子文件夹、文件、子文件
func visit(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)
	return nil
}
func GetAllFiles(dirname string) {
	//递归遍历目录
	err := filepath.Walk(dirname, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}

//获取当前目录所有非文件夹的文件及其大小
func GetAllFilesAndSize() {
	dirname := "." + string(filepath.Separator)
	d, err := os.Open(dirname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//defer 代码块会在函数调用链表中增加一个函数调用，这个函数调用不是普通的函数调用，而是会在函数正常返回，
	//也就是 return 之后添加一个函数调用。因此，defer 通常用来释放函数内部变量
	//若上面操作失败，函数 return，就会调用 defer 函数释放句柄
	defer d.Close()
	//Readdir 读取目录并返回排好序的文件以及子目录名
	//-1 表示读取目录中的所有目录项
	fi, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, fi := range fi {
		//忽略不是普通的文件
		if fi.Mode().IsRegular() {
			fmt.Println(fi.Name(), fi.Size(), "bytes")
		}
	}
}

//获取正在执行的文件所在的目录
func GetEXEDirectory() {
	//filepath.Abs() 检测路径是否是绝对路径，是则直接返回，不是则会添加当前工作路径到参数path前，然后返回
	//filepath.Dir() 返回指定路径中除最后一个元素以外的所有元素
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		//打印输出错误码，退出应用程序，defer 函数不会执行
		log.Fatal(err)
	}

	fmt.Println(dir)
}

//获取当前文件所在的路径
func GetCurrentDirectory() {
	//获取当前文件路径
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)
}

// https://blog.csdn.net/JineD/article/details/116273979
func GetCurrentDirectory2() {
	//返回当前运行的程序路径，vscode 直接运行会输出临时路径
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
}

func GetCurrentDirectory3() {
	fmt.Println(filepath.Abs("./"))
}

//调用第三方库 https://github.com/kardianos/osext
//vscode 直接运行也会输出临时路径
func GetCurrentDirectory4() {
	//获取当前正在执行的文件路径
	folderEXEPath, err := osext.Executable()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(folderEXEPath)
	//获取当前目录
	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(folderPath)
}

//修改文件名
func ChangeFileName() {
	originalPath := "./test.txt"
	newPath := "test_new.txt"
	//重命名文件
	err := os.Rename(originalPath, newPath)
	//移动文件与重命名同理
	//os.Rename("D:/Netch/xz/test.txt", "D:/Netch/test.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func ChangeFileName2(originalPath string, newPath string) {
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}

//修改文件夹名称
func ChangeFolderName() {
	originalPath := "test"
	newPath := "test_new"
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}

//判断文件、文件夹是否存在
func Exists(name string) bool {
	//os.Stat(name) 返回描述文件的FileInfo类型值，用来获取文件属性
	if _, err := os.Stat(name); err != nil {
		//若文件不存在则 err != nil，故 IsNotExist 为 false
		if os.IsNotExist(err) {
			return false
		}
	}
	return true

	/* _, err := os.Stat(name)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true */
}

//判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

//创建文件夹
func Mkdir(dir string) {
	exist := Exists(dir)   //判断文件存在
	existdir := IsDir(dir) //判断是否是文件夹
	if exist {
		if existdir {
			fmt.Println(dir + " 文件夹已存在")
		} else {
			fmt.Println(dir + " 已存在与文件夹同名的文件")
		}
	} else {
		//文件夹名称，权限
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			fmt.Println(dir + " 文件夹创建失败：" + err.Error())
		} else {
			fmt.Println(dir + " 文件夹创建成功！")
		}
	}
}

//删除文件及文件夹
func RemoveDir(name string) {
	exist := Exists(name)
	if exist {
		//os.RemoveAll 是遍历删除，文件及文件及均可使用
		err := os.RemoveAll(name)
		if err != nil {
			fmt.Println(name + " 删除失败：" + err.Error())
		} else {
			fmt.Println(name + " 删除成功！")
		}
	} else {
		fmt.Println(name + " 文件或文件夹不存在！")
	}
}

//判断文件的读写权限
func FilePermission(filename string) {
	//写权限
	file, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		//返回一个布尔值说明该错误是否表示因权限不足要求被拒绝
		if os.IsPermission(err) {
			log.Println("Error: Write permission denied.")
		}
	}
	file.Close()

	//读权限
	file, err = os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Read permission denied.")
		}
	}
	file.Close()
}

func main() {
	//获取指定目录中的所有文件
	//my_folder := "D:\\Netch"
	//GetDirectoryAllFile(my_folder)

	//获取所有文件夹和文件
	//GetAllFiles(my_folder)

	//获取当前目录中的所有非文件夹的文件及其大小
	//GetAllFilesAndSize()

	//获取可执行文件路径
	//GetEXEDirectory()

	//获取当前文件所在的路径
	/* GetCurrentDirectory()
	GetCurrentDirectory2()
	GetCurrentDirectory3()
	GetCurrentDirectory4()*/

	//文件重命名
	//ChangeFileName()
	//originalPath := "./test.txt"
	//newPath := "test_new.txt"
	//ChangeFileName2(newPath, originalPath)

	//文件夹重命名
	//ChangeFolderName()

	//判断文件是否存在
	//originalPath := "demo"
	//result := Exists(originalPath)
	//fmt.Println(result)

	//判断是否为文件夹
	//originalPath := "xc"
	//originalPath := "D:\\Netch"
	//result := IsDir(originalPath)
	//fmt.Println(result)

	//创建文件夹
	//originalPath := "D:\\Netch\\xz"
	//Mkdir(originalPath)

	//删除文件及文件夹
	originalPath := "D:\\Netch\\xz\\we"
	RemoveDir(originalPath)

	//判断文件的读写权限
	//FilePermission("D:\\go\\src\\Project\\FILE\\file.go")
}
