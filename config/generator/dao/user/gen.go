//go:generate go run .

package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type StructField struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Nullable bool   `yaml:"nullable"`
	Number   int    `yaml:"number"`
}

type StructInfo struct {
	Name    string                 `yaml:"name"`
	Package string                 `yaml:"package"`
	Fields  map[string]StructField `yaml:"structure"`
	Primary []string               `yaml:"primary"`
	Index   []string               `yaml:"index"`
}

type methodType struct {
	Script string
}

const daoTemplateCode = `
package {{.Package}}

import (
	"github.com/jinzhu/gorm"
	"github.com/game-core/gocrafter/config/database"
	"github.com/game-core/gocrafter/domain/entity/user/{{.Package}}"
	{{.RepositoryImportPath}}
)

type {{.Package}}Dao struct {
	Read  *gorm.DB
	Write *gorm.DB
}

func New{{.Name}}Dao(conn *database.SqlHandler) {{.Package}}Repository.{{.RepositoryInterface}} {
	return &{{.Package}}Dao{
		Read:  conn.User.ReadConn,
		Write: conn.User.WriteConn,
	}
}

{{range $methodName, $methodType := .Methods }}
	{{.Script}}
{{end}}
`

func generateDao(yamlFilePath string, outputBaseDir string) error {
	structInfo, err := getStructInfo(yamlFilePath)
	if err != nil {
		return err
	}

	outputDir := filepath.Join(outputBaseDir, structInfo.Package)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating output directory %s: %v", outputDir, err)
	}

	outputFileName := filepath.Join(outputDir, fmt.Sprintf("%s_dao.gen.go", structInfo.Package))
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("outputFileName file %s create error: %v", outputFileName, err)
	}
	defer outputFile.Close()

	methods := generateMethods(structInfo)

	daoTmpl, err := template.New("daoTemplate").Parse(daoTemplateCode)
	if err != nil {
		return fmt.Errorf("error parsing DAO template: %v", err)
	}

	if err := daoTmpl.ExecuteTemplate(outputFile, "daoTemplate", struct {
		Name                 string
		Package              string
		Methods              map[string]methodType
		RepositoryImportPath string
		RepositoryInterface  string
	}{
		Name:                 structInfo.Name,
		Package:              structInfo.Package,
		Methods:              methods,
		RepositoryImportPath: fmt.Sprintf("%s \"github.com/game-core/gocrafter/domain/repository/user/%s\"", structInfo.Package+"Repository", structInfo.Package),
		RepositoryInterface:  fmt.Sprintf("%sRepository", structInfo.Name),
	}); err != nil {
		return fmt.Errorf("template error: %v", err)
	}

	fmt.Printf("Created %s Dao in %s\n", structInfo.Name, outputFileName)

	return nil
}

func generateMethods(structInfo *StructInfo) map[string]methodType {
	methods := make(map[string]methodType)

	// FindByID
	if len(structInfo.Primary) > 0 {
		methods["FindByID"] = methodType{
			Script: generateFindByID(structInfo),
		}
	}

	// FindByIndex
	for _, index := range structInfo.Index {
		indexFields := strings.Split(index, ",")
		methods[fmt.Sprintf("FindBy%s", strings.Join(indexFields, "And"))] = methodType{
			Script: generateFindByIndex(structInfo, indexFields),
		}
	}

	// List
	methods["List"] = methodType{
		Script: generateList(structInfo),
	}

	// Create
	methods["Create"] = methodType{
		Script: generateCreate(structInfo),
	}

	// Update
	methods["Update"] = methodType{
		Script: generateUpdate(structInfo),
	}

	// Delete
	methods["Delete"] = methodType{
		Script: generateDelete(structInfo),
	}

	return methods
}

func generateFindByID(structInfo *StructInfo) string {
	return fmt.Sprintf(
		`func (d *%sDao) FindByID(ID int64) (*%s.%s, error) {
			entity := &%s.%s{}
			res := d.Read.Where("id = ?", ID).Find(entity)
			if err := res.Error; err != nil {
				return nil, err
			}
		
			return entity, nil
		}
		`,
		structInfo.Package,
		structInfo.Package,
		structInfo.Name,
		structInfo.Package,
		structInfo.Name,
	)
}

