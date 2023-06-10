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

const Fieldbanner_home_redis = "LISTBANNER_FRONTEND_ISBPANEL"

func Bannerhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_banner)
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
	log.Println("BANNER CLIENT origin : ", hostname)
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

	var obj entities.Model_banner
	var arraobj []entities.Model_banner
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldbanner_home_redis)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		banner_url, _ := jsonparser.GetString(value, "banner_url")
		banner_urlwebsite, _ := jsonparser.GetString(value, "banner_urlwebsite")
		banner_posisi, _ := jsonparser.GetString(value, "banner_posisi")
		banner_device, _ := jsonparser.GetString(value, "banner_device")

		obj.Banner_url = banner_url
		obj.Banner_urlwebsite = banner_urlwebsite
		obj.Banner_device = banner_device
		obj.Banner_posisi = banner_posisi
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Get_AllBanner()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldbanner_home_redis, result, 60*time.Hour)
		log.Println("BANNER MYSQL")
		return c.JSON(result)
	} else {
		log.Println("BANNER CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
