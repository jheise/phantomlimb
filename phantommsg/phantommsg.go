package phantommsg

type PhantomResponse struct {
	ResponseCode int
	Data         string
}

type PhantomRequest struct {
	PluginName string
	Arguments  []string
	Response   chan *PhantomResponse
}

func NewPhantomRequest(plugin string, args []string) *PhantomRequest {
	request := new(PhantomRequest)
	request.PluginName = plugin
	request.Arguments = args
	request.Response = make(chan *PhantomResponse)

	return request
}
