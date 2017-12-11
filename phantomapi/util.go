package phantomapi

type PhantomApi struct {
	Host string
}

func NewPhantomApi(host string) *PhantomApi{
	api := new(PhantomApi)
	api.Host = host
	return api
}
