package model

import (
	"personal/TODO-list/database"
	"time"
)

type TknWeb struct {
	Id     int
	UserId int
	Email  string
	WatId  string
	WatIat time.Time
	WatExp time.Time
}

func VerifyAccessToken(token string) interface{} {
	var dbtoken TknWeb
	res := database.Db.
		Where("wat_id = ?", token).
		Take(&dbtoken)
	if res.Error == nil {
		return dbtoken
	}
	return nil
}

func GetWebTokenByUsername(userId int) (tknWeb TknWeb, err error) {
	err = database.Db.
		Where("user_id = ?", userId).
		First(&tknWeb).Error
	if err != nil {
		return TknWeb{}, err
	}
	return tknWeb, nil
}

func CreateToken(token TknWeb) error {
	return database.Db.Create(&token).Error
}

func DeleteTokenByUsername(userId int) error {
	var tknWeb TknWeb
	err := database.Db.
		Where("user_id = ?", userId).
		Delete(&tknWeb).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteTokenByWatId(watId string) error {
	err := database.Db.
		Where("wat_id = ?", watId).
		Delete(&TknWeb{}).Error
	if err != nil {
		return err
	}

	return nil
}
