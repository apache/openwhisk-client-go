package whisk

import "net/http"

type Info struct {
	Whisk   string `json:"whisk,omitempty"`
	Version string `json:"version,omitempty"`
	Build   string `json:"build,omitempty"`
}

type InfoService struct {
	client *Client
}

func (s *InfoService) Get() (*Info, *http.Response, error) {
	// make a request to c.BaseURL / namespaces

	req, err := http.NewRequest("GET", s.client.BaseURL.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	info := new(Info)
	resp, err := s.client.Do(req, &info)
	if err != nil {
		return nil, resp, err
	}

	return info, resp, nil
}
