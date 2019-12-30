schema:
	go-bindata -ignore=\.go -pkg=schema -o=../../api/app/schema/bindata.go ../../api/app/schema/... ../../internal/graphql/schema/...

run:
	go run main.go

test:
	@read -p "Enter Function Name:" func; \
	go test -timeout 30s github.com/TylerGrey/studyhub/api/app/resolvers -v -run $$func