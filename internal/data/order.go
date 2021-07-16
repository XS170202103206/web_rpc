package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"web_rpc/internal/biz"
	"web_rpc/internal/data/model"
)
//实现biz中的接口
type orderRepo struct {
	//data *Data
	db *gorm.DB
	log  *log.Helper
}
// NewOrderRepo .
func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		db: data.db,
		log:  log.NewHelper(logger),
	}
}

func (o *orderRepo) CreateOrder(_ context.Context, order *biz.Order) error {
	var m model.OrderModel
	m.OrderNo = order.OrderNo
	m.Amount = order.Amount
	m.Status = order.Status
	m.FileUrl = order.FileUrl
	m.UserName = order.UserName
	if err := o.db.Create(&m).Error; err != nil {
		return err
	}
	return nil
	//panic("implement me")
}

func (o *orderRepo) UpdateOrder(_ context.Context, order *biz.Order) error {
	var m model.OrderModel
	//m.ID = order.ID
    m.UserName = order.UserName
    m.OrderNo = order.OrderNo
	m.Amount = order.Amount
	m.Status = order.Status
	m.FileUrl = order.FileUrl

	//result := o.db.Model(&m).Update(&m)
	result := o.db.Where("id = ?", order.ID).Updates(&m)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return  errors.New("update failed")
	}
	return nil
	//panic("implement me")
}

func (o *orderRepo) GetOrder(_ context.Context, id uint) (*biz.Order, error) {
	var m model.OrderModel
	if err := o.db.Where("id = ?", id).First(&m).Error; err != nil {
		return nil, err
	}
	return toBiz(m), nil
	//panic("implement me")
}

func toBiz(m model.OrderModel) *biz.Order {
	return &biz.Order{
		ID :      m.ID,
		OrderNo:  m.OrderNo,
		UserName: m.UserName,
		Amount:   m.Amount,
		Status:   m.Status,
		FileUrl:  m.FileUrl,
	}
}

func (o *orderRepo) ListOrder(_ context.Context, query, sequence, by string) ([]*biz.Order, error) {
	var ms []*model.OrderModel
	err := o.db.Where("user_name LIKE ?", "%"+query+"%").
		Order(sequence + " " + by).
	    Find(&ms).Error
    //err := o.db.Model(&ms).Error
	if err != nil {
		return nil, err
	}
	return toBizSlice(ms), nil
	//panic("implement me")
}

func (o *orderRepo) GetListOrder(_ context.Context) ([]*model.OrderModel, error) {
	var m model.OrderModel
	var result  []*model.OrderModel
	o.db.Model(m).Find(&result)
	//return toBizMap(result), nil
	return result, nil
}

func toBizSlice(ms []*model.OrderModel) []*biz.Order {
	s := make([]*biz.Order, len(ms))
	for i, v := range ms {
		s[i] = toBiz(*v)
	}
	return s
}

//func toBizMap(ms []*model.OrderModel) []*biz.Order {
//	s := make([]*biz.Order, len(ms))
//	for i, v := range ms {
//		s[i] = toBiz(*v)
//	}
//	return s
//}
