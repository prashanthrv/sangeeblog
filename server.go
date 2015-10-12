package main

import (
	"net/http"

	"github.com/prashanthrv/sangeeblog/Godeps/_workspace/src/github.com/golang/glog"
	"github.com/prashanthrv/sangeeblog/Godeps/_workspace/src/github.com/gorilla/context"

	"github.com/prashanthrv/sangeeblog/controllers"
	"github.com/prashanthrv/sangeeblog/system"

	"github.com/prashanthrv/sangeeblog/Godeps/_workspace/src/github.com/zenazn/goji"
	"github.com/prashanthrv/sangeeblog/Godeps/_workspace/src/github.com/zenazn/goji/graceful"
	"github.com/prashanthrv/sangeeblog/Godeps/_workspace/src/github.com/zenazn/goji/web"
	//"golang.org/x/net/http2"
)

func main() {
	// filename := flag.String("config", "config.toml", "Path to configuration file")
	//
	// flag.Parse()
	filename := "config.toml"
	defer glog.Flush()

	var application = &system.Application{}

	application.Init(&filename)
	application.LoadTemplates()

	// Setup static files
	static := web.New()
	publicPath := application.Config.Get("general.public_path").(string)
	static.Get("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir(publicPath))))

	http.Handle("/assets/", static)

	// Apply middleware
	goji.Use(application.ApplyTemplates)
	goji.Use(application.ApplySessions)
	//goji.Use(application.ApplyDbMap)
	goji.Use(application.ApplyGormDB)
	//goji.Use(application.ApplyAuth)
	goji.Use(application.ApplyIsXhr)
	goji.Use(application.ApplyCsrfProtection)
	goji.Use(context.ClearHandler)

	controller := &controllers.MainController{}

	// Couple of files - in the real world you would use nginx to serve them.
	goji.Get("/robots.txt", http.FileServer(http.Dir(publicPath)))
	goji.Get("/favicon.ico", http.FileServer(http.Dir(publicPath+"/images")))

	// Home page
	goji.Get("/", application.Route(controller, "Blog"))

	// Sign In routes
	goji.Get("/signin", application.Route(controller, "SignIn"))
	goji.Post("/signin", application.Route(controller, "SignInPost"))

	// Sign Up routes
	goji.Get("/signup", application.Route(controller, "SignUp"))
	goji.Post("/signup", application.Route(controller, "SignUpPost"))

	// KTHXBYE
	goji.Get("/logout", application.Route(controller, "Logout"))

	goji.Get("/blog", application.Route(controller, "Blog"))
	goji.Get("/post/:postid", application.Route(controller, "Post"))
	goji.Get("/category/:categoryid", application.Route(controller, "Categories"))
	goji.Get("/page/:pageid", application.Route(controller, "Pages"))

	graceful.PostHook(func() {
		application.Close()
	})
	goji.Serve()
	//http2.ConfigureServer(http, &http2.Server{})
}
