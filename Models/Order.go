package Models

import "domino/domino/Config"

func CreateOrder(order *Order) (err error) {
	if err = Config.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderByAgen(orders *[]Order, agen_id string) (err error) {
	if err = Config.DB.Where("user_id = ?", agen_id).Order("id desc").Find(&orders).Error; err != nil {
		return err
	}
	return nil
}
