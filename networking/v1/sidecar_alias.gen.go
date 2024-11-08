// Code generated by protoc-gen-alias. DO NOT EDIT.
package v1

import "github.com/boanlab/api/networking/v1alpha3"

// `Sidecar` describes the configuration of the sidecar proxy that mediates
// inbound and outbound communication of the workload instance to which it is
// attached.
//
// <!-- crd generation tags
// +cue-gen:Sidecar:groupName:networking.istio.io
// +cue-gen:Sidecar:versions:v1beta1,v1alpha3,v1
// +cue-gen:Sidecar:annotations:helm.sh/resource-policy=keep
// +cue-gen:Sidecar:labels:app=istio-pilot,chart=istio,heritage=Tiller,release=istio
// +cue-gen:Sidecar:subresource:status
// +cue-gen:Sidecar:scope:Namespaced
// +cue-gen:Sidecar:resource:categories=istio-io,networking-istio-io
// +cue-gen:Sidecar:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=networking.istio.io/v1alpha3
// +genclient
// +k8s:deepcopy-gen=true
// -->
type Sidecar = v1alpha3.Sidecar

// `IstioIngressListener` specifies the properties of an inbound
// traffic listener on the sidecar proxy attached to a workload instance.
type IstioIngressListener = v1alpha3.IstioIngressListener

// `IstioEgressListener` specifies the properties of an outbound traffic
// listener on the sidecar proxy attached to a workload instance.
type IstioEgressListener = v1alpha3.IstioEgressListener

// `WorkloadSelector` specifies the criteria used to determine if the
// `Gateway`, `Sidecar`, `EnvoyFilter`, `ServiceEntry`, or `DestinationRule`
// configuration can be applied to a proxy. The matching criteria
// includes the metadata associated with a proxy, workload instance
// info such as labels attached to the pod/VM, or any other info that
// the proxy provides to Istio during the initial handshake. If
// multiple conditions are specified, all conditions need to match in
// order for the workload instance to be selected. Currently, only
// label based selection mechanism is supported.
type WorkloadSelector = v1alpha3.WorkloadSelector

// `OutboundTrafficPolicy` sets the default behavior of the sidecar for
// handling unknown outbound traffic from the application.
type OutboundTrafficPolicy = v1alpha3.OutboundTrafficPolicy
type OutboundTrafficPolicy_Mode = v1alpha3.OutboundTrafficPolicy_Mode

// In `REGISTRY_ONLY` mode, unknown outbound traffic will be dropped.
// Traffic destinations must be explicitly declared into the service registry through `ServiceEntry` configurations.
//
// Note: Istio [does not offer an outbound traffic security policy](https://istio.io/latest/docs/ops/best-practices/security/#understand-traffic-capture-limitations).
// This option does not act as one, or as any form of an outbound firewall.
// Instead, this option exists primarily to offer users a way to detect missing `ServiceEntry` configurations by explicitly failing.
const OutboundTrafficPolicy_REGISTRY_ONLY OutboundTrafficPolicy_Mode = v1alpha3.OutboundTrafficPolicy_REGISTRY_ONLY

// In `ALLOW_ANY` mode, any traffic to unknown destinations will be allowed.
// Unknown destination traffic will have limited functionality, however, such as reduced observability.
// This mode allows users that do not have all possible egress destinations registered through `ServiceEntry` configurations to still connect
// to arbitrary destinations.
const OutboundTrafficPolicy_ALLOW_ANY OutboundTrafficPolicy_Mode = v1alpha3.OutboundTrafficPolicy_ALLOW_ANY

// Port describes the properties of a specific port of a service.
type SidecarPort = v1alpha3.SidecarPort

// `CaptureMode` describes how traffic to a listener is expected to be
// captured. Applicable only when the listener is bound to an IP.
type CaptureMode = v1alpha3.CaptureMode

// The default capture mode defined by the environment.
const CaptureMode_DEFAULT CaptureMode = v1alpha3.CaptureMode_DEFAULT

// Capture traffic using IPtables redirection.
const CaptureMode_IPTABLES CaptureMode = v1alpha3.CaptureMode_IPTABLES

// No traffic capture. When used in an egress listener, the application is
// expected to explicitly communicate with the listener port or Unix
// domain socket. When used in an ingress listener, care needs to be taken
// to ensure that the listener port is not in use by other processes on
// the host.
const CaptureMode_NONE CaptureMode = v1alpha3.CaptureMode_NONE
