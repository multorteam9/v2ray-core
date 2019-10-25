package blackhole_test

import (
	"context"
	"testing"

	"v2ray.com/core/v4/common"
	"v2ray.com/core/v4/common/buf"
	"v2ray.com/core/v4/common/serial"
	"v2ray.com/core/v4/proxy/blackhole"
	"v2ray.com/core/v4/transport"
	"v2ray.com/core/v4/transport/pipe"
)

func TestBlackholeHTTPResponse(t *testing.T) {
	handler, err := blackhole.New(context.Background(), &blackhole.Config{
		Response: serial.ToTypedMessage(&blackhole.HTTPResponse{}),
	})
	common.Must(err)

	reader, writer := pipe.New(pipe.WithoutSizeLimit())

	var mb buf.MultiBuffer
	var rerr error
	go func() {
		b, e := reader.ReadMultiBuffer()
		mb = b
		rerr = e
	}()

	link := transport.Link{
		Reader: reader,
		Writer: writer,
	}
	common.Must(handler.Process(context.Background(), &link, nil))
	common.Must(rerr)
	if mb.IsEmpty() {
		t.Error("expect http response, but nothing")
	}
}
