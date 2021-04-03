.DEFAULT_GOAL := swagger

swagger_win:
		go get -u github.com/go-swagger/go-swagger/cmd/swagger
		swagger generate spec -o ./swagger.yaml --scan-models

swagger_osx:
		which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger
		GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models