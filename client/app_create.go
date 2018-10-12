package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
	"k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// CreateAppData defines the request body for the CreateApp API.
type CreateAppData struct {
	Meta      AppMetaData                 `json:"resource"`
	Resources []unstructured.Unstructured `json:"kubernetes"`
}

// AppMetaData defines the metadata for the application used in the CreateApp API.
type AppMetaData struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// RunApp creates and runs a single-deployment application
func (client *Client) RunApp(cluster string, namespace string, name string) error {
	url := client.buildURL("v2", "kubernetes", "clusters/%s/applications", cluster)

	request, err := client.buildRunAppRequest(namespace, name)
	if err != nil {
		return err
	}

	response, err := request.Post(url)
	if err != nil {
		return err
	}

	err = response.CheckStatusCode()

	return err
}

func (client *Client) buildRunAppRequest(namespace string, name string) (*rest.Request, error) {
	request := rest.NewRequest(client.Token())

	deployment, err := makeDeployment(namespace, name)
	if err != nil {
		return nil, err
	}

	resource, err := NewKubernetesResourceFromDeployment(deployment)
	if err != nil {
		return nil, err
	}

	data := CreateAppData{
		Meta: AppMetaData{
			Name:      name,
			Namespace: namespace,
		},
		Resources: []unstructured.Unstructured{
			resource.Resource,
		},
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	request.SetBody(body)

	return request, nil
}

func makeDeployment(namespace string, name string) (*v1.Deployment, error) {
	replicas := int32(1)

	labels := make(map[string]string)
	labels["app"] = name

	podSpec, err := makePodSpec(name)
	if err != nil {
		return nil, err
	}

	deployment := v1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: v1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: *podSpec,
			},
		},
	}

	return &deployment, nil
}

func makePodSpec(name string) (*corev1.PodSpec, error) {
	spec := corev1.PodSpec{
		Containers: []corev1.Container{
			{
				Name:  name,
				Image: "nginx:1.7.9",
			},
		},
	}

	return &spec, nil
}
