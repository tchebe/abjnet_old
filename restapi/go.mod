module github.com/tchebe/abjnet/restapi

go 1.14

require (
	github.com/emicklei/go-restful/v3 v3.1.0
	github.com/gorilla/schema v1.1.0
	github.com/micro/go-micro/v2 v2.8.0
	github.com/tchebe/abjnet/payment_service v0.0.0-20201014032030-3e65499d6157
	github.com/tchebe/abjnet/prestation_service v0.0.0-20201014032030-3e65499d6157
	github.com/tchebe/abjnet/product_service v0.0.0-20201014032030-3e65499d6157
	github.com/tchebe/abjnet/souscription_service v0.0.0-20201014032030-3e65499d6157
	github.com/tchebe/abjnet/user_service v0.0.0-20201013223527-4e595b60441a
)

replace github.com/zjjt/abjnet/payment_service => github.com/tchebe/abjnet/payment_service v0.0.0-20200531141652-e91fe0365427

replace github.com/zjjt/abjnet/prestation_service => github.com/tchebe/abjnet/prestation_service v0.0.0-20200601212057-d89d44de0468

replace github.com/zjjt/abjnet/product_service => github.com/tchebe/abjnet/product_service v0.0.0-20200804064938-decd38f6bccb

replace github.com/zjjt/abjnet/souscription_service => github.com/tchebe/abjnet/souscription_service v0.0.0-20200531141652-e91fe0365427

replace github.com/zjjt/abjnet/user_service => github.com/tchebe/abjnet/user_service v0.0.0-20200531233639-a49046ecb633
