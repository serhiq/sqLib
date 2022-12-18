package handlers

import (
	"broker/data"
	"broker/repositories"
	"fmt"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type ImageResponse struct {
	Url   string    `json:"url"`
	Tags  []*string `json:"tags"`
	Title string    `json:"title"`
}

type Images []ImageResponse

var actualImageCache = false
var cacheImages = Images{}

func GetAllImages(repo *repositories.ImageRepository) iris.Handler {
	return func(ctx iris.Context) {
		descriptions, err := repo.GetAll()

		if err != nil {
			ctx.StopWithStatus(iris.StatusInternalServerError)
		}

		images := Images{}

		for _, file := range descriptions {
			tags := []*string{}

			for _, tag := range file.Tags {
				tags = append(tags, &tag.Title)
			}
			item := ImageResponse{"http://localhost:7777/images/" + file.FileName, tags, file.FileName}
			images = append(images, item)
		}

		ctx.JSON(images)
		ctx.StatusCode(200)
	}
}
func GetAllTags(repo *repositories.TagRepository) iris.Handler {
	return func(ctx iris.Context) {
		tags, err := repo.GetAll()

		if err != nil {
			ctx.StopWithStatus(iris.StatusInternalServerError)
		}

		ctx.JSON(tags)
		ctx.StatusCode(200)
	}
}

func DeleteImages(repo *repositories.ImageRepository) iris.Handler {
	return func(ctx iris.Context) {

		var request DeleteImageRequest
		err := ctx.ReadJSON(&request)
		if err != nil {
			ctx.StopWithStatus(iris.StatusBadRequest)
		}

		for _, fileName := range request.Ids {
			err := repo.DeleteByFileName(fileName)
			if err != nil {

				if err != nil {
					return
				}
				continue

			}

			src := "./uploads/" + fileName
			os.Remove(src)

		}

		ctx.StatusCode(200)
	}
}

func GetImageById() iris.Handler {
	return func(ctx iris.Context) {

		id := ctx.Params().Get("title")

		src := "./uploads/" + id

		file, err := ioutil.ReadFile(src)
		if err != nil {
			ctx.StopWithStatus(iris.StatusNotFound)

			fmt.Println(err)
		}

		ctx.Write(file)
		ctx.StatusCode(200)
	}
}

func GetImageInfoById(repo *repositories.ImageRepository) iris.Handler {
	return func(ctx iris.Context) {

		fileName := ctx.Params().Get("title")

		info, err := repo.GetByFilename(fileName)

		if err != nil {
			ctx.StopWithStatus(iris.StatusNotFound)
		}

		tags := []*string{}

		for _, tag := range info.Tags {
			tags = append(tags, &tag.Title)
		}

		response := ImageResponse{"http://localhost:7777/images/" + info.FileName, tags, info.FileName}

		ctx.JSON(response)
	}
}
func UpdateTags(imgRepo *repositories.ImageRepository, tagRepo *repositories.TagRepository) iris.Handler {
	return func(ctx iris.Context) {

		fileName := ctx.Params().Get("title")

		var request UpdateImageTagsRequest
		err := ctx.ReadJSON(&request)
		if err != nil {
			ctx.StopWithStatus(iris.StatusBadRequest)
		}
		img, err := imgRepo.GetByFilename(fileName)

		if err != nil {
			ctx.StopWithStatus(iris.StatusNotFound)
		}

		var newTagArray []*data.Tag

		for _, tag := range request.Tags {
			newTag, _ := tagRepo.GetOrCreate(tag)
			if newTag != nil {
				newTagArray = append(newTagArray, newTag)
			}

		}
		imgRepo.UpdateTags(strconv.FormatUint(uint64(img.ID), 10), newTagArray)

		ctx.StatusCode(200)
	}
}

func GetAllImagesFromFS() iris.Handler {
	return func(ctx iris.Context) {

		if !actualImageCache {

			src := "./uploads/"

			files, err := ioutil.ReadDir(src)
			if err != nil {
				log.Fatal(err)
			}

			SortTimeDescend(files)

			cacheImages = Images{}

			for _, file := range files {
				item := ImageResponse{"http://localhost:7777/images/" + file.Name(), []*string{}, file.Name()}
				cacheImages = append(cacheImages, item)
			}
			actualImageCache = true
		}

		ctx.JSON(cacheImages)
		ctx.StatusCode(200)
	}
}

func SortTimeDescend(files []os.FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime().After(files[j].ModTime())
	})
}

func PostImages(tagRepo *repositories.TagRepository, imgRepo *repositories.ImageRepository) iris.Handler {
	return func(ctx iris.Context) {
		actualImageCache = false

		strTags := ctx.FormValue("tags")
		tagsStr := getTagList(strTags)
		var tags = []*data.Tag{}

		for _, tag := range tagsStr {
			tag, err := tagRepo.GetOrCreate(tag)
			if err == nil {
				tags = append(tags, tag)
			}
		}

		f, fh, err := ctx.FormFile("file")

		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			return
		}
		defer f.Close()

		fileName := generateFileName(fh.Filename)

		_, err = ctx.SaveFormFile(fh, filepath.Join("./uploads/", fileName))
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		}
		img := data.ImagesDescription{
			FileName: fileName,
			Tags:     tags,
		}

		err = imgRepo.Insert(&img)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		}
		ctx.StatusCode(201)
	}
}

func generateFileName(fileName string) string {
	ext := afterStr(fileName, ".")
	name := strconv.FormatInt(time.Now().Unix(), 10)
	return fmt.Sprintf("%s%s%s", name+randSeq(3), ".", ext)
}

func afterStr(value string, a string) string {
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func getTagList(strTags string) []string {
	var result = []string{}
	check := make(map[string]int)

	arr := strings.Split(strTags, ",")

	for _, r := range arr {
		check[r] = 1
	}

	for name := range check {
		if name != "" {
			result = append(result, name)

		}
	}
	return result
}

func AddTag(tagRepo *repositories.TagRepository, imgRepo *repositories.ImageRepository) iris.Handler {
	return func(ctx iris.Context) {

		var request BulkOperationImageRequest
		err := ctx.ReadJSON(&request)
		if err != nil {
			ctx.StopWithStatus(iris.StatusBadRequest)
		}

		for _, fileName := range request.Ids {
			img, err := imgRepo.GetByFilename(fileName)
			if err != nil {
				continue
			}
			var needAdd []*data.Tag

			for _, tag := range request.Tags {
				if !haveSameTag(img.Tags, tag) {
					newTag, _ := tagRepo.GetOrCreate(tag)
					if newTag != nil {
						needAdd = append(needAdd, newTag)
					}
				}
			}
			newTagArray := append(img.Tags, needAdd...)
			imgRepo.UpdateTags(strconv.FormatUint(uint64(img.ID), 10), newTagArray)

		}

		ctx.StatusCode(200)
	}
}

func haveSameTag(tags []*data.Tag, tag string) bool {
	for _, t := range tags {
		if t.Title == tag {
			return true
		}
	}
	return false
}

type BulkOperationImageRequest struct {
	Ids  []string `json:"ids"`
	Tags []string `json:"tags"`
}

type DeleteImageRequest struct {
	Ids []string `json:"ids"`
}

type UpdateImageTagsRequest struct {
	Tags []string `json:"tags"`
}
