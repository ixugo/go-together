package model

type Article struct {
	Img         string   `json:"img,omitempty"`         // 封面图
	Title       string   `json:"title,omitempty"`       // 标题
	Description string   `json:"description,omitempty"` // 描述
	CreateAt    string   `json:"create_at,omitempty"`   // 创建时间
	Tags        []string `json:"tags,omitempty"`        // 标签
	Category    string   `json:"category,omitempty"`    // 分类
	Link        string   `json:"link,omitempty"`        // 超链接
}
