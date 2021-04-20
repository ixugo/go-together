package ierr

// 业务错误
var (
	GetBlog = NewError(10000, "查询博客列表失败")
)
