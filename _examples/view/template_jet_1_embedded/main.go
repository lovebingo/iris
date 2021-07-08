// Package main shows how to use jet templates embedded in your application with ease using the Iris built-in Jet view engine.
// This example is a customized fork of https://github.com/CloudyKit/jet/tree/master/examples/asset_packaging, so you can
// notice the differences side by side. For example, you don't have to use any external package inside your application,
// Iris manually builds the template loader for binary data when Asset and AssetNames are available via tools like the go-bindata.
package main

import (
	"os"

	"github.com/kataras/iris/v12"
)

// $ go get -u github.com/go-bindata/go-bindata/...
// # OR: go get -u github.com/go-bindata/go-bindata/v3/go-bindata
// # to save it to your go.mod file
//
// $ go-bindata -fs -prefix "views" ./views/...
// $ go run .
//
// System files are not used, you can optionally delete the folder and run the example now.
func main() {
	app := iris.New()
	tmpl := iris.Jet(AssetFile(), ".jet")
	app.RegisterView(tmpl)

	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.jet")
	})

	addr := ":8080"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	app.Listen(addr)
}
