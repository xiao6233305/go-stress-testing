package model

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

/**
 返回get 或者post 对应的大量的数据
 */
func ParseMultFile(path string) (data []string, err error) {
	if path == "" {
		err = errors.New("路径不能为空")
		return
	}
	file, err := os.Open(path)
	if err != nil {
		err = errors.New("打开文件失败:" + err.Error())
		return
	}
	defer func() {
		_ = file.Close()
	}()
	buf := bufio.NewReader(file)
	for{
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if len(line)>0{
			data = append(data,line)
		}
		if err == io.EOF {
			break
		}
	}
	return
}
