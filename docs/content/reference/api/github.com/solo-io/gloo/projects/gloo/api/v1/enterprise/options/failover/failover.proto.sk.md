
---
title: "failover.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `failover.options.gloo.solo.io` 
#### Types:


- [Failover](#failover)
- [Simple](#simple)
- [PrioritizedLocality](#prioritizedlocality)
- [Address](#address)
- [Endpoint](#endpoint)
- [HealthCheckConfig](#healthcheckconfig)
- [LbEndpoint](#lbendpoint)
- [LocalityLbEndpoints](#localitylbendpoints)
  



##### Source File: [github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/failover/failover.proto](https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/enterprise/options/failover/failover.proto)





---
### Failover



```yaml
"simple": .failover.options.gloo.solo.io.Simple

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `simple` | [.failover.options.gloo.solo.io.Simple](../failover.proto.sk/#simple) |  |  |




---
### Simple



```yaml
"locality": .envoy.config.core.v3.Locality
"prioritizedLocalities": []failover.options.gloo.solo.io.Simple.PrioritizedLocality

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `locality` | [.envoy.config.core.v3.Locality](../../../../../../../../../../../envoy/config/core/v3/base.proto.sk/#locality) | Identifies location of where the parent upstream hosts run. |  |
| `prioritizedLocalities` | [[]failover.options.gloo.solo.io.Simple.PrioritizedLocality](../failover.proto.sk/#prioritizedlocality) |  |  |




---
### PrioritizedLocality



```yaml
"endpoints": []failover.options.gloo.solo.io.LocalityLbEndpoints

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `endpoints` | [[]failover.options.gloo.solo.io.LocalityLbEndpoints](../failover.proto.sk/#localitylbendpoints) |  |  |




---
### Address

 
Represents a single instance of an upstream

```yaml
"addr": string
"port": int

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `addr` | `string` | Address (hostname or IP). |  |
| `port` | `int` | Port the instance is listening on. |  |




---
### Endpoint

 
Upstream host identifier.

```yaml
"address": .failover.options.gloo.solo.io.Address
"healthCheckConfig": .failover.options.gloo.solo.io.Endpoint.HealthCheckConfig
"hostname": string
"upstreamSslConfig": .core.gloo.solo.io.UpstreamSslConfig

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `address` | [.failover.options.gloo.solo.io.Address](../failover.proto.sk/#address) | The upstream host address. .. attention:: The form of host address depends on the given cluster type. For STATIC or EDS, it is expected to be a direct IP address (or something resolvable by the specified :ref:`resolver <envoy_api_field_config.core.v3.SocketAddress.resolver_name>` in the Address). For LOGICAL or STRICT DNS, it is expected to be hostname, and will be resolved via DNS. |  |
| `healthCheckConfig` | [.failover.options.gloo.solo.io.Endpoint.HealthCheckConfig](../failover.proto.sk/#healthcheckconfig) | The optional health check configuration is used as configuration for the health checker to contact the health checked host. .. attention:: This takes into effect only for upstream clusters with :ref:`active health checking <arch_overview_health_checking>` enabled. |  |
| `hostname` | `string` | The hostname associated with this endpoint. This hostname is not used for routing or address resolution. If provided, it will be associated with the endpoint, and can be used for features that require a hostname, like :ref:`auto_host_rewrite <envoy_api_field_config.route.v3.RouteAction.auto_host_rewrite>`. |  |
| `upstreamSslConfig` | [.core.gloo.solo.io.UpstreamSslConfig](../../../../core/ssl.proto.sk/#upstreamsslconfig) |  |  |




---
### HealthCheckConfig

 
The optional health check configuration.

```yaml
"portValue": int
"hostname": string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `portValue` | `int` | Optional alternative health check port value. By default the health check address port of an upstream host is the same as the host's serving address port. This provides an alternative health check port. Setting this with a non-zero value allows an upstream host to have different health check address port. |  |
| `hostname` | `string` | By default, the host header for L7 health checks is controlled by cluster level configuration (see: :ref:`host <envoy_api_field_config.core.v3.HealthCheck.HttpHealthCheck.host>` and :ref:`authority <envoy_api_field_config.core.v3.HealthCheck.GrpcHealthCheck.authority>`). Setting this to a non-empty value allows overriding the cluster level configuration for a specific endpoint. |  |




---
### LbEndpoint

 
An Endpoint that Envoy can route traffic to.
[#next-free-field: 6]

```yaml
"endpoint": .failover.options.gloo.solo.io.Endpoint
"loadBalancingWeight": .google.protobuf.UInt32Value

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `endpoint` | [.failover.options.gloo.solo.io.Endpoint](../failover.proto.sk/#endpoint) |  |  |
| `loadBalancingWeight` | [.google.protobuf.UInt32Value](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/u-int-32-value) | The optional load balancing weight of the upstream host; at least 1. Envoy uses the load balancing weight in some of the built in load balancers. The load balancing weight for an endpoint is divided by the sum of the weights of all endpoints in the endpoint's locality to produce a percentage of traffic for the endpoint. This percentage is then further weighted by the endpoint's locality's load balancing weight from LocalityLbEndpoints. If unspecified, each host is presumed to have equal weight in a locality. |  |




---
### LocalityLbEndpoints

 
A group of endpoints belonging to a Locality.
One can have multiple LocalityLbEndpoints for a locality, but this is
generally only done if the different groups need to have different load
balancing weights or different priorities.

```yaml
"locality": .envoy.config.core.v3.Locality
"lbEndpoints": []failover.options.gloo.solo.io.LbEndpoint
"loadBalancingWeight": .google.protobuf.UInt32Value

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `locality` | [.envoy.config.core.v3.Locality](../../../../../../../../../../../envoy/config/core/v3/base.proto.sk/#locality) | Identifies location of where the upstream hosts run. |  |
| `lbEndpoints` | [[]failover.options.gloo.solo.io.LbEndpoint](../failover.proto.sk/#lbendpoint) | The group of endpoints belonging to the locality specified. |  |
| `loadBalancingWeight` | [.google.protobuf.UInt32Value](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/u-int-32-value) | Optional: Per priority/region/zone/sub_zone weight; at least 1. The load balancing weight for a locality is divided by the sum of the weights of all localities at the same priority level to produce the effective percentage of traffic for the locality. Locality weights are only considered when :ref:`locality weighted load balancing <arch_overview_load_balancing_locality_weighted_lb>` is configured. These weights are ignored otherwise. If no weights are specified when locality weighted load balancing is enabled, the locality is assigned no load. |  |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->