package entities

type Mahasiswa struct {
	Id int64
	NamaLengkap string `validate:"required" label:"Nama Lengkap"`
	JenisKelamin string `validate:"required" label:"Jenis Kelamin"`
	TempatLahir string `validate:"required" label:"Tempat Lahir"`
	TanggalLahir string `validate:"required" label:"Tanggal Lahir"`
	Alamat string `validate:"required"`
}