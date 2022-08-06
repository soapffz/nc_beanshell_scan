package nc

import (
	"fmt"
	"nc_beanshell_scan/module/queue"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gookit/color"
)

type FinScan struct {
	UrlQueue *queue.Queue
	Ch       chan []string
	Wg       sync.WaitGroup
	Thread   int
	Output   []string
}

func NewScan(urls []string, thread int, output []string) *FinScan {
	s := &FinScan{
		UrlQueue: queue.NewQueue(),
		Ch:       make(chan []string, thread),
		Wg:       sync.WaitGroup{},
		Thread:   thread,
		Output:   output,
	}
	for _, url := range urls {
		s.UrlQueue.Push([]string{url, "0"})
	}
	return s
}

func (s *FinScan) StartScan() {
	filename := DateNowFormatStr()

	for i := 0; i <= s.Thread; i++ {
		s.Wg.Add(1)
		go func() {
			defer s.Wg.Done()
			s.ncbeanshellScan()
		}()
	}
	s.Wg.Wait()
	if len(s.Output) != 0 {
		outputfile(filename, s.Output)
	}

}

func (s *FinScan) ncbeanshellScan() {
	// 创建存储用的数组
	for s.UrlQueue.Len() != 0 {
		dataface := s.UrlQueue.Pop()
		switch dataface := dataface.(type) {
		case []string:
			url := dataface
			var data *resps
			data, err := endpointrequest(url)
			if err != nil {
				url[0] = strings.ReplaceAll(url[0], "https://", "http://")
				data, err = endpointrequest(url)
				if err != nil {
					continue
				}
			}
			// fmt.Println(data.statuscode)
			if data.statuscode == 200 {
				data, random_str, nil := rcerequest(data.url)
				if err != nil {
					continue
				}
				if strings.Contains(data.body, "BeanShell") && strings.Contains(data.body, random_str) {
					s.Output = append(s.Output, data.url)
					color.Green.Printf("[+] " + data.url + " 存在beanshell rce漏洞！\n")
				}
			}
		default:
			continue
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
