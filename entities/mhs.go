package entities

//menyimpan sesuai tabel database
type Mahasiswa struct {
	Pert  string
	No    int64
	Nama  string
	NPM   string
	Kelas string
	UTS   float64
	UAS   float64
	Total float64
}