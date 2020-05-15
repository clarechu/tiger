package utils

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)


func ReadFile(filePath string, old, new string) (output []byte, err error) {
	output = make([]byte, 0)
	in, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer in.Close()
	br := bufio.NewReader(in)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return output, err
		}
		if strings.Contains(string(line), old) {
			newLine := strings.Replace(string(line), old, new, -1)
			newByte := []byte(newLine)
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
		} else {
			output = append(output, line...)
			output = append(output, []byte("\n")...)
		}

	}
	return output, nil
}

func WriteToFile(fileName string, outPut []byte) (err error) {
	var file *os.File
	if checkFileIsExist(fileName) { //如果文件存在
/*		file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Println(err.Error())
		}*/
		err := os.Remove(fileName)
		if err != nil {
			log.Println(err.Error())
		}
		file, err = os.Create(fileName) //创建文件
		if err != nil {
			return err
		}
	} else {
		file, err = os.Create(fileName) //创建文件
		if err != nil {
			return err
		}
	}
	defer file.Close()
	file.Write(outPut)
	return err
}

/**
 * checkFileIsExist 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}