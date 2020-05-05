package rabbithole

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// RuntimeParameter represents a vhost-scoped parameter.
// Value is interface{} to support creating parameters directly from types such as
// FederationInfo and ShovelInfo.
type RuntimeParameter struct {
	Name      string      `json:"name"`
	Vhost     string      `json:"vhost"`
	Component string      `json:"component"`
	Value     interface{} `json:"value"`
}

// RuntimeParameterValue represents arbitrary parameter data.
type RuntimeParameterValue map[string]interface{}

//
// GET /api/parameters
//

// ListRuntimeParameters returns all runtime parameters.
func (c *Client) ListRuntimeParameters() (params []RuntimeParameter, err error) {
	req, err := newGETRequest(c, "parameters")
	if err != nil {
		return []RuntimeParameter{}, err
	}

	if err = executeAndParseRequest(c, req, &params); err != nil {
		return []RuntimeParameter{}, err
	}

	return params, nil
}

//
// GET /api/parameters/{component}
//

// ListRuntimeParametersFor returns all runtime parameters for a component in all vhosts.
func (c *Client) ListRuntimeParametersFor(component string) (params []RuntimeParameter, err error) {
	req, err := newGETRequest(c, "parameters/"+url.PathEscape(component))
	if err != nil {
		return []RuntimeParameter{}, err
	}

	if err = executeAndParseRequest(c, req, &params); err != nil {
		return []RuntimeParameter{}, err
	}

	return params, nil
}

//
// GET /api/parameters/{component}/{vhost}
//

// ListRuntimeParametersIn returns all runtime parameters for a component in a vhost.
func (c *Client) ListRuntimeParametersIn(component, vhost string) (p []RuntimeParameter, err error) {
	req, err := newGETRequest(c, "parameters/"+url.PathEscape(component)+"/"+url.PathEscape(vhost))
	if err != nil {
		return []RuntimeParameter{}, err
	}

	if err = executeAndParseRequest(c, req, &p); err != nil {
		return []RuntimeParameter{}, err
	}

	return p, nil
}

//
// GET /api/parameters/{component}/{vhost}/{name}
//

// GetRuntimeParameter returns a runtime parameter.
func (c *Client) GetRuntimeParameter(component, vhost, name string) (p *RuntimeParameter, err error) {
	req, err := newGETRequest(c, "parameters/"+url.PathEscape(component)+"/"+url.PathEscape(vhost)+"/"+url.PathEscape(name))
	if err != nil {
		return nil, err
	}

	if err = executeAndParseRequest(c, req, &p); err != nil {
		return nil, err
	}

	return p, nil
}

//
// PUT /api/parameters/{component}/{vhost}/{name}
//

// PutRuntimeParameter creates a runtime parameter.
func (c *Client) PutRuntimeParameter(component, vhost, name string, value interface{}) (res *http.Response, err error) {
	param := RuntimeParameter{
		name,
		vhost,
		component,
		value,
	}

	body, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	req, err := newRequestWithBody(c, "PUT", "parameters/"+url.PathEscape(component)+"/"+url.PathEscape(vhost)+"/"+url.PathEscape(name), body)
	if err != nil {
		return nil, err
	}

	if res, err = executeRequest(c, req); err != nil {
		return nil, err
	}

	return res, nil
}

//
// DELETE /api/parameters/{component}/{vhost}/{name}
//

// DeleteRuntimeParameter removes a runtime parameter.
func (c *Client) DeleteRuntimeParameter(component, vhost, name string) (res *http.Response, err error) {
	req, err := newRequestWithBody(c, "DELETE", "parameters/"+url.PathEscape(component)+"/"+url.PathEscape(vhost)+"/"+url.PathEscape(name), nil)
	if err != nil {
		return nil, err
	}

	if res, err = executeRequest(c, req); err != nil {
		return nil, err
	}

	return res, nil
}
