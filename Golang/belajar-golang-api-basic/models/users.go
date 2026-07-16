package models

type User struct {
	ID 		 	int		`json:"id"` 
	Nama	 	string	`json:"nama"`
	PosisiID 	int		`json:"posisi_id"`
	Perusahaan	string 	`json:"perusahaan"`
}

