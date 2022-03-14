package Repository

import (
	_ "github.com/go-sql-driver/mysql"
	manager "github.com/plouiserre/exposecongressman/Manager"
	models "github.com/plouiserre/exposecongressman/Models"
)

type CongressmanRepository struct {
	LogManager *manager.LogManager
}

func (cr CongressmanRepository) GetCongressmansMandates(congressmanId int64) *models.MandatesModel {
	var rb RepositoryBase
	var mandates models.MandatesModel

	db := rb.InitDB()

	rows, err := db.Query("select * FROM PROCESSDEPUTES.Mandate where CongressmanId=?;", congressmanId)
	if err != nil {
		cr.LogManager.WriteErrorLog("Erreur pour la requête des mandats d'un député " + err.Error())
	} else {
		result, _ := mandates.RowsScanGetEntities(rows, cr.LogManager)
		mandates = result.(models.MandatesModel)
	}
	return &mandates
}
