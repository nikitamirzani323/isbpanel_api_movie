package models

import (
	"context"
	"time"

	"bitbucket.org/isbtotogroup/isbpanel_api_movie/config"
	"bitbucket.org/isbtotogroup/isbpanel_api_movie/db"
	"bitbucket.org/isbtotogroup/isbpanel_api_movie/entities"
	"bitbucket.org/isbtotogroup/isbpanel_api_movie/helpers"
	"github.com/gofiber/fiber/v2"
)

func Get_AllBanner() (helpers.Response, error) {
	var obj entities.Model_banner
	var arraobj []entities.Model_banner
	var res helpers.Response
	msg := "Data Not Found"
	render_page := time.Now()
	ctx := context.Background()
	con := db.CreateCon()

	sql_select := `SELECT
		urlbanner, urlwebsite, posisibanner, devicebanner   
		FROM ` + config.DB_tbl_mst_banner + `  
		WHERE statusbanner = 'Y'
		ORDER BY displaybanner 
	`
	row, err := con.QueryContext(ctx, sql_select)
	defer row.Close()
	helpers.ErrorCheck(err)

	for row.Next() {
		var urlbanner_db, urlwebsite_db, posisibanner_db, devicebanner_db string
		err = row.Scan(&urlbanner_db, &urlwebsite_db, &posisibanner_db, &devicebanner_db)
		if err != nil {
			return res, err
		}
		obj.Banner_url = urlbanner_db
		obj.Banner_urlwebsite = urlwebsite_db
		obj.Banner_posisi = posisibanner_db
		obj.Banner_device = devicebanner_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(render_page).String()

	return res, nil
}
