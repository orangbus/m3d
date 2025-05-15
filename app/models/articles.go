package models

type Articles struct {
	TableId
	SiteName string `json:"site_name"`
	Title    string `json:"title"`
	CateName string `json:"cate_name"`
	Content  string `json:"content"`
	Url      string `json:"url"`
	TableTime
}
