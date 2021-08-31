module github.com/lambda-platform/dataform

go 1.15

require (
	github.com/PaesslerAG/gval v1.1.0
	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/iancoleman/strcase v0.1.3 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/labstack/echo/v4 v4.3.0
	github.com/lambda-platform/agent v0.1.13
	github.com/lambda-platform/lambda v0.0.1
	github.com/thedevsaddam/govalidator v1.9.10
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

//replace github.com/lambda-platform/lambda v0.0.1 => ../lambda
//replace github.com/lambda-platform/agent v0.0.1 => ../agent

