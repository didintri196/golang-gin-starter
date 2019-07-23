package models

import (
	"golang3/db"
	"golang3/forms"

	"gopkg.in/mgo.v2/bson"
)

type Mahasiswa struct {
	ID   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Nim  string        `json:"nim" bson:"nim"`
	Nama string        `json:"nama" bson:"nama"`
}

type MahasiswaModel struct{}

//CreateModel
func (m *MahasiswaModel) Create(data forms.MahasiswaCommand) (err error) {
	db := db.DBConnection{}
	db.NewConnection()
	defer db.Close()
	collection := db.Session.DB("MahasiswaService").C("data_mahasiswa")
	err = collection.Insert(bson.M{"nim": data.Nim, "nama": data.Nama})
	return err
}

//Menampilkan Seluruh Data
func (m *MahasiswaModel) Find() (mhs []Mahasiswa, err error) {
	db := db.DBConnection{}
	db.NewConnection()
	defer db.Close()
	collection := db.Session.DB("MahasiswaService").C("data_mahasiswa")
	err = collection.Find(bson.M{}).All(&mhs)
	return mhs, err
}

//Menampilkan By ID
func (m *MahasiswaModel) FindId(id string) (mhs Mahasiswa, err error) {
	db := db.DBConnection{}
	db.NewConnection()
	defer db.Close()
	collection := db.Session.DB("MahasiswaService").C("data_mahasiswa")
	err = collection.FindId(bson.ObjectIdHex(id)).One(&mhs)
	return mhs, err
}

//Menampilkan By NIM
func (m *MahasiswaModel) FindNim(nim string) (mhs []Mahasiswa, err error) {
	db := db.DBConnection{}
	db.NewConnection()
	defer db.Close()
	collection := db.Session.DB("MahasiswaService").C("data_mahasiswa")
	err = collection.Find(bson.M{"nim": nim}).All(&mhs)
	return mhs, err
}
