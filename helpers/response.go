package helpers

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
	Time    string      `json:"time"`
}
type ResponsePaging struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Record      interface{} `json:"record"`
	Perpage     int         `json:"perpage"`
	Totalrecord int         `json:"totalrecord"`
	Time        string      `json:"time"`
}
type ResponseMovieGenre struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Record      interface{} `json:"record"`
	Perpage     int         `json:"perpage"`
	Totalrecord int         `json:"totalrecord"`
	Genre       string      `json:"genre"`
	Time        string      `json:"time"`
}
type ResponseKeluaran struct {
	Status        int         `json:"status"`
	Message       string      `json:"message"`
	Pasaran       string      `json:"pasaran_nama"`
	Livedraw      string      `json:"pasaran_livedraw"`
	Pasarandiundi string      `json:"pasaran_diundi"`
	Pasaranjadwal string      `json:"pasaran_jadwal"`
	Pasaran_title string      `json:"pasaran_title"`
	Pasaran_descp string      `json:"pasaran_descp"`
	Record        interface{} `json:"record"`
	Paito_minggu  interface{} `json:"paito_minggu"`
	Paito_senin   interface{} `json:"paito_senin"`
	Paito_selasa  interface{} `json:"paito_selasa"`
	Paito_rabu    interface{} `json:"paito_rabu"`
	Paito_kamis   interface{} `json:"paito_kamis"`
	Paito_jumat   interface{} `json:"paito_jumat"`
	Paito_sabtu   interface{} `json:"paito_sabtu"`
	List_pasaran  interface{} `json:"list_pasaran"`
	Time          string      `json:"time"`
}
type Responseproviderslot struct {
	Status             int    `json:"status"`
	Message            string `json:"message"`
	Providerslot_name  string `json:"providerslot_name"`
	Providerslot_image string `json:"providerslot_image"`
	Providerslot_title string `json:"providerslot_title"`
	Providerslot_descp string `json:"providerslot_descp"`
}
type Responsemobilemovie struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Slider  interface{} `json:"slider"`
	Genre   interface{} `json:"genre"`
	Time    string      `json:"time"`
}

type ResponseAdmin struct {
	Status   int         `json:"status"`
	Message  string      `json:"message"`
	Record   interface{} `json:"record"`
	Listrule interface{} `json:"listruleadmin"`
	Time     string      `json:"time"`
}
type ErrorResponse struct {
	Field string
	Tag   string
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}
