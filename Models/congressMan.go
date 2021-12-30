package Models

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CongressMan struct {
	Id              int64  `json:"Id"`
	Uid             string `json:"Uid"`
	Civility        string `json:"Civility"`
	FirstName       string `json:"FirstName"`
	LastName        string `json:"LastName"`
	Alpha           string `json:"Alpha"`
	Trigramme       string `json:"Trigramme"`
	BirthDate       string `json:"BirthDate"`
	BirthCity       string `json:"BirthCity"`
	BirthDepartment string `json:"BirthDepartment"`
	BirthCountry    string `json:"BirthCountry"`
	Jobtitle        string `json:"Jobtitle"`
	CatSocPro       string `json:"CatSocPro"`
	FamSocPro       string `json:"FamSocPro"`
}

type CongressMans []CongressMan

func InitDB() (db *sql.DB) {
	db, err := sql.Open("mysql", "ProcessDeputesData:ASimpleP@ssW0rd@/PROCESSDEPUTES")

	//TODO mettre un système de log si il y a une erreur de connexion
	if err != nil {
		fmt.Println("Erreur connexion " + err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db
}

func AllCongressMans() *CongressMans {
	var congressMans CongressMans
	db := InitDB()

	//TODO implémenter l'erreur 500
	rows, err := db.Query("select * FROM PROCESSDEPUTES.CongressMan;")

	if err != nil {
		//TODO mettre un système de log fichier
		fmt.Println("Erreur requête " + err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var congressman CongressMan
		err := rows.Scan(&congressman.Id, &congressman.Uid, &congressman.Civility, &congressman.FirstName,
			&congressman.LastName, &congressman.Alpha, &congressman.Trigramme, &congressman.BirthDate,
			&congressman.BirthCity, &congressman.BirthDepartment, &congressman.BirthCountry,
			&congressman.Jobtitle, &congressman.CatSocPro, &congressman.FamSocPro)

		if err != nil {
			//TODO mettre un système de log fichier
			fmt.Println("Erreur récupération du résultat " + err.Error())
		}

		congressMans = append(congressMans, congressman)
	}

	return &congressMans
}

func GetCongressMan(id int) *CongressMan {
	var congressman CongressMan
	db := InitDB()

	//TODO implémenter l'erreur 500
	row, err := db.Query("select * FROM PROCESSDEPUTES.CongressMan where CongressManId=?;", id)

	if err != nil {
		//TODO mettre un système de log fichier
		fmt.Println("Erreur requête " + err.Error())
	}
	if row.Next() {
		errScan := row.Scan(&congressman.Id, &congressman.Uid, &congressman.Civility, &congressman.FirstName,
			&congressman.LastName, &congressman.Alpha, &congressman.Trigramme, &congressman.BirthDate,
			&congressman.BirthCity, &congressman.BirthDepartment, &congressman.BirthCountry,
			&congressman.Jobtitle, &congressman.CatSocPro, &congressman.FamSocPro)

		fmt.Println(congressman)

		if errScan != nil {
			//TODO mettre un système de log fichier
			fmt.Println("Erreur récupération du résultat " + errScan.Error())
		}

	}
	row.Close()

	return &congressman
}

//TODO voir si on met un type de retour
func InsertCongressMan(congressman *CongressMan) int64 {
	db := InitDB()
	var lid int64

	if congressman == nil {
		fmt.Println("No Data send to insert")
	} else {
		queryCongressMan := "INSERT INTO PROCESSDEPUTES.Congressman(CongressManUid, Civility, FirstName, LastName, Alpha, Trigramme, BirthDate, BirthCity, BirthDepartment, BirthCountry, JobTitle, CatSocPro, FamSocPro) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"
		stmt, errPrepare := db.Prepare(queryCongressMan)
		if errPrepare != nil {
			fmt.Println("Erreur récupération du résultat " + errPrepare.Error())
		} else {
			res, errExec := stmt.Exec(congressman.Uid, congressman.Civility, congressman.FirstName,
				congressman.LastName, congressman.Alpha, congressman.Trigramme, congressman.BirthDate,
				congressman.BirthCity, congressman.BirthDepartment, congressman.BirthCountry,
				congressman.Jobtitle, congressman.CatSocPro, congressman.FamSocPro)
			if errExec != nil {
				fmt.Println("Congressman Repository : Erreur exécution requête " + errExec.Error())
			} else {
				var errGetLastId error
				lid, errGetLastId = res.LastInsertId()
				if errGetLastId != nil {
					fmt.Println("Erreur lors de la récupération de l'id enregistré " + errGetLastId.Error())
				}
			}
		}
	}
	defer db.Close()
	
	return lid
}