package entities

type Model_moviegenre struct {
	Movie_genre string `json:"movie_genre"`
	Movie_slug  string `json:"movie_slug"`
}
type Model_movie struct {
	Movie_slug      string `json:"movie_slug"`
	Movie_title     string `json:"movie_title"`
	Movie_thumbnail string `json:"movie_thumbnail"`
}
type Model_moviedetailwebsite struct {
	Movie_type           string      `json:"movie_type"`
	Movie_title          string      `json:"movie_title"`
	Movie_descp          string      `json:"movie_descp"`
	Movie_slug           string      `json:"movie_slug"`
	Movie_year           int         `json:"movie_year"`
	Movie_view           int         `json:"movie_view"`
	Movie_img            string      `json:"movie_img"`
	Movie_src            string      `json:"movie_src"`
	Movie_genre          interface{} `json:"movie_genre"`
	Movie_video          interface{} `json:"movie_video"`
	Movie_listvideogenre interface{} `json:"movie_listvideogenre"`
	Movie_listvideonew   interface{} `json:"movie_listvideonew"`
	Movie_listseason     interface{} `json:"movie_listseason"`
}
type Model_moviebanner struct {
	Moviebanner_urlimg         string `json:"moviebanner_urlimg"`
	Moviebanner_urldestination string `json:"moviebanner_urldestination"`
}
type Model_movievideo struct {
	Movie_title string `json:"movie_title"`
	Movie_src   string `json:"movie_src"`
}
type Model_movieseason struct {
	Season_id    int    `json:"season_id"`
	Season_title string `json:"season_title"`
}
type Model_movieepisode struct {
	Episode_id    int    `json:"episode_id"`
	Episode_title string `json:"episode_title"`
	Episode_src   string `json:"episode_src"`
}

// MOBILE
type Model_movielist struct {
	Movie_id    int    `json:"movie_id"`
	Movie_type  string `json:"movie_type"`
	Movie_title string `json:"movie_title"`
	Movie_label string `json:"movie_label"`
	Movie_descp string `json:"movie_descp"`
	Movie_year  int    `json:"movie_year"`
	Movie_view  int    `json:"movie_view"`
	Movie_img   string `json:"movie_img"`
}
type Model_moviedetail struct {
	Movie_id          int         `json:"movie_id"`
	Movie_type        string      `json:"movie_type"`
	Movie_title       string      `json:"movie_title"`
	Movie_label       string      `json:"movie_label"`
	Movie_descp       string      `json:"movie_descp"`
	Movie_year        int         `json:"movie_year"`
	Movie_view        int         `json:"movie_view"`
	Movie_img         string      `json:"movie_img"`
	Movie_genre       string      `json:"movie_genre"`
	Movie_src         string      `json:"movie_src"`
	Movie_favorite    string      `json:"movie_favorite"`
	Movie_totalsource int         `json:"movie_totalsource"`
	Movie_totalbanner int         `json:"movie_totalbanner"`
	Movie_video       interface{} `json:"movie_video"`
	Movie_banner      interface{} `json:"movie_banner"`
}
type Model_mobilemoviecategory struct {
	Movie_idcategory int         `json:"movie_idcategory"`
	Movie_category   string      `json:"movie_category"`
	Movie_list       interface{} `json:"movie_list"`
}
type Model_mobilemoviecomment struct {
	Movie_idcomment int    `json:"movie_idcomment"`
	Movie_name      string `json:"movie_name"`
	Movie_comment   string `json:"movie_comment"`
	Movie_create    string `json:"movie_create"`
}
type Model_mobileuser struct {
	User_username   string      `json:"user_username"`
	User_name       string      `json:"user_name"`
	User_coderef    string      `json:"user_coderef"`
	User_point      int         `json:"user_point"`
	Listclaim       interface{} `json:"listclaim"`
	Listclaimdetail interface{} `json:"listclaimdetail"`
}
type Model_mobilelistclaim struct {
	Claim_id    int    `json:"claim_id"`
	Claim_name  string `json:"claim_name"`
	Claim_point int    `json:"claim_point"`
}

type Controller_clientgenre struct {
	Client_hostname string `json:"client_hostname" validate:"required"`
}
type Controller_clientmovie struct {
	Client_hostname string `json:"client_hostname" validate:"required"`
	Movie_search    string `json:"movie_search"`
	Movie_tipe      string `json:"movie_tipe"`
	Movie_perpage   int    `json:"movie_perpage"`
	Movie_page      int    `json:"movie_page"`
}
type Controller_clientmoviesearch struct {
	Client_hostname string `json:"client_hostname" validate:"required"`
	Movie_search    string `json:"movie_search"`
}
type Controller_clientmoviegenre struct {
	Client_hostname string `json:"client_hostname" validate:"required"`
	Movie_page      int    `json:"movie_page"`
	Slug            string `json:"slug"`
}
type Controller_clientmoviedetail struct {
	Client_hostname string `json:"client_hostname" validate:"required"`
	Slug            string `json:"slug"`
}
type Controller_season struct {
	Client_hostname string `json:"client_hostname" validate:"required"`
	Movie_id        int    `json:"movie_id" validate:"required"`
}
type Controller_episode struct {
	Client_hostname string `json:"client_hostname" validate:"required"`
	Season_id       int    `json:"season_id" validate:"required"`
}

type Controller_clientmobilemovie struct {
	Client_type     string `form:"type" validate:"required"`
	Client_username string `form:"username" `
	Client_search   string `form:"search" `
}
type Controller_clientmobilegenremovie struct {
	Client_genre int `form:"genre" `
}
type Controller_clientmobiledetailmobile struct {
	Client_idmovie  int    `form:"idmovie" validate:"required"`
	Client_username string `form:"username" validate:"required"`
}
type Controller_mobileseason struct {
	Movie_id int `json:"movie_id" form:"movie_id" validate:"required"`
}
type Controller_mobileepisode struct {
	Season_id int `json:"season_id" form:"season_id" validate:"required"`
}
type Controller_clientmobilecomment struct {
	Movie_id int `json:"movie_id" form:"movie_id" validate:"required"`
}
type Controller_clientmobilesavecomment struct {
	Moviecomment_movieid  int    `json:"moviecomment_movieid" form:"moviecomment_movieid" validate:"required"`
	Moviecomment_username string `json:"moviecomment_username" form:"moviecomment_username" validate:"required"`
	Moviecomment_comment  string `json:"moviecomment_comment" form:"moviecomment_comment" validate:"required"`
}
type Controller_clientmobilesaverate struct {
	Movierate_movieid  int    `form:"movierate_movieid" validate:"required"`
	Movierate_rating   string `form:"movierate_rating" validate:"required"`
	Movierate_username string `form:"movierate_username" validate:"required"`
}
type Controller_clientmobilesavefavorite struct {
	Moviefavorite_movieid  int    `form:"moviefavorite_movieid" validate:"required"`
	Moviefavorite_username string `form:"moviefavorite_username" validate:"required"`
}
type Controller_clientmobilesavereport struct {
	Moviereport_movieid  int    `form:"moviereport_movieid" validate:"required"`
	Moviereport_username string `form:"moviereport_username" validate:"required"`
	Moviereport_info     string `form:"moviereport_info" validate:"required"`
}
type Controller_clientmobileuser struct {
	Client_username string `form:"username" validate:"required"`
}
type Controller_clientmobilesaveclaim struct {
	Claim_idclaim     int    `form:"claim_idclaim" validate:"required"`
	Claim_username    string `form:"claim_username" validate:"required"`
	Claim_point       int    `form:"claim_point" validate:"required"`
	Claim_pointbefore int    `form:"claim_pointbefore" validate:"required"`
}
