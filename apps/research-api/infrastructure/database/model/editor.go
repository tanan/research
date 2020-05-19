package model

type Editor struct {
	EditorId int    `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Icon     []byte `gorm:"column:icon"`
}

func (a Editor) TableName() string {
	return "research.editor"
}