func generateFindByIndex(structInfo *StructInfo, indexFields []string) string {
	params := make([]struct{ Name, Type string }, len(indexFields))
	paramStrings := make([]string, len(indexFields))
	scriptStrings := make([]string, len(indexFields))

	for i, field := range indexFields {
		params[i] = struct{ Name, Type string }{field, structInfo.Fields[field].Type}
		paramStrings[i] = fmt.Sprintf("%s %s", field, structInfo.Fields[field].Type)
		scriptStrings[i] = fmt.Sprintf("Where(\"%s = ?\", %s)", structInfo.Fields[field].Name, field)
	}

	return fmt.Sprintf(
		`func (d *%sDao) FindBy%s(%s) (*%s.%s, error) {
			entity := &%s.%s{}
			res := d.Read.%s.Find(entity)
			if err := res.Error; err != nil {
				return nil, err
			}
		
			return entity, nil
		}
		`,
		structInfo.Package,
		strings.Join(indexFields, "And"),
		strings.Join(paramStrings, ","),
		structInfo.Package,
		structInfo.Name,
		structInfo.Package,
		structInfo.Name,
		strings.Join(scriptStrings, "."),
	)
}

func generateList(structInfo *StructInfo) string {
	return fmt.Sprintf(
		`func (d *%sDao) List(limit int64) (*%s.%ss, error) {
			entity := &%s.%ss{}
			res := d.Read.Limit(limit).Find(entity)
			if err := res.Error; err != nil {
				return nil, err
			}
		
			return entity, nil
		}
		`,
		structInfo.Package,
		structInfo.Package,
		structInfo.Name,
		structInfo.Package,
		structInfo.Name,
	)
}

func generateCreate(structInfo *StructInfo) string {
	return fmt.Sprintf(
		`func (d *%sDao) Create(entity *%s.%s, tx *gorm.DB) (*%s.%s, error) {
			var conn *gorm.DB
			if tx != nil {
				conn = tx
			} else {
				conn = d.Write
			}
		
			res := conn.Model(&%s.%s{}).Create(entity)
			if err := res.Error; err != nil {
				return nil, err
			}
		
			return entity, nil
		}
		`,
		structInfo.Package,
		structInfo.Package,
		structInfo.Name,
		structInfo.Package,
		structInfo.Name,
		structInfo.Package,
		structInfo.Name,
	)
}

func generateUpdate(structInfo *StructInfo) string {
	return fmt.Sprintf(
		`func (d *%sDao) Update(entity *%s.%s, tx *gorm.DB) (*%s.%s, error) {
			var conn *gorm.DB
			if tx != nil {
				conn = tx
			} else {
				conn = d.Write
			}
		
			res := conn.Model(&%s.%s{}).Where("id = ?", entity.ID).Update(entity)
			if err := res.Error; err != nil {
				return nil, err
			}
		
			return entity, nil
		}
		`,
		structInfo.Package,
		structInfo.Package,
		structInfo.Name,
		structInfo.Package,
		structInfo.Name,
		structInfo.Package,
		structInfo.Name,
	)
}

func generateDelete(structInfo *StructInfo) string {
	return fmt.Sprintf(
		`func (d *%sDao) Delete(entity *%s.%s, tx *gorm.DB) error {
			var conn *gorm.DB
			if tx != nil {
				conn = tx
			} else {
				conn = d.Write
			}
		
			res := conn.Model(&%s.%s{}).Where("id = ?", entity.ID).Delete(entity)
			if err := res.Error; err != nil {
				return err
			}
		
			return nil
		}
		`,
		structInfo.Package,
		structInfo.Package,
		structInfo.Name,
		structInfo.Package,
		structInfo.Name,
	)
}

func getStructInfo(yamlFilePath string) (*StructInfo, error) {
	yamlData, err := ioutil.ReadFile(yamlFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file %s: %v", yamlFilePath, err)
	}

	var structInfo StructInfo
	if err := yaml.Unmarshal(yamlData, &structInfo); err != nil {
		return nil, fmt.Errorf("error unmarshalling YAML in file %s: %v", yamlFilePath, err)
	}

	return &structInfo, nil
}

func main() {
	outputBaseDir := "../../../infra/dao/user"
	yamlFiles, err := filepath.Glob("../../../docs/entity/user/*.yaml")
	if err != nil {
		log.Fatalf("Error finding YAML files: %v", err)
	}

	for _, yamlFile := range yamlFiles {
		err := generateDao(yamlFile, outputBaseDir)
		if err != nil {
			log.Printf("Error generating dao from YAML file %s: %v", yamlFile, err)
		}
	}
}
