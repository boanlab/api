// Code generated by protoc-gen-alias. DO NOT EDIT.
package v1

import "github.com/boanlab/api/networking/v1alpha3"

// ServiceEntry enables adding additional entries into Istio's internal
// service registry.
//
// <!-- crd generation tags
// +cue-gen:ServiceEntry:groupName:networking.istio.io
// +cue-gen:ServiceEntry:versions:v1beta1,v1alpha3,v1
// +cue-gen:ServiceEntry:annotations:helm.sh/resource-policy=keep
// +cue-gen:ServiceEntry:labels:app=istio-pilot,chart=istio,heritage=Tiller,release=istio
// +cue-gen:ServiceEntry:subresource:status
// +cue-gen:ServiceEntry:scope:Namespaced
// +cue-gen:ServiceEntry:resource:categories=istio-io,networking-istio-io,shortNames=se,plural=serviceentries
// +cue-gen:ServiceEntry:printerColumn:name=Hosts,type=string,JSONPath=.spec.hosts,description="The hosts associated with the ServiceEntry"
// +cue-gen:ServiceEntry:printerColumn:name=Location,type=string,JSONPath=.spec.location,description="Whether the service is external to the
// mesh or part of the mesh (MESH_EXTERNAL or MESH_INTERNAL)"
// +cue-gen:ServiceEntry:printerColumn:name=Resolution,type=string,JSONPath=.spec.resolution,description="Service resolution mode for the hosts
// (NONE, STATIC, or DNS)"
// +cue-gen:ServiceEntry:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// +cue-gen:ServiceEntry:preserveUnknownFields:false
// +cue-gen:ServiceEntry:spec:required
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=networking.istio.io/v1alpha3
// +genclient
// +k8s:deepcopy-gen=true
// istiostatus-override: ServiceEntryStatus: github.com/boanlab/api/networking/v1alpha3
// -->
// +kubebuilder:validation:XValidation:message="only one of WorkloadSelector or Endpoints can be set",rule="(has(self.workloadSelector)?1:0)+(has(self.endpoints)?1:0)<=1"
// +kubebuilder:validation:XValidation:message="CIDR addresses are allowed only for NONE/STATIC resolution types",rule="!(has(self.addresses) && self.addresses.exists(k, k.contains('/')) && (has(self.resolution) && self.resolution != 'STATIC' && self.resolution != 'NONE'))"
// +kubebuilder:validation:XValidation:message="NONE mode cannot set endpoints",rule="(!has(self.resolution) || self.resolution == 'NONE') ? !has(self.endpoints) : true"
// +kubebuilder:validation:XValidation:message="DNS_ROUND_ROBIN mode cannot have multiple endpoints",rule="(has(self.resolution) && self.resolution == 'DNS_ROUND_ROBIN') ? (!has(self.endpoints) || size(self.endpoints) == 1) : true"
type ServiceEntry = v1alpha3.ServiceEntry

// Location specifies whether the service is part of Istio mesh or
// outside the mesh.  Location determines the behavior of several
// features, such as service-to-service mTLS authentication, policy
// enforcement, etc. When communicating with services outside the mesh,
// Istio's mTLS authentication is disabled, and policy enforcement is
// performed on the client-side as opposed to server-side.
type ServiceEntry_Location = v1alpha3.ServiceEntry_Location

// Signifies that the service is external to the mesh. Typically used
// to indicate external services consumed through APIs.
const ServiceEntry_MESH_EXTERNAL ServiceEntry_Location = v1alpha3.ServiceEntry_MESH_EXTERNAL

// Signifies that the service is part of the mesh. Typically used to
// indicate services added explicitly as part of expanding the service
// mesh to include unmanaged infrastructure (e.g., VMs added to a
// Kubernetes based service mesh).
const ServiceEntry_MESH_INTERNAL ServiceEntry_Location = v1alpha3.ServiceEntry_MESH_INTERNAL

// Resolution determines how the proxy will resolve the IP addresses of
// the network endpoints associated with the service, so that it can
// route to one of them. The resolution mode specified here has no impact
// on how the application resolves the IP address associated with the
// service. The application may still have to use DNS to resolve the
// service to an IP so that the outbound traffic can be captured by the
// Proxy. Alternatively, for HTTP services, the application could
// directly communicate with the proxy (e.g., by setting HTTP_PROXY) to
// talk to these services.
type ServiceEntry_Resolution = v1alpha3.ServiceEntry_Resolution

// Assume that incoming connections have already been resolved (to a
// specific destination IP address). Such connections are typically
// routed via the proxy using mechanisms such as IP table REDIRECT/
// eBPF. After performing any routing related transformations, the
// proxy will forward the connection to the IP address to which the
// connection was bound.
const ServiceEntry_NONE ServiceEntry_Resolution = v1alpha3.ServiceEntry_NONE

// Use the static IP addresses specified in endpoints (see below) as the
// backing instances associated with the service.
const ServiceEntry_STATIC ServiceEntry_Resolution = v1alpha3.ServiceEntry_STATIC

// Attempt to resolve the IP address by querying the ambient DNS,
// asynchronously. If no endpoints are specified, the proxy
// will resolve the DNS address specified in the hosts field, if
// wildcards are not used. If endpoints are specified, the DNS
// addresses specified in the endpoints will be resolved to determine
// the destination IP address.  DNS resolution cannot be used with Unix
// domain socket endpoints.
const ServiceEntry_DNS ServiceEntry_Resolution = v1alpha3.ServiceEntry_DNS

// Attempt to resolve the IP address by querying the ambient DNS,
// asynchronously. Unlike `DNS`, `DNS_ROUND_ROBIN` only uses the
// first IP address returned when a new connection needs to be initiated
// without relying on complete results of DNS resolution, and connections
// made to hosts will be retained even if DNS records change frequently
// eliminating draining connection pools and connection cycling.
// This is best suited for large web scale services that
// must be accessed via DNS. The proxy will resolve the DNS address
// specified in the hosts field, if wildcards are not used. DNS resolution
// cannot be used with Unix domain socket endpoints.
const ServiceEntry_DNS_ROUND_ROBIN ServiceEntry_Resolution = v1alpha3.ServiceEntry_DNS_ROUND_ROBIN

// ServicePort describes the properties of a specific port of a service.
type ServicePort = v1alpha3.ServicePort
type ServiceEntryStatus = v1alpha3.ServiceEntryStatus

// minor abstraction to allow for adding hostnames if relevant
type ServiceEntryAddress = v1alpha3.ServiceEntryAddress
