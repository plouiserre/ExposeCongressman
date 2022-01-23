package Service

import (
	"github.com/plouiserre/exposecongressman/Manager"
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
