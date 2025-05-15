package models

type Movies struct {
	ApiId        int32  `gorm:"column:api_id;not null;comment:接口id" json:"api_id"`                        // 接口id
	Status       int32  `gorm:"column:status;not null;default:1;comment:可用状态，默认1，可用，2：不可用" json:"status"` // 可用状态，默认1，可用，2：不可用
	VodID        int32  `gorm:"column:vod_id;not null;comment:视频ID" json:"vod_id"`                        // 视频ID
	VodName      string `gorm:"column:vod_name;not null;comment:标题" json:"vod_name"`                      // 标题
	TypeID       int32  `gorm:"column:type_id;not null;comment:分类id" json:"type_id"`                      // 分类id
	VodTag       string `gorm:"column:vod_tag;not null;comment:类型标签" json:"vod_tag"`                      // 类型标签
	TypeName     string `gorm:"column:type_name;not null;comment:分类名称" json:"type_name"`                  // 分类名称
	VodPic       string `gorm:"column:vod_pic;not null;comment:图片" json:"vod_pic"`                        // 图片
	VodArea      string `gorm:"column:vod_area;not null;comment:地区" json:"vod_area"`                      // 地区
	VodLang      string `gorm:"column:vod_lang;not null;comment:语言" json:"vod_lang"`                      // 语言
	VodYear      string `gorm:"column:vod_year;not null;comment:年份" json:"vod_year"`                      // 年份
	VodActor     string `gorm:"column:vod_actor;not null;comment:演员" json:"vod_actor"`                    // 演员
	VodDirector  string `gorm:"column:vod_director;not null;comment:导演" json:"vod_director"`              // 导演
	VodRemarks   string `gorm:"column:vod_remarks;not null;comment:更新集数" json:"vod_remarks"`              // 更新集数
	VodContent   string `gorm:"column:vod_content;comment:介绍" json:"vod_content"`                         // 介绍
	VodHits      int32  `gorm:"column:vod_hits;not null;comment:点击数" json:"vod_hits"`                     // 点击数
	VodHitsDay   int32  `gorm:"column:vod_hits_day;not null;comment:今日点击" json:"vod_hits_day"`            // 今日点击
	VodHitsWeek  int32  `gorm:"column:vod_hits_week;not null;comment:本周点击" json:"vod_hits_week"`          // 本周点击
	VodHitsMonth int32  `gorm:"column:vod_hits_month;not null;comment:本月" json:"vod_hits_month"`          // 本月

	VodPlayFrom string `gorm:"column:vod_play_from;not null;comment:播放来源" json:"vod_play_from"`     // 播放来源
	VodPlayNote string `gorm:"column:vod_play_note;not null;comment:播放链接切割依据" json:"vod_play_note"` // 播放链接切割依据
	VodPlayURL  string `gorm:"column:vod_play_url;comment:播放地址" json:"vod_play_url"`                // 播放地址
	VodDownNote string `gorm:"column:vod_down_note;not null;comment:播放类型" json:"vod_down_note"`     // 播放类型
}
