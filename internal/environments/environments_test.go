package environments

import (
	"testing"

	consulApi "github.com/hashicorp/consul/api"
	"github.com/stretchr/testify/assert"
	"github.com/trento-project/trento/internal/consul/mocks"
	"github.com/trento-project/trento/internal/hosts"
)

func TestNewEnvironment(t *testing.T) {
	name := "env1"
	env := NewEnvironment("env1")
	assert.Equal(t, env.Name, name)
	assert.NotNil(t, env.Landscapes)
}

func TestNewLandscape(t *testing.T) {
	name := "land1"
	land := NewLandscape("land1")
	assert.Equal(t, land.Name, name)
	assert.NotNil(t, land.SAPSystems)
}

func TestNewSystem(t *testing.T) {
	name := "sys1"

	sys := NewSystem("sys1", 0)
	assert.Equal(t, sys.Name, name)
}

func TestNewDefaultEnvironment(t *testing.T) {
	env := NewDefaultEnvironment()
	assert.Equal(t, env.Name, defaultName)
	assert.NotNil(t, env.Landscapes)
}

func TestNewDefaultLandscape(t *testing.T) {
	land := NewDefaultLandscape()
	assert.Equal(t, land.Name, defaultName)
	assert.NotNil(t, land.SAPSystems)
}

func TestEnvironmentHealth(t *testing.T) {
	consulInst := new(mocks.Client)
	health := new(mocks.Health)

	node1HealthChecks := consulApi.HealthChecks{
		&consulApi.HealthCheck{
			Status: consulApi.HealthPassing,
		},
	}

	node2HealthChecks := consulApi.HealthChecks{
		&consulApi.HealthCheck{
			Status: consulApi.HealthPassing,
		},
	}

	node3HealthChecks := consulApi.HealthChecks{
		&consulApi.HealthCheck{
			Status: consulApi.HealthCritical,
		},
	}

	node4HealthChecks := consulApi.HealthChecks{
		&consulApi.HealthCheck{
			Status: consulApi.HealthWarning,
		},
	}

	consulInst.On("Health").Return(health)
	health.On("Node", "node1", (*consulApi.QueryOptions)(nil)).Return(node1HealthChecks, nil, nil)
	health.On("Node", "node2", (*consulApi.QueryOptions)(nil)).Return(node2HealthChecks, nil, nil)
	health.On("Node", "node3", (*consulApi.QueryOptions)(nil)).Return(node3HealthChecks, nil, nil)
	health.On("Node", "node4", (*consulApi.QueryOptions)(nil)).Return(node4HealthChecks, nil, nil)

	node1 := consulApi.Node{
		Node: "node1",
		Meta: map[string]string{
			"meta1": "value1",
			"meta2": "value2",
		},
	}

	node2 := consulApi.Node{
		Node: "node2",
		Meta: map[string]string{
			"meta3": "value3",
			"meta4": "value4",
		},
	}

	node3 := consulApi.Node{
		Node: "node3",
		Meta: map[string]string{
			"meta5": "value5",
			"meta6": "value6",
		},
	}

	node4 := consulApi.Node{
		Node: "node4",
		Meta: map[string]string{
			"meta7": "value7",
			"meta8": "value8",
		},
	}

	host1 := hosts.NewHost(node1, consulInst)
	host2 := hosts.NewHost(node2, consulInst)
	host3 := hosts.NewHost(node3, consulInst)
	host4 := hosts.NewHost(node4, consulInst)

	e := Environment{
		Name: "env1",
		Landscapes: map[string]*Landscape{
			"land1": {
				Name: "land1",
				SAPSystems: map[string]*SAPSystem{
					"sys1": {
						Name: "sys1",
						Hosts: hosts.HostList{
							&host1,
							&host2,
						},
					},
				},
			},
			"land2": {
				Name: "land2",
				SAPSystems: map[string]*SAPSystem{
					"sys2": {
						Name: "sys2",
						Hosts: hosts.HostList{
							&host3,
							&host4,
						},
					},
				},
			},
		},
	}

	expectedHealth := &EnvironmentHealth{
		Health: consulApi.HealthCritical,
		HealthMap: map[string]string{
			"land1": consulApi.HealthPassing,
			"land2": consulApi.HealthCritical,
		},
	}

	h := e.Health()

	assert.Equal(t, h, *expectedHealth)
}

