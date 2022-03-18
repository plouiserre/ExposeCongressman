package Repository

import (
	_ "github.com/go-sql-driver/mysql"
	manager "github.com/plouiserre/exposecongressman/Manager"
	models "github.com/plouiserre/exposecongressman/Models"
)

type CongressmanRepository struct {
	LogManager *manager.LogManager
}

func (cr CongressmanRepository) GetCongressmansMandates(congressmanId int) *models.MandatesModel {
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

func (cr CongressmanRepository) GetCongressmansDepartments(departmentId int) models.CongressmanDepartmentsModel {
	var rb RepositoryBase
	var congressmansDepartment models.CongressmanDepartmentsModel

	db := rb.InitDB()

	rows, err := db.Query("select cr.CongressmanId, cr.CongressmanUid, cr.Civility, cr.FirstName, cr.LastName, cr.Alpha, cr.BirthDate, el.Region, el.Department, el.Departmentnum FROM PROCESSDEPUTES.CongressMan cr INNER JOIN PROCESSDEPUTES.Mandate mt ON mt.CongressmanId = cr.CongressmanId INNER JOIN PROCESSDEPUTES.Election el ON mt.MandateId = el.MandateId AND el.DepartmentNum = ?;", departmentId)
	if err != nil {
		cr.LogManager.WriteErrorLog("Erreur pour la requête des députés d'un département " + err.Error())
	} else {
		for rows.Next() {
			var congressmanDepartment models.CongressmanDepartmentModel
			err := rows.Scan(&congressmanDepartment.Id, &congressmanDepartment.Uid, &congressmanDepartment.Civility,
				&congressmanDepartment.FirstName, &congressmanDepartment.LastName, &congressmanDepartment.Alpha,
				&congressmanDepartment.BirthDate, &congressmanDepartment.Region, &congressmanDepartment.Department,
				&congressmanDepartment.Departmentnum)

			if err != nil {
				cr.LogManager.WriteErrorLog("Erreur récupération du résultat " + err.Error())
			}

			congressmansDepartment = append(congressmansDepartment, congressmanDepartment)
		}
	}
	return congressmansDepartment
}

func (cr CongressmanRepository) GetCongressmansJobs(jobs string) *models.CongressmansModel {
	var rb RepositoryBase
	var congressmans models.CongressmansModel

	db := rb.InitDB()

	rows, err := db.Query("select * FROM PROCESSDEPUTES.Congressman where JobTitle=?;", jobs)
	if err != nil {
		cr.LogManager.WriteErrorLog("Erreur pour la requête des mandats d'un député " + err.Error())
	} else {
		result, _ := congressmans.RowsScanGetEntities(rows, cr.LogManager)
		congressmans = result.(models.CongressmansModel)
	}
	return &congressmans
}
