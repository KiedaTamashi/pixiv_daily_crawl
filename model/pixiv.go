package model

type PixivResponse struct {
	Error        bool        `json:"error"`
	ErrorMessage string      `json:"message"`
	Pages        []pixivPage `json:"body"`
}

type pixivPage struct {
	Urls struct {
		Original string `json:"original"`
	} `json:"urls"`
}

type PixivOneIdImage struct {
	Id     string
	Images []*PixivImage
}

type PixivImage struct {
	Data     []byte
	Filename string
}
