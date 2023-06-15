package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"bitbucket.org/isbtotogroup/isbpanel_api_movie/config"
	"bitbucket.org/isbtotogroup/isbpanel_api_movie/db"
	"bitbucket.org/isbtotogroup/isbpanel_api_movie/entities"
	"bitbucket.org/isbtotogroup/isbpanel_api_movie/helpers"
	"github.com/gofiber/fiber/v2"
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
func Fetch_movieHome(search, tipe string, perpage_client, page int) (helpers.ResponsePaging, error) {
	var obj entities.Model_movie
	var arraobj []entities.Model_movie
	var res helpers.ResponsePaging
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	perpage := perpage_client
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
		sql_select += "ORDER BY random() LIMIT " + strconv.Itoa(perpage)
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

		movie_src := ""

		//GENRE
		type s_genre struct {
			Genre_id int
		}
		var objgen s_genre
		var arrobjgen []s_genre
		var temp_idgenre = ""
		var objmoviegenre entities.Model_moviegenre
		var arraobjmoviegenre []entities.Model_moviegenre
		sql_selectmoviegenre := `SELECT 
			B.idgenre, B.nmgenre, B.slug  
			FROM ` + config.DB_tbl_trx_moviegenre + ` as A 
			JOIN ` + config.DB_tbl_mst_movie_genre + ` as B ON B.idgenre = A.idgenre 
			WHERE A.movieid = $1   
		`
		row_moviegenre, err := con.QueryContext(ctx, sql_selectmoviegenre, movieid_db)
		helpers.ErrorCheck(err)
		for row_moviegenre.Next() {
			var (
				idgenre_db          int
				nmgenre_db, slug_db string
			)
			err := row_moviegenre.Scan(&idgenre_db, &nmgenre_db, &slug_db)
			helpers.ErrorCheck(err)
			objmoviegenre.Movie_genre = nmgenre_db
			objmoviegenre.Movie_slug = slug_db
			arraobjmoviegenre = append(arraobjmoviegenre, objmoviegenre)

			objgen.Genre_id = idgenre_db
			arrobjgen = append(arrobjgen, objgen)
		}
		defer row_moviegenre.Close()
		for i := 0; i < len(arrobjgen); i++ {
			if i == len(arrobjgen)-1 {
				temp_idgenre += strconv.Itoa(arrobjgen[i].Genre_id)
			} else {
				temp_idgenre += strconv.Itoa(arrobjgen[i].Genre_id) + ","
			}
		}

		//SOURCE
		var objmoviesource entities.Model_movievideo
		var arraobjmoviesource []entities.Model_movievideo
		var objmovieseries entities.Model_movieseason
		var arrobjmovieseries []entities.Model_movieseason
		if movietype_db == "movie" {
			movie_src = _GetVideoSingleRandom(movieid_db)

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
			defer row_moviesource.Close()
		} else {
			sql_selectmovieseason := `SELECT 
				id, title   
				FROM ` + config.DB_tbl_mst_series_season + ` 
				WHERE poster_id=$1   
				ORDER BY position ASC   
			`
			// fmt.Println(sql_selectmovieseason)
			row_movieseason, err := con.QueryContext(ctx, sql_selectmovieseason, movieid_db)
			helpers.ErrorCheck(err)
			for row_movieseason.Next() {
				var (
					id_db    int
					title_db string
				)
				err := row_movieseason.Scan(&id_db, &title_db)
				helpers.ErrorCheck(err)
				objmovieseries.Season_id = id_db
				objmovieseries.Season_title = title_db
				arrobjmovieseries = append(arrobjmovieseries, objmovieseries)
			}
			defer row_movieseason.Close()
		}

		// MOVIE NEW
		var objmovienew entities.Model_movie
		var arraobjmovienew []entities.Model_movie
		sql_movienew := `SELECT 
				movietitle , COALESCE(posted_id,0) , urlthumbnail, slug 
				FROM ` + config.DB_tbl_trx_movie + ` 
				WHERE enabled = 1 
				ORDER BY createdatemovie DESC LIMIT 15   
			`

		row_movienew, err_movienew := con.QueryContext(ctx, sql_movienew)
		helpers.ErrorCheck(err_movienew)

		for row_movienew.Next() {
			var (
				posted_id_db                            int
				movietitle_db, urlthumbnail_db, slug_db string
			)

			err := row_movienew.Scan(&movietitle_db, &posted_id_db, &urlthumbnail_db, &slug_db)
			helpers.ErrorCheck(err)
			path_image := ""
			if urlthumbnail_db == "" {
				poster_image, poster_extension := _GetMedia(posted_id_db)
				path_image = "https://duniafilm.b-cdn.net/uploads/cache/poster_thumb/uploads/" + poster_extension + "/" + poster_image
			} else {
				path_image = urlthumbnail_db
			}

			objmovienew.Movie_title = movietitle_db
			objmovienew.Movie_thumbnail = path_image
			objmovienew.Movie_slug = slug_db
			arraobjmovienew = append(arraobjmovienew, objmovienew)
		}
		defer row_movienew.Close()

		// MOVIE GENRE
		var objlistmoviegenre entities.Model_movie
		var arraobjlistmoviegenre []entities.Model_movie
		sql_listmoviegenre := ""
		sql_listmoviegenre += "SELECT "
		sql_listmoviegenre += "distinct on (B.slug) B.slug, B.movietitle , COALESCE(B.posted_id,0) , B.urlthumbnail "
		sql_listmoviegenre += "FROM " + config.DB_tbl_trx_moviegenre + " as A "
		sql_listmoviegenre += "JOIN " + config.DB_tbl_trx_movie + " as B ON A.movieid = B.movieid "
		sql_listmoviegenre += "WHERE B.enabled = 1 "
		sql_listmoviegenre += "AND A.idgenre in (" + fmt.Sprint(temp_idgenre) + ") "
		sql_listmoviegenre += "LIMIT 15 "
		// fmt.Println(sql_listmoviegenre)
		row_listmoviegenre, err_listmoviegenre := con.QueryContext(ctx, sql_listmoviegenre)
		helpers.ErrorCheck(err_listmoviegenre)

		for row_listmoviegenre.Next() {
			var (
				posted_id_db                            int
				movietitle_db, urlthumbnail_db, slug_db string
			)

			err := row_listmoviegenre.Scan(&slug_db, &movietitle_db, &posted_id_db, &urlthumbnail_db)
			helpers.ErrorCheck(err)
			path_image := ""
			if urlthumbnail_db == "" {
				poster_image, poster_extension := _GetMedia(posted_id_db)
				path_image = "https://duniafilm.b-cdn.net/uploads/cache/poster_thumb/uploads/" + poster_extension + "/" + poster_image
			} else {
				path_image = urlthumbnail_db
			}

			objlistmoviegenre.Movie_title = movietitle_db
			objlistmoviegenre.Movie_thumbnail = path_image
			objlistmoviegenre.Movie_slug = slug_db
			arraobjlistmoviegenre = append(arraobjlistmoviegenre, objlistmoviegenre)
		}
		defer row_listmoviegenre.Close()

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
		obj.Movie_listvideonew = arraobjmovienew
		obj.Movie_listvideogenre = arraobjlistmoviegenre
		obj.Movie_listseason = arrobjmovieseries
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
		WHERE season_id=$1   
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
