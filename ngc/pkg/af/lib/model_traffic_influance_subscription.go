// Copyright 2019 Intel Corporation and Smart-Edge.com, Inc. All rights reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package af

// SubscribedEvent : The possible value is CHANGE_OF_DNAI - the AF requests to
// be notified when the UP path changes for the PDU isession.
type SubscribedEvent string

// URI identifying the referenced resource
// URI string formatted accordingding to IETF RFC 3986
type Link string

// A string used to indicate the features supported by an API that is used
// (subclause 6.6 in 3GPP TS 29.500).
// The string shall contain a bitmask indicating supported features in
// hexadecimal representation.
// Each character in the string shall take a value of "0" to "9" or "A" to "F"
// and shall represent the support of 4 features as described in table 5.2.2-3.
// The most significant character representing the highest-numbered features
// shall appear first in the string,
// and the character representing features 1 to 4 shall appear last
// in the string.
// The list of features and their numbering (starting with 1)
// are defined separately for each API.
// Possible features for traffic influencing are
// Notification_websocket( takes vlue of 1) and
// Notification_test_event(takes value of 2)
// pattern: '^[A-Fa-f0-9]*$'
type SupportedFeatures string

// List of SubscribedEvent
const (
	//UP_PATH_CHANGE SubscribedEvent = "UP_PATH_CHANGE" >> causing lint error
	UpPathChange SubscribedEvent = "UP_PATH_CHANGE"
)

// TrafficInfluSub is Traffic Influence Subscription structure
type TrafficInfluSub struct {
	// Identifies a service on behalf of which the AF is issuing the request.
	AfServiceID string `json:"afServiceId,omitempty"`
	// Identifies an application.
	AfAppID string `json:"afAppId,omitempty"`
	// Identifies an NEF Northbound interface transaction, generated by the AF.
	AfTransID string `json:"afTransId,omitempty"`
	// Identifies data network name
	Dnn string `json:"dnn,omitempty"`
	// Network slice identifier
	Snssai Snssai `json:"snssai,omitempty"` //p
	// string containing a local identifier followed by \"@\" and
	// a domain identifier.
	// Both the local identifier and the domain identifier shall be encoded as
	// strings that do not contain any \"@\" characters.
	// See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information.
	ExternalGroupID string `json:"externalGroupId,omitempty"`
	// Identifies the requirement to be notified of the event(s).
	SubscribedEvents []SubscribedEvent `json:"subscribedEvents,omitempty"`
	//Generic Public Servie Identifiers asssociated wit the UE
	Gpsi string `json:"gpsi,omitempty"`
	// string identifying a Ipv4 address formatted in the \"dotted decimal\"
	//notation as defined in IETF RFC 1166.
	Ipv4Addr string `json:"ipv4Addr,omitempty"`
	// string identifying a Ipv6 address formatted according to clause 4
	// in IETF RFC 5952.
	Ipv6Addr string `json:"ipv6Addr,omitempty"`
	// string identifying mac address of UE
	MacAddr string `json:"macAddr,omitempty"`
	// Identifies the type of notification regarding UP path management event.
	// Possible values are:
	// EARLY - early notification of UP path reconfiguration.
	// EARLY_LATE - early and late notification of UP path reconfiguration.
	// This value shall only be present in the subscription to the
	// DNAI change event.
	// LATE - late notification of UP path reconfiguration.
	DnaiChgType string `json:"dnaiChgType,omitempty"`
	// URL where notifications shall be sent
	NotificationDestination Link `json:"notificationDestination,omitempty"`
	// URL of created subscription resource
	Self Link `json:"self,omitempty"`
	// Identifies IP packet filters.
	TrafficFilters []FlowInfo `json:"trafficFilters,omitempty"`
	// Identifies Ethernet packet filters.
	EthTrafficFilters []EthFlowDescription `json:"ethTrafficFilters,omitempty"`
	// Identifies the N6 traffic routing requirement.
	TrafficRoutes []RouteToLocation `json:"trafficRoutes,omitempty"`
	// Settings for temporary validity of the subscription
	TempValidities []TemporalValidity `json:"tempValidities,omitempty"`
	// Identifies a geographic zone that the AF request applies only to the
	// traffic of UE(s) located in this specific zone.
	ValidGeoZoneIDs []string `json:"validGeoZoneIds,omitempty"`
	// String identifying supported features per Traffic Influence service
	SuppFeat SupportedFeatures `json:"suppFeat,omitempty"`
	// Configuration used for sending notifications though web sockets
	WebsockNotifConfig WebsockNotifConfig `json:"websockNotifConfig,omitempty"`
	// Set to true by the AF to request the NEF to send a test notification. //p
	//Set to false or omitted otherwise.
	RequestTestNotification bool `json:"requestTestNotification,omitempty"`
	// Identifies whether an application can be relocated once a location of
	// the application has been selected.
	AppReloInd bool `json:"appReloInd,omitempty"`
	// Identifies whether the AF request applies to any UE.
	AnyUeInd bool `json:"anyUeInd,omitempty"`
}

