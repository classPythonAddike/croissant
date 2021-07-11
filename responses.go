package croissant

/*
File Response class

A class which can be used to return files throught croissant.Response.File

Simply provide the name of the file to be returned into the Filename field.
*/
type FileResponse struct {
	Filename string
}

/*
Response class

A class which all route handlers must return.

The handler must return either HTML, json, a file, or none as response.
For example, if a handler wanted to return `<h1>Hello, World!</h1>`, it would look like -

func Handle(r *croissant.Request, b croissant.FormBody) croissant.Response {
	return croissant.Response {
		Html: "<h1>Hello, World!</h1>",
	}
}
*/
type Response struct {
	Html string
	File FileResponse
	Json interface{}

	StatusCode int
}