func TestLandscapeHealth(t *testing.T) {
	consulInst := new(mocks.Client)
	health := new(mocks.Health)

	node1HealthChecks := consulApi.HealthChecks{
		&consulApi.HealthCheck{
			Status: consulApi.HealthPassing,
		},
	}

	node2HealthChecks := consulApi.HealthChecks{
		&consulApi.HealthCheck{
			Status: consulApi.HealthPassing,
		},
	}

	node3HealthChecks := consulApi.HealthChecks{
		&consulApi.HealthCheck{
			Status: consulApi.HealthPassing,
		},
	}

	node4HealthChecks := consulApi.HealthChecks{
		&consulApi.HealthCheck{
			Status: consulApi.HealthWarning,
		},
	}

	consulInst.On("Health").Return(health)
	health.On("Node", "node1", (*consulApi.QueryOptions)(nil)).Return(node1HealthChecks, nil, nil)
	health.On("Node", "node2", (*consulApi.QueryOptions)(nil)).Return(node2HealthChecks, nil, nil)
	health.On("Node", "node3", (*consulApi.QueryOptions)(nil)).Return(node3HealthChecks, nil, nil)
	health.On("Node", "node4", (*consulApi.QueryOptions)(nil)).Return(node4HealthChecks, nil, nil)

	node1 := consulApi.Node{
		Node: "node1",
		Meta: map[string]string{
			"meta1": "value1",
			"meta2": "value2",
		},
	}

	node2 := consulApi.Node{
		Node: "node2",
		Meta: map[string]string{
			"meta3": "value3",
			"meta4": "value4",
		},
	}

	node3 := consulApi.Node{
		Node: "node3",
		Meta: map[string]string{
			"meta5": "value5",
			"meta6": "value6",
		},
	}

	node4 := consulApi.Node{
		Node: "node4",
		Meta: map[string]string{
			"meta7": "value7",
			"meta8": "value8",
		},
	}

	host1 := hosts.NewHost(node1, consulInst)
	host2 := hosts.NewHost(node2, consulInst)
	host3 := hosts.NewHost(node3, consulInst)
	host4 := hosts.NewHost(node4, consulInst)

	l := Landscape{
		Name: "land1",
		SAPSystems: map[string]*SAPSystem{
			"sys1": {
				Name: "sys1",
				Hosts: hosts.HostList{
					&host1,
					&host2,
				},
			},
			"sys2": {
				Name: "sys2",
				Hosts: hosts.HostList{
					&host3,
					&host4,
				},
			},
		},
	}

	expectedHealth := &EnvironmentHealth{
		Health: consulApi.HealthWarning,
		HealthMap: map[string]string{
			"sys1": consulApi.HealthPassing,
			"sys2": consulApi.HealthWarning,
		},
	}

	h := l.Health()

	assert.Equal(t, h, *expectedHealth)
}

func TestSAPSystemHealth(t *testing.T) {
	consulInst := new(mocks.Client)
	health := new(mocks.Health)

	node1HealthChecks := consulApi.HealthChecks{
		&consulApi.HealthCheck{
			Status: consulApi.HealthPassing,
		},
	}

	node2HealthChecks := consulApi.HealthChecks{
		&consulApi.HealthCheck{
			Status: consulApi.HealthWarning,
		},
	}

	node3HealthChecks := consulApi.HealthChecks{
		&consulApi.HealthCheck{
			Status: consulApi.HealthCritical,
		},
	}

	node4HealthChecks := consulApi.HealthChecks{
		&consulApi.HealthCheck{
			Status: consulApi.HealthWarning,
		},
	}

	consulInst.On("Health").Return(health)
	health.On("Node", "node1", (*consulApi.QueryOptions)(nil)).Return(node1HealthChecks, nil, nil)
	health.On("Node", "node2", (*consulApi.QueryOptions)(nil)).Return(node2HealthChecks, nil, nil)
	health.On("Node", "node3", (*consulApi.QueryOptions)(nil)).Return(node3HealthChecks, nil, nil)
	health.On("Node", "node4", (*consulApi.QueryOptions)(nil)).Return(node4HealthChecks, nil, nil)

	node1 := consulApi.Node{
		Node: "node1",
		Meta: map[string]string{
			"meta1": "value1",
			"meta2": "value2",
		},
	}

	node2 := consulApi.Node{
		Node: "node2",
		Meta: map[string]string{
			"meta3": "value3",
			"meta4": "value4",
		},
	}

	node3 := consulApi.Node{
		Node: "node3",
		Meta: map[string]string{
			"meta5": "value5",
			"meta6": "value6",
		},
	}

	node4 := consulApi.Node{
		Node: "node4",
		Meta: map[string]string{
			"meta7": "value7",
			"meta8": "value8",
		},
	}

	host1 := hosts.NewHost(node1, consulInst)
	host2 := hosts.NewHost(node2, consulInst)
	host3 := hosts.NewHost(node3, consulInst)
	host4 := hosts.NewHost(node4, consulInst)

	s := SAPSystem{
		Name: "sys1",
		Hosts: hosts.HostList{
			&host1,
			&host2,
			&host3,
			&host4,
		},
	}

	expectedHealth := &EnvironmentHealth{
		Health: consulApi.HealthCritical,
		HealthMap: map[string]string{
			"node1": consulApi.HealthPassing,
			"node2": consulApi.HealthWarning,
			"node3": consulApi.HealthCritical,
			"node4": consulApi.HealthWarning,
		},
	}

	h := s.Health()

	assert.Equal(t, h, *expectedHealth)
}
