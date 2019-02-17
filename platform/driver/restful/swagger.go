package restful

//go:generate go get -v github.com/go-swagger/go-swagger/cmd/swagger
//go:generate rm -rf ./models ./restapi/operations
//go:generate swagger generate server --target . --spec swagger.yml --exclude-main --principal models.Session
