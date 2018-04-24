package client

// App defines the response body of the InspectApp API.
type App struct {
	Resource  AppResource  `json:"resource"`
	Cluster   AppCluster   `json:"cluster"`
	Namespace AppNamespace `json:"namespace"`
}

// AppResource contains the resource definition of an application.
type AppResource struct {
	ID          string `json:"uuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	State       string `json:"status"`
}

// AppCluster is the Kubernetes cluster in which the application is deployed.
type AppCluster struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// AppNamespace is the Kubernetes namespace in which the applicaiton is deployed.
type AppNamespace struct {
	ID   string `json:"uuid"`
	Name string `json:"name"`
}
