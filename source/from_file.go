package source

import (
	"bufio"
	"devops-dns-server/config"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

// 从 文件 中获取ip地址 比如从/etc/hosts文件中

var (
	NameIP = sync.Map{}
)

func IsExist(name string) bool {
	_, err := os.Stat(name)
	return err == nil || os.IsExist(err)
}

func readFile(filename string) error {
	var (
		err  error
		file *os.File
		buf  *bufio.Reader
		line []byte
	)

	if !IsExist(filename) {
		return fmt.Errorf("%s not exist", filename)
	}

	if file, err = os.Open(filename); err != nil {
		return fmt.Errorf("read %s error", filename)
	}

	buf = bufio.NewReader(file)
	for {
		line, _, err = buf.ReadLine()
		if err == io.EOF {
			break
		}
		ipName := strings.Split(string(line), " ")
		if ipName[0] == "" {
			continue
		}
		NameIP.Store(ipName[1], ipName[0])
	}
	return nil
}

func WatchFile() {
	var (
		//watcher  *fsnotify.Watcher
		err      error
		watch    string
		filename string
		interval int64
		ticker   *time.Ticker
	)

	filename = config.GetConfig().String("fromFile::filepath")
	if filename == "" {
		return
	}

	// 启动的时候读取一次该文件 然后再监控文件的修改
	if err = readFile(filename); err != nil {
		log.Println(err)
		return
	}

	watch = config.GetConfig().String("fromFile::watch")
	if watch != "yes" {
		return
	}

	if interval, err = config.GetConfig().Int64("fromFile::interval"); err != nil {
		log.Fatal("interval must be a number")
	}

	// 定时读取该文件
	ticker = time.NewTicker(time.Duration(interval) * time.Second)

	go func() {
		for _ = range ticker.C {
			if err = readFile(filename); err != nil {
				log.Println(err)
			}
		}
	}()

}

func FromFile(hostname string) string {
	if ip, ok := NameIP.Load(hostname); ok {
		return ip.(string)
	}
	return ""
}
