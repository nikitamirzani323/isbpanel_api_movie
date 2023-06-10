package controllers

import (
	"log"
	"strconv"
	"time"

	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/entities"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/helpers"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/models"
	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const Fieldmovie_home_redis = "LISTMOVIE_FRONTEND_ISBPANEL"
const Fieldseason_home_redis = "LISTSEASON_FRONTEND_ISBPANEL"
const Fieldepisode_home_redis = "LISTEPISODE_FRONTEND_ISBPANEL"

const Fieldmovie_mobile_redis = "LISTMOVIE-MOBILE"
const Fieldmoviegenre_mobile_redis = "LISTMOVIEGENRE-MOBILE"
const Fieldmoviedetail_mobile_redis = "LISTMOVIEDETAIL-MOBILE"
const Fieldfrontpagemovie_mobile_redis = "LISTFRONTPAGE-MOBILE"
const Fieldseason_mobile_redis = "LISTSEASON_MOBILE"
const Fieldepisode_mobile_redis = "LISTSEASONEPISODE_MOBILE"
const Fieldmoviecomment_mobile_redis = "LISTMOVIECOMMENT_MOBILE"
const Fielduser_mobile_redis = "LISTUSER_MOBILE"

func Moviehome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_clientmovie)
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
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	log.Println("Client TOKEN : ", temp_decp)
	log.Println("Client BODYPARSE : ", client.Client_hostname)

	// flag_client := models.Get_Domain(temp_decp)

	// if !flag_client {
	// 	c.Status(fiber.StatusBadRequest)
	// 	return c.JSON(fiber.Map{
	// 		"status":  fiber.StatusBadRequest,
	// 		"message": "NOT REGISTER",
	// 		"record":  nil,
	// 	})
	// }
	var obj entities.Model_moviecategory
	var arraobj []entities.Model_moviecategory
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldmovie_home_redis)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		movie_category, _ := jsonparser.GetString(value, "movie_category")
		movie_list, _, _, _ := jsonparser.Get(value, "movie_list")
		var objchild entities.Model_movie
		var arraobjchild []entities.Model_movie
		jsonparser.ArrayEach(movie_list, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			movie_title, _ := jsonparser.GetString(value, "movie_title")
			movie_thumbnail, _ := jsonparser.GetString(value, "movie_thumbnail")
			movie_slug, _ := jsonparser.GetString(value, "movie_slug")
			// movie_description, _ := jsonparser.GetString(value, "movie_description")
			// movie_video, _, _, _ := jsonparser.Get(value, "movie_video")
			// var objmoviesrc entities.Model_movievideo
			// var arraobjmoviesrc []entities.Model_movievideo
			// jsonparser.ArrayEach(movie_video, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			// 	movie_src, _ := jsonparser.GetString(value, "movie_src")
			// 	objmoviesrc.Movie_src = movie_src
			// 	arraobjmoviesrc = append(arraobjmoviesrc, objmoviesrc)
			// })

			objchild.Movie_title = movie_title
			objchild.Movie_thumbnail = movie_thumbnail
			objchild.Movie_slug = movie_slug
			// objchild.Movie_video = arraobjmoviesrc
			arraobjchild = append(arraobjchild, objchild)
		})
		obj.Movie_category = movie_category
		obj.Movie_list = arraobjchild
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_movieHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldmovie_home_redis, result, time.Minute*120)
		log.Println("MOVIE MYSQL")
		return c.JSON(result)
	} else {
		log.Println("MOVIE CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Movieseason(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_season)
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
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	log.Println("Client TOKEN : ", temp_decp)
	log.Println("Client BODYPARSE : ", client.Client_hostname)
	// flag_client := models.Get_Domain(temp_decp)
	// if !flag_client {
	// 	c.Status(fiber.StatusBadRequest)
	// 	return c.JSON(fiber.Map{
	// 		"status":  fiber.StatusBadRequest,
	// 		"message": "NOT REGISTER",
	// 		"record":  nil,
	// 	})
	// }

	var obj entities.Model_movieseason
	var arraobj []entities.Model_movieseason
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldseason_home_redis + "_" + strconv.Itoa(client.Movie_id))
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		season_id, _ := jsonparser.GetInt(value, "season_id")
		season_title, _ := jsonparser.GetString(value, "season_title")

		obj.Season_id = int(season_id)
		obj.Season_title = season_title
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.SeasonMovie(client.Movie_id)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldseason_home_redis+"_"+strconv.Itoa(client.Movie_id), result, time.Minute*1)
		log.Println("MOVIE SEASON MYSQL")
		return c.JSON(result)
	} else {
		log.Println("MOVIE SEASON CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Movieepisode(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_episode)
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
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	log.Println("Client TOKEN : ", temp_decp)
	log.Println("Client BODYPARSE : ", client.Client_hostname)

	var obj entities.Model_movieepisode
	var arraobj []entities.Model_movieepisode
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldepisode_home_redis + "_" + strconv.Itoa(client.Season_id))
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		episode_id, _ := jsonparser.GetInt(value, "episode_id")
		episode_title, _ := jsonparser.GetString(value, "episode_title")
		episode_src, _ := jsonparser.GetString(value, "episode_src")

		obj.Episode_id = int(episode_id)
		obj.Episode_title = episode_title
		obj.Episode_src = episode_src
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.EpisodeMovie(client.Season_id)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldepisode_home_redis+"_"+strconv.Itoa(client.Season_id), result, time.Minute*1)
		log.Println("MOVIE SEASON MYSQL")
		return c.JSON(result)
	} else {
		log.Println("MOVIE SEASON CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
