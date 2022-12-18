package main

import (
	"broker/database"
	"broker/handlers"
	"broker/repositories"
	"broker/utils"
	"github.com/kataras/iris/v12"
	"log"
	"os"
	"strconv"
)

func main() {

	db, err := database.InitDb()
	if err != nil {
		log.Panic("Can't connect to Mysql")
	}

	userRepo := repositories.CreateUserRepo(db.Gorm)
	tagRepository := repositories.CreateTagRepository(db.Gorm)
	imgRepository := repositories.CreateImageRepository(db.Gorm)

	//clearTables(tagRepository, imgRepository)
	ac := utils.MakeAccessLog()
	defer ac.Close() // Close the underline file.
	ac.AddOutput(os.Stdout)
	ac.RequestBody = true

	app := iris.New()

	crs := func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")

		if ctx.Method() == iris.MethodOptions {
			ctx.Header("Access-Control-Methods",
				"POST, PUT, PATCH, DELETE")

			ctx.Header("Access-Control-Allow-Headers",
				"Access-Control-Allow-Origin,Content-Type, authorization")

			ctx.Header("Access-Control-Max-Age",
				"86400")

			ctx.StatusCode(iris.StatusNoContent)
			return
		}

		ctx.Next()
	}

	app.UseRouter(ac.Handler, crs)

	//app.UseRouter(crs)

	app.PartyFunc("/", handlers.AuthRouter(userRepo))
	app.PartyFunc("/", handlers.TestRouter())
	app.PartyFunc("/", handlers.ImageAndTagsRouter(imgRepository, tagRepository))

	app.Listen(":7777")
}

func clearTables(r *repositories.TagRepository, rImg *repositories.ImageRepository) {
	tagsFromDb, _ := (r).GetAll()
	for _, tag := range tagsFromDb {
		r.Delete(strconv.FormatUint(uint64(tag.ID), 10))
	}
	imagesFromDb, _ := rImg.GetAll()
	for _, description := range imagesFromDb {
		rImg.Delete(strconv.FormatUint(uint64(description.ID), 10))
	}
}
