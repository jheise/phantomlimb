package phantomapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (self *PhantomApi) ScreenShot(target string) ([]byte, error){
	endpoint := fmt.Sprintf("http://%s/phantom/v1/screenshot/%s", self.Host, url.PathEscape(target))
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
