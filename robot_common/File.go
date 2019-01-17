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
