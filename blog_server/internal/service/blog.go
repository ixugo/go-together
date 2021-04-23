package service

import (
	"context"
	"log"
	"together/blog_server/internal/dao"
	"together/global"
	pb "together/proto"

	"github.com/gocolly/colly/v2"
)

func getWebsite(url string) []*pb.GetListReply_Data {
	menus, err := dao.SelectBlogMenusByUrl(url, context.Background())
	if err != nil {
		log.Println(err)
		return nil
	}
	if len(menus) == 0 {
		switch url {
		case global.BlogServer.IxugoDomain:
			return getIxugo(url)
		case global.BlogServer.WangboDomain:
			return getWangbo(url)
		}
		return nil
	}
	return menus
}

func getIxugo(url string) []*pb.GetListReply_Data {
	// TODO 识别链接中的域名作为参数填入下方
	a := assets.New()
	menus := make([]*pb.GetListReply_Data, 0, 10)
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
				Link:        url + h.ChildAttr(prefix+" h2 a", "href"),
			}
			menus = append(menus, &art)
		})
	})
	a.Visit(url)
	err := dao.InsertBlog(url, menus, context.Background())
	if err != nil {
		return nil
	}
	return menus
}

func getWangbo(url string) []*pb.GetListReply_Data {
	// TODO 识别链接中的域名作为参数填入下方
	a := assets.New()
	menus := make([]*pb.GetListReply_Data, 0, 10)
	a.OnHTML(".recent-posts", func(e *colly.HTMLElement) {
		e.ForEach(".recent-post-item", func(i int, h *colly.HTMLElement) {
			const website = "https://chenyunxin.cn"
			art := pb.GetListReply_Data{
				Img:         h.ChildAttr(".post_cover a img", "data-original"),
				Title:       h.ChildAttr(".post_cover a", "title"),
				Description: "",
				CreateAt:    h.ChildText(".recent-post-info div time"),
				Tags:        []string{},
				Category:    h.ChildText(".article-meta__categories"),
				Link:        website + h.ChildAttr(".post_cover a", "href"),
			}
			menus = append(menus, &art)
		})
	})
	a.Visit(url)
	err := dao.InsertBlog(url, menus, context.Background())
	if err != nil {
		return nil
	}
	return menus
}
