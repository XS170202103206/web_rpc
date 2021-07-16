package model

import "gorm.io/gorm"

type OrderModel struct {
	gorm.Model
	OrderNo string `json:"order_no" gorm:"column:order_no" comment:""`
	UserName string `json:"username" gorm:"column:username" comment:""`
	Amount float64 `json:"amount" gorm:"column:amount" comment:""`
	Status string `json:"status" gorm:"column:status" comment:""`
	FileUrl string `json:"file_url" gorm:"column:file_url" comment:""`
}