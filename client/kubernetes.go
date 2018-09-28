package client

import (
	"encoding/json"
	"errors"

	appv1beta1 "github.com/kubernetes-sigs/application/pkg/apis/app/v1beta1"
	"k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// KubernetesResource represents a generic representation of the Kubernetes Resource.
type KubernetesResource struct {
	ResourceActions []string                  `json:"resource_actions"`
	Resource        unstructured.Unstructured `json:"kubernetes"`
}

// GetKind returns the kind of the Kubernetes Resource.
func (k *KubernetesResource) GetKind() string {
	return k.Resource.GetKind()
}

// ToAppCrd converts the generic Kubernetes Resource into an Application CRD.
func (k *KubernetesResource) ToAppCrd() (*appv1beta1.Application, error) {
	kind := k.GetKind()

	if kind != "Application" {
		return nil, errors.New("Resource is not a valid Application CRD")
	}

	data, err := k.Resource.MarshalJSON()
	if err != nil {
		return nil, err
	}

	var app appv1beta1.Application

	err = json.Unmarshal(data, &app)

	return &app, err
}

// ToDeployment converts the generic Kubernetes Resource into a Deployment.
func (k *KubernetesResource) ToDeployment() (*v1.Deployment, error) {
	kind := k.GetKind()

	if kind != "Deployment" {
		return nil, errors.New("Resource is not a valid Deployment")
	}

	data, err := k.Resource.MarshalJSON()
	if err != nil {
		return nil, err
	}

	var deployment v1.Deployment

	err = json.Unmarshal(data, &deployment)

	return &deployment, err
}
