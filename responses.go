package croissant

type FileResponse struct {
	Filename string
}

type Response struct {
	Html string
	File FileResponse
	Json interface{}

	StatusCode int
}
