module github.com/tchebe/abjnet/email_service

go 1.14

require (
	github.com/360EntSecGroup-Skylar/excelize/v2 v2.2.0
	github.com/go-mail/mail v2.3.1+incompatible
	github.com/micro/go-micro/v2 v2.8.0
	github.com/tchebe/abjnet/payment_service v0.0.0-20201013223527-4e595b60441a
	github.com/tchebe/abjnet/prestation_service v0.0.0-20201014032030-3e65499d6157
	github.com/tchebe/abjnet/souscription_service v0.0.0-20201014032030-3e65499d6157
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/mail.v2 v2.3.1 // indirect
)

replace github.com/zjjt/abjnet/payment_service => github.com/tchebe/abjnet/payment_service v0.0.0-20200531233639-a49046ecb633

replace github.com/zjjt/abjnet/prestation_service => github.com/tchebe/abjnet/prestation_service v0.0.0-20200531233639-a49046ecb633

replace github.com/zjjt/abjnet/product_service => github.com/tchebe/abjnet/product_service v0.0.0-20200804064938-decd38f6bccb

replace github.com/zjjt/abjnet/souscription_service => github.com/tchebe/abjnet/souscription_service v0.0.0-20200531233639-a49046ecb633

replace github.com/zjjt/abjnet/user_service => github.com/tchebe/abjnet/user_service v0.0.0-20200531233639-a49046ecb633
