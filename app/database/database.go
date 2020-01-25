package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/windstep/go-web-skeleton/config"
)

func Connect() (*gorm.DB, error) {
	cfg := config.GetInstance()
	return gorm.Open(cfg.DBDriver, cfg.DBUrl)
}
