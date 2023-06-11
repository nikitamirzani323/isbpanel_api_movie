package controllers

import (
	"fmt"
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

const Fieldmovie_genre_redis = "LISTGENRE_FRONTEND_ISBPANEL"
const Fieldmovie_home_redis = "LISTMOVIE_FRONTEND_ISBPANEL"
const Fieldmoviegenre_home_redis = "LISTMOVIEGENRE_FRONTEND_ISBPANEL"
const Fieldmoviedetail_home_redis = "LISTMOVIEDETAIL_FRONTEND_ISBPANEL"
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

func Moviegenre(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_clientgenre)
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

	var obj entities.Model_moviegenre
	var arraobj []entities.Model_moviegenre
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldmovie_genre_redis)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		movie_genre, _ := jsonparser.GetString(value, "movie_genre")
		movie_slug, _ := jsonparser.GetString(value, "movie_slug")

		obj.Movie_genre = movie_genre
		obj.Movie_slug = movie_slug
		arraobj = append(arraobj, obj)

	})
	if !flag {
		result, err := models.Fetch_genre()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldmovie_genre_redis, result, time.Minute*720)
		fmt.Println("GENRE MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("GENRE CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
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

	var obj entities.Model_movie
	var arraobj []entities.Model_movie
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldmovie_home_redis + "_" + strconv.Itoa(client.Movie_page) + "_" + client.Movie_tipe + "_" + client.Movie_search)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	perpage_RD, _ := jsonparser.GetInt(jsonredis, "perpage")
	totalrecord_RD, _ := jsonparser.GetInt(jsonredis, "totalrecord")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		movie_title, _ := jsonparser.GetString(value, "movie_title")
		movie_thumbnail, _ := jsonparser.GetString(value, "movie_thumbnail")
		movie_slug, _ := jsonparser.GetString(value, "movie_slug")

		obj.Movie_title = movie_title
		obj.Movie_thumbnail = movie_thumbnail
		obj.Movie_slug = movie_slug
		arraobj = append(arraobj, obj)

	})
	if !flag {
		result, err := models.Fetch_movieHome(client.Movie_search, client.Movie_tipe, client.Movie_page)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldmovie_home_redis+"_"+strconv.Itoa(client.Movie_page)+"_"+client.Movie_tipe+"_"+client.Movie_search, result, time.Minute*720)
		fmt.Println("MOVIE MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("MOVIE CACHE")
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     message_RD,
			"record":      arraobj,
			"perpage":     perpage_RD,
			"totalrecord": totalrecord_RD,
			"time":        time.Since(render_page).String(),
		})
	}
}
func MoviehomeByGenre(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_clientmoviegenre)
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

	var obj entities.Model_movie
	var arraobj []entities.Model_movie
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldmoviegenre_home_redis + "_" + strconv.Itoa(client.Movie_page) + "_" + client.Slug)
	jsonredis := []byte(resultredis)
	genre_RD, _ := jsonparser.GetString(jsonredis, "genre")
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	perpage_RD, _ := jsonparser.GetInt(jsonredis, "perpage")
	totalrecord_RD, _ := jsonparser.GetInt(jsonredis, "totalrecord")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		movie_title, _ := jsonparser.GetString(value, "movie_title")
		movie_thumbnail, _ := jsonparser.GetString(value, "movie_thumbnail")
		movie_slug, _ := jsonparser.GetString(value, "movie_slug")

		obj.Movie_title = movie_title
		obj.Movie_thumbnail = movie_thumbnail
		obj.Movie_slug = movie_slug
		arraobj = append(arraobj, obj)

	})
	if !flag {
		result, err := models.Fetch_movieByGenre(client.Slug, client.Movie_page)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldmoviegenre_home_redis+"_"+strconv.Itoa(client.Movie_page)+"_"+client.Slug, result, time.Minute*720)
		fmt.Println("MOVIE-GENRE MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("MOVIE-GENRE CACHE")
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     message_RD,
			"record":      arraobj,
			"perpage":     perpage_RD,
			"totalrecord": totalrecord_RD,
			"genre":       genre_RD,
			"time":        time.Since(render_page).String(),
		})
	}
}
func MoviehomeByDetail(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_clientmoviedetail)
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

	var obj entities.Model_moviedetailwebsite
	var arraobj []entities.Model_moviedetailwebsite
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldmoviedetail_home_redis + "_" + client.Slug)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		movie_type, _ := jsonparser.GetString(value, "movie_type")
		movie_title, _ := jsonparser.GetString(value, "movie_title")
		movie_descp, _ := jsonparser.GetString(value, "movie_descp")
		movie_img, _ := jsonparser.GetString(value, "movie_img")
		movie_src, _ := jsonparser.GetString(value, "movie_src")
		movie_slug, _ := jsonparser.GetString(value, "movie_slug")
		movie_year, _ := jsonparser.GetInt(value, "movie_year")
		movie_view, _ := jsonparser.GetInt(value, "movie_view")

		var objmoviegenre entities.Model_moviegenre
		var arraobjmoviegenre []entities.Model_moviegenre
		record_moviegenre_RD, _, _, _ := jsonparser.Get(value, "movie_genre")
		jsonparser.ArrayEach(record_moviegenre_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			movie_genre, _ := jsonparser.GetString(value, "movie_genre")
			movie_slug, _ := jsonparser.GetString(value, "movie_slug")
			objmoviegenre.Movie_genre = movie_genre
			objmoviegenre.Movie_slug = movie_slug
			arraobjmoviegenre = append(arraobjmoviegenre, objmoviegenre)
		})

		var objmoviesource entities.Model_movievideo
		var arraobjmoviesource []entities.Model_movievideo
		record_moviesource_RD, _, _, _ := jsonparser.Get(value, "movie_video")
		jsonparser.ArrayEach(record_moviesource_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			movie_title, _ := jsonparser.GetString(value, "movie_title")
			movie_src, _ := jsonparser.GetString(value, "movie_src")

			objmoviesource.Movie_title = movie_title
			objmoviesource.Movie_src = movie_src
			arraobjmoviesource = append(arraobjmoviesource, objmoviesource)
		})

		obj.Movie_type = movie_type
		obj.Movie_title = movie_title
		obj.Movie_descp = movie_descp
		obj.Movie_src = movie_src
		obj.Movie_img = movie_img
		obj.Movie_year = int(movie_year)
		obj.Movie_view = int(movie_view)
		obj.Movie_slug = movie_slug
		obj.Movie_genre = arraobjmoviegenre
		obj.Movie_video = arraobjmoviesource
		arraobj = append(arraobj, obj)

	})
	if !flag {
		result, err := models.Fetch_movieDetail(client.Slug)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldmoviedetail_home_redis+"_"+client.Slug, result, time.Minute*720)
		fmt.Println("MOVIE-DETAIL MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("MOVIE-DETAIL CACHE")
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
