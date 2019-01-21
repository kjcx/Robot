package robot_common

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)
type Account struct {
	UserName string
	Passwd string
}
func File(fileName string) []*Account{
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return nil
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	var size = stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)
	lines := []*Account{}
	for {
		account := &Account{}
		line, err := buf.ReadString('\n')
		//fmt.Println(line)
		line = strings.TrimSpace(line)
		//fmt.Println(line)
		line_arr := strings.Split(line,",")
		fmt.Println(line_arr)
		account.UserName = line_arr[0]
		account.Passwd = line_arr[1]
		lines = append(lines,account)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return lines
			}
		}
	}
	return lines
}

// fileName:文件名字(带全路径)
// content: 写入的内容
func AppendToFile(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, io.SeekEnd)
		fmt.Println("offset",n)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()
	return err
}
