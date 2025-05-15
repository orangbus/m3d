package bootstrap

import (
	"github.com/orangbus/m3d/pkg/config"
)

func SetUp() {
	config.LoadConfig()
	SetupDatabase()
	//search.NewSearch()
}
