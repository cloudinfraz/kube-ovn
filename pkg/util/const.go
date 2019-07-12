package util

const (
	MacAddressAnnotation = "ovn.kubernetes.io/mac_address"
	IpAddressAnnotation  = "ovn.kubernetes.io/ip_address"
	CidrAnnotation       = "ovn.kubernetes.io/cidr"
	GatewayAnnotation    = "ovn.kubernetes.io/gateway"
	IpPoolAnnotation     = "ovn.kubernetes.io/ip_pool"

	IngressRateAnnotation = "ovn.kubernetes.io/ingress_rate"
	EgressRateAnnotation  = "ovn.kubernetes.io/egress_rate"

	PortNameAnnotation = "ovn.kubernetes.io/port_name"

	LogicalSwitchAnnotation = "ovn.kubernetes.io/logical_switch"

	ProtocolTCP = "tcp"
	ProtocolUDP = "udp"

	NodeNic           = "ovn0"
	NodeAllowPriority = "3000"

	IngressExceptDropPriority = "2002"
	IngressAllowPriority      = "2001"
	IngressDefaultDrop        = "2000"

	EgressExceptDropPriority = "2002"
	EgressAllowPriority      = "2001"
	EgressDefaultDrop        = "2000"

	SubnetAllowPriority = "1001"
	DefaultDropPriority = "1000"

	// Refer to http://www.openvswitch.org/support/dist-docs/ovn-architecture.7.html "Tunnel Encapsulations"
	// outer IP header + outer UDP header + Geneve header + OVN header + inner eth header
	GeneveHeaderLength = 20 + 8 + 8 + 8 + 14
)
