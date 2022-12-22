package main

type AuthSuccess struct {
	Mail  string
	Token string
}

type ImageResponse struct {
	Url   string    `json:"url"`
	Tags  []*string `json:"tags"`
	Title string    `json:"title"`
}

type Images []ImageResponse
