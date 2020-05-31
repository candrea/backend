package model

import (
	//"gopkg.in/mgo.v2/bson"
)

// Usuario - Model
type Usuario struct {
//	ID         bson.ObjectId `bson:"_id" json:"id"`
	Name      string  `bson:"nombre"   json:"nombre"`
	Email     string `bson:"email" json:"email"`
	Clave     string  `bson:"clave" json:"clave"`
	Rol     string  `bson:"rol" json:"rol"`
	Telefono string `bson:"telefono" bson:"telefono"`
}

type Reciclaje struct {
	//	ID         bson.ObjectId `bson:"_id" json:"id"`
		Tipo      string  `json:"tipo"`
		User_id     string  `json:"user_id"`
		Latitud    string  `json:"latitud"`
		Longitud     string  `json:"longitud"`
		Fecha string `json:"fecha"`
	}

//  - for request
type UsuarioID struct {
//	ID         bson.ObjectId `bson:"_id" json:"id"`

}
