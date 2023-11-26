package rpc

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ClientsManager interface {
	GetConf() *Config
	GetDB() *gorm.DB
}

type ClientsHandler struct {
	ClientConf *Config
	DB         *gorm.DB
}

func (h ClientsHandler) GetConf() *Config {
	if h.ClientConf == nil {
		h.ClientConf = MustInitConfig()
	}
	return h.ClientConf
}

func (h ClientsHandler) GetDB() *gorm.DB {
	if h.DB == nil {
		var db *gorm.DB
		var err error
		switch h.ClientConf.DB.Type {
		case "mysql":
			db, err = gorm.Open(mysql.Open(h.ClientConf.DB.DSN), &gorm.Config{})
			if err != nil {
				logrus.Panic("can't init db, type: %v", h.ClientConf.DB.Type)
			}
		default:
			logrus.Panic("can't find db tpye: %v", h.ClientConf.DB.Type)
		}
		h.DB = db
	}
	return h.DB
}

var (
	cliManager = &ClientsHandler{}
)

func GetCliMgr() ClientsManager {
	if cliManager.ClientConf == nil {
		cliManager.ClientConf = MustInitConfig()
	}
	return cliManager
}
