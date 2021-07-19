package discovery

import (
	"fmt"

	"github.com/trento-project/trento/internal/cluster"
	"github.com/trento-project/trento/internal/consul"
	"github.com/trento-project/trento/internal/hosts"
)

const ClusterDiscoveryId string = "ha_cluster_discovery"

// This Discover handles any Pacemaker Cluster type
type ClusterDiscovery struct {
	id        string
	discovery BaseDiscovery
	Cluster   cluster.Cluster
}

func NewClusterDiscovery(client consul.Client) ClusterDiscovery {
	r := ClusterDiscovery{}
	r.id = ClusterDiscoveryId
	r.discovery = NewDiscovery(client)
	return r
}

func (c ClusterDiscovery) GetId() string {
	return c.id
}

// Execute one iteration of a discovery and store the result in the Consul KVStore.
func (d ClusterDiscovery) Discover() (string, error) {
	cluster, err := cluster.NewCluster()

	if err != nil {
		return fmt.Errorf("%v", err).Error(), err
	}

	d.Cluster = cluster

	err = d.Cluster.Store(d.discovery.client)
	if err != nil {
		return fmt.Errorf("%v", err).Error(), err
	}

	err = storeClusterMetadata(d.discovery.client, cluster.Name, cluster.Id)
	if err != nil {
		return fmt.Errorf("%v", err).Error(), err
	}

	return fmt.Sprintf("Cluster with name: %s successfully discovered", cluster.Name), nil
}

func storeClusterMetadata(client consul.Client, clusterName string, clusterId string) error {
	metadata := hosts.Metadata{
		Cluster:   clusterName,
		ClusterId: clusterId,
	}
	err := metadata.Store(client)
	if err != nil {
		return err
	}

	return nil
}
