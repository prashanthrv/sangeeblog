package controllers

import (
	"net/http"

	//"github.com/golang/glog"

	"html/template"

	"github.com/prashanthrv/sangeeblog/helpers"
	"github.com/prashanthrv/sangeeblog/models"
	"github.com/prashanthrv/sangeeblog/system"
	"github.com/zenazn/goji/web"
//	"fmt"
)
type MainController struct {
	system.Controller
}

// Home page route
func (controller *MainController) Index(c web.C, r *http.Request) (string, int) {
	t := controller.GetTemplate(c)

	widgets := helpers.Parse(t, "home", nil)

	// With that kind of flags template can "figure out" what route is being rendered
	c.Env["IsIndex"] = true

	c.Env["Title"] = "Default Project - free Go website project template"
	c.Env["Content"] = template.HTML(widgets)

	return helpers.Parse(t, "main", c.Env), http.StatusOK
}

// Sign in route
// func (controller *MainController) SignIn(c web.C, r *http.Request) (string, int) {
// 	t := controller.GetTemplate(c)
// 	session := controller.GetSession(c)
//
// 	// With that kind of flags template can "figure out" what route is being rendered
// 	c.Env["IsSignIn"] = true
//
// 	c.Env["Flash"] = session.Flashes("auth")
// 	var widgets = controller.Parse(t, "auth/signin", c.Env)
//
// 	c.Env["Title"] = "Default Project - Sign In"
// 	c.Env["Content"] = template.HTML(widgets)
//
// 	return controller.Parse(t, "main", c.Env), http.StatusOK
// }
//
// // Sign In form submit route. Logs user in or set appropriate message in session if login was not succesful
// func (controller *MainController) SignInPost(c web.C, r *http.Request) (string, int) {
// 	email, password := r.FormValue("email"), r.FormValue("password")
//
// 	session := controller.GetSession(c)
// 	dbMap := controller.GetDbMap(c)
//
// 	user, err := helpers.Login(dbMap, email, password)
//
// 	if err != nil {
// 		session.AddFlash("Invalid Email or Password", "auth")
// 		return controller.SignIn(c, r)
// 	}
//
// 	session.Values["UserId"] = user.Id
//
// 	return "/", http.StatusSeeOther
// }
//
// // Sign up route
// func (controller *MainController) SignUp(c web.C, r *http.Request) (string, int) {
// 	t := controller.GetTemplate(c)
// 	session := controller.GetSession(c)
//
// 	// With that kind of flags template can "figure out" what route is being rendered
// 	c.Env["IsSignUp"] = true
//
// 	c.Env["Flash"] = session.Flashes("auth")
//
// 	var widgets = controller.Parse(t, "auth/signup", c.Env)
//
// 	c.Env["Title"] = "Default Project - Sign Up"
// 	c.Env["Content"] = template.HTML(widgets)
//
// 	return controller.Parse(t, "main", c.Env), http.StatusOK
// }
//
// // Sign Up form submit route. Registers new user or shows Sign Up route with appropriate messages set in session
// func (controller *MainController) SignUpPost(c web.C, r *http.Request) (string, int) {
// 	email, password := r.FormValue("email"), r.FormValue("password")
//
// 	session := controller.GetSession(c)
// 	dbMap := controller.GetDbMap(c)
//
// 	user := models.GetUserByEmail(dbMap, email)
//
// 	if user != nil {
// 		session.AddFlash("User exists", "auth")
// 		return controller.SignUp(c, r)
// 	}
//
// 	user = &models.User{
// 		Username: email,
// 		Email:    email,
// 	}
// 	user.HashPassword(password)
//
// 	if err := models.InsertUser(dbMap, user); err != nil {
// 		session.AddFlash("Error whilst registering user.")
// 		glog.Errorf("Error whilst registering user: %v", err)
// 		return controller.SignUp(c, r)
// 	}
//
// 	session.Values["UserId"] = user.Id
//
// 	return "/", http.StatusSeeOther
// }
//
// // This route logs user out
// func (controller *MainController) Logout(c web.C, r *http.Request) (string, int) {
// 	session := controller.GetSession(c)
//
// 	session.Values["UserId"] = nil
//
// 	return "/", http.StatusSeeOther
// }

func (controller *MainController) Blog(c web.C, r *http.Request) (string, int) {
		gormDB := controller.GetGormDB(c)
		posts := models.GetPosts(gormDB)
		categories := models.GetCategories(gormDB)
		pages := models.GetPages(gormDB)
    ctx := map[string]interface{}{
            "title": "Prashanth",
            "description":  "Im the One.",
						"bottomline": "Bottom Message",
						"posts": posts,
						"categories": categories,
						"pages": pages,
    			}
		finalHTML := helpers.RenderFileInLayout("Index","default",ctx)
		return finalHTML, http.StatusOK
}
func (controller *MainController) Post(c web.C, r *http.Request) (string, int) {
		postid := c.URLParams["postid"]
		gormDB := controller.GetGormDB(c)
		post := models.GetPost(gormDB,postid)
		categories := models.GetCategories(gormDB)
		pages := models.GetPages(gormDB)
    ctx := map[string]interface{}{
            "title": "Prashanth",
            "description":  "Im the One.",
						"bottomline": "Bottom Message",
						"post": post,
						"categories": categories,
						"pages": pages,
    			}
		finalHTML := helpers.RenderFileInLayout("Post","default",ctx)
		return finalHTML, http.StatusOK
}
func (controller *MainController) Categories(c web.C, r *http.Request) (string, int) {
		categoryid := c.URLParams["categoryid"]
		gormDB := controller.GetGormDB(c)
		posts := models.GetPostsByID(gormDB,categoryid)
		categories := models.GetCategories(gormDB)
		pages := models.GetPages(gormDB)
    ctx := map[string]interface{}{
            "title": "Prashanth",
            "description":  "Im the One.",
						"bottomline": "Bottom Message",
						"posts": posts,
						"categories": categories,
						"pages": pages,
    			}
		finalHTML := helpers.RenderFileInLayout("Index","default",ctx)
		return finalHTML, http.StatusOK
}
func (controller *MainController) Pages(c web.C, r *http.Request) (string, int) {
		pageid := c.URLParams["pageid"]
		gormDB := controller.GetGormDB(c)
		posts := models.GetPostsByPage(gormDB,pageid)
		categories := models.GetCategories(gormDB)
		pages := models.GetPages(gormDB)
    ctx := map[string]interface{}{
            "title": "Prashanth",
            "description":  "Im the One.",
						"bottomline": "Bottom Message",
						"posts": posts,
						"categories": categories,
						"pages": pages,
    			}
		finalHTML := helpers.RenderFileInLayout("Index","default",ctx)
		return finalHTML, http.StatusOK
}
