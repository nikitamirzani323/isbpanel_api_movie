package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/config"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/db"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/entities"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/nleeper/goment"
)

func Fetch_genre() (helpers.Response, error) {
	var obj entities.Model_moviegenre
	var arraobj []entities.Model_moviegenre
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
		nmgenre, slug 
		FROM ` + config.DB_tbl_mst_movie_genre + ` 
		ORDER BY genredisplay ASC   
	`

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)

	for row.Next() {
		var (
			nmgenre_db, slug_db string
		)

		err := row.Scan(&nmgenre_db, &slug_db)
		helpers.ErrorCheck(err)

		obj.Movie_genre = nmgenre_db
		obj.Movie_slug = slug_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_movieHome(search, tipe string, page int) (helpers.ResponsePaging, error) {
	var obj entities.Model_movie
	var arraobj []entities.Model_movie
	var res helpers.ResponsePaging
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	perpage := 40
	totalrecord := 200
	// offset := page

	sql_select := ""
	sql_select += "SELECT "
	sql_select += "movietitle , COALESCE(posted_id,0) , urlthumbnail, slug "
	sql_select += "FROM " + config.DB_tbl_trx_movie + " "
	sql_select += "WHERE enabled = 1 "
	switch tipe {
	case "NEW":
		sql_select += "ORDER BY createdatemovie DESC LIMIT " + strconv.Itoa(perpage)
	case "UPDATE":
		sql_select += "ORDER BY updatedatemovie DESC LIMIT " + strconv.Itoa(perpage)
	case "RANDOM":
		sql_select += "ORDER BY random() LIMIT 500 "
	default:
		sql_select += "ORDER BY random() DESC LIMIT " + strconv.Itoa(perpage)
	}

	fmt.Println(sql_select)
	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)

	for row.Next() {
		var (
			posted_id_db                            int
			movietitle_db, urlthumbnail_db, slug_db string
		)

		err := row.Scan(&movietitle_db, &posted_id_db, &urlthumbnail_db, &slug_db)
		helpers.ErrorCheck(err)
		path_image := ""
		if urlthumbnail_db == "" {
			poster_image, poster_extension := _GetMedia(posted_id_db)
			path_image = "https://duniafilm.b-cdn.net/uploads/cache/poster_thumb/uploads/" + poster_extension + "/" + poster_image
		} else {
			path_image = urlthumbnail_db
		}

		// movie_url, _, _ := _GetVideo(movieid_db, "")

		obj.Movie_title = movietitle_db
		obj.Movie_thumbnail = path_image
		obj.Movie_slug = slug_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Totalrecord = totalrecord
	res.Perpage = perpage
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_movieByGenre(genre string, page int) (helpers.ResponseMovieGenre, error) {
	var obj entities.Model_movie
	var arraobj []entities.Model_movie
	var res helpers.ResponseMovieGenre
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	idgenre, nmgenre := _GetIdGenre(genre)
	perpage := 40
	totalrecord := 200
	offset := page

	sql_select := ""
	sql_select += "SELECT "
	sql_select += "B.movietitle , COALESCE(B.posted_id,0) , B.urlthumbnail, B.slug "
	sql_select += "FROM " + config.DB_tbl_trx_moviegenre + " as A "
	sql_select += "JOIN " + config.DB_tbl_trx_movie + " as B ON A.movieid = B.movieid "
	sql_select += "WHERE B.enabled = 1 "
	sql_select += "AND A.idgenre = " + strconv.Itoa(idgenre) + " "
	sql_select += "ORDER BY B.createdatemovie DESC OFFSET " + strconv.Itoa(offset) + "LIMIT " + strconv.Itoa(perpage)

	fmt.Println(sql_select)
	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)

	for row.Next() {
		var (
			posted_id_db                            int
			movietitle_db, urlthumbnail_db, slug_db string
		)

		err := row.Scan(&movietitle_db, &posted_id_db, &urlthumbnail_db, &slug_db)
		helpers.ErrorCheck(err)
		path_image := ""
		if urlthumbnail_db == "" {
			poster_image, poster_extension := _GetMedia(posted_id_db)
			path_image = "https://duniafilm.b-cdn.net/uploads/cache/poster_thumb/uploads/" + poster_extension + "/" + poster_image
		} else {
			path_image = urlthumbnail_db
		}

		// movie_url, _, _ := _GetVideo(movieid_db, "")

		obj.Movie_title = movietitle_db
		obj.Movie_thumbnail = path_image
		obj.Movie_slug = slug_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Totalrecord = totalrecord
	res.Perpage = perpage
	res.Genre = nmgenre
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_movieDetail(slug string) (helpers.Response, error) {
	var obj entities.Model_moviedetailwebsite
	var arraobj []entities.Model_moviedetailwebsite
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
		movieid, movietitle, movietype, COALESCE(posted_id,0) , urlthumbnail, slug, description, 
		views, year  
		FROM ` + config.DB_tbl_trx_movie + ` 
		WHERE slug=$1 
		LIMIT 1 
	`

	row, err := con.QueryContext(ctx, sql_select, slug)
	helpers.ErrorCheck(err)

	for row.Next() {
		var (
			movieid_db, posted_id_db, views_db, year_db                           int
			movietype_db, movietitle_db, urlthumbnail_db, slug_db, description_db string
		)

		err := row.Scan(&movieid_db, &movietitle_db, &movietype_db, &posted_id_db, &urlthumbnail_db, &slug_db, &description_db, &views_db, &year_db)
		helpers.ErrorCheck(err)

		path_image := ""
		if urlthumbnail_db == "" {
			poster_image, poster_extension := _GetMedia(posted_id_db)
			path_image = "https://duniafilm.b-cdn.net/uploads/cache/poster_thumb/uploads/" + poster_extension + "/" + poster_image
		} else {
			path_image = urlthumbnail_db
		}

		movie_src := _GetVideoSingleRandom(movieid_db)

		//GENRE
		var objmoviegenre entities.Model_moviegenre
		var arraobjmoviegenre []entities.Model_moviegenre
		sql_selectmoviegenre := `SELECT 
			B.nmgenre, B.slug  
			FROM ` + config.DB_tbl_trx_moviegenre + ` as A 
			JOIN ` + config.DB_tbl_mst_movie_genre + ` as B ON B.idgenre = A.idgenre 
			WHERE A.movieid = $1   
		`
		row_moviegenre, err := con.QueryContext(ctx, sql_selectmoviegenre, movieid_db)
		helpers.ErrorCheck(err)
		for row_moviegenre.Next() {
			var (
				nmgenre_db, slug_db string
			)
			err := row_moviegenre.Scan(&nmgenre_db, &slug_db)
			helpers.ErrorCheck(err)
			objmoviegenre.Movie_genre = nmgenre_db
			objmoviegenre.Movie_slug = slug_db
			arraobjmoviegenre = append(arraobjmoviegenre, objmoviegenre)
		}

		//SOURCE
		var objmoviesource entities.Model_movievideo
		var arraobjmoviesource []entities.Model_movievideo
		if movietype_db == "movie" {
			sql_selectmoviesource := `SELECT 
				url 
				FROM ` + config.DB_tbl_mst_movie_source + ` 
				WHERE poster_id = $1  
			`
			row_moviesource, err := con.QueryContext(ctx, sql_selectmoviesource, movieid_db)
			helpers.ErrorCheck(err)
			nosource := 0
			for row_moviesource.Next() {
				nosource = nosource + 1
				var (
					url_db string
				)
				err := row_moviesource.Scan(&url_db)
				helpers.ErrorCheck(err)
				objmoviesource.Movie_title = "STREAM-" + strconv.Itoa(nosource)
				objmoviesource.Movie_src = url_db
				arraobjmoviesource = append(arraobjmoviesource, objmoviesource)
			}
		}

		obj.Movie_type = movietype_db
		obj.Movie_title = movietitle_db
		obj.Movie_descp = description_db
		obj.Movie_img = path_image
		obj.Movie_src = movie_src
		obj.Movie_year = year_db
		obj.Movie_view = views_db
		obj.Movie_slug = slug_db
		obj.Movie_genre = arraobjmoviegenre
		obj.Movie_video = arraobjmoviesource
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
func SeasonMovie(idmovie int) (helpers.Response, error) {
	var obj entities.Model_movieseason
	var arraobj []entities.Model_movieseason
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_season := `SELECT 
		id, title   
		FROM ` + config.DB_tbl_mst_series_season + ` 
		WHERE poster_id=?  
		ORDER BY position ASC      
	`
	row, err := con.QueryContext(ctx, sql_season, idmovie)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			id_db    int
			title_db string
		)

		err = row.Scan(&id_db, &title_db)
		helpers.ErrorCheck(err)

		obj.Season_id = id_db
		obj.Season_title = title_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
func EpisodeMovie(idseason int) (helpers.Response, error) {
	var obj entities.Model_movieepisode
	var arraobj []entities.Model_movieepisode
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_season := `SELECT 
		id, title   
		FROM ` + config.DB_tbl_mst_series_episode + ` 
		WHERE season_id=?  
		ORDER BY position ASC      
	`
	row, err := con.QueryContext(ctx, sql_season, idseason)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			id_db    int
			title_db string
		)

		err = row.Scan(&id_db, &title_db)
		helpers.ErrorCheck(err)

		obj.Episode_id = id_db
		obj.Episode_title = title_db
		obj.Episode_src = _GetVideoEpisode(id_db)
		arraobj = append(arraobj, obj)
		msg = "Success"
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}

func Update_movieview(username string, idmovie int) bool {
	flag := false

	viewmovielast := _GetMovie(idmovie, "")
	viewmovienow := viewmovielast + 1

	sql_update := `
			UPDATE 
			` + config.DB_tbl_trx_movie + `  
			SET views=$1  
			WHERE movieid=$2  
		`

	flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_trx_movie, "UPDATE", viewmovienow, idmovie)

	if flag_update {
		flag = true
		log.Println(msg_update)
	} else {
		log.Println(msg_update)
	}

	return flag
}
func Save_moviecomment(username, comment string, idmovie int) bool {
	tglnow, _ := goment.New()
	flag := false

	field_column := config.DB_tbl_trx_comment + tglnow.Format("YYYY")
	idrecord_counter := Get_counter(field_column)
	sql_insert := `
			insert into
			` + config.DB_tbl_trx_comment + ` (
				idcomment , idposter, username, comment, statusread, createcomment
			) values (
				$1,$2,$3,$4,$5,$6 
			)
		`
	flag_insert, msg_insert := Exec_SQL(sql_insert, config.DB_tbl_trx_comment, "INSERT", tglnow.Format("YY")+strconv.Itoa(idrecord_counter), idmovie, username, comment, "Y", tglnow.Format("YYYY-MM-DD HH:mm:ss"))

	if flag_insert {
		flag = true
		log.Println(msg_insert)
	} else {
		log.Println(msg_insert)
	}

	return flag
}
func Save_movierate(username, rating string, idmovie int) bool {
	tglnow, _ := goment.New()
	flag := false

	flag_check := CheckDBTwoField(config.DB_tbl_trx_rate, "username", username, "idposter", strconv.Itoa(idmovie))
	if !flag_check {
		field_column := config.DB_tbl_trx_rate + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		sql_insert := `
			insert into
			` + config.DB_tbl_trx_rate + ` (
				idrate , username, idposter, ratingposter, createrate
			) values (
				$1,$2,$3,$4,$5 
			)
		`
		flag_insert, msg_insert := Exec_SQL(sql_insert, config.DB_tbl_trx_rate, "INSERT",
			tglnow.Format("YY")+strconv.Itoa(idrecord_counter),
			username, idmovie, rating, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			flag = true
			log.Println(msg_insert)
		} else {
			log.Println(msg_insert)
		}
	}

	return flag
}
func Save_moviefavorite(username string, idmovie int) bool {
	tglnow, _ := goment.New()
	flag := false

	flag_check := CheckDBTwoField(config.DB_tbl_trx_favorite, "username", username, "idposter", strconv.Itoa(idmovie))
	if !flag_check {
		field_column := config.DB_tbl_trx_favorite + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		sql_insert := `
			insert into
			` + config.DB_tbl_trx_favorite + ` (
				idfavorite , idposter, username, createfavorite
			) values (
				$1,$2,$3,$4 
			)
		`
		flag_insert, msg_insert := Exec_SQL(sql_insert, config.DB_tbl_trx_favorite, "INSERT",
			tglnow.Format("YY")+tglnow.Format("MM")+strconv.Itoa(idrecord_counter),
			idmovie, username, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			flag = true
			log.Println(msg_insert)
		} else {
			log.Println(msg_insert)
		}
	}

	return flag
}
func Delete_moviefavorite(username string, idmovie int) bool {
	flag := false

	flag_check := CheckDBTwoField(config.DB_tbl_trx_favorite, "username", username, "idposter", strconv.Itoa(idmovie))
	if flag_check {
		sql_delete := `
			DELETE FROM 
			` + config.DB_tbl_trx_favorite + ` 
			WHERE idposter=$1 AND username=$2 
		`
		flag_delete, msg_delete := Exec_SQL(sql_delete, config.DB_tbl_trx_favorite, "DELETE", idmovie, username)

		if flag_delete {
			flag = true
			log.Println(msg_delete)
		} else {
			log.Println(msg_delete)
		}
	}

	return flag
}
func Save_moviereport(username, info string, idmovie int) bool {
	tglnow, _ := goment.New()
	flag := false

	info_1 := strings.ToUpper(info)
	info_final := strings.Replace(info_1, " ", "_", -1)

	field_column := config.DB_tbl_mst_report + tglnow.Format("YYYY")
	idrecord_counter := Get_counter(field_column)
	sql_insert := `
			insert into
			` + config.DB_tbl_mst_report + ` (
				idreport , idmovie, username, inforeport, poinreport, statusreport, createdatereport 
			) values (
				$1,$2,$3,$4,$5,$6,$7 
			)
		`
	flag_insert, msg_insert := Exec_SQL(sql_insert, config.DB_tbl_trx_favorite, "INSERT",
		tglnow.Format("YY")+tglnow.Format("MM")+strconv.Itoa(idrecord_counter),
		idmovie, username, info_final, 0, "PROCESS", tglnow.Format("YYYY-MM-DD HH:mm:ss"))

	if flag_insert {
		flag = true
		log.Println(msg_insert)
	} else {
		log.Println(msg_insert)
	}

	return flag
}
func Fetch_usermovie(username string) (helpers.Response, error) {
	var obj entities.Model_mobileuser
	var arraobj []entities.Model_mobileuser
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()
	tglnow, _ := goment.New()

	sql_select := `SELECT 
		username , nmuser, coderef, 
		point_in , point_out 
		FROM ` + config.DB_tbl_trx_user + ` 
		WHERE username=$1  
	`

	row, err := con.QueryContext(ctx, sql_select, username)
	helpers.ErrorCheck(err)

	for row.Next() {
		var (
			point_in_db, point_out_db          int
			username_db, nmuser_db, coderef_db string
		)

		err := row.Scan(&username_db, &nmuser_db, &coderef_db, &point_in_db, &point_out_db)
		helpers.ErrorCheck(err)

		if coderef_db == "" {
			flag_coderef := true
			numbertemp := ""
			for {
				numbergenerate := helpers.GenerateNumber(7)
				flag_coderef = CheckDB(config.DB_tbl_trx_user, "coderef", numbergenerate)
				if !flag_coderef {
					numbertemp = numbergenerate
					break
				}
			}
			if !flag_coderef {
				sql_update := `
					UPDATE 
					` + config.DB_tbl_trx_user + ` 
					SET coderef=$1, updatedateuser=$2  
					WHERE username=$3  
				`
				log.Println(username)
				log.Println(numbertemp)
				flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_trx_user, "UPDATE",
					numbertemp, tglnow.Format("YYYY-MM-DD HH:mm:ss"), username)

				if flag_update {
					coderef_db = numbertemp
					log.Println(msg_update)

				} else {
					log.Println(msg_update)
				}
			}
		}

		var objclaim entities.Model_mobilelistclaim
		var arraobjclaim []entities.Model_mobilelistclaim
		sql_selectlistclaim := `SELECT 
		idlistclaim , nmlistclaim, pointlistclaim 
		FROM ` + config.DB_tbl_mst_listclaim + ` 
		WHERE statuslistclaim='Y' 
		ORDER BY pointlistclaim ASC 
	`

		row_listclaim, err_listclaim := con.QueryContext(ctx, sql_selectlistclaim)
		helpers.ErrorCheck(err_listclaim)
		for row_listclaim.Next() {
			var (
				idlistclaim_db, pointlistclaim_db int
				nmlistclaim_db                    string
			)

			err := row_listclaim.Scan(&idlistclaim_db, &nmlistclaim_db, &pointlistclaim_db)
			helpers.ErrorCheck(err)

			objclaim.Claim_id = idlistclaim_db
			objclaim.Claim_name = nmlistclaim_db
			objclaim.Claim_point = pointlistclaim_db
			arraobjclaim = append(arraobjclaim, objclaim)
		}
		obj.User_username = username_db
		obj.User_name = nmuser_db
		obj.User_coderef = coderef_db
		obj.User_point = point_in_db - point_out_db
		obj.Listclaim = arraobjclaim
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
func Save_userclaim(username string, idlistclaim, point, pointbefore int) bool {
	tglnow, _ := goment.New()
	flag := false

	field_column := config.DB_tbl_mst_listclaim_user + tglnow.Format("YYYY")
	idrecord_counter := Get_counter(field_column)
	sql_insert := `
			insert into
			` + config.DB_tbl_mst_listclaim_user + ` (
				idclaimuser , idlistclaim, poinlistclaimtemp, username, pointbefore, statusclaimuser, 
				createclaimuser , createdateclaimuser
			) values (
				$1,$2,$3,$4,$5,$6,
				$7,$8 
			)
		`
	flag_insert, msg_insert := Exec_SQL(sql_insert, config.DB_tbl_mst_listclaim_user, "INSERT",
		tglnow.Format("YY")+tglnow.Format("MM")+strconv.Itoa(idrecord_counter),
		idlistclaim, username, point, pointbefore, "PROCESS", username, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

	if flag_insert {
		flag = true
		log.Println(msg_insert)
	} else {
		log.Println(msg_insert)
	}

	return flag
}

func Save_moviepoint(username, nmpoint string, idmovie, point int) bool {
	tglnow, _ := goment.New()
	flag := false

	field_column := config.DB_tbl_mst_point + tglnow.Format("YYYY")
	idrecord_counter := Get_counter(field_column)
	sql_insert := `
			insert into
			` + config.DB_tbl_mst_point + ` (
				idpoint , username, nmpoint, posted_id, point, createdatepoint
			) values (
				$1,$2,$3,$4,$5,$6 
			)
		`
	flag_insert, msg_insert := Exec_SQL(sql_insert, config.DB_tbl_mst_point, "INSERT",
		tglnow.Format("YY")+tglnow.Format("MM")+strconv.Itoa(idrecord_counter),
		username, nmpoint, idmovie, point, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

	if flag_insert {
		flag = true
		log.Println(msg_insert)

	} else {
		log.Println(msg_insert)
	}

	return flag
}
func Get_moviepoint(username, nmpoint string) string {
	con := db.CreateCon()
	ctx := context.Background()
	createdatepoint := ""

	sql_select := `
			SELECT
			createdatepoint      
			FROM ` + config.DB_tbl_mst_point + ` 
			WHERE username  = $1  
			AND nmpoint = $2  
			ORDER BY idpoint DESC LIMIT 1 
		`

	row := con.QueryRowContext(ctx, sql_select, username, nmpoint)
	switch e := row.Scan(&createdatepoint); e {
	case sql.ErrNoRows:
	case nil:
	default:
	}

	return createdatepoint
}
func Update_pointmember(username string) bool {
	tglnow, _ := goment.New()
	flag := false

	point_total := _GetPoint_In(username)

	sql_update := `
		UPDATE 
		` + config.DB_tbl_trx_user + ` 
		SET point_in=$1, updatedateuser=$2  
		WHERE username=$3  
	`
	log.Println(username)
	log.Println(point_total)
	flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_trx_user, "UPDATE",
		point_total, tglnow.Format("YYYY-MM-DD HH:mm:ss"), username)

	if flag_update {
		flag = true
		log.Println(msg_update)

	} else {
		log.Println(msg_update)
	}

	return flag
}
func _GetMedia(idrecord int) (string, string) {
	con := db.CreateCon()
	ctx := context.Background()
	url := ""
	extension := ""

	sql_select := `SELECT
		url, extension   
		FROM ` + config.DB_tbl_mst_mediatable + `  
		WHERE idmediatable = $1  
	`
	row := con.QueryRowContext(ctx, sql_select, idrecord)
	switch e := row.Scan(&url, &extension); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}
	return url, extension
}
func _GetVideoSingleRandom(idrecord int) string {
	con := db.CreateCon()
	ctx := context.Background()
	totalsource := 0
	source := ""
	sql_select := ""
	sql_select += "SELECT "
	sql_select += "url "
	sql_select += "FROM " + config.DB_tbl_mst_movie_source + " "
	sql_select += "WHERE poster_id = $1 "
	sql_select += "ORDER BY RANDOM() DESC LIMIT 1 "

	row_select, err_select := con.QueryContext(ctx, sql_select, idrecord)
	helpers.ErrorCheck(err_select)
	for row_select.Next() {
		totalsource = totalsource + 1
		var url_db string

		err := row_select.Scan(&url_db)
		helpers.ErrorCheck(err)
		source = url_db
	}
	return source
}
func _GetVideoEpisode(idrecord int) string {
	con := db.CreateCon()
	ctx := context.Background()
	url := ""

	sql_select := `SELECT
		url   
		FROM ` + config.DB_tbl_mst_movie_source + `  
		WHERE episode_id = $1  
	`
	row := con.QueryRowContext(ctx, sql_select, idrecord)
	switch e := row.Scan(&url); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}
	return url
}
func _GetMovie(idrecord int, tipe string) int {
	con := db.CreateCon()
	ctx := context.Background()
	views := 0

	sql_select := `SELECT
		views    
		FROM ` + config.DB_tbl_trx_movie + `  
		WHERE movieid = $1  
	`
	row := con.QueryRowContext(ctx, sql_select, idrecord)
	switch e := row.Scan(&views); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}
	return views
}
func _GetPoint_In(username string) int {
	con := db.CreateCon()
	ctx := context.Background()
	totalpoint := 0

	sql_select := `SELECT
		SUM(point) as totalpoint     
		FROM ` + config.DB_tbl_mst_point + `  
		WHERE username = $1  
	`
	row := con.QueryRowContext(ctx, sql_select, username)
	switch e := row.Scan(&totalpoint); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}
	return totalpoint
}
func _GetFavorite(idrecord int, username string) string {
	con := db.CreateCon()
	ctx := context.Background()
	idfavorite := 0
	favorite := "N"

	sql_select := `SELECT
		idfavorite   
		FROM ` + config.DB_tbl_trx_favorite + `  
		WHERE idposter = $1 AND username=$2  
	`
	row := con.QueryRowContext(ctx, sql_select, idrecord, username)
	switch e := row.Scan(&idfavorite); e {
	case sql.ErrNoRows:
	case nil:
		favorite = "Y"
	default:
		helpers.ErrorCheck(e)
	}
	return favorite
}
func _GetBanner() (interface{}, int) {
	var obj entities.Model_moviebanner
	var arraobj []entities.Model_moviebanner
	con := db.CreateCon()
	ctx := context.Background()
	totalbanner := 0

	sql_select := ""
	sql_select += "SELECT "
	sql_select += "urlbanner,urlwebsite "
	sql_select += "FROM " + config.DB_tbl_mst_banner + " "
	sql_select += "WHERE statusbanner = 'Y' "
	sql_select += "AND devicebanner = 'DEVICE' "
	sql_select += "ORDER BY displaybanner ASC "

	row_select, err_select := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err_select)
	for row_select.Next() {
		totalbanner = totalbanner + 1
		var urlimg_db, urldestination string

		err_select = row_select.Scan(&urlimg_db, &urldestination)
		helpers.ErrorCheck(err_select)
		obj.Moviebanner_urlimg = urlimg_db
		obj.Moviebanner_urldestination = urldestination
		arraobj = append(arraobj, obj)
	}
	return arraobj, totalbanner
}
func _GetIdGenre(slug string) (int, string) {
	con := db.CreateCon()
	ctx := context.Background()
	idgenre := 0
	nmgenre := ""

	sql_select := `SELECT
		idgenre,nmgenre    
		FROM ` + config.DB_tbl_mst_movie_genre + `  
		WHERE slug = $1  
	`
	row := con.QueryRowContext(ctx, sql_select, slug)
	switch e := row.Scan(&idgenre, &nmgenre); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}
	return idgenre, nmgenre
}
