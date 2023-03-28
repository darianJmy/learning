package main

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/go-openapi/spec"
	"log"
	"net/http"
)

type UserResource struct {
	users map[string]User
}

type User struct {
	ID   string `xml:"id" json:"id" description:"identifier of the user"`
	Name string `xml:"name" json:"name" description:"name of the user" default:"john"`
	Age  int    `xml:"age" json:"age" description:"age of the user" default:"21"`
}

func (u *UserResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	tags := []string{"users"}

	ws.Route(ws.GET("/").To(u.findAllUsers).
		Doc("get all users").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]User{}).
		Returns(200, "OK", []User{}))

	ws.Route(ws.GET("/{user-id}").To(u.findUser).
		Doc("get a user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(User{}).
		Returns(200, "OK", User{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{user-id}").To(u.updateUser).
		Doc("update a user").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(User{}))

	ws.Route(ws.PUT("").To(u.createUser).
		Doc("create a user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(User{}))

	ws.Route(ws.DELETE("/{user-id}").To(u.removeUser).
		Doc("delete a user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")))

	return ws
}

func (u *UserResource) findAllUsers(req *restful.Request, resp *restful.Response) {
	list := []User{}
	for _, each := range u.users {
		list = append(list, each)
	}
	resp.WriteEntity(list)
}

func (u *UserResource) findUser(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("user-id")
	usr := u.users[id]
	if len(usr.ID) == 0 {
		resp.WriteErrorString(404, "User could no be found.")
	} else {
		resp.WriteEntity(usr)
	}
}

func (u *UserResource) updateUser(req *restful.Request, resp *restful.Response) {
	usr := new(User)
	err := req.ReadEntity(&usr)
	if err == nil {
		u.users[usr.ID] = *usr
		resp.WriteEntity(usr)
	} else {
		resp.WriteError(500, err)
	}
}

func (u *UserResource) createUser(req *restful.Request, resp *restful.Response) {
	usr := User{ID: req.PathParameter("user-id")}
	err := req.ReadEntity(&usr)
	if err == nil {
		u.users[usr.ID] = usr
		resp.WriteHeaderAndEntity(201, usr)
	} else {
		resp.WriteError(500, err)
	}
}

func (u *UserResource) removeUser(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("user-id")
	delete(u.users, id)
}

func main() {
	u := UserResource{map[string]User{}}
	restful.DefaultContainer.Add(u.WebService())

	config := restfulspec.Config{
		WebServices:                   restful.RegisteredWebServices(),
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject,
	}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))

	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("/Users/emicklei/Projects/swagger-ui/dist"))))

	// Optionally, you may need to enable CORS for the UI to work.
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		CookiesAllowed: false,
		Container:      restful.DefaultContainer}
	restful.DefaultContainer.Filter(cors.Filter)
	log.Printf("Get the API using http://localhost:8080/apidocs.json")
	log.Printf("Open Swagger UI using http://localhost:8080/apidocs/?url=http://localhost:8080/apidocs.json")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "UserService",
			Description: "Resource for managing Users",
			Contact: &spec.ContactInfo{
				ContactInfoProps: spec.ContactInfoProps{
					Name:  "john",
					Email: "john@doe.rp",
					URL:   "http://johndoe.org",
				},
			},
			License: &spec.License{
				LicenseProps: spec.LicenseProps{
					Name: "MIT",
					URL:  "http://mit.org",
				},
			},
			Version: "1.0.0",
		},
	}
	swo.Tags = []spec.Tag{spec.Tag{TagProps: spec.TagProps{
		Name:        "users",
		Description: "Managing users"}}}
}
