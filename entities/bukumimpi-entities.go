package entities

type Model_bukumimpi struct {
	Bukumimpi_type  string `json:"bukumimpi_type"`
	Bukumimpi_name  string `json:"bukumimpi_name"`
	Bukumimpi_nomor string `json:"bukumimpi_nomor"`
}
type Model_tafsirmimpi struct {
	Tafsirmimpi_mimpi     string `json:"tafsirmimpi_mimpi"`
	Tafsirmimpi_artimimpi string `json:"tafsirmimpi_artimimpi"`
	Tafsirmimpi_angka2d   string `json:"tafsirmimpi_angka2d"`
	Tafsirmimpi_angka3d   string `json:"tafsirmimpi_angka3d"`
	Tafsirmimpi_angka4d   string `json:"tafsirmimpi_angka4d"`
}

type Controller_clienrequest struct {
	Client_Device string `json:"client_device"`
	Tipe          string `json:"tipe"`
	Nama          string `json:"nama"`
}
type Controller_clientafsirmimpirequest struct {
	Search string `json:"search"`
}
