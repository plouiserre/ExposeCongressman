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
	IModel         models.IModel
	RepositoryBase repository.RepositoryBase
}

//TODO à supprimer après
func (entityService *EntityService) InitLogManager() Manager.LogManager {
	logManager := Manager.LogManager{}
	logManager.InitLog()

	return logManager
}

func (entityService EntityService) GetAll() (*models.EntityModel, bool) {
	entities, noError := entityService.RepositoryBase.GetAll(entityService.Models)

	return entities, noError
}

func (entityService EntityService) GetById(id int) (*models.EntityModel, bool) {
	entity, noError := entityService.RepositoryBase.GetById(entityService.IModel, id)

	return entity, noError
}

func (entityService EntityService) CreateEntity(entity *models.EntityModel) (int64, bool) {
	lid, noError := entityService.RepositoryBase.CreateEntity(entityService.IModel, entity)

	return lid, noError
}

func (entityService EntityService) UpdateEntity(entity *models.EntityModel, id int64) (int64, bool) {
	id, noError := entityService.RepositoryBase.UpdateEntity(entityService.IModel, entity, id)

	return id, noError
}

func (entityService EntityService) DeleteEntity(id int) (int64, bool) {
	nbDelete, noError := entityService.RepositoryBase.DeleteEntity(entityService.IModel, id)

	return nbDelete, noError
}
