GO111MODULE=off = gj get -u github.com/go-swagger/go-swagger/cmd/swagger
swagger:
	GO111MODULE=off swagger generate spec -o ./swagger2.yaml --scan-models