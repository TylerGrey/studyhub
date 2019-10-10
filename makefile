schema:
	go-bindata -ignore=\.go -pkg=schema -o=api/schema/bindata.go api/schema/...

run:
	echo $(cat configs/.env.local)
	go run cmd/hub-api/main.go