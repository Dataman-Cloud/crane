package registry

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// the following code borrowed from vmvare/Harbor project
func (registry *Registry) RegistryAPIGet(path, username string) ([]byte, error) {
	url := fmt.Sprintf("%s/v2/%s", registry.Config.RegistryAddr, path)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		return result, nil
	} else if response.StatusCode == http.StatusUnauthorized {
		authenticate := response.Header.Get("WWW-Authenticate")
		str := strings.Split(authenticate, " ")[1]
		var service string
		var scope string
		strs := strings.Split(str, ",")
		for _, s := range strs {
			if strings.Contains(s, "service") {
				service = s
			} else if strings.Contains(s, "scope") {
				scope = s
			}
		}
		service = strings.Split(service, "\"")[1]
		scope = strings.Split(scope, "\"")[1]
		token, err := GenTokenForUI(registry.Config, username, service, scope)
		if err != nil {
			return nil, err
		}
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		request.Header.Add("Authorization", "Bearer "+token)
		client := &http.Client{}
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			if len(via) >= 10 {
				return fmt.Errorf("too many redirects")
			}
			for k, v := range via[0].Header {
				if _, ok := req.Header[k]; !ok {
					req.Header[k] = v
				}
			}
			return nil
		}
		response, err = client.Do(request)
		if err != nil {
			return nil, err
		}

		if response.StatusCode != http.StatusOK {
			errMsg := fmt.Sprintf("Unexpected return code from registry: %d", response.StatusCode)
			log.Printf(errMsg)
			return nil, fmt.Errorf(errMsg)
		}
		result, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()
		return result, nil
	} else {
		return nil, errors.New(string(result))
	}
}
