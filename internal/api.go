package api

import (
	"chat/internal/rest"
	"chat/logger"
	"chat/pkg/config"
	"chat/pkg/persistence"
	"chat/repository"
	"chat/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

func setConfiguration(configPath string) {
	config.Setup(configPath)
	gin.SetMode(config.GetConfig().Server.Mode)
}

func setupDbAndServices() *services.Services {
	conf := config.GetConfig()
	mongoClient, err := persistence.NewClient(conf.Database.Url, conf.Database.Username, conf.Database.Password)
	if err != nil {
		logger.Error(err)

	}

	fmt.Println(mongoClient)
	db := mongoClient.Database(conf.Database.DbName)
	repos := repository.NewRepositories(db)

	return services.NewServices(services.Deps{
		Repos: repos,
	})
}
func Run(configPath string) {
	if configPath == "" {
		configPath = "data/config.yml"
	}
	setConfiguration(configPath)
	setupServices := setupDbAndServices()
	conf := config.GetConfig()
	handler := rest.NewHandler(setupServices)
	web := handler.Setup()
	fmt.Println("Go API REST Running on port " + conf.Server.Port)
	fmt.Println("==================>")
	_ = web.Run(":" + conf.Server.Port)
}
