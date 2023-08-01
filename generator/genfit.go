package generator

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ApiDef struct {
	Info struct {
		Description string `json:"description"`
		Version     string `json:"version"`
	} `json:"info"`

	Host string `json:"host"`

	Tags []struct {
		Name         string `json:"name"`
		Description  string `json:"description"`
		ExternalDocs struct {
			Description string `json:"description"`
			Url         string `json:"url"`
		} `json:"externalDocs"`
	} `json:"tags"`

	Paths map[string]struct {
		Get    RequestDef `json:"get"`
		Post   RequestDef `json:"post"`
		Delete RequestDef `json:"delete"`
	} `json:"paths"`

	Security struct {
		Oauth2 struct {
			Type             string            `json:"type"`
			AuthorizationUrl string            `json:"authorizationUrl"`
			Flow             string            `json:"flow"`
			Scopes           map[string]string `json:"scopes"`
		} `json:"oauth2"`
	} `json:"securityDefinitions"`
}

func (def *ApiDef) GenerateTag(f *jen.File, tag string) error {
	for url, reqDef := range def.Paths {
		if !Contains(reqDef.Get.Tags, tag) {
			continue
		}
		reqDef.Get.printSignature()
		reqDef.Get.generateFunc(f, url)
	}
	return nil
}

func (def *ApiDef) GenerateTagFiles() error {
	for _, tag := range def.Tags {
		f := jen.NewFile("gfit")
		f.HeaderComment(tag.Name)
		f.HeaderComment(tag.Description)
		f.HeaderComment(fmt.Sprintf("%s: %s", tag.ExternalDocs.Description, tag.ExternalDocs.Url))
		f.HeaderComment(fmt.Sprintf("code generated %s", time.Now().Format(time.DateTime)))
		if err := def.GenerateTag(f, tag.Name); err != nil {
			return err
		}
		fn := filepath.Join(GfitDir, fmt.Sprintf("%s.go", strcase.ToSnake(tag.Name)))
		file, err := os.Create(fn)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(file, "%#v", f)
		if err != nil {
			return err
		}
	}
	return nil
}

func (def *ApiDef) GenerateConstants() error {
	f := jen.NewFile("gfit")
	f.HeaderComment(fmt.Sprintf("auto generated fitbit client library: version %s", def.Info.Version))
	f.HeaderComment(def.Info.Description)
	f.HeaderComment(fmt.Sprintf("code generated %s", time.Now().Format(time.DateTime)))

	f.PackageComment("Package gfit")

	tokenUrl := ""
	for e, p := range def.Paths {
		if p.Post.OperationId != "oauthToken" {
			continue
		}
		tokenUrl = e
		break
	}

	f.Const().Defs(
		jen.Id("Host").Op("=").Lit(fmt.Sprintf("https://%s", def.Host)),
		jen.Id("AuthorizationUrl").Op("=").Lit(def.Security.Oauth2.AuthorizationUrl),
		jen.Id("TokenUrl").Op("=").Lit(tokenUrl),
	)
	f.Type().Id("AuthScope").String()
	f.Const().DefsFunc(func(g *jen.Group) {
		for scope, desc := range def.Security.Oauth2.Scopes {
			id := strcase.ToCamel(scope)
			g.Commentf("%s ->%s", id, desc)
			g.Id(id).Id("AuthScope").Op("=").Lit(scope)
		}
	})

	fn := filepath.Join(GfitDir, "gfit.go")
	file, err := os.Create(fn)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(file, "%#v", f)
	if err != nil {
		return err
	}
	return nil
}

type RequestDef struct {
	Tags        []string          `json:"tags"`
	Summary     string            `json:"summary"`
	Description string            `json:"description"`
	OperationId string            `json:"operationId"`
	Produces    []string          `json:"produces"`
	Parameters  []Parameter       `json:"parameters"`
	Responses   map[string]string `json:"responses"`
	Security    []struct {
		Oauth2 []string `json:"oauth2"`
	} `json:"security"`
}

