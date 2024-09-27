package database

import (
	"fiber-api/app/service"
	"fiber-api/internal/encrypt"
	"fiber-api/pkg/db"
	"io/ioutil"
	"os"
	"strings"
)

var lock = "app.lock"

// 初始化
func Init() {

	if _, err := os.Stat(lock); err == nil {
		return
	}

	sqlcontent, err := ioutil.ReadFile("database/init.sql")
	if err != nil {
		panic(err)
	}

	sqls := strings.Split(string(sqlcontent), ";")
	for _, sql := range sqls {
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}
		if err := db.GormClient.Exec(sql).Error; err != nil {
			panic(err)
		}
	}

	file, _ := os.Create(lock)
	defer file.Close()

	// 存在超级管理员不执行生成初始超级管理员
	admin := (&service.Admin{}).DetailByUsername("admin")
	if admin.Id > 0 {
		return
	}
	// 生成初始超级管理员
	(&service.Admin{}).Create(&service.Admin{
		Username: "admin",
		Nickname: "超级管理员",
		Gender:   1,
		Password: encrypt.Generate("123456"),
	})
}
