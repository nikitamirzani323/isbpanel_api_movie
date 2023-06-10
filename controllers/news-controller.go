package controllers

import (
	"log"
	"time"

	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/entities"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/helpers"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/models"
	"github.com/buger/jsonparser"
	"github.com/gofiber/fiber/v2"
)

const Fieldnews_home_redis = "LISTNEWS_FRONTEND_ISBPANEL"
const Fieldnewsmovie_home_redis = "LISTNEWSMOVIES_FRONTEND_ISBPANEL"

func Newshome(c *fiber.Ctx) error {
	client_origin := c.Request().Body()
	data_origin := []byte(client_origin)
	hostname, _ := jsonparser.GetString(data_origin, "client_hostname")
	log.Println("Request Body : ", string(data_origin))
	log.Println("NEWS Client origin : ", hostname)
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

	var obj entities.Model_news
	var arraobj []entities.Model_news
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldnews_home_redis)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		news_title, _ := jsonparser.GetString(value, "news_title")
		news_descp, _ := jsonparser.GetString(value, "news_descp")
		news_url, _ := jsonparser.GetString(value, "news_url")
		news_image, _ := jsonparser.GetString(value, "news_image")

		obj.News_title = news_title
		obj.News_descp = news_descp
		obj.News_url = news_url
		obj.News_image = news_image
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_newsHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldnews_home_redis, result, 60*time.Minute)
		log.Println("NEWS MYSQL")
		return c.JSON(result)
	} else {
		log.Println("NEWS CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Newsmoviehome(c *fiber.Ctx) error {
	client_origin := c.Request().Body()
	data_origin := []byte(client_origin)
	hostname, _ := jsonparser.GetString(data_origin, "client_hostname")
	log.Println("Request Body : ", string(data_origin))
	log.Println("MOVIE Client origin : ", hostname)
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
	var obj entities.Model_news
	var arraobj []entities.Model_news
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldnewsmovie_home_redis)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		news_title, _ := jsonparser.GetString(value, "news_title")
		news_descp, _ := jsonparser.GetString(value, "news_descp")
		news_url, _ := jsonparser.GetString(value, "news_url")
		news_image, _ := jsonparser.GetString(value, "news_image")

		obj.News_title = news_title
		obj.News_descp = news_descp
		obj.News_url = news_url
		obj.News_image = news_image
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_newsMovieHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldnewsmovie_home_redis, result, 60*time.Minute)
		log.Println("NEWS MOVIE MYSQL")
		return c.JSON(result)
	} else {
		log.Println("NEWS MOVIE CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
