package controllers

import (
	"log"
	"time"

	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/entities"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/helpers"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/models"
	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

const fielddomain_redis = "LISTDOMAIN_FRONTEND_ISBPANEL"
const Field_home_redis = "LISTPASARAN_FRONTEND_ISBPANEL"
const Field_keluaran_redis = "LISTKELUARAN_FRONTEND_ISBPANEL"

func Pasaranhome(c *fiber.Ctx) error {
	client_origin := c.Request().Body()
	data_origin := []byte(client_origin)
	hostname, _ := jsonparser.GetString(data_origin, "client_hostname")
	log.Println("Request Body : ", string(data_origin))
	log.Println("PASARAN Client origin : ", hostname)
	flag_client := _domainsecurity(hostname)
	log.Println("STATUS DOMAIN : ", flag_client)
	if !flag_client {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "NOT REGISTER",
			"record":  nil,
		})
	}

	var obj entities.Model_pasaran
	var arraobj []entities.Model_pasaran
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Field_home_redis)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		pasaran_id, _ := jsonparser.GetString(value, "pasaran_id")
		pasaran_name, _ := jsonparser.GetString(value, "pasaran_name")
		pasaran_url, _ := jsonparser.GetString(value, "pasaran_url")
		pasaran_diundi, _ := jsonparser.GetString(value, "pasaran_diundi")
		pasaran_jamjadwal, _ := jsonparser.GetString(value, "pasaran_jamjadwal")
		pasaran_datekeluaran, _ := jsonparser.GetString(value, "pasaran_datekeluaran")
		pasaran_slug, _ := jsonparser.GetString(value, "pasaran_slug")
		pasaran_meta_title, _ := jsonparser.GetString(value, "pasaran_meta_title")
		pasaran_meta_descp, _ := jsonparser.GetString(value, "pasaran_meta_descp")
		pasaran_keluaran, _ := jsonparser.GetString(value, "pasaran_keluaran")
		pasaran_dateprediksi, _ := jsonparser.GetString(value, "pasaran_dateprediksi")
		pasaran_bbfsprediksi, _ := jsonparser.GetString(value, "pasaran_bbfsprediksi")
		pasaran_nomorprediksi, _ := jsonparser.GetString(value, "pasaran_nomorprediksi")

		obj.Pasaran_id = pasaran_id
		obj.Pasaran_name = pasaran_name
		obj.Pasaran_url = pasaran_url
		obj.Pasaran_diundi = pasaran_diundi
		obj.Pasaran_jamjadwal = pasaran_jamjadwal
		obj.Pasaran_slug = pasaran_slug
		obj.Pasaran_meta_title = pasaran_meta_title
		obj.Pasaran_meta_descp = pasaran_meta_descp
		obj.Pasaran_datekeluaran = pasaran_datekeluaran
		obj.Pasaran_keluaran = pasaran_keluaran
		obj.Pasaran_dateprediksi = pasaran_dateprediksi
		obj.Pasaran_bbfsprediksi = pasaran_bbfsprediksi
		obj.Pasaran_nomorprediksi = pasaran_nomorprediksi
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_pasaranHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Field_home_redis, result, 60*time.Minute)
		log.Println("PASARAN MYSQL")
		return c.JSON(result)
	} else {
		log.Println("PASARAN CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Keluaranhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_keluaran)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	client_origin := c.Request().Body()
	data_origin := []byte(client_origin)
	hostname, _ := jsonparser.GetString(data_origin, "client_hostname")
	log.Println("Request Body : ", string(data_origin))
	log.Println("KELUARAN Client origin : ", hostname)
	flag_client := _domainsecurity(hostname)
	log.Println("STATUS DOMAIN : ", flag_client)
	if !flag_client {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "NOT REGISTER",
			"record":  nil,
		})
	}

	var obj entities.Model_keluaran
	var arraobj []entities.Model_keluaran
	var obj_pasaransimple entities.Model_pasaransimple
	var arraobj_pasaransimple []entities.Model_pasaransimple
	var obj_minggu entities.Model_keluaranpaitominggu
	var arraobj_minggu []entities.Model_keluaranpaitominggu
	var obj_senin entities.Model_keluaranpaitosenin
	var arraobj_senin []entities.Model_keluaranpaitosenin
	var obj_selasa entities.Model_keluaranpaitoselasa
	var arraobj_selasa []entities.Model_keluaranpaitoselasa
	var obj_rabu entities.Model_keluaranpaitorabu
	var arraobj_rabu []entities.Model_keluaranpaitorabu
	var obj_kamis entities.Model_keluaranpaitokamis
	var arraobj_kamis []entities.Model_keluaranpaitokamis
	var obj_jumat entities.Model_keluaranpaitojumat
	var arraobj_jumat []entities.Model_keluaranpaitojumat
	var obj_sabtu entities.Model_keluaranpaitosabtu
	var arraobj_sabtu []entities.Model_keluaranpaitosabtu
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Field_keluaran_redis + "_" + client.Pasaran_id)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	pasaran_nama, _ := jsonparser.GetString(jsonredis, "pasaran_nama")
	pasaran_livedraw, _ := jsonparser.GetString(jsonredis, "pasaran_livedraw")
	pasaran_diundi, _ := jsonparser.GetString(jsonredis, "pasaran_diundi")
	pasaran_jadwal, _ := jsonparser.GetString(jsonredis, "pasaran_jadwal")
	pasaran_title, _ := jsonparser.GetString(jsonredis, "pasaran_title")
	pasaran_descp, _ := jsonparser.GetString(jsonredis, "pasaran_descp")
	list_pasaran_RD, _, _, _ := jsonparser.Get(jsonredis, "list_pasaran")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	paito_minggu_RD, _, _, _ := jsonparser.Get(jsonredis, "paito_minggu")
	paito_senin_RD, _, _, _ := jsonparser.Get(jsonredis, "paito_senin")
	paito_selasa_RD, _, _, _ := jsonparser.Get(jsonredis, "paito_selasa")
	paito_rabu_RD, _, _, _ := jsonparser.Get(jsonredis, "paito_rabu")
	paito_kamis_RD, _, _, _ := jsonparser.Get(jsonredis, "paito_kamis")
	paito_jumat_RD, _, _, _ := jsonparser.Get(jsonredis, "paito_jumat")
	paito_sabtu_RD, _, _, _ := jsonparser.Get(jsonredis, "paito_sabtu")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		keluaran_datekeluaran, _ := jsonparser.GetString(value, "keluaran_datekeluaran")
		keluaran_periode, _ := jsonparser.GetString(value, "keluaran_periode")
		keluaran_nomor, _ := jsonparser.GetString(value, "keluaran_nomor")

		obj.Keluaran_datekeluaran = keluaran_datekeluaran
		obj.Keluaran_periode = keluaran_periode
		obj.Keluaran_nomor = keluaran_nomor
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(list_pasaran_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		pasaran_name, _ := jsonparser.GetString(value, "pasaran_name")
		pasaran_url, _ := jsonparser.GetString(value, "pasaran_url")
		pasaran_slug, _ := jsonparser.GetString(value, "pasaran_slug")

		obj_pasaransimple.Pasaran_name = pasaran_name
		obj_pasaransimple.Pasaran_url = pasaran_url
		obj_pasaransimple.Pasaran_slug = pasaran_slug
		arraobj_pasaransimple = append(arraobj_pasaransimple, obj_pasaransimple)
	})
	jsonparser.ArrayEach(paito_minggu_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		keluaran_nomor, _ := jsonparser.GetString(value, "keluaran_nomor")
		obj_minggu.Keluaran_nomor = keluaran_nomor
		arraobj_minggu = append(arraobj_minggu, obj_minggu)
	})
	jsonparser.ArrayEach(paito_senin_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		keluaran_nomor, _ := jsonparser.GetString(value, "keluaran_nomor")
		obj_senin.Keluaran_nomor = keluaran_nomor
		arraobj_senin = append(arraobj_senin, obj_senin)
	})
	jsonparser.ArrayEach(paito_selasa_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		keluaran_nomor, _ := jsonparser.GetString(value, "keluaran_nomor")
		obj_selasa.Keluaran_nomor = keluaran_nomor
		arraobj_selasa = append(arraobj_selasa, obj_selasa)
	})
	jsonparser.ArrayEach(paito_rabu_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		keluaran_nomor, _ := jsonparser.GetString(value, "keluaran_nomor")
		obj_rabu.Keluaran_nomor = keluaran_nomor
		arraobj_rabu = append(arraobj_rabu, obj_rabu)
	})
	jsonparser.ArrayEach(paito_kamis_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		keluaran_nomor, _ := jsonparser.GetString(value, "keluaran_nomor")
		obj_kamis.Keluaran_nomor = keluaran_nomor
		arraobj_kamis = append(arraobj_kamis, obj_kamis)
	})
	jsonparser.ArrayEach(paito_jumat_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		keluaran_nomor, _ := jsonparser.GetString(value, "keluaran_nomor")
		obj_jumat.Keluaran_nomor = keluaran_nomor
		arraobj_jumat = append(arraobj_jumat, obj_jumat)
	})
	jsonparser.ArrayEach(paito_sabtu_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		keluaran_nomor, _ := jsonparser.GetString(value, "keluaran_nomor")
		obj_sabtu.Keluaran_nomor = keluaran_nomor
		arraobj_sabtu = append(arraobj_sabtu, obj_sabtu)
	})
	if !flag {
		result, err := models.Fetch_keluaran(client.Pasaran_id)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Field_keluaran_redis+"_"+client.Pasaran_id, result, 60*time.Minute)
		log.Println("KELUARAN MYSQL")
		return c.JSON(result)
	} else {
		log.Println("KELUARAN CACHE")
		return c.JSON(fiber.Map{
			"status":           fiber.StatusOK,
			"message":          message_RD,
			"pasaran_nama":     pasaran_nama,
			"pasaran_livedraw": pasaran_livedraw,
			"pasaran_diundi":   pasaran_diundi,
			"pasaran_jadwal":   pasaran_jadwal,
			"pasaran_title":    pasaran_title,
			"pasaran_descp":    pasaran_descp,
			"list_pasaran":     arraobj_pasaransimple,
			"record":           arraobj,
			"paito_minggu":     arraobj_minggu,
			"paito_senin":      arraobj_senin,
			"paito_selasa":     arraobj_selasa,
			"paito_rabu":       arraobj_rabu,
			"paito_kamis":      arraobj_kamis,
			"paito_jumat":      arraobj_jumat,
			"paito_sabtu":      arraobj_sabtu,
			"time":             time.Since(render_page).String(),
		})
	}
}
func _domainsecurity(nmdomain string) bool {
	log.Printf("Domain Client : %s", nmdomain)
	resultredis, flag_domain := helpers.GetRedis(fielddomain_redis)
	flag := false
	if len(nmdomain) > 0 {
		if !flag_domain {
			result, temp_flag, err := models.Get_AllDomain(nmdomain)
			if err != nil {
				flag = false
			}
			log.Println("DOMAIN MYSQL")
			helpers.SetRedis(fielddomain_redis, result, 24*time.Hour)
			flag = temp_flag
		} else {
			jsonredis_domain := []byte(resultredis)
			record_RD, _, _, _ := jsonparser.Get(jsonredis_domain, "record")
			jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				domain, _ := jsonparser.GetString(value, "domain_name")
				if nmdomain == domain {
					flag = true
					log.Println("DOMAIN CACHE")
				}
			})
		}
	}
	return flag
}