func (r *RequestDef) FuncId() string {
	return strcase.ToCamel(r.OperationId)
}

func (r *RequestDef) defineUrl(url string) *jen.Statement {
	for _, p := range r.Parameters {
		url = p.ModifyUrl(url)
	}
	url = strings.Join([]string{"%s", url}, "")
	return jen.Id("url").
		Op(":=").
		Qual("fmt", "Sprintf").
		CallFunc(func(g *jen.Group) {
			g.Id("\n").Lit(url)
			g.Id("\n").Id("Host")
			for _, p := range r.Parameters {
				p.JenId(g)
			}
			g.Id("\n")
		})
}

func (r *RequestDef) generateFunc(f *jen.File, url string) {
	f.Comment(r.FuncId())
	f.Comment(r.Description)
	f.Func().Id(r.FuncId()).
		ParamsFunc(func(g *jen.Group) {
			g.Id("client").Op("*").Qual("net/http", "Client")
			for _, param := range r.Parameters {
				param.JenArg(g)
			}
		}).
		Params(jen.Index().Byte(), jen.Error()).
		Block(
			r.defineUrl(url),
			jen.Return(jen.Id("makeRequest").Call(jen.Id("client"), jen.Id("url"))),
		)
	f.Line()
}

func (r *RequestDef) printSignature() {
	s := jen.Func().Id(r.FuncId()).
		ParamsFunc(func(g *jen.Group) {
			g.Id("client").Op("*").Qual("net/http", "Client")
			for _, param := range r.Parameters {
				param.JenArg(g)
			}
		})
	fmt.Printf("%#v", s)
	fmt.Println()
}

type Parameter struct {
	Name        string   `json:"name"`
	In          string   `json:"in"` // "path", "query"
	Description string   `json:"description"`
	Required    bool     `json:"required"`
	Type        string   `json:"type"`   // "boolean", "integer", "string"
	Format      string   `json:"format"` // "date"
	Enum        []string `json:"enum"`
}

func (p Parameter) IdStr() string {
	return strcase.ToLowerCamel(p.Name)
}

func (p Parameter) JenId(g *jen.Group) {
	if p.Format == "date" {
		g.Id("\n").Id(p.IdStr()).Dot("Format").Call(jen.Lit("yyyy-MM-dd"))
	} else {
		g.Id("\n").Id(p.IdStr())
	}
}

func (p Parameter) JenType(s *jen.Statement) {
	switch p.Type {
	case "boolean":
		s.Bool()
	case "integer":
		s.Int()
	case "string":
		if p.Format == "date" {
			s.Qual("time", "Time")
		} else {
			s.String()
		}
	default:
		log.Fatalln(p.Type)
	}
}

func (p Parameter) JenArg(g *jen.Group) {
	s := g.Id(p.IdStr())
	switch p.Type {
	case "boolean":
		s.Bool()
	case "integer":
		s.Int()
	case "string":
		if p.Format == "date" {
			s.Qual("time", "Time")
		} else {
			s.String()
		}
	default:
		log.Fatalln(p.Type)
	}
}

func (p Parameter) FormatString() string {
	switch p.Type {
	case "boolean":
		return "%t"
	case "integer":
		return "%d"
	case "string":
		return "%s"
	default:
		log.Fatalln(p.Type)
		return "%s"
	}
}

func (p Parameter) ModifyUrl(url string) string {

	switch p.In {
	case "path":
		url = strings.Replace(
			url,
			strings.Join([]string{"{", p.Name, "}"}, ""),
			p.FormatString(),
			-1,
		)
	case "query":
		url = strings.Join(
			[]string{
				url,
				strings.Join([]string{"&", p.Name, "=", p.FormatString()}, ""),
			},
			"",
		)
	}
	return url
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

const ApiDefFile = "fitbit-web-api-swagger.json"
const GfitDir = "gfit"

func (def *ApiDef) Generate() error {
	def.GenerateConstants()
	def.GenerateTagFiles()
	return nil
}
