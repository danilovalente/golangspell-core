package domain

import (
	"os"

	"github.com/danilovalente/golangspell-core/appcontext"
)

//Renderer defines the features delivered by the Code Template Renderer
type Renderer interface {
	//RenderFile renders a template file
	RenderFile(sourcePath string, info os.FileInfo) error

	//RenderPath renders an object (file os directory) in the templates directory
	RenderPath(sourcePath string, info os.FileInfo, err error) error

	//RenderTemplate renders all templates in the template directory providing the respective variables
	//commandName: specifies the name of the command for which the template will be rendered
	//globalVariables: defines the list of variables (value) which should be provided for rendering all files
	//specificVariables: defines the list of variables (value) which should be provided for rendering specific file names (key)
	RenderTemplate(commandName string, globalVariables map[string]interface{}, specificVariables map[string]map[string]interface{}) error
}

//GetRenderer returns the current component registered to provide the code rendering features
func GetRenderer() Renderer {
	return appcontext.Current.Get(appcontext.Renderer).(Renderer)
}