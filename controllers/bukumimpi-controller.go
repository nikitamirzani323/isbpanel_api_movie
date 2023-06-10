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

const Field_bukumimpihome_redis = "LISTBUKUMIMPI_FRONTEND_ISBPANEL"
const Field_tafsirmimpihome_redis = "LISTTAFSIRMIMPI_FRONTEND_ISBPANEL"

func Bukumimpihome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_clienrequest)
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
	log.Println("BUKUMIMPI Client origin : ", hostname)

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

	var obj entities.Model_bukumimpi
	var arraobj []entities.Model_bukumimpi
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Field_bukumimpihome_redis + "-" + client.Tipe + "-" + client.Nama)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		bukumimpi_type, _ := jsonparser.GetString(value, "bukumimpi_type")
		bukumimpi_name, _ := jsonparser.GetString(value, "bukumimpi_name")
		bukumimpi_nomor, _ := jsonparser.GetString(value, "bukumimpi_nomor")

		obj.Bukumimpi_type = bukumimpi_type
		obj.Bukumimpi_name = bukumimpi_name
		obj.Bukumimpi_nomor = bukumimpi_nomor
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_bukumimpiHome(client.Tipe, client.Nama)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Field_bukumimpihome_redis+"-"+client.Tipe+"-"+client.Nama, result, 1*time.Minute)
		log.Printf("BUKUMIMPI MYSQL %s - %s\n", client.Tipe, client.Nama)
		return c.JSON(result)
	} else {
		log.Printf("BUKUMIMPI CACHE %s - %s\n", client.Tipe, client.Nama)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func TafsirMimpihome(c *fiber.Ctx) error {
	client := new(entities.Controller_clientafsirmimpirequest)
	if err := c.BodyParser(client); err != nil {
		return err
	}

	client_origin := c.Request().Body()
	data_origin := []byte(client_origin)
	hostname, _ := jsonparser.GetString(data_origin, "client_hostname")
	log.Println("Request Body : ", string(data_origin))
	log.Println("TAFSIR MIMPI Client origin : ", hostname)
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

	var obj entities.Model_tafsirmimpi
	var arraobj []entities.Model_tafsirmimpi
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Field_tafsirmimpihome_redis + "-" + client.Search)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		tafsirmimpi_mimpi, _ := jsonparser.GetString(value, "tafsirmimpi_mimpi")
		tafsirmimpi_artimimpi, _ := jsonparser.GetString(value, "tafsirmimpi_artimimpi")
		tafsirmimpi_angka2d, _ := jsonparser.GetString(value, "tafsirmimpi_angka2d")
		tafsirmimpi_angka3d, _ := jsonparser.GetString(value, "tafsirmimpi_angka3d")
		tafsirmimpi_angka4d, _ := jsonparser.GetString(value, "tafsirmimpi_angka4d")

		obj.Tafsirmimpi_mimpi = tafsirmimpi_mimpi
		obj.Tafsirmimpi_artimimpi = tafsirmimpi_artimimpi
		obj.Tafsirmimpi_angka4d = tafsirmimpi_angka4d
		obj.Tafsirmimpi_angka3d = tafsirmimpi_angka3d
		obj.Tafsirmimpi_angka2d = tafsirmimpi_angka2d
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_tafsirmimpiHome(client.Search)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Field_tafsirmimpihome_redis+"-"+client.Search, result, 1*time.Minute)
		log.Printf("TAFSIR MIMPI MYSQL %s\n", client.Search)
		return c.JSON(result)
	} else {
		log.Printf("TAFSIR MIMPI CACHE %s\n", client.Search)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
