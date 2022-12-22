package handlers

import (
	"broker/repositories"
	"github.com/kataras/iris/v12"
)

const maxSize = 8 * iris.MB

func AuthRouter(userRepo *repositories.UserRepo) func(iris.Party) {
	return func(app iris.Party) {
		app.Post("register", SignUp(*userRepo))
		app.Post("login", SignIn(*userRepo))
		app.Use(Verify()) // protect the next routes with JWT.

		app.Get("test", func(ctx iris.Context) {
			ctx.Writef("you touch protected %s ))", "Gateway")
		})

		app.Post("signout", SignOut)
	}
}

func TestRouter() func(iris.Party) {
	return func(app iris.Party) {
		app.Get("/", func(ctx iris.Context) {
			ctx.Writef("I'm live %s ))", "Gateway")
		})
	}
}

func ImageAndTagsRouter(imageRepo *repositories.ImageRepository, tagRepo *repositories.TagRepository) func(iris.Party) {
	return func(app iris.Party) {
		app.Get("images", GetAllImages(imageRepo))
		//app.Get("images", GetAllImagesFromFS())
		app.Get("images/{title}", GetImageById())
		app.Get("images/{title}/info", GetImageInfoById(imageRepo))
		app.Get("tags", GetAllTags(tagRepo))
		app.Use(Verify()) // protect the next routes with JWT.

		app.Post("images/{title}/update_tags", UpdateTags(imageRepo, tagRepo))
		app.Post("images/", iris.LimitRequestBodySize(maxSize+1<<20), PostImages(tagRepo, imageRepo))
		app.Post("images/actions/add_tag", AddTag(tagRepo, imageRepo))
		app.Post("images/actions/delete", DeleteImages(imageRepo))


		app.Use(Verify()) // protect the next routes with JWT.
	}
}
