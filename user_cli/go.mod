module github.com/tchebe/abjnet/user_cli

go 1.14

require (
	github.com/coreos/etcd v3.3.20+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/micro/go-micro/v2 v2.6.0
	replace github.com/zjjt/abjnet/user_service v0.0.0-20200508133603-c1790a700d4e
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/tools v0.0.0-20200507205054-480da3ebd79c // indirect
	google.golang.org/genproto v0.0.0-20200507105951-43844f6eee31 // indirect
	google.golang.org/grpc v1.29.1 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
	honnef.co/go/tools v0.0.1-2020.1.3 // indirect
)

replace github.com/zjjt/abjnet/user_service => github.com/tchebe/abjnet/user_service v0.0.0-20200508133603-c1790a700d4e
