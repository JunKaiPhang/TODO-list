package model

import (
	"personal/TODO-list/database"
	"time"
)

type ApiLog struct {
	Id         int
	ApiUrlPath string
	ApiInput   string
	ApiOutput  string
	IpAddr     string
	CreatedAt  time.Time
}

func LogAPI(apiLog ApiLog) (err error) {
	err = database.Db.Create(&apiLog).Error
	if err != nil {
		return err
	}

	return nil
}
