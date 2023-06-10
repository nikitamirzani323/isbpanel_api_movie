package models

import (
	"context"
	"database/sql"
	"time"

	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/config"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/db"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/entities"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/helpers"
	"github.com/gofiber/fiber/v2"
)

func Get_AllDomain(client_domain string) (helpers.Response, bool, error) {
	var obj entities.Model_domain
	var arraobj []entities.Model_domain
	var res helpers.Response
	msg := "Data Not Found"
	render_page := time.Now()
	ctx := context.Background()
	con := db.CreateCon()
	statusdomain := ""
	flag := false

	sql_select := `SELECT
		nmdomain  
		FROM ` + config.DB_tbl_mst_domain + `  
		WHERE statusdomain = 'RUNNING'
	`
	row, err := con.QueryContext(ctx, sql_select)
	defer row.Close()
	helpers.ErrorCheck(err)

	for row.Next() {
		var nmdomain_db string
		err = row.Scan(&nmdomain_db)
		if err != nil {
			return res, false, err
		}
		obj.Domain_name = nmdomain_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}

	sql_select_detail := `SELECT
		statusdomain  
		FROM ` + config.DB_tbl_mst_domain + `  
		WHERE nmdomain = $1 
		AND statusdomain = 'RUNNING' 
	`
	rowdetail := con.QueryRowContext(ctx, sql_select_detail, client_domain)
	switch e := rowdetail.Scan(&statusdomain); e {
	case sql.ErrNoRows:
		flag = false
	case nil:
		flag = true
	default:
		flag = false
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(render_page).String()

	return res, flag, nil
}
