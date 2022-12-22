package main

import (
	"bytes"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"path/filepath"
)

const base = "http://localhost:7777"

func touchProtected(client *resty.Client, token string) int {
	endpoint := base + "/test"
	resp, _ := client.R().
		EnableTrace().
		SetHeader("Authorization", CreateAccessToken(token)).
		Get(endpoint)

	return resp.StatusCode()
}

func singUp(client *resty.Client, email string, password string, result *AuthSuccess) (*resty.Response, error) {
	endpoint := base + "/register"

	return client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"mail":"` + email + `", "password":"` + password + `"}`).
		SetResult(result).
		Post(endpoint)
}

func login(client *resty.Client, email string, password string, result *AuthSuccess) (*resty.Response, error) {
	endpoint := base + "/login"
	return client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"mail":"` + email + `", "password":"` + password + `"}`).
		SetResult(result).
		Post(endpoint)
}

func getAllImages(client *resty.Client) (int, Images) {
	result := Images{}

	endpoint := base + "/images"
	resp, _ := client.R().
		EnableTrace().
		SetResult(&result).
		Get(endpoint)

	return resp.StatusCode(), result
}

func getImage(client *resty.Client, response ImageResponse) int {
	endpoint := base + "/images/" + response.Title
	resp, _ := client.R().
		EnableTrace().
		Get(endpoint)

	return resp.StatusCode()
}

func getImageInfo(client *resty.Client, response ImageResponse) (int, ImageResponse) {
	result := ImageResponse{}

	endpoint := base + "/images/" + response.Title + "/info"
	resp, _ := client.R().
		EnableTrace().
		//ForceContentType("application/json").
		SetResult(&result).
		Get(endpoint)

	return resp.StatusCode(), result
}

func uploadFiles(client *resty.Client, token string) int {
	endpoint := base + "/images"

	profileImgBytes, _ := ioutil.ReadFile(filepath.Join(getTestDataPath(), "test-img.png"))

	resp, _ := client.R().
		SetHeader("Authorization", CreateAccessToken(token)).
		SetFileReader("file", "test-img.png", bytes.NewReader(profileImgBytes)).
		SetFormData(map[string]string{
			"tags": "Jeevanandam",
		}).
		Post(endpoint)

	return resp.StatusCode()
}

func deleteImages(client *resty.Client, response ImageResponse, token string) int {
	endpoint := base + "/images/actions/delete/"
	resp, _ := client.R().
		EnableTrace().
		SetHeader("Authorization", CreateAccessToken(token)).
		SetBody(`{"ids":["` + response.Title + `"]}`).
		Post(endpoint)

	return resp.StatusCode()
}
