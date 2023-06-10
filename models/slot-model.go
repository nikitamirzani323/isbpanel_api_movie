package models

import (
	"context"
	"strconv"
	"strings"
	"time"

	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/config"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/db"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/entities"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/helpers"
	"github.com/gofiber/fiber/v2"
)

func Fetch_providerslotHome() (helpers.Response, error) {
	var obj entities.Model_providerslot
	var arraobj []entities.Model_providerslot
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
		nmproviderslot , providerslot_slug, providerslot_image, 
		providerslot_title , providerslot_descp 
		FROM ` + config.DB_tbl_mst_providerslot + ` 
		WHERE providerslot_status = 'Y'  
		ORDER BY providerslot_display ASC    
	`

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			nmproviderslot_db, providerslot_slug_db, providerslot_image_db string
			providerslot_title_db, providerslot_descp_db                   string
		)

		err = row.Scan(&nmproviderslot_db, &providerslot_slug_db, &providerslot_image_db,
			&providerslot_title_db, &providerslot_descp_db)

		helpers.ErrorCheck(err)

		obj.Providerslot_name = nmproviderslot_db
		obj.Providerslot_slug = providerslot_slug_db
		obj.Providerslot_image = providerslot_image_db
		obj.Providerslot_title = providerslot_title_db
		obj.Providerslot_descp = providerslot_descp_db
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
func Fetch_providerslotDetail(slug string) (helpers.Responseproviderslot, error) {
	var res helpers.Responseproviderslot
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()

	providerslot_name := ""
	providerslot_image := ""
	providerslot_title := ""
	providerslot_descp := ""

	sql_select := `SELECT 
		nmproviderslot , providerslot_image, 
		providerslot_title , providerslot_descp 
		FROM ` + config.DB_tbl_mst_providerslot + ` 
		WHERE providerslot_status = 'Y'  
		AND providerslot_slug = $1    
	`

	row, err := con.QueryContext(ctx, sql_select, slug)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			nmproviderslot_db, providerslot_image_db     string
			providerslot_title_db, providerslot_descp_db string
		)

		err = row.Scan(&nmproviderslot_db, &providerslot_image_db,
			&providerslot_title_db, &providerslot_descp_db)

		helpers.ErrorCheck(err)

		providerslot_name = nmproviderslot_db
		providerslot_image = providerslot_image_db
		providerslot_title = providerslot_title_db
		providerslot_descp = providerslot_descp_db

		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Providerslot_name = providerslot_name
	res.Providerslot_image = providerslot_image
	res.Providerslot_title = providerslot_title
	res.Providerslot_descp = providerslot_descp

	return res, nil
}
func Fetch_prediksislot(idprovider string, limit int) (helpers.Response, error) {
	var obj entities.Model_prediksislot
	var arraobj []entities.Model_prediksislot
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := ""
	sql_select += ""
	sql_select += "SELECT "
	sql_select += "A.nmgameslot , A.gameslot_image, A.gameslot_prediksi "
	sql_select += "FROM " + config.DB_tbl_trx_prediksislot + " as A  "
	sql_select += "JOIN " + config.DB_tbl_mst_providerslot + " as B ON B.idproviderslot = A.idproviderslot  "
	if idprovider != "" {
		sql_select += "WHERE LOWER(B.providerslot_slug) ='" + strings.ToLower(idprovider) + "'   "
	}
	if limit > 0 {
		sql_select += "ORDER BY A.gameslot_prediksi DESC LIMIT " + strconv.Itoa(limit) + "   "
	} else {
		sql_select += "ORDER BY A.gameslot_prediksi DESC    "
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			nmgameslot_db, gameslot_image_db string
			gameslot_prediksi_db             int
		)

		err = row.Scan(&nmgameslot_db, &gameslot_image_db, &gameslot_prediksi_db)

		helpers.ErrorCheck(err)

		prediksi_class := "bg-primary"
		if gameslot_prediksi_db > 70 {
			prediksi_class = "bg-success"
		}
		if gameslot_prediksi_db > 49 && gameslot_prediksi_db < 71 {
			prediksi_class = "bg-accent"
		}

		obj.Prediksislot_name = nmgameslot_db
		obj.Prediksislot_image = gameslot_image_db
		obj.Prediksislot_prediksi = gameslot_prediksi_db
		obj.Prediksislot_prediksi_class = prediksi_class
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
