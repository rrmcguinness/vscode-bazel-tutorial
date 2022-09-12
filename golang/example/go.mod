module example

go 1.19

require (
	github.com/google/uuid v1.3.0
	github.com/stretchr/testify v1.8.0
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.0.0-20220909164309-bea034e7d591 // indirect
	golang.org/x/sys v0.0.0-20220909162455-aba9fc2a8ff2 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220909194730-69f6226f97e5 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	example/client => ./client
	example/model => ./model
	example/server => ./server
	example/service => ./service
)
