package Models

//TODO trouver une alternative plus propre quand tout sera factoriser où plus tard dans le projet
type EntityModel struct {
	Mandates MandatesModel
	Mandate  MandateModel
	Deputies DeputiesModel
}
