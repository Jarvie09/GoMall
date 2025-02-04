package model

type Goods struct {
	//gorm.Model
	goodsId    int     `gorm:"uniqueIndex"`
	goodsName  string  `gorm:"type:varchar(255)"`
	goodsPrice float64 `gorm:"type:decimal(10,2)"`
	goodsNum   int     `gorm:"type:int(11)"`
	goodsUrl   string  `gorm:"type:varchar(255)"`
}
