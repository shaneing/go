/*
Package varlink provides varlink client and server implementations. See http://varlink.org
for more information about varlink.

Example varlink interface definition in a org.example.this.varlink file:
	interface org.example.this

	method Ping(in: string) -> (out: string)

Generated Go module in a orgexamplethis/orgexamplethis.go file:
	// Generated with varlink-generator -- github.com/varlink/go/cmd/varlink-generator
	package orgexamplethis

	import "github.com/varlink/go/varlink"

	type orgexamplethisInterface interface {
		Ping(c VarlinkCall, in string) error
	}

	type VarlinkCall struct{ varlink.Call }

	func (c *VarlinkCall) ReplyPing(out string) error {
		var out struct {
			Out string `json:"out,omitempty"`
		}
		out.Out = out
		return c.Reply(&out)
	}

	func (s *VarlinkInterface) Ping(c VarlinkCall, in string) error {
		return c.ReplyMethodNotImplemented("Ping")
	}

	[...]

Example service:
	import ("orgexamplethis")

	type Data struct {
		orgexamplethis.VarlinkInterface
		data string
	}

	data := Data{data: "test"}

	func (d *Data) Ping(call orgexamplethis.VarlinkCall, ping string) error {
		return call.ReplyPing(ping)
	}

	service, _ = varlink.NewService(
	        "Example",
	        "This",
	        "1",
	         "https://example.org/this",
	)

	service.RegisterInterface(orgexamplethis.VarlinkNew(&data))
	err := service.Listen("unix:@orgexamplethis", 0)
*/
package varlink