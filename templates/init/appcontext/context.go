package appcontext

//List of consts containing the names of the available components in the Application Context - appcontext.Current (Add your component names here as constants)
const (
	Logger = "Logger"
)

//Component is the Base interface for all Components
type Component interface{}

//ApplicationContext is the type defining a map of Components
type ApplicationContext struct {
	components map[string]Component
}

//Current keeps all components available, initialized in the application startup
var Current ApplicationContext

//Add a component By Name
func (applicationContext *ApplicationContext) Add(componentName string, component Component) {
	applicationContext.components[componentName] = component
}

//Get a component By Name
func (applicationContext *ApplicationContext) Get(componentName string) Component {
	return applicationContext.components[componentName]
}

//Delete a component By Name
func (applicationContext *ApplicationContext) Delete(componentName string) {
	delete(applicationContext.components, componentName)
}

//Count returns the count of components registered
func (applicationContext *ApplicationContext) Count() int {
	return len(applicationContext.components)
}

//CreateApplicationContext creates a new ApplicationContext instance
func CreateApplicationContext() ApplicationContext {
	return ApplicationContext{components: make(map[string]Component)}
}

func init() {
	Current = CreateApplicationContext()
}
