package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	restful "github.com/emicklei/go-restful/v3"
	"github.com/gorilla/schema"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/web"
	productP "github.com/zjjt/abjnet/product_service/proto/product"
	souscriptionP "github.com/zjjt/abjnet/souscription_service/proto/souscription"
	userP "github.com/zjjt/abjnet/user_service/proto/user"
)

//Api is a struct used in the rest api
type Api struct{}

var (
	userC         userP.UserService
	productC      productP.ProductService
	souscriptionC souscriptionP.SouscriptionService
	decoder       *schema.Decoder
)

type userdetails struct {
	Email    string
	Password string
}

//Login logs in the user and returns a token
func (s *Api) Login(req *restful.Request, res *restful.Response) {
	log.Println("attempting to log in via rest api")
	err := req.Request.ParseForm()
	if err != nil {
		res.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	user := new(userdetails)
	err = decoder.Decode(user, req.Request.PostForm)
	if err != nil {
		log.Println("here l38")
		res.WriteError(http.StatusBadRequest, fmt.Errorf("Mauvais identifiants de connexion1"))
		return
	}
	response, err := userC.Auth(context.TODO(), &userP.User{Email: user.Email, Password: user.Password})
	if err != nil {
		log.Println("here l44")
		res.WriteError(http.StatusBadRequest, errors.New("Mauvais identifiants de connexion"))
	}
	res.WriteEntity(response)

}
func (s *Api) ListContracts(req *restful.Request, res *restful.Response) {
	log.Println("attempting to fetch list of contracts via rest api")
	//extract the token from the headers
	tokenheader := req.HeaderParameter("Authorization")
	if tokenheader == "" {
		res.WriteErrorString(http.StatusForbidden, "Not Authorized")
		return
	}
	splitted := strings.Split(tokenheader, " ")
	if len(splitted) != 2 {
		res.WriteErrorString(http.StatusForbidden, "Not Authorized")
		return
	}
	token := splitted[1]
	_, err := userC.ValidateToken(context.Background(), &userP.Token{
		Token: token,
	})
	if err != nil {
		res.WriteErrorString(http.StatusForbidden, "Not Authorized")
		return
	}
	ctx := metadata.Set(context.Background(), "Token", token)
	log.Println("Authenticated with token ", token)
	response, err := productC.GetAll(ctx, &productP.Request{})
	if err != nil {
		theerror := fmt.Sprintf("Une erreur est survenue lors de la recuperation des produits %v", err)
		res.WriteError(http.StatusBadRequest, errors.New(theerror))
	}
	res.WriteEntity(response)
}

func (s *Api) Souscrire(req *restful.Request, res *restful.Response) {
	log.Println("attempting to suscribe to a contract via rest api")
	//extract the token from the headers
	tokenheader := req.HeaderParameter("Authorization")
	if tokenheader == "" {
		res.WriteErrorString(http.StatusForbidden, "Not Authorized")
		return
	}
	splitted := strings.Split(tokenheader, " ")
	if len(splitted) != 2 {
		res.WriteErrorString(http.StatusForbidden, "Not Authorized")
		return
	}
	token := splitted[1]
	_, err := userC.ValidateToken(context.Background(), &userP.Token{
		Token: token,
	})
	if err != nil {
		res.WriteErrorString(http.StatusForbidden, "Not Authorized")
		return
	}
	ctx := metadata.Set(context.Background(), "Token", token)
	err = req.Request.ParseForm()
	if err != nil {
		res.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	sub := new(souscriptionP.Souscription)
	err = decoder.Decode(sub, req.Request.PostForm)
	if err != nil {
		res.WriteError(http.StatusBadRequest, fmt.Errorf("une erreur est survenue lors de la souscription"))
		return
	}
	response, err := souscriptionC.Subscribe(ctx, sub)
	if err != nil {
		res.WriteError(http.StatusBadRequest, errors.New("Un probleme a été rencontré lors de la souscription"))
	}
	res.WriteEntity(response)
}

// func jwtAuthentication(req *restful.Request, res *restful.Response, chain *restful.FilterChain){
// 	//extract the token from the headers
// 	token:=req.HeaderParameter("Authorization")
// 	if token == ""
// }

func main() {
	//create rest service
	service := web.NewService(
		web.Name("abjnet.api.api"),
	)
	service.Init()
	//setup user,product and souscription server client
	userC = userP.NewUserService("abjnet.service.user", client.DefaultClient)
	productC = productP.NewProductService("abjnet.service.product", client.DefaultClient)
	souscriptionC = souscriptionP.NewSouscriptionService("abjnet.service.souscription", client.DefaultClient)
	//create RESTFUL handler
	decoder = schema.NewDecoder()
	api := new(Api)
	ws := new(restful.WebService)
	wc := restful.NewContainer()
	wc.EnableContentEncoding(true)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Path("/api")
	ws.Route(ws.POST("/login").Consumes("application/x-www-form-urlencoded").To(api.Login))
	ws.Route(ws.GET("/listeproduit").To(api.ListContracts))
	ws.Route(ws.POST("/souscription").Consumes("application/x-www-form-urlencoded").To(api.Souscrire))
	wc.Add(ws)
	//register handler
	service.Handle("/", wc)
	//run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
