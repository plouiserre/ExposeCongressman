package Models

type CongressmanDepartmentModel struct {
	Id            int64
	Uid           string
	Civility      string
	FirstName     string
	LastName      string
	Alpha         string
	BirthDate     string
	Region        string
	Department    string
	Departmentnum int64
}

type CongressmanDepartmentsModel []CongressmanDepartmentModel
