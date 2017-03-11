package vtm

import (
	// "encoding/json"
	"fmt"
)

// Pools is a collection of pools
type Pools struct {
	Pools []Pool `json:"children"`
}

type Pool struct {
	Name string `json:"name"`
	Href string `json:"href"`
}

// ListPools retrieves an array of the pool names currently registered in stingray
func (r *marathonClient) ListPools() ([]string, error) {
	pools := new(Pools)
	err := r.apiGet(marathonAPIPools, nil, pools)
	if err != nil {
		return nil, err
	}

	fmt.Println(pools)

	var list []string
	for _, pool := range pools.Pools {
		list = append(list, pool.Name)
	}

	return list, nil
}

// Application is the definition for an application in marathon
// type Application struct {
// 	ID                         string              `json:"id,omitempty"`
// 	Cmd                        *string             `json:"cmd,omitempty"`
// 	Args                       *[]string           `json:"args,omitempty"`
// 	Constraints                *[][]string         `json:"constraints,omitempty"`
// 	CPUs                       float64             `json:"cpus,omitempty"`
// 	Disk                       *float64            `json:"disk,omitempty"`
// 	Env                        *map[string]string  `json:"env,omitempty"`
// 	Executor                   *string             `json:"executor,omitempty"`
// 	Instances                  *int                `json:"instances,omitempty"`
// 	Mem                        *float64            `json:"mem,omitempty"`
// 	Ports                      []int               `json:"ports"`
// 	RequirePorts               *bool               `json:"requirePorts,omitempty"`
// 	BackoffSeconds             *float64            `json:"backoffSeconds,omitempty"`
// 	BackoffFactor              *float64            `json:"backoffFactor,omitempty"`
// 	MaxLaunchDelaySeconds      *float64            `json:"maxLaunchDelaySeconds,omitempty"`
// 	TaskKillGracePeriodSeconds *float64            `json:"taskKillGracePeriodSeconds,omitempty"`
// 	Deployments                []map[string]string `json:"deployments,omitempty"`
// 	Dependencies               []string            `json:"dependencies"`
// 	TasksRunning               int                 `json:"tasksRunning,omitempty"`
// 	TasksStaged                int                 `json:"tasksStaged,omitempty"`
// 	TasksHealthy               int                 `json:"tasksHealthy,omitempty"`
// 	TasksUnhealthy             int                 `json:"tasksUnhealthy,omitempty"`
// 	User                       string              `json:"user,omitempty"`
// 	Uris                       *[]string           `json:"uris,omitempty"`
// 	Version                    string              `json:"version,omitempty"`
// 	Labels                     *map[string]string  `json:"labels,omitempty"`
// 	AcceptedResourceRoles      []string            `json:"acceptedResourceRoles,omitempty"`
// }

// NewDockerApplication creates a default docker application
// func NewDockerApplication() *Application {
// 	application := new(Application)
// 	return application
// }

// String returns the json representation of this application
// func (r *Application) String() string {
// 	s, err := json.MarshalIndent(r, "", "  ")
// 	if err != nil {
// 		return fmt.Sprintf(`{"error": "error decoding type into json: %s"}`, err)
// 	}
//
// 	return string(s)
// }

// Application retrieves the application configuration from marathon
// 		name: 		the id used to identify the application
// func (r *marathonClient) Application(name string) (*Application, error) {
// 	var wrapper struct {
// 		Application *Application `json:"app"`
// 	}
//
// 	if err := r.apiGet(buildURI(name), nil, &wrapper); err != nil {
// 		return nil, err
// 	}
//
// 	return wrapper.Application, nil
// }

// CreateApplication creates a new application in Marathon
// 		application:		the structure holding the application configuration
// func (r *marathonClient) CreateApplication(application *Application) (*Application, error) {
// 	result := new(Application)
// 	if err := r.apiPost(marathonAPIApps, application, result); err != nil {
// 		return nil, err
// 	}
//
// 	return result, nil
// }

// DeleteApplication deletes an application from marathon
// 		name: 		the id used to identify the application
//		force:		used to force the delete operation in case of blocked deployment
// func (r *marathonClient) DeleteApplication(name string, force bool) (*DeploymentID, error) {
// 	uri := buildURIWithForceParam(name, force)
// 	// step: check of the application already exists
// 	deployID := new(DeploymentID)
// 	if err := r.apiDelete(uri, nil, deployID); err != nil {
// 		return nil, err
// 	}
//
// 	return deployID, nil
// }

// UpdateApplication updates an application in Marathon
// 		application:		the structure holding the application configuration
// func (r *marathonClient) UpdateApplication(application *Application, force bool) (*DeploymentID, error) {
// 	result := new(DeploymentID)
// 	uri := buildURIWithForceParam(application.ID, force)
// 	if err := r.apiPut(uri, application, result); err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// func buildURI(path string) string {
// 	return fmt.Sprintf("%s/%s", marathonAPIApps, trimRootPath(path))
// }
