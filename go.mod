module github.com/tnnmuhandiram/terraform-gcp-poc

go 1.13

require (
	cloud.google.com/go v0.38.0
	github.com/gorilla/mux v1.7.3
	github.com/gruntwork-io/gruntwork-cli v0.5.1
	github.com/gruntwork-io/terratest v0.22.1
	github.com/hashicorp/vault/api v1.0.4
	github.com/jinzhu/gorm v1.9.16
	github.com/joho/godotenv v1.3.0
	github.com/jstemmer/go-junit-report v0.0.0-20190106144839-af01ea7f8024
	github.com/magiconair/properties v1.8.1
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.5.1
	github.com/tnnmuhandiram/terraform v0.0.0-20191024052749-3cb7c0818392
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	google.golang.org/api v0.9.1-0.20190821000710-329ecc3c9c34
)

replace github.com/satori/go.uuid v1.2.0 => github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
