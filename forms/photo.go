package forms

type PhotoCommand struct {
	Nama   string `json:"nama" binding:"required"`
	Gambar string `json:"gambar" binding:"required"`
}
