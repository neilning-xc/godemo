package service

import "godemo/model"

func ListUser() ([]*model.User, error) {
	users := make([]*model.User, 0)
	if err := model.Db.Where(&model.User{}).Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}
