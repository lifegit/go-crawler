package ipManager

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

var IpDoMain *IpDomain

func init() {
	IpDoMain = New()
}

type IpDomain struct {
	pointer int
	domain  []domain
}
type domain struct {
	url      string
	errCount int
}

func New() *IpDomain {
	return &IpDomain{
		pointer: -1,
		domain: []domain{
			{url: "http://ip.3322.net"},
			{url: "http://ident.me/"},
			{url: "http://whatismyip.akamai.com/"},
		},
	}
}

func (i *IpDomain) GetIp() (string, error) {
	i.pointer++
	if i.pointer >= len(i.domain)-1 {
		i.pointer = 0
	}

	return i.domain[i.pointer].url, nil
}

/*
	验证代理ip是否可用
	通过传入一个代理ip，然后使用它去访问一个url看看是否访问成功，以此为依据进行判断当前代理ip是否有效。
	参数：proxy_addr 要验证的ip
	返回：ip 验证通过的ip、status 状态（200表示成功）
*/
func (i *IpDomain) CheckIpProxyThorn(proxyAddr string, httpUrl string) (ip string, status int) {
	//访问查看ip的一个网址
	// httpUrl := "http://ip.3322.net"
	proxy, err := url.Parse(proxyAddr)

	httpClient := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			Proxy:                 http.ProxyURL(proxy),
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * time.Duration(5),
		},
	}
	res, err := httpClient.Get(httpUrl)
	if err != nil {
		//fmt.Println("错误信息：",err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Println(err)
		return
	}
	c, _ := ioutil.ReadAll(res.Body)

	return string(c), res.StatusCode
}
