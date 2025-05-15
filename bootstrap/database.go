package bootstrap

import (
	"errors"
	"github.com/orangbus/m3d/app/models"
	"github.com/orangbus/m3d/pkg/config"
	"github.com/orangbus/m3d/pkg/database"
	"github.com/orangbus/m3d/pkg/debug"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

func SetupDatabase() {
	var dbConfig gorm.Dialector
	switch config.GetString("db.driver") {
	case "sqlite":
		db_path := getDatabasePath()
		dbConfig = sqlite.Open(db_path)
	case "mysql":
		dbConfig = mysql.Open(config.GetMysqlUrl())
		database.Connect(dbConfig)
	default:
		debug.Panic("数据库类型错误", errors.New("数据库类型错误"))
	}

	database.Connect(dbConfig)
	var article models.Articles
	var video models.Video
	database.DB.AutoMigrate(&article, &video)
}

func getDatabasePath() string {
	basePath, err := os.Getwd()
	if err != nil {
		return "database.sqlite"
	}
	return filepath.Join(basePath, "database.sqlite")
}
