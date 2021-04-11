package assets

import (
	"fmt"
	"sync"

	"github.com/gocolly/colly/v2"
)

type Assets interface {
	OnHTML(string, func(*colly.HTMLElement))
	Visit(s string) error
	GetColly() *colly.Collector
	New(domain ...string) Assets
}

type assets struct {
	c *colly.Collector
}

var _assets *assets

// 创建资源对象
var once sync.Once

func GetInstance() Assets {
	once.Do(func() {
		_assets = &assets{
			c: new(),
		}
	})
	return _assets
}

func new() *colly.Collector {
	return colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Mobile Safari/537.36"),
		colly.MaxDepth(2),
	)
}

func (a *assets) New(domain ...string) Assets {
	t := a.c.Clone()
	t.OnError(onError)
	t.OnRequest(onRequest)
	colly.AllowedDomains(domain...)(t)
	return &assets{
		c: t,
	}
}

func (a *assets) GetColly() *colly.Collector {
	return a.c
}

func (a *assets) GetAssets() {
	// return
}

func (a *assets) OnHTML(selector string, f func(*colly.HTMLElement)) {
	a.c.OnHTML(selector, f)
}

func (a *assets) Visit(s string) error {
	return a.c.Visit(s)
}

// 统一 error 处理
func onError(resp *colly.Response, err error) {
	fmt.Println(resp.Request.URL, err)
}

// 请求前记录日志
func onRequest(r *colly.Request) {
	fmt.Println("Visiting", r.URL)
}
