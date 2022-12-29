package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	c := resty.New()

	//   send protected request => 401
	assert.Equal(t, http.StatusUnauthorized, touchProtected(c, ""))

	//create user => 201
	singUpResult := AuthSuccess{}
	singUpResponse, _ := singUp(c, "email", "12345", &singUpResult)
	assert.Equal(t, http.StatusOK, singUpResponse.StatusCode())
	assert.Equal(t, "email", singUpResult.Mail)
	assert.NotNil(t, "token", singUpResult.Token)

	//   send protected request => 200
	assert.Equal(t, http.StatusOK, touchProtected(c, singUpResult.Token))
}

func TestLoginUser(t *testing.T) {
	c := resty.New()

	//   send protected request => 401
	assert.Equal(t, http.StatusUnauthorized, touchProtected(c, ""))

	//create user => 201
	singUpResult := AuthSuccess{}
	singUpResponse, _ := singUp(c, "testEmail", "12345", &singUpResult)
	assert.Equal(t, http.StatusOK, singUpResponse.StatusCode())

	// login => 200
	singInResult := AuthSuccess{}
	singInResponse, _ := login(c, "email", "12345", &singInResult)
	assert.Equal(t, http.StatusOK, singInResponse.StatusCode())
	assert.Equal(t, "email", singInResult.Mail)
	assert.NotNil(t, "token", singInResult.Token)

	//   send protected request => 200
	assert.Equal(t, http.StatusOK, touchProtected(c, singInResult.Token))
}

func TestUploadImages(t *testing.T) {
	c := resty.New()

	//create user => 201
	singUpResult := AuthSuccess{}
	singUpResponse, _ := singUp(c, "testEmail", "12345", &singUpResult)
	assert.Equal(t, http.StatusOK, singUpResponse.StatusCode())

	assert.Equal(t, http.StatusCreated, uploadFiles(c, singUpResult.Token))
}

func TestUploadAndDeleteImages(t *testing.T) {
	c := resty.New()

	//   send protected request => 401
	assert.Equal(t, http.StatusUnauthorized, uploadFiles(c, ""))

	//create user => 201
	singUpResult := AuthSuccess{}
	singUpResponse, _ := singUp(c, "testEmail", "12345", &singUpResult)
	assert.Equal(t, http.StatusOK, singUpResponse.StatusCode())

	assert.Equal(t, http.StatusCreated, uploadFiles(c, singUpResult.Token))
	//assertEqual(t, true, strings.Contains(resp.String(), "File Uploaded successfully"))

	code, images := getAllImages(c)
	assert.Equal(t, http.StatusOK, code)

	code, singleImage := getImageInfo(c, images[0])

	assert.Equal(t, images[0].Url, singleImage.Url)
	assert.Equal(t, images[0].Title, singleImage.Title)

	assert.Equal(t, http.StatusOK, code)

	code = getImage(c, images[0])
	assert.Equal(t, http.StatusOK, code)

	code = deleteImages(c, images[0], "")
	assert.Equal(t, http.StatusUnauthorized, code)

	code = deleteImages(c, images[0], singUpResult.Token)
	assert.Equal(t, http.StatusOK, code)

	code = getImage(c, images[0])
	assert.Equal(t, http.StatusNotFound, code)
}
