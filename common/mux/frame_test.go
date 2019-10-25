package mux_test

import (
	"testing"

	"v2ray.com/core/v4/common"
	"v2ray.com/core/v4/common/buf"
	"v2ray.com/core/v4/common/mux"
	"v2ray.com/core/v4/common/net"
)

func BenchmarkFrameWrite(b *testing.B) {
	frame := mux.FrameMetadata{
		Target:        net.TCPDestination(net.DomainAddress("www.v2ray.com"), net.Port(80)),
		SessionID:     1,
		SessionStatus: mux.SessionStatusNew,
	}
	writer := buf.New()
	defer writer.Release()

	for i := 0; i < b.N; i++ {
		common.Must(frame.WriteTo(writer))
		writer.Clear()
	}
}
