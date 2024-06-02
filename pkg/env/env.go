package env

import (
	"github.com/goccha/envar"
)

// DeployEnv represents the deployment environment.
// Deprecated
func NewDeployEnv(name string, value int) envar.DeployEnv {
	return envar.NewDeployEnv(name, value)
}

const (
	Production  = envar.Production
	Demo        = envar.Demo
	Staging     = envar.Staging
	QA          = envar.QA
	Development = envar.Development
	Local       = envar.Local
)

// GetDeployEnv returns the current deployment environment.
// Deprecated
func GetDeployEnv() envar.DeployEnv {
	return envar.GetDeployEnv()
}

// Is returns true if the current deployment environment is greater than or equal to the specified deployment environment.
// Deprecated
func Is(deployEnv envar.DeployEnv) bool {
	return GetDeployEnv() >= deployEnv
}
