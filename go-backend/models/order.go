package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	FirstName  string      `json:"-"`
	LastName   string      `json:"-"`
	Email      string      `json:"email"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderId"`
	Name       string      `json:"name" gorm:"-" `
	Total      int         `json:"total" gorm:"-" `
}

type OrderItem struct {
	Id           uint    `json:"id"`
	OrderId      uint    `json:"order_id"`
	ProductTitle string  `json:"product_title"`
	Price        float32 `json:"price"`
	Quantity     uint    `json:"quantity"`
}

func (order *Order) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Order{}).Count(&total)
	return total
}

func (order *Order) Take(db *gorm.DB, limit int, offset int) interface{} {
	var orders []Order
	db.Preload("OrderItems").Offset(offset).Limit(limit).Find(&orders)
	for i, _ := range orders {
		orders[i].Name = orders[i].FirstName + " " + orders[i].LastName
		var total float32 = 0
		for _, orderItem := range orders[i].OrderItems {
			total += float32(orderItem.Price) * float32(orderItem.Quantity)
		}
		orders[i].Total = int(total)
	}
	return orders
}