// TrafficInfluSubPatch Traffic Influence Subscription Patch structure
type TrafficInfluSubPatch struct {
	// Identifies IP packet filters.
	TrafficFilters []FlowInfo `json:"trafficFilters,omitempty"`
	// Identifies whether an application can be relocated once a location of
	// the application has been selected.
	AppReloInd bool `json:"appReloInd,omitempty"`
	// Identifies Ethernet packet filters.
	EthTrafficFilters []EthFlowDescription `json:"ethTrafficFilters,omitempty"`
	// Identifies the N6 traffic routing requirement.
	TrafficRoutes []RouteToLocation `json:"trafficRoutes,omitempty"`
	// Settings for temporary validity of the subscription
	TempValidities []TemporalValidity `json:"tempValidities,omitempty"`
	// Identifies a geographic zone that the AF request applies only to the
	// traffic of UE(s) located in this specific zone.
	ValidGeoZoneIDs []string `json:"validGeoZoneIds,omitempty"`
}

// TemporalValidity Indicates the time interval(s) during which the AF request is to be applied
type TemporalValidity struct {
	// string with format \"date-time\" as defined in OpenAPI.
	StartTime string `json:"startTime,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	StopTime string `json:"stopTime,omitempty"`
}

// EthFlowDescription Identifies an Ethernet flow
type EthFlowDescription struct {
	DestMacAddr string `json:"destMacAddr,omitempty"`
	//EtherType number
	EthType string `json:"ethType"`
	// Defines a packet filter of an IP flow.
	FDesc string `json:"fDesc,omitempty"`
	// Possible values are DOWNLINK - The corresponding filter applies for
	// traffic to the UE.
	// UPLINK - The corresponding filter applies for traffic from the UE.
	// BIDIRECTIONAL The corresponding filter applies for traffic both to and
	// from the UE.
	// UNSPECIFIED - The corresponding filter applies for traffic to the UE
	// (downlink), but has no specific direction declared.
	// The service data flow detection shall apply the filter for uplink traffic
	// as if the filter was bidirectional.
	FDir string `json:"fDir,omitempty"`
	// Source mac address
	SourceMacAddr string `json:"sourceMacAddr,omitempty"`
	// Vlan tags
	VlanTags []string `json:"vlanTags,omitempty"`
}

// WebsockNotifConfig Websocket noticcation configuration
type WebsockNotifConfig struct {
	// string formatted according to IETF RFC 3986 identifying a
	// referenced resource.
	WebsocketURI string `json:"websocketUri,omitempty"`
	// Set by the AF to indicate that the Websocket delivery is requested.
	RequestWebsocketURI bool `json:"requestWebsocketUri,omitempty"`
}

// Snssai Network slice identifier
type Snssai struct {
	Sst int32  `json:"sst"`
	Sd  string `json:"sd,omitempty"`
}

// RouteToLocation Route to location structure
type RouteToLocation struct {
	// Data network access identifier
	Dnai string `json:"dnai"`
	// Dnai route profile identifier
	RouteProfID string `json:"routeProfId,omitempty"`
	// Additional route information about the route to Dnai
	RouteInfo *RouteInformation `json:"routeInfo,omitempty"`
}

// RouteInformation Route information struct
type RouteInformation struct {
	// string identifying a Ipv4 address formatted in the \"dotted decimal\"
	// notation as defined in IETF RFC 1166.
	Ipv4Addr string `json:"ipv4Addr,omitempty"`
	// string identifying a Ipv6 address formatted according to clause 4 in
	// IETF RFC 5952.
	// The mixed Ipv4 Ipv6 notation according to clause 5 of IETF RFC 5952 shall
	// not be used.
	Ipv6Addr string `json:"ipv6Addr,omitempty"`
	// Port number
	PortNumber int32 `json:"portNumber"`
}

// ProblemDetails Problem details struct
type ProblemDetails struct {
	// problem type
	Type Link `json:"type,omitempty"`
	// A short, human-readable summary of the problem type.
	// It should not change from occurrence to occurrence of the problem.
	Title string `json:"title,omitempty"`
	// A human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail,omitempty"`
	// URL to problem instance
	Instance Link `json:"instance,omitempty"`
	// A machine-readable application error cause specific to this occurrence
	// of the problem.
	// This IE should be present and provide application-related error
	// information, if available.
	Cause string `json:"cause,omitempty"`
	// Description of invalid parameters, for a request rejected due to
	// invalid parameters.
	InvalidParams []InvalidParam `json:"invalidParams,omitempty"`
	// The HTTP status code for this occurrence of the problem.
	Status int32 `json:"status,omitempty"`
}

// InvalidParam Invalid Parameters struct
type InvalidParam struct {
	// Attribute''s name encoded as a JSON Pointer, or header''s name.
	Param string `json:"param"`
	// A human-readable reason, e.g. \"must be a positive integer\".
	Reason string `json:"reason,omitempty"`
}

// FlowInfo Flow information struct
type FlowInfo struct {
	// Indicates the packet filters of the IP flow. Refer to subclause 5.3.8 of
	//3GPP TS 29.214 for encoding.
	// It shall contain UL and/or DL IP flow description.
	FlowDescriptions []string `json:"flowDescriptions,omitempty"`
	// Indicates the IP flow.
	FlowID int32 `json:"flowId"`
}
