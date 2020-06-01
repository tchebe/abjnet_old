package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	restful "github.com/emicklei/go-restful/v3"
	"github.com/gorilla/schema"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/web"
	paymentP "github.com/zjjt/abjnet/payment_service/proto/payment"
	prestationP "github.com/zjjt/abjnet/prestation_service/proto/prestation"
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
	paymentC      paymentP.PaymentService
	prestationC   prestationP.PrestationService
	decoder       *schema.Decoder
)

type createUser struct {
	Nom      string
	Prenom   string
	Email    string
	Password string
}
type userdetails struct {
	Email    string
	Password string
}

func wrapReturns(vs ...interface{}) []interface{} {
	return vs
}

//CreateUser creates a user in db
func (s *Api) CreateUser(req *restful.Request, res *restful.Response) {
	err := req.Request.ParseForm()
	if err != nil {
		res.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	user := new(createUser)
	err = decoder.Decode(user, req.Request.PostForm)
	if err != nil {
		res.WriteError(http.StatusBadRequest, fmt.Errorf("Reverifiez les parametres de creation de l'utilisateur"))
		return
	}
	response, err := userC.Create(context.TODO(), &userP.User{Nom: user.Nom, Prenom: user.Prenom, Email: user.Email, Password: user.Password})
	if err != nil {
		res.WriteError(http.StatusBadRequest, errors.New("Impossible de creer cet utilisateur"))
	}
	res.WriteEntity(response)
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
		theerror := fmt.Sprintf("erreur pendant le decodage des parametres de login %v", err)
		log.Println(theerror)
		res.WriteError(http.StatusBadRequest, fmt.Errorf("Mauvais identifiants de connexion1"))
		return
	}
	response, err := userC.Auth(context.TODO(), &userP.User{Email: user.Email, Password: user.Password})
	if err != nil {
		theerror := fmt.Sprintf("Mauvais identifiants de connexion %v", err)
		log.Println(theerror)
		res.WriteError(http.StatusBadRequest, errors.New("Mauvais identifiants de connexion"))
	}
	res.WriteEntity(response)

}

//ListContracts affiche la liste des contrats
func (s *Api) ListContracts(req *restful.Request, res *restful.Response) {
	log.Println("attempting to fetch list of contracts via rest api")
	//extract the token from the headers
	var token string
	if os.Getenv("DISABLE_AUTH") != "true" {
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
		token = splitted[1]
		_, err := userC.ValidateToken(context.Background(), &userP.Token{
			Token: token,
		})
		if err != nil {
			res.WriteErrorString(http.StatusForbidden, "Not Authorized")
			return
		}
	}
	ctx := metadata.Set(context.Background(), "Token", token)
	log.Println("Authenticated with token ", token)
	response, err := productC.GetAll(ctx, &productP.Request{})
	if err != nil {
		theerror := fmt.Sprintf("Une erreur est survenue lors de la recuperation des produits %v", err)
		log.Println(theerror)
		res.WriteError(http.StatusBadRequest, errors.New("Une erreur est survenue lors de la recuperation des produits"))
	}
	res.WriteEntity(response)
}

