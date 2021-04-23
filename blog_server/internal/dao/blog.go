package dao

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v7"
	"together/blog_server/internal/model"
	pb "together/proto"
	"together/utils"
)

var EsClient *elasticsearch.Client

func SelectBlogMenusByUrl(url string, context context.Context) (menus []*pb.GetListReply_Data, err error) {
	blogSearchBody := make(utils.BodyMap)
	blogSearchBody.SetBodyMap("query", func(bm utils.BodyMap) {
		bm.SetBodyMap("match", func(bm utils.BodyMap) {
			bm.Set("url.keyword", url)
		})
	})

	res, err := EsClient.Search(
		EsClient.Search.WithContext(context),
		EsClient.Search.WithIndex("blog"),
		EsClient.Search.WithBody(blogSearchBody.BufferBody()),
		EsClient.Search.WithTrackTotalHits(true),
		EsClient.Search.WithPretty(),
	)

	if err != nil || res.IsError() {
		return nil, err
	}
	dataResp := make(map[string]interface{})
	if err := json.NewDecoder(res.Body).Decode(&dataResp); err != nil {
		return nil, err
	}
	total := dataResp["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)
	if total == 0 {
		return make([]*pb.GetListReply_Data, 0, 0), nil
	}
	menus = make([]*pb.GetListReply_Data, 0, 10)
	for _, hit := range dataResp["hits"].(map[string]interface{})["hits"].([]interface{}) {
		for _, value := range hit.(map[string]interface{})["_source"].(map[string]interface{})["menus"].([]interface{}) {
			var data *pb.GetListReply_Data
			marshal, _ := json.Marshal(value)
			err := json.Unmarshal(marshal, &data)
			if err != nil {
				return nil, errors.New("JSON系列化错误")
			}
			menus = append(menus, data)
		}
	}
	return menus, nil
}

func InsertBlog(url string, menus []*pb.GetListReply_Data, context context.Context) error {
	esBlogReq := &model.EsBlogReq{
		Url:   url,
		Menus: menus,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(esBlogReq)
	if err != nil {
		return err
	}
	res, err := EsClient.Index("blog", &buf, EsClient.Index.WithContext(context))
	if err != nil || res.IsError() {
		return err
	}
	return nil
}
