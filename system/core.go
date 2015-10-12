package system

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"crypto/sha256"

	"fmt"
	"github.com/prashanthrv/sangeeblog/Godeps/_workspace/src/github.com/golang/glog"
	"github.com/prashanthrv/sangeeblog/Godeps/_workspace/src/github.com/gorilla/sessions"
	"github.com/prashanthrv/sangeeblog/Godeps/_workspace/src/github.com/jinzhu/gorm"
	"github.com/prashanthrv/sangeeblog/Godeps/_workspace/src/github.com/pelletier/go-toml"
	"github.com/prashanthrv/sangeeblog/Godeps/_workspace/src/github.com/prashanthrv/raymond"
	"github.com/prashanthrv/sangeeblog/Godeps/_workspace/src/github.com/zenazn/goji/web"
	"github.com/prashanthrv/sangeeblog/Godeps/_workspace/src/gopkg.in/gorp.v1"
	"github.com/prashanthrv/sangeeblog/models"
	"time"
)

type CsrfProtection struct {
	Key    string
	Cookie string
	Header string
	Secure bool
}

type Application struct {
	Config         *toml.TomlTree
	Template       *template.Template
	Store          *sessions.CookieStore
	DbMap          *gorp.DbMap
	GormDB         *gorm.DB
	CsrfProtection *CsrfProtection
}

func (application *Application) Init(filename *string) {

	config, err := toml.LoadFile(*filename)
	if err != nil {
		glog.Fatalf("TOML load failed: %s\n", err)
	}

	hash := sha256.New()
	io.WriteString(hash, config.Get("cookie.mac_secret").(string))
	application.Store = sessions.NewCookieStore(hash.Sum(nil))
	application.Store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		Secure:   config.Get("cookie.secure").(bool),
	}
	//dbConfig := config.Get("database").(*toml.TomlTree)
	// application.DbMap = models.GetDbMap(
	// 	dbConfig.Get("user").(string),
	// 	dbConfig.Get("password").(string),
	// 	dbConfig.Get("hostname").(string),
	// 	dbConfig.Get("port").(string),
	// 	dbConfig.Get("database").(string))

	application.GormDB = models.GetDB("gohan")

	application.CsrfProtection = &CsrfProtection{
		Key:    config.Get("csrf.key").(string),
		Cookie: config.Get("csrf.cookie").(string),
		Header: config.Get("csrf.header").(string),
		Secure: config.Get("cookie.secure").(bool),
	}

	application.Config = config
	//Helpers register
	raymond.RegisterHelper("time", func(input, format string) raymond.SafeString {
		fmt.Println(input)
		t, err := time.Parse("2006-01-02 15:04:05.00000000 +0000 UTC", input)
		if err != nil {
			t, err = time.Parse("2006-01-02 15:04:05.000000000 +0000 UTC", input)
			if err != nil {
				panic(err)
			}
		}
		return raymond.SafeString(t.Format("January 2006"))
	})
}

func (application *Application) LoadTemplates() error {
	var templates []string

	fn := func(path string, f os.FileInfo, err error) error {
		if f.IsDir() != true && strings.HasSuffix(f.Name(), ".html") {
			templates = append(templates, path)
		}
		return nil
	}

	err := filepath.Walk(application.Config.Get("general.template_path").(string), fn)

	if err != nil {
		return err
	}

	application.Template = template.Must(template.ParseFiles(templates...))
	return nil
}

func (application *Application) Close() {
	glog.Info("Bye!")
}

func (application *Application) Route(controller interface{}, route string) interface{} {
	fn := func(c web.C, w http.ResponseWriter, r *http.Request) {
		c.Env["Content-Type"] = "text/html"

		methodValue := reflect.ValueOf(controller).MethodByName(route)
		methodInterface := methodValue.Interface()
		method := methodInterface.(func(c web.C, r *http.Request) (string, int))

		body, code := method(c, r)

		if session, exists := c.Env["Session"]; exists {
			err := session.(*sessions.Session).Save(r, w)
			if err != nil {
				glog.Errorf("Can't save session: %v", err)
			}
		}

		switch code {
		case http.StatusOK:
			if _, exists := c.Env["Content-Type"]; exists {
				w.Header().Set("Content-Type", c.Env["Content-Type"].(string))
			}
			io.WriteString(w, body)
		case http.StatusSeeOther, http.StatusFound:
			http.Redirect(w, r, body, code)
		}
	}
	return fn
}