//Cotisations affiche l'etat des cotisation
func (s *Api) Cotisations(req *restful.Request, res *restful.Response) {
	log.Println("attempting to fetch list of contracts via rest api")
	//extract the token from the headers
	var token string
	if os.Getenv("DISABLE_AUTH") != "true" {
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
		token = splitted[1]
		_, err := userC.ValidateToken(context.Background(), &userP.Token{
			Token: token,
		})
		if err != nil {
			res.WriteErrorString(http.StatusForbidden, "Not Authorized")
			return
		}
	}
	ctx := metadata.Set(context.Background(), "Token", token)
	log.Println("Authenticated with token ", token)
	err := req.Request.ParseForm()
	if err != nil {
		res.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	police := new(productP.Police)
	err = decoder.Decode(police, req.Request.PostForm)
	if err != nil {
		theerror := fmt.Sprintf("erreur pendant le decodage des parametres de Cotisations %v", err)
		log.Println(theerror)
		res.WriteError(http.StatusBadRequest, fmt.Errorf("Mauvais numero de police"))
		return
	}
	response, err := productC.GetCotisations(ctx, &productP.Police{Police: police.Police})
	if err != nil {
		theerror := fmt.Sprintf("Une erreur est survenue lors de la recuperation de l'etat des cotisation  %v", err)
		log.Println(theerror)
		res.WriteError(http.StatusBadRequest, errors.New("Une erreur est survenue lors de la recuperation des produits"))
	}
	res.WriteEntity(response)
}

//Souscrire creates a subscription in the system
func (s *Api) Souscrire(req *restful.Request, res *restful.Response) {
	log.Println("attempting to suscribe to a contract via rest api")
	//extract the token from the headers
	var token string
	if os.Getenv("DISABLE_AUTH") != "true" {
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
		token = splitted[1]
		_, err := userC.ValidateToken(context.Background(), &userP.Token{
			Token: token,
		})
		if err != nil {
			res.WriteErrorString(http.StatusForbidden, "Not Authorized")
			return
		}
	}
	ctx := metadata.Set(context.Background(), "Token", token)
	err := req.Request.ParseForm()
	if err != nil {
		res.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	sub := new(souscriptionP.Souscription)

	err = decoder.Decode(sub, req.Request.PostForm)
	if err != nil {
		theerror := fmt.Sprintf("une erreur est survenue lors de la souscription %v d", err)
		log.Println(theerror)
		res.WriteError(http.StatusBadRequest, errors.New(theerror))
		return
	}
	//les statuts peuvent etre CREE,TRAITEMENT
	sub.Etattraitement = "CREE"
	response, err := souscriptionC.Subscribe(ctx, sub)
	if err != nil {
		theerror := fmt.Sprintf("une erreur est survenue lors de la souscription %v s", err)
		log.Println(theerror)
		res.WriteError(http.StatusBadRequest, errors.New("Un probleme a été rencontré lors de la souscription"))
	}
	res.WriteEntity(response)
}

//Payer creates a payment in the system
func (s *Api) Payer(req *restful.Request, res *restful.Response) {
	log.Println("attempting to pay a premium via rest api")
	//extract the token from the headers
	var token string
	if os.Getenv("DISABLE_AUTH") != "true" {
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
		token = splitted[1]
		_, err := userC.ValidateToken(context.Background(), &userP.Token{
			Token: token,
		})
		if err != nil {
			res.WriteErrorString(http.StatusForbidden, "Not Authorized")
			return
		}
	}
	ctx := metadata.Set(context.Background(), "Token", token)
	err := req.Request.ParseForm()
	if err != nil {
		res.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	pay := new(paymentP.Payment)

	err = decoder.Decode(pay, req.Request.PostForm)
	log.Println("pay payload is ", pay)
	if err != nil {
		theerror := fmt.Sprintf("une erreur est survenue lors du paiement %v d", err)
		log.Println(theerror)
		res.WriteError(http.StatusBadRequest, errors.New(theerror))
		return
	}
	log.Printf("Pay response %v\n%v", wrapReturns(paymentC.Pay(ctx, pay))...)
	response, err := paymentC.Pay(ctx, pay)
	if err != nil {
		theerror := fmt.Sprintf("une erreur est survenue lors du paiement %v s", err)
		log.Println(theerror)
		res.WriteError(http.StatusBadRequest, errors.New("Un probleme a été rencontré lors du paiement"))
	}
	res.WriteEntity(response)
}

//ValeurRachat returns the rachat value
func (s *Api) ValeurRachat(req *restful.Request, res *restful.Response) {
	log.Println("attempting to make a prestation via rest api")
	//extract the token from the headers
	var token string
	if os.Getenv("DISABLE_AUTH") != "true" {
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
		token = splitted[1]
		_, err := userC.ValidateToken(context.Background(), &userP.Token{
			Token: token,
		})
		if err != nil {
			res.WriteErrorString(http.StatusForbidden, "Not Authorized")
			return
		}
	}
	ctx := metadata.Set(context.Background(), "Token", token)
	err := req.Request.ParseForm()
	if err != nil {
		res.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	police := new(prestationP.Request)

	err = decoder.Decode(police, req.Request.PostForm)
	if err != nil {
		theerror := fmt.Sprintf("une erreur est survenue lors de la consultation de la valeur de rachat %v d", err)
		log.Println(theerror)
		res.WriteError(http.StatusBadRequest, errors.New(theerror))
		return
	}

	response, err := prestationC.ValeurRachat(ctx, police)
	if err != nil {
		theerror := fmt.Sprintf("une erreur est survenue lors de la consultation de la valeur de rachat %v s", err)
		log.Println(theerror)
		res.WriteError(http.StatusBadRequest, errors.New("Un probleme a été rencontré lors de la consultation de la valeur de rachat"))
	}
	res.WriteEntity(response)
}

//Prester creates a prestation in the system
func (s *Api) Prester(req *restful.Request, res *restful.Response) {
	log.Println("attempting to make a prestation via rest api")
	var token string
	if os.Getenv("DISABLE_AUTH") != "true" {
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
		token = splitted[1]
		_, err := userC.ValidateToken(context.Background(), &userP.Token{
			Token: token,
		})
		if err != nil {
			res.WriteErrorString(http.StatusForbidden, "Not Authorized")
			return
		}
	}
	ctx := metadata.Set(context.Background(), "Token", token)
	err := req.Request.ParseForm()
	if err != nil {
		res.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	presta := new(prestationP.Prestation)

	err = decoder.Decode(presta, req.Request.PostForm)
	if err != nil {
		theerror := fmt.Sprintf("une erreur est survenue lors du paiement %v d", err)
		log.Println(theerror)
		res.WriteError(http.StatusBadRequest, errors.New(theerror))
		return
	}
	//les statuts peuvent etre CREE,TRAITEMENT
	presta.Etattraitement = "CREE"
	response, err := prestationC.Rachat(ctx, presta)
	if err != nil {
		theerror := fmt.Sprintf("une erreur est survenue lors de la demande de prestation %v s", err)
		log.Println(theerror)
		res.WriteError(http.StatusBadRequest, errors.New("Un probleme a été rencontré lors de la demande de prestation"))
	}
	res.WriteEntity(response)
}

// func jwtAuthentication(req *restful.Request, res *restful.Response, chain *restful.FilterChain){
// 	//extract the token from the headers
// 	token:=req.HeaderParameter("Authorization")
// 	if token == ""
// }
func init() {
	/*if os.Getenv("ENV") != "PROD" || os.Getenv("ENV") != "TEST" {
		if err := godotenv.Load("../.env"); err != nil {
			log.Fatalf("Couldnt load .env file %v", err)
		}
	}*/

}
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
	paymentC = paymentP.NewPaymentService("abjnet.service.payment", client.DefaultClient)
	prestationC = prestationP.NewPrestationService("abjnet.service.prestation", client.DefaultClient)

	//create RESTFUL handler
	decoder = schema.NewDecoder()
	api := new(Api)
	ws := new(restful.WebService)
	wc := restful.NewContainer()
	wc.EnableContentEncoding(true)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Path("/api")
	ws.Route(ws.POST("/createuser").Consumes("application/x-www-form-urlencoded").To(api.CreateUser))
	ws.Route(ws.POST("/login").Consumes("application/x-www-form-urlencoded").To(api.Login))
	ws.Route(ws.GET("/listeproduit").To(api.ListContracts))
	ws.Route(ws.POST("/souscription").Consumes("application/x-www-form-urlencoded").To(api.Souscrire))
	ws.Route(ws.POST("/cotisations").Consumes("application/x-www-form-urlencoded").To(api.Cotisations))
	ws.Route(ws.POST("/prime").Consumes("application/x-www-form-urlencoded").To(api.Payer))
	ws.Route(ws.POST("/valeurrachat").Consumes("application/x-www-form-urlencoded").To(api.ValeurRachat))
	ws.Route(ws.POST("/prestation").Consumes("application/x-www-form-urlencoded").To(api.Prester))
	wc.Add(ws)
	//register handler
	service.Handle("/", wc)
	//run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
