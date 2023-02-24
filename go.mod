module github.com/tnnmuhandiram/terraform-gcp-poc

go 1.13

require (
	cloud.google.com/go v0.51.0 // indirect
	cloud.google.com/go/storage v1.0.0
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/gorilla/mux v1.7.3
	github.com/gruntwork-io/gruntwork-cli v0.5.1
	github.com/gruntwork-io/terratest v0.22.1
	github.com/hashicorp/vault/api v1.0.4
	github.com/jinzhu/gorm v1.9.16
	github.com/joho/godotenv v1.3.0
	github.com/jstemmer/go-junit-report v0.9.1
	github.com/kr/pretty v0.2.0 // indirect
	github.com/magiconair/properties v1.8.1
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.4.0
	github.com/tnnmuhandiram/terraform v0.0.0-20191024052749-3cb7c0818392
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
	golang.org/x/oauth2 v0.0.0-20191202225959-858c2ad4c8b6
	golang.org/x/sys v0.0.0-20200622214017-ed371f2e16b4 // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	google.golang.org/api v0.15.0
	google.golang.org/protobuf v1.24.0 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace github.com/satori/go.uuid v1.2.0 => github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
