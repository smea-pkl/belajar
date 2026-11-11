module pkl

go 1.26.0

tool (
	connectrpc.com/connect/cmd/protoc-gen-connect-go
	github.com/google/wire/cmd/wire
	google.golang.org/protobuf/cmd/protoc-gen-go
)

require (
	github.com/stretchr/testify v1.12.1
	gorm.io/gorm v1.30.0
)

require (
	connectrpc.com/connect v1.21.1 // indirect
	github.com/google/subcommands v1.2.0 // indirect
	github.com/google/wire v0.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/mod v0.20.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/tools v0.24.1 // indirect
	google.golang.org/protobuf v1.36.12 // indirect
)

require (
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	gorm.io/driver/sqlite v1.6.0
)
