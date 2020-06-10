// Copyright 2018 Envoyproxy Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package xds

import (
	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	listener "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	route "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	hcm "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"

	"github.com/envoyproxy/go-control-plane/pkg/conversion"
	"github.com/golang/protobuf/ptypes"
	"github.com/solo-io/solo-kit/pkg/api/v1/control-plane/cache"
	"github.com/solo-io/solo-kit/pkg/api/v1/control-plane/util"
)

type EnvoyResource struct {
	ProtoMessage cache.ResourceProto
}

var _ cache.Resource = &EnvoyResource{}

func NewEnvoyResource(r cache.ResourceProto) *EnvoyResource {
	return &EnvoyResource{ProtoMessage: r}
}

// Resource types in xDS v2.
const (
	typePrefix   = cache.TypePrefix + "/envoy.api.v2."
	EndpointType = typePrefix + "ClusterLoadAssignment"
	ClusterType  = typePrefix + "Cluster"
	RouteType    = typePrefix + "RouteConfiguration"
	ListenerType = typePrefix + "Listener"
)

var (
	// ResponseTypes are supported response types.
	ResponseTypes = []string{
		EndpointType,
		ClusterType,
		RouteType,
		ListenerType,
	}
)

func (e *EnvoyResource) Self() cache.XdsResourceReference {
	return cache.XdsResourceReference{
		Name: e.Name(),
		Type: e.Type(),
	}
}

// GetResourceName returns the resource name for a valid xDS response type.
func (e *EnvoyResource) Name() string {
	switch v := e.ProtoMessage.(type) {
	case *endpoint.ClusterLoadAssignment:
		return v.GetClusterName()
	case *cluster.Cluster:
		return v.GetName()
	case *route.RouteConfiguration:
		return v.GetName()
	case *listener.Listener:
		return v.GetName()
	// adding cases below as temporary solution to enable incremental changes
	case *v2.ClusterLoadAssignment:
		return v.GetClusterName()
	case *v2.Cluster:
		return v.GetName()
	case *v2.RouteConfiguration:
		return v.GetName()
	case *v2.Listener:
		return v.GetName()
	default:
		return ""
	}
}

func (e *EnvoyResource) ResourceProto() cache.ResourceProto {
	return e.ProtoMessage
}

func (e *EnvoyResource) Type() string {
	switch e.ProtoMessage.(type) {
	case *endpoint.ClusterLoadAssignment:
		return EndpointType
	case *cluster.Cluster:
		return ClusterType
	case *route.RouteConfiguration:
		return RouteType
	case *listener.Listener:
		return ListenerType
	default:
		return ""
	}
}

func (e *EnvoyResource) References() []cache.XdsResourceReference {
	out := make(map[cache.XdsResourceReference]bool)
	res := e.ProtoMessage
	if res == nil {
		return nil
	}
	switch v := res.(type) {
	case *endpoint.ClusterLoadAssignment:
		// no dependencies
	case *cluster.Cluster:
		// for EDS type, use cluster name or ServiceName override
		if v.GetType() == cluster.Cluster_EDS {
			rr := cache.XdsResourceReference{
				Type: EndpointType,
			}
			if v.EdsClusterConfig != nil && v.EdsClusterConfig.ServiceName != "" {
				rr.Name = v.EdsClusterConfig.ServiceName
			} else {
				rr.Name = v.Name
			}
			out[rr] = true
		}
	case *route.RouteConfiguration:
		// References to clusters in both routes (and listeners) are not included
		// in the result, because the clusters are retrieved in bulk currently,
		// and not by name.
	case *listener.Listener:
		// extract route configuration names from HTTP connection manager
		for _, chain := range v.FilterChains {
			for _, filter := range chain.Filters {
				if filter.Name != util.HTTPConnectionManager {
					continue
				}

				config := &hcm.HttpConnectionManager{}

				switch filterConfig := filter.ConfigType.(type) {
				case *listener.Filter_HiddenEnvoyDeprecatedConfig:
					if conversion.StructToMessage(filterConfig.HiddenEnvoyDeprecatedConfig, config) != nil {
						continue

					}
				case *listener.Filter_TypedConfig:
					if ptypes.UnmarshalAny(filterConfig.TypedConfig, config) != nil {
						continue
					}
				}

				if rds, ok := config.RouteSpecifier.(*hcm.HttpConnectionManager_Rds); ok && rds != nil && rds.Rds != nil {
					rr := cache.XdsResourceReference{
						Type: RouteType,
						Name: rds.Rds.RouteConfigName,
					}
					out[rr] = true
				}
			}
		}
	}

	var references []cache.XdsResourceReference
	for k, _ := range out {
		references = append(references, k)
	}
	return references
}

// GetResourceReferences returns the names for dependent resources (EDS cluster
// names for CDS, RDS routes names for LDS).
func GetResourceReferences(resources map[string]cache.Resource) map[string]bool {
	out := make(map[string]bool)
	for _, res := range resources {
		if res == nil {
			continue
		}
		switch v := res.ResourceProto().(type) {
		case *endpoint.ClusterLoadAssignment:
			// no dependencies
		case *cluster.Cluster:
			// for EDS type, use cluster name or ServiceName override
			if v.GetType() == cluster.Cluster_EDS {
				if v.EdsClusterConfig != nil && v.EdsClusterConfig.ServiceName != "" {
					out[v.EdsClusterConfig.ServiceName] = true
				} else {
					out[v.Name] = true
				}
			}
		case *route.RouteConfiguration:
			// References to clusters in both routes (and listeners) are not included
			// in the result, because the clusters are retrieved in bulk currently,
			// and not by name.
		case *listener.Listener:
			// extract route configuration names from HTTP connection manager
			for _, chain := range v.FilterChains {
				for _, filter := range chain.Filters {
					if filter.Name != util.HTTPConnectionManager {
						continue
					}

					config := &hcm.HttpConnectionManager{}

					switch filterConfig := filter.ConfigType.(type) {
					case *listener.Filter_HiddenEnvoyDeprecatedConfig:
						if conversion.StructToMessage(filterConfig.HiddenEnvoyDeprecatedConfig, config) != nil {
							continue

						}
					case *listener.Filter_TypedConfig:
						if ptypes.UnmarshalAny(filterConfig.TypedConfig, config) != nil {
							continue
						}
					}

					if rds, ok := config.RouteSpecifier.(*hcm.HttpConnectionManager_Rds); ok && rds != nil && rds.Rds != nil {
						out[rds.Rds.RouteConfigName] = true
					}

				}
			}
		}
	}
	return out
}

// GetResourceName returns the resource name for a valid xDS response type.
func GetResourceName(res cache.ResourceProto) string {
	switch v := res.(type) {
	case *endpoint.ClusterLoadAssignment:
		return v.GetClusterName()
	case *cluster.Cluster:
		return v.GetName()
	case *route.RouteConfiguration:
		return v.GetName()
	case *listener.Listener:
		return v.GetName()
	default:
		return ""
	}
}
