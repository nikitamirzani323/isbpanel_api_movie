package entities

type Model_pasaran struct {
	Pasaran_id            string `json:"pasaran_id"`
	Pasaran_name          string `json:"pasaran_name"`
	Pasaran_url           string `json:"pasaran_url"`
	Pasaran_diundi        string `json:"pasaran_diundi"`
	Pasaran_jamjadwal     string `json:"pasaran_jamjadwal"`
	Pasaran_datekeluaran  string `json:"pasaran_datekeluaran"`
	Pasaran_slug          string `json:"pasaran_slug"`
	Pasaran_meta_title    string `json:"pasaran_meta_title"`
	Pasaran_meta_descp    string `json:"pasaran_meta_descp"`
	Pasaran_keluaran      string `json:"pasaran_keluaran"`
	Pasaran_dateprediksi  string `json:"pasaran_dateprediksi"`
	Pasaran_bbfsprediksi  string `json:"pasaran_bbfsprediksi"`
	Pasaran_nomorprediksi string `json:"pasaran_nomorprediksi"`
}
type Model_pasaransimple struct {
	Pasaran_name string `json:"pasaran_name"`
	Pasaran_url  string `json:"pasaran_url"`
	Pasaran_slug string `json:"pasaran_slug"`
}
type Model_keluaran struct {
	Keluaran_datekeluaran string `json:"keluaran_datekeluaran"`
	Keluaran_periode      string `json:"keluaran_periode"`
	Keluaran_nomor        string `json:"keluaran_nomor"`
}
type Model_keluaranpaitominggu struct {
	Keluaran_nomor interface{} `json:"keluaran_nomor"`
}
type Model_keluaranpaitosenin struct {
	Keluaran_nomor interface{} `json:"keluaran_nomor"`
}
type Model_keluaranpaitoselasa struct {
	Keluaran_nomor interface{} `json:"keluaran_nomor"`
}
type Model_keluaranpaitorabu struct {
	Keluaran_nomor interface{} `json:"keluaran_nomor"`
}
type Model_keluaranpaitokamis struct {
	Keluaran_nomor interface{} `json:"keluaran_nomor"`
}
type Model_keluaranpaitojumat struct {
	Keluaran_nomor interface{} `json:"keluaran_nomor"`
}
type Model_keluaranpaitosabtu struct {
	Keluaran_nomor interface{} `json:"keluaran_nomor"`
}
type Controller_pasaran struct {
	Client_hostname string `json:"client_hostname" validate:"required"`
}
type Controller_keluaran struct {
	Pasaran_id string `json:"pasaran_id" validate:"required"`
}
