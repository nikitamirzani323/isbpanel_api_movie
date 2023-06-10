package controllers

import (
	"log"
	"strconv"

	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/config"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/entities"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/helpers"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/nleeper/goment"
)

func CheckLoginmobile(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Loginmobile)
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

	result := models.Loginmobile_Model(client.Username, client.Name, "LOGIN")
	version := models.Mobileversion_Model()

	if result {
		flag_insert := false
		tglnow, _ := goment.New()
		tglcreate := models.Get_moviepoint(client.Username, "POINT_LOGIN")
		log.Println(tglcreate)
		if tglcreate == "" {
			flag_insert = true
		} else {
			tglcreate2, _ := goment.New(tglcreate)
			tglstart := tglnow.Format("YYYY-MM-DD") + " 00:00:00"
			tglend := tglnow.Format("YYYY-MM-DD") + " 23:59:59"
			tgldbpoint := tglcreate2.Format("YYYY-MM-DD HH:mm:ss")

			log.Printf("data server : %s", tgldbpoint)
			if tgldbpoint >= tglstart && tgldbpoint <= tglend {
				flag_insert = false
			} else {
				flag_insert = true
			}
		}
		if flag_insert { // UPDATE POINT PER DAY
			flag_point := models.Save_moviepoint(client.Username, "POINT_LOGIN", 0, config.POINT_LOGIN)
			log.Printf("POINT_LOGIN STATUS : %t", flag_point)
		}

		dataclient := client.Username + "==" + client.Name + "==" + client.Device
		dataclient_encr, keymap := helpers.Encryption(dataclient)
		dataclient_encr_final := dataclient_encr + "|" + strconv.Itoa(keymap)
		t, err := helpers.GenerateNewAccessToken(dataclient_encr_final)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{
			"status":  "Y",
			"version": version,
			"message": "updated",
			"token":   t,
		})
	} else {
		result := models.Loginmobile_Model(client.Username, client.Name, "INSERT")
		log.Println("RESULT", result)
		return c.JSON(fiber.Map{
			"status":  "N",
			"version": version,
			"message": "",
		})
	}
}
