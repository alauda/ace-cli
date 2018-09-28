package client

import (
	"errors"

	appv1beta1 "github.com/kubernetes-sigs/application/pkg/apis/app/v1beta1"
	"k8s.io/api/apps/v1"
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

// ExtractDeployments returns the list of Deployments from the Application.
func (app *App) ExtractDeployments() ([]*v1.Deployment, error) {
	var deployments []*v1.Deployment

	for _, r := range *app {
		if r.GetKind() == "Deployment" {
			deployment, err := r.ToDeployment()
			if err != nil {
				return deployments, err
			}

			deployments = append(deployments, deployment)
		}
	}

	return deployments, nil
}
