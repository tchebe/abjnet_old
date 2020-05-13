package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
	"github.com/gorilla/schema"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/web"
	pb "github.com/zjjt/abjnet/user_service/proto/user"
)

//User is a struct used in the rest api
type User struct{}

var (
	cl      pb.UserService
	decoder *schema.Decoder
)

type userdetails struct {
	Email    string
	Password string
}

//Login logs in the user and returns a token
func (s *User) Login(req *restful.Request, res *restful.Response) {
	log.Println("attempting to log in via rest")
	if err := req.Request.ParseForm(); err != nil {
		res.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	user := new(userdetails)
	if err := decoder.Decode(user, req.Request.PostForm); err != nil {
		log.Println("here l38")
		res.WriteError(http.StatusBadRequest, fmt.Errorf("Mauvais identifiants de connexion1"))
		return
	}
	response, err := cl.Auth(context.TODO(), &pb.User{Email: user.Email, Password: user.Password})
	if err != nil {
		log.Println("here l44")
		res.WriteError(http.StatusBadRequest, errors.New("Mauvais identifiants de connexion"))
	}
	res.WriteEntity(response)

}
func main() {
	//create rest service
	service := web.NewService(
		web.Name("abjnet.service.user"),
	)
	service.Init()
	//setup user server client
	cl = pb.NewUserService("abjnet.service.srv.user", client.DefaultClient)
	//create RESTFUL handler
	decoder = schema.NewDecoder()
	userd := new(User)
	ws := new(restful.WebService)
	wc := restful.NewContainer()
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Path("/user")
	ws.Route(ws.POST("/login").Consumes("application/x-www-form-urlencoded").To(userd.Login))
	wc.Add(ws)
	//register handler
	service.Handle("/", wc)
	//run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
