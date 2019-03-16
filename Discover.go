package standard

const RegistryConfigName = "Registry"
const AppConfigName = "App"
const WeightConfigName = "Weight"
const Calls = "Calls"

type Discover interface {
	IsServer() (isok bool)
	IsClient() (isok bool)
	Start(addr string, conf map[string]interface{}) (isok bool)
	Restart() (isok bool)
	Stop()
	Wait()
	AddExternalApp(app string, conf map[string]interface{}) (isok bool)
}

type Caller interface {
	Get(app, path string, headers ...string) (result []byte)
	Post(app, path string, data interface{}, headers ...string) (result []byte)
	Put(app, path string, data interface{}, headers ...string) (result []byte)
	Delete(app, path string, data interface{}, headers ...string) (result []byte)
	Head(app, path string, data interface{}, headers ...string) (result []byte)
	Do(method, app, path string, data interface{}, headers ...string) (result []byte)
	DoWithNode(method, app, withNode, path string, data interface{}, headers ...string) (result []byte, addr string)
}
