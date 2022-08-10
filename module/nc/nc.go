package nc

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gookit/color"
	"github.com/panjf2000/ants/v2"
)

var url_l []string

func StartScan(urls []string, thread int, output []string) {
	filename := DateNowFormatStr()
	defer ants.Release()
	var wg sync.WaitGroup
	// Use the pool with a function,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		Scan(i)
		wg.Done()
	})
	defer p.Release()
	// Submit tasks one by one.
	for _, url := range urls {
		wg.Add(1)
		_ = p.Invoke(url)
	}
	wg.Wait()
	if len(url_l) != 0 {
		outputfile(filename, url_l)
	}
}

func Scan(i interface{}) {
	url, ok := i.(string)
	if !ok {
		return
	}
	var data *resps
	data, err := endpointrequest(url)
	if err != nil {
		url = strings.ReplaceAll(url, "https://", "http://")
		data, err = endpointrequest(url)
		if err != nil {
			return
		}
	}
	// fmt.Println(data.statuscode)
	if data.statuscode == 200 {
		data, random_str, nil := rcerequest(data.url)
		if err != nil {
			return
		}
		if strings.Contains(data.body, "BeanShell") && strings.Contains(data.body, random_str) {
			url_l = append(url_l, data.url)
			color.Green.Printf("[+] " + data.url + " 存在beanshell rce漏洞！\n")
		}
	}
}

func outputfile(filename string, url_l []string) {
	//运行完后写文件到日志
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("创建日志文件失败，请检查：" + err.Error())
	} else {
		defer f.Close()
		for _, v := range url_l {
			f.Write([]byte(v + "\n"))
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}
}

func DateNowFormatStr() string {
	tm := time.Now()
	return tm.Format("log_20060102_150405.txt")
}
