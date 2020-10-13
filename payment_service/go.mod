module github.com/tchebe/abjnet/payment_service

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/gorm v1.9.12
	github.com/joho/godotenv v1.3.0
	github.com/micro/go-micro/v2 v2.8.0
	github.com/satori/go.uuid v1.2.0
	github.com/zjjt/abjnet/user_service v0.0.0-20200531124055-aab7d7162a78
)

replace github.com/zjjt/abjnet/user_service => github.com/tchebe/abjnet/user_service v0.0.0-20200531124055-aab7d7162a78