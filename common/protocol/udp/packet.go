package udp

import (
	"v2ray.com/core/v4/common/buf"
	"v2ray.com/core/v4/common/net"
)

// Packet is a UDP packet together with its source and destination address.
type Packet struct {
	Payload *buf.Buffer
	Source  net.Destination
	Target  net.Destination
}
