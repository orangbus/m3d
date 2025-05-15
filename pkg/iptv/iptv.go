package iptv

// #EXTM3U
// #EXTINF:-1 logo='..png' group-title=名称 name=名称,
// http://xxx.m3u8
type Iptv struct {
	Logo       string `json:"logo"`
	GroupTitle string `json:"group-title"`
	Url        string `json:"url"`
}
