package client

import (
	"errors"

	appv1beta1 "github.com/kubernetes-sigs/application/pkg/apis/app/v1beta1"
)

// App is composed of a list of Kubernetes resources.
type App []KubernetesResource

// ExtractAppCrd returns the Application CRD from the resource list.
func (app *App) ExtractAppCrd() (*appv1beta1.Application, error) {
	for _, r := range *app {
		if r.GetKind() == "Application" {
			return r.ToAppCrd()
		}
	}

	return nil, errors.New("No Application CRD found")
}
