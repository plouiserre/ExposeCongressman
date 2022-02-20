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
	Models models.IModels
	//TODO trouver un autre nom
	IGetByIdEntity models.IGetByIdEntity
	ICreateEntity  models.ICreateEntity
	IUpdateEntity  models.IUpdateEntity
	RepositoryBase repository.RepositoryBase
}

//TODO à supprimer après
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

func (entityService EntityService) GetAll() (*models.EntityModel, bool) {
	entities, noError := entityService.RepositoryBase.GetAll(entityService.Models)

	return entities, noError
}

func (entityService EntityService) GetById(id int) (*models.EntityModel, bool) {
	entity, noError := entityService.RepositoryBase.GetById(entityService.IGetByIdEntity, id)

	return entity, noError
}

func (entityService EntityService) CreateEntity(entity *models.EntityModel) (int64, bool) {
	lid, noError := entityService.RepositoryBase.CreateEntity(entityService.ICreateEntity, entity)

	return lid, noError
}

func (entityService EntityService) UpdateEntity(entity *models.EntityModel, id int) bool {
	noError := entityService.RepositoryBase.UpdateEntity(entityService.IUpdateEntity, entity, id)

	return noError
}

func (entityService EntityService) DeleteEntity(repo repository.IRepository, id int) (int64, bool) {
	nbDelete, noError := repo.DeleteEntity(id)

	return nbDelete, noError
}
