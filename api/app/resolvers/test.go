package resolvers

import (
	"fmt"

	"github.com/TylerGrey/studyhub/api/app/resolvers/args"
)

// FileUpload ...
func (r *Resolver) FileUpload(args struct {
	Title string
	File  args.FileInput
}) (bool, error) {
	/*
		curl http://localhost:8080/graphql \
		-F operations='{ "query": "mutation fileUpload($file: Upload!, $title: String!) { fileUpload(file: $file, title: $title) }", "variables": { "file": null, "title": "TEST" } }' \
		-F map='{ "file": ["variables.file"] }' \
		-F file=@img_ci_01.png
	*/
	fmt.Println("성공!!!!")
	fmt.Println(args.File)
	fmt.Println(args.Title)

	return true, nil
}
