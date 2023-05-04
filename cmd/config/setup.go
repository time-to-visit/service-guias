package config

import (
	"flag"
	"service-user/cmd/handler"
	"service-user/internal/domain/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	var configPath = ""
	configPath = *flag.String("config", "", "")

	if configPath == "" {
		configPath = "./data/config.yml"
	}

	setConfiguration(configPath)
}

func setConfiguration(configPath string) {
	Setup(configPath)
}

func Run(configPath string) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	conf := GetConfig()
	setupDB(conf)

	//ioc
	ioc := genIoc(conf)
	e = handler.NewHandlerActivities(e, ioc["activities"].(usecase.ActivitiesUseCase), AuthVerify)
	e = handler.NewHandlerCategories(e, ioc["categories"].(usecase.CategoriesUseCase), AuthVerify)
	e = handler.NewHandlerContainerLevel(e, ioc["container-level"].(usecase.ContainerLevelsUseCase), AuthVerify)
	e = handler.NewHandlerGuide(e, ioc["guides"].(usecase.GuidesUseCase), AuthVerify)
	e = handler.NewHandlerLevel(e, ioc["levels"].(usecase.LevelsUseCase), AuthVerify)
	e = handler.NewHandlerObjetive(e, ioc["objectives"].(usecase.ObjectivesUseCase), AuthVerify)
	e = handler.NewHandlerQuestion(e, ioc["question"].(usecase.QuestionUseCase), AuthVerify)
	e = handler.NewHandlerSection(e, ioc["section"].(usecase.SectionUseCase), AuthVerify)

	e.Logger.Fatal(e.Start(":" + conf.Server.Port))
}
