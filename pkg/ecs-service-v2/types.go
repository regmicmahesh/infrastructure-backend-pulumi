package ecsservicev2

type ContainerDefinitions []ContainerDefinition

type ContainerDefinition struct {
	Name             string           `json:"name"`
	Image            string           `json:"image"`
	Links            []string         `json:"links"`
	CPU              int              `json:"cpu"`
	PortMappings     []PortMapping    `json:"portMappings"`
	Essential        bool             `json:"essential"`
	Environment      []Environment    `json:"environment"`
	EnvironmentFiles []any            `json:"environmentFiles"`
	MountPoints      []MountPoint     `json:"mountPoints"`
	VolumesFrom      []any            `json:"volumesFrom"`
	Secrets          []Secret         `json:"secrets"`
	Ulimits          []any            `json:"ulimits"`
	LogConfiguration LogConfiguration `json:"logConfiguration"`
	HealthCheck      *HealthCheck     `json:"healthCheck"`
	Command          []string         `json:"command"`
	DependsOn        []DependsOn       `json:"dependsOn"`
}

type DependsOn struct {
	ContainerName string `json:"containerName"`
	Condition     string `json:"condition"`
}

type HealthCheck struct {
	Command []string `json:"command"`
}

type Environments []Environment

type Environment struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type MountPoints []MountPoint

type MountPoint struct {
	SourceVolume  string `json:"sourceVolume"`
	ContainerPath string `json:"containerPath"`
	ReadOnly      bool   `json:"readOnly"`
}

type PortMappings []PortMapping

type PortMapping struct {
	Name          string `json:"name"`
	ContainerPort int    `json:"containerPort"`
	HostPort      int    `json:"hostPort"`
	Protocol      string `json:"protocol"`
}

type Secrets []Secret

type Secret struct {
	Name      string `json:"name"`
	ValueFrom string `json:"valueFrom"`
}
type Options struct {
	AwslogsGroup        string `json:"awslogs-group"`
	AwslogsRegion       string `json:"awslogs-region"`
	AwslogsStreamPrefix string `json:"awslogs-stream-prefix"`
}
type LogConfiguration struct {
	LogDriver     string  `json:"logDriver"`
	Options       Options `json:"options"`
	SecretOptions []any   `json:"secretOptions"`
}
