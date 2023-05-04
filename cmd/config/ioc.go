package config

import (
	"service-user/internal/domain/repository"
	"service-user/internal/domain/usecase"
	"service-user/internal/infra/storage"
)

func genIoc(config *Configuration) map[string]interface{} {
	ioc := make(map[string]interface{})
	//database
	db := GetDB()
	repoAct := repository.NewRepositoryActivities(db)
	repoCat := repository.NewRepositoryCategories(db)
	repoCont := repository.NewRepositoryContainerLevels(db)
	repoGuide := repository.NewRepositoryGuides(db)
	repoLevel := repository.NewRepositoryLevels(db)
	repoObjective := repository.NewRepositoryObjectives(db)
	repoQuestion := repository.NewRepositoryQuestion(db)
	repoSection := repository.NewRepositorySection(db)

	clientStorage := storage.InitStorage(GetStorageClient(), config.Credential.Gcbucket)

	//ioc
	ioc["activities"] = usecase.NewActivitiesUseCase(repoAct, clientStorage)
	ioc["categories"] = usecase.NewCateogriesUseCase(repoCat, clientStorage)
	ioc["container-level"] = usecase.NewContainerLevelsUseCase(repoCont, clientStorage)
	ioc["guides"] = usecase.NewGuidesUseCase(repoGuide, clientStorage)
	ioc["levels"] = usecase.NewLevelsUseCase(repoLevel, clientStorage)
	ioc["objectives"] = usecase.NewObjectivesUseCase(repoObjective)
	ioc["question"] = usecase.NewQuestionUseCase(repoQuestion, clientStorage)
	ioc["section"] = usecase.NewSectionUseCase(repoSection, clientStorage)

	return ioc
}
