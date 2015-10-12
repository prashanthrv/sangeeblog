package helpers

import (
	"bytes"
	"github.com/prashanthrv/sangeeblog/Godeps/_workspace/src/github.com/golang/glog"
	"github.com/prashanthrv/sangeeblog/Godeps/_workspace/src/github.com/pelletier/go-toml"
	"github.com/prashanthrv/sangeeblog/Godeps/_workspace/src/github.com/prashanthrv/raymond"
	"html/template"
	"io/ioutil"
	"path"
)

type HelperController struct {
	Config *toml.TomlTree
}

func (application *HelperController) Init(filename *string) {
	config, err := toml.LoadFile(*filename)
	if err != nil {
		glog.Fatalf("TOML load failed: %s\n", err)
	}
	application.Config = config
}

func Parse(t *template.Template, name string, data interface{}) string {
	var doc bytes.Buffer
	t.ExecuteTemplate(&doc, name, data)
	return doc.String()
}

func RenderFile(filename string, ctx interface{}) string {
	var application = HelperController{}
	configfilename := "config.toml"
	application.Init(&configfilename)
	template_path := application.Config.Get("general.template_path").(string)
	data, err := ioutil.ReadFile(path.Join(template_path, filename+".hbs"))
	if err != nil {
		glog.Warningf("Can't get template: %v", err)
	}
	tpl, err := raymond.Parse(string(data))
	if err != nil {
		panic(err)
	}
	result, err := tpl.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return result
}

func RenderFileInLayout(filename string, layout string, ctx interface{}) string {
	var application = HelperController{}
	configfilename := "config.toml"
	application.Init(&configfilename)
	template_path := application.Config.Get("general.template_path").(string)
	data, err := ioutil.ReadFile(path.Join(template_path, filename+".hbs"))
	if err != nil {
		glog.Warningf("Can't get template: %v", err)
	}
	tpl, err := raymond.Parse(string(data))
	if err != nil {
		panic(err)
	}
	result, err := tpl.Exec(ctx)
	if err != nil {
		panic(err)
	}
	allCtx, ok := ctx.(map[string]interface{})
	if ok == false {
		//allCtx := map[string]interface{}{}
		glog.Warningf("Empty context")
	}
	allCtx["body"] = result

	layoutData, err := ioutil.ReadFile(path.Join(template_path, layout+".hbs"))
	if err != nil {
		glog.Warningf("Can't get layout template: %v", err)
	}
	layoutTpl, err := raymond.Parse(string(layoutData))
	if err != nil {
		panic(err)
	}
	finalResult, err := layoutTpl.Exec(allCtx)
	if err != nil {
		panic(err)
	}
	return finalResult
}
