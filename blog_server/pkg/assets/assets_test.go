package assets

import (
	"fmt"
	"testing"
	"together/model"

	"github.com/gocolly/colly/v2"
)

func TestCollly(t *testing.T) {
	i := GetInstance()
	a := i.New()

	data := make([]*model.Article, 0, 10)
	a.OnHTML("main section", func(e *colly.HTMLElement) {
		e.ForEach("article", func(i int, h *colly.HTMLElement) {
			const prefix = "header div"
			fmt.Printf("=============%d=============\n", i)

			art := model.Article{
				Img:      "",
				Title:    h.ChildText(prefix + " h2 a"),
				Tags:     h.ChildTexts(prefix + " header a"),
				Link:     h.ChildAttr(prefix+" h2 a", "href"),
				CreateAt: h.ChildText(prefix + " footer time"),
			}
			data = append(data, &art)
		})
	})

	page := 1
	a.OnHTML("main nav", func(h *colly.HTMLElement) {
		h.ForEach("a", func(i int, h *colly.HTMLElement) {
			page += 1
			if h.Text != fmt.Sprint(page) {
				return
			}
			fmt.Println("翻页", h.Text)

			href := h.Attr("href")
			fmt.Println("下一页", href)

			h.Request.Visit("localhost" + href)
		})
	})
	a.Visit("localhost")

	for _, v := range data {
		fmt.Println(v)
	}
}
