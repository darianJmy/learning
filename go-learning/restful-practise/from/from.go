package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"github.com/gorilla/schema"
	"io"
	"log"
	"net/http"
)

type Profile struct {
	Name string
	Age  int
}

var decoder *schema.Decoder

func main() {
	decoder = schema.NewDecoder()
	ws := new(restful.WebService)
	ws.Route(ws.POST("/profiles").Consumes("application/x-www-form-urlencoded").To(postAddress))
	ws.Route(ws.GET("/profiles").To(addressForm))
	restful.Add(ws)
	http.ListenAndServe(":8080", nil)
}

func postAddress(req *restful.Request, resp *restful.Response) {
	err := req.Request.ParseForm()
	if err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		log.Println(1)
		return
	}
	p := new(Profile)
	err = decoder.Decode(p, req.Request.PostForm)
	if err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		log.Println(2)
		return
	}
	io.WriteString(resp.ResponseWriter, fmt.Sprintf("<html><body>Name=%s, Age=%d</body></html>", p.Name, p.Age))
}

func addressForm(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp.ResponseWriter,
		`<html>
		<body>
		<h1>Enter Profile</h1>
		<form method="post">
		    <label>Name:</label>
			<input type="text" name="Name"/>
			<label>Age:</label>
		    <input type="text" name="Age"/>
			<input type="Submit" />
		</form>
		</body>
		</html>`)
}
