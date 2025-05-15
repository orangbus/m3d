package models

// 显示需要
type MovieCate struct {
	TypeID   int    `json:"type_id"`   // 分类id
	TypePid  int    `json:"type_pid"`  // 分类id
	TypeName string `json:"type_name"` // 分类名称
}
