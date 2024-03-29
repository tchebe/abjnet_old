module github.com/tchebe/abjnet/user_cli

go 1.14

require (
	github.com/coreos/etcd v3.3.20+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/micro/go-micro/v2 v2.8.0
	github.com/tchebe/abjnet/user_service v0.0.0-20201013223527-4e595b60441a
	github.com/zjjt/abjnet/user_service v0.0.0-20200508133603-c1790a700d4e // indirect
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	google.golang.org/grpc v1.33.0 // indirect
	google.golang.org/grpc/examples v0.0.0-20201013205100-7745e521ff61 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace github.com/zjjt/abjnet/user_service => github.com/tchebe/abjnet/user_service v0.0.0-20200508133603-c1790a700d4e
