go test ./...

go test ./... -coverprofile=coverage.out

go tool cover -func=coverage.out 

mockery --dir=service --name=AuthInterface --filename=iauth.go --output=mocks/repomocks  
--outpkg=repomocks