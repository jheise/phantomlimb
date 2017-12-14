package phantomapi

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	// "net/url"
)

func (self *PhantomApi) ScreenShot(target string) ([]byte, error){
	site := b64.StdEncoding.EncodeToString([]byte(target))
	endpoint := fmt.Sprintf("http://%s/phantom/v1/screenshot/%s", self.Host, site)
	fmt.Printf("calling %s\n", endpoint)
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}
