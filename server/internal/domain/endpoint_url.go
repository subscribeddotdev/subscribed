package domain

import "net/url"

type EndpointURL struct {
	value *url.URL
}

func NewEndpointURL(rawURL string) (EndpointURL, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return EndpointURL{}, err
	}

	return EndpointURL{value: u}, nil
}

func (e EndpointURL) String() string {
	return e.value.String()
}

func (e EndpointURL) IsEmpty() bool {
	return e.value == nil || e.value.String() == ""
}
