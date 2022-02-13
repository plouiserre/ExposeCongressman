package Service

import (
	"github.com/plouiserre/exposecongressman/Manager"
	models "github.com/plouiserre/exposecongressman/Models"
	repository "github.com/plouiserre/exposecongressman/Repository"
)

type Entity int

const (
	Congressman Entity = iota
	Mandate
	Deputy
)

type EntityService struct {
	Entity
}

func (entityService *EntityService) InitRepository(entityType int) (repository.IRepository, Manager.LogManager) {
	logManager := Manager.LogManager{}
	logManager.InitLog()

	entityService.Entity = Entity(entityType)

	if entityService.Entity == Congressman {
		congressmanRepository := repository.CongressmanRepository{
			LogManager: &logManager,
		}
		return congressmanRepository, logManager
	} else if entityService.Entity == Mandate {
		mandateRepository := repository.MandateRepository{
			LogManager: &logManager,
		}
		return mandateRepository, logManager
	} else {
		deputyRepository := repository.DeputyRepository{
			LogManager: &logManager,
		}
		return deputyRepository, logManager
	}
}

func (entityService EntityService) GetAll(repo repository.IRepository) (*models.EntityModel, bool) {
	entities, noError := repo.GetAll()

	return entities, noError
}

func (entityService EntityService) GetById(id int, repo repository.IRepository) (*models.EntityModel, bool) {
	entity, noError := repo.GetById(id)

	return entity, noError
}

func (entityService EntityService) CreateEntity(repo repository.IRepository, entity *models.EntityModel) (int64, bool) {
	lid, noError := repo.CreateEntity(entity)

	return lid, noError
}

func (entityService EntityService) UpdateEntity(repo repository.IRepository, entity *models.EntityModel, id int) bool {
	noError := repo.UpdateEntity(entity, id)

	return noError
}
