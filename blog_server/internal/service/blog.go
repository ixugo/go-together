package service

import (
	"together/global"
	pb "together/proto"

	"github.com/gocolly/colly/v2"
)

func getWebsite(domain string) []*pb.GetListReply_Data {
	switch domain {
	case global.BlogServer.IxugoURL:
		return getIxugo(domain)
	}
	return nil
}

func getIxugo(domain string) []*pb.GetListReply_Data {
	a := assets.New()
	data := make([]*pb.GetListReply_Data, 0, 10)
	a.OnHTML("main section", func(e *colly.HTMLElement) {
		e.ForEach("article", func(i int, h *colly.HTMLElement) {
			const prefix = "header div"
			art := pb.GetListReply_Data{
				Img:         "",
				Title:       h.ChildText(prefix + " h2 a"),
				Description: "",
				CreateAt:    h.ChildText(prefix + " footer time"),
				Tags:        h.ChildTexts(prefix + " header a"),
				Category:    "",
				Link:        h.ChildAttr(prefix+" h2 a", "href"),
			}
			data = append(data, &art)
		})
	})
	a.Visit(domain)

	return data
}
