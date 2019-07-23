package models

import (
	"golang3/db"
	"golang3/forms"

	"gopkg.in/mgo.v2/bson"
)

type PhotoModel struct {
	ID     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Nama   string        `json:"nama" bson:"nama"`
	Gambar string        `json:"gambar" bson:"gambar"`
}

//CreateModel
func (m *PhotoModel) Create(data forms.PhotoCommand) (err error) {
	db := db.DBConnection{}
	db.NewConnection()
	defer db.Close()
	collection := db.Session.DB("MahasiswaService").C("data_photo")
	err = collection.Insert(bson.M{"nama": data.Nama, "gambar": data.Gambar})
	return err
}
