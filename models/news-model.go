package models

import (
	"context"
	"time"

	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/config"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/db"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/entities"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/helpers"
	"github.com/gofiber/fiber/v2"
)

func Fetch_newsHome() (helpers.Response, error) {
	var obj entities.Model_news
	var arraobj []entities.Model_news
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
		title_news , descp_news, 
		url_news , img_news 
		FROM ` + config.DB_tbl_trx_news + ` 
		WHERE idcatenews != '2112'  
		ORDER BY idnews DESC LIMIT 15     
	`

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			title_news_db, descp_news_db, url_news_db, img_news_db string
		)

		err = row.Scan(&title_news_db, &descp_news_db, &url_news_db, &img_news_db)

		helpers.ErrorCheck(err)

		obj.News_title = title_news_db
		obj.News_descp = descp_news_db
		obj.News_url = url_news_db
		obj.News_image = img_news_db
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
func Fetch_newsMovieHome() (helpers.Response, error) {
	var obj entities.Model_news
	var arraobj []entities.Model_news
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
		title_news , descp_news, 
		url_news , img_news 
		FROM ` + config.DB_tbl_trx_news + ` 
		WHERE idcatenews = '2112'  
		ORDER BY idnews DESC LIMIT 15    
	`

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			title_news_db, descp_news_db, url_news_db, img_news_db string
		)

		err = row.Scan(&title_news_db, &descp_news_db, &url_news_db, &img_news_db)

		helpers.ErrorCheck(err)

		obj.News_title = title_news_db
		obj.News_descp = descp_news_db
		obj.News_url = url_news_db
		obj.News_image = img_news_db
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
