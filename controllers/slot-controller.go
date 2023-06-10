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

const Fieldproviderslot_home_redis = "LISTPROVIDERSLOT_FRONTEND_ISBPANEL"
const Fieldprediksislot_home_redis = "LISTPREDIKSISLOT_FRONTEND_ISBPANEL"

func Providerslothome(c *fiber.Ctx) error {
	client_origin := c.Request().Body()
	data_origin := []byte(client_origin)
	hostname, _ := jsonparser.GetString(data_origin, "client_hostname")
	log.Println("Request Body : ", string(data_origin))
	log.Println("SLOT Client origin : ", hostname)
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

	var obj entities.Model_providerslot
	var arraobj []entities.Model_providerslot
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldproviderslot_home_redis)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		providerslot_slug, _ := jsonparser.GetString(value, "providerslot_slug")
		providerslot_name, _ := jsonparser.GetString(value, "providerslot_name")
		providerslot_image, _ := jsonparser.GetString(value, "providerslot_image")
		providerslot_title, _ := jsonparser.GetString(value, "providerslot_title")
		providerslot_descp, _ := jsonparser.GetString(value, "providerslot_descp")

		obj.Providerslot_slug = providerslot_slug
		obj.Providerslot_name = providerslot_name
		obj.Providerslot_image = providerslot_image
		obj.Providerslot_title = providerslot_title
		obj.Providerslot_descp = providerslot_descp
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_providerslotHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldproviderslot_home_redis, result, 60*time.Minute)
		log.Println("PROVIDER SLOT MYSQL")
		return c.JSON(result)
	} else {
		log.Println("PROVIDER SLOT CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Prediksislotdetail(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_providerslot)
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
	log.Println("PREDIKSI SLOT Client origin : ", hostname)
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

	field_redis := Fieldproviderslot_home_redis + "_" + client.Providerslot_slug
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	providerslot_name, _ := jsonparser.GetString(jsonredis, "providerslot_name")
	providerslot_image, _ := jsonparser.GetString(jsonredis, "providerslot_image")
	providerslot_title, _ := jsonparser.GetString(jsonredis, "providerslot_title")
	providerslot_descp, _ := jsonparser.GetString(jsonredis, "providerslot_descp")

	if !flag {
		result, err := models.Fetch_providerslotDetail(client.Providerslot_slug)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(field_redis, result, 60*time.Minute)
		log.Println("PROVIDER SLOT DETAIL MYSQL")
		return c.JSON(result)
	} else {
		log.Println("PROVIDER SLOT DETAIL CACHE")
		return c.JSON(fiber.Map{
			"status":             fiber.StatusOK,
			"message":            message_RD,
			"providerslot_name":  providerslot_name,
			"providerslot_image": providerslot_image,
			"providerslot_title": providerslot_title,
			"providerslot_descp": providerslot_descp,
		})
	}
}
func Prediksislothome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_prediksislot)
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
	log.Println("PREDIKSI SLOT DASHBOARD Client origin : ", hostname)
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
	var obj entities.Model_prediksislot
	var arraobj []entities.Model_prediksislot
	render_page := time.Now()
	field_redis := ""
	if client.Providerslot_id != "" {
		field_redis = Fieldprediksislot_home_redis + "_" + client.Providerslot_id
	} else {
		field_redis = Fieldprediksislot_home_redis
	}
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		prediksislot_name, _ := jsonparser.GetString(value, "prediksislot_name")
		prediksislot_image, _ := jsonparser.GetString(value, "prediksislot_image")
		prediksislot_prediksi_class, _ := jsonparser.GetString(value, "prediksislot_prediksi_class")
		prediksislot_prediksi, _ := jsonparser.GetInt(value, "prediksislot_prediksi")

		obj.Prediksislot_name = prediksislot_name
		obj.Prediksislot_image = prediksislot_image
		obj.Prediksislot_prediksi_class = prediksislot_prediksi_class
		obj.Prediksislot_prediksi = int(prediksislot_prediksi)
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_prediksislot(client.Providerslot_id, client.Prediksislot_limit)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(field_redis, result, 60*time.Minute)
		log.Println("PREDIKSI SLOT MYSQL")
		return c.JSON(result)
	} else {
		log.Println("PREDIKSI SLOT CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
