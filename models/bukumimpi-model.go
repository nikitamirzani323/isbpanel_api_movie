package models

import (
	"context"
	"strings"
	"time"

	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/config"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/db"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/entities"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/helpers"
	"github.com/gofiber/fiber/v2"
)

func Fetch_bukumimpiHome(tipe, nama string) (helpers.Response, error) {
	var obj entities.Model_bukumimpi
	var arraobj []entities.Model_bukumimpi
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()
	sql_select := ""
	if tipe == "" {
		sql_select += ""
		sql_select += "SELECT "
		sql_select += "typebukumimpi, nmbukumimpi, nmrbukumimpi "
		sql_select += "FROM " + config.DB_tbl_trx_bukumimpi + " "
		if nama != "" {
			sql_select += "WHERE LOWER(nmbukumimpi) LIKE '%" + strings.ToLower(nama) + "%' "
			sql_select += "ORDER BY nmbukumimpi ASC LIMIT 500 "
		} else {
			sql_select += "ORDER BY random() LIMIT 500 "
		}

	} else {
		sql_select += ""
		sql_select += "SELECT "
		sql_select += "typebukumimpi, nmbukumimpi, nmrbukumimpi "
		sql_select += "FROM " + config.DB_tbl_trx_bukumimpi + " "
		if nama != "" {
			sql_select += "WHERE LOWER(nmbukumimpi) LIKE '%" + strings.ToLower(nama) + "%' "
			sql_select += "AND typebukumimpi='" + tipe + "' "
		} else {
			sql_select += "WHERE typebukumimpi='" + tipe + "' "
		}
		sql_select += "ORDER BY nmbukumimpi ASC LIMIT 500 "
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			typebukumimpi_db, nmbukumimpi_db, nmrbukumimpi_db string
		)

		err = row.Scan(&typebukumimpi_db, &nmbukumimpi_db, &nmrbukumimpi_db)

		helpers.ErrorCheck(err)

		obj.Bukumimpi_type = typebukumimpi_db
		obj.Bukumimpi_name = nmbukumimpi_db
		obj.Bukumimpi_nomor = nmrbukumimpi_db
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
func Fetch_tafsirmimpiHome(search string) (helpers.Response, error) {
	var obj entities.Model_tafsirmimpi
	var arraobj []entities.Model_tafsirmimpi
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()
	sql_select := ""
	if search == "" {
		sql_select += ""
		sql_select += "SELECT "
		sql_select += "mimpi, artimimpi, angka2d, angka3d, angka4d "
		sql_select += "FROM " + config.DB_tbl_mst_tafsirmimpi + " "
		sql_select += "WHERE statustafsirmimpi = 'Y' "
		sql_select += "ORDER BY random() LIMIT 500  "
	} else {
		sql_select += ""
		sql_select += "SELECT "
		sql_select += "mimpi, artimimpi, angka2d, angka3d, angka4d "
		sql_select += "FROM " + config.DB_tbl_mst_tafsirmimpi + " "
		sql_select += "WHERE statustafsirmimpi = 'Y' "
		sql_select += "AND LOWER(mimpi) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "OR LOWER(artimimpi) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "ORDER BY mimpi ASC LIMIT 500  "
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			mimpi_db, artimimpi_db, angka2d_db, angka3d_db, angka4d_db string
		)

		err = row.Scan(&mimpi_db, &artimimpi_db, &angka2d_db, &angka3d_db, &angka4d_db)

		helpers.ErrorCheck(err)

		obj.Tafsirmimpi_mimpi = mimpi_db
		obj.Tafsirmimpi_artimimpi = artimimpi_db
		obj.Tafsirmimpi_angka4d = angka4d_db
		obj.Tafsirmimpi_angka3d = angka3d_db
		obj.Tafsirmimpi_angka2d = angka2d_db
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
