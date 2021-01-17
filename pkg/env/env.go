package env

import (
	"github.com/goccha/envar"
	"os"
	"strings"
)

const PrefixDefine = "_DEFINE_DEPLOY_ENV_"

func init() {
	prefix := len(PrefixDefine)
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, PrefixDefine) {
			name := env[prefix:]
			value := envar.Int(env)
			NewDeployEnv(name, value)
		}
	}
}

type DeployEnv int

var _nameMap = make(map[string]DeployEnv)
var _valueMap = make(map[DeployEnv]string)

func NewDeployEnv(name string, value int) DeployEnv {
	name = strings.ToLower(name)
	_nameMap[name] = DeployEnv(value)
	_valueMap[DeployEnv(value)] = name
	return DeployEnv(value)
}

const (
	Production  DeployEnv = 10000
	Demo        DeployEnv = 5000
	Staging     DeployEnv = 3000
	QA          DeployEnv = 2000
	Development DeployEnv = 1000
	Local       DeployEnv = 0
)

func GetDeployEnv() DeployEnv {
	v := envar.Get("DEPLOY_ENV").String("local")
	return deploymentOf(v)
}

func Is(deployEnv DeployEnv) bool {
	return GetDeployEnv() >= deployEnv
}

func deploymentOf(value string) DeployEnv {
	value = strings.ToLower(value)
	if v, ok := _nameMap[value]; ok {
		return v
	}
	switch value {
	case "development":
		return Development
	case "qa":
		return QA
	case "demo":
		return Demo
	case "staging":
		return Staging
	case "production":
		return Production
	default:
		return Local
	}
}
func (d DeployEnv) String() string {
	if v, ok := _valueMap[d]; ok {
		return v
	}
	switch d {
	case Development:
		return "development"
	case QA:
		return "qa"
	case Demo:
		return "demo"
	case Staging:
		return "staging"
	case Production:
		return "production"
	default:
		return "local"
	}
}
