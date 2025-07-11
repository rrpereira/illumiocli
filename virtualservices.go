package illumiocli

import "fmt"

type VirtualServiceServicePort struct {
	Port  int `json:"port,omitempty"`
	Proto int `json:"proto,omitempty"`
}

type VirtualServiceLabel struct {
	Href  string `json:"href,omitempty"`
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type VirtualServiceServiceAddress struct {
	Fqdn        string `json:"fqdn,omitempty"`
	Description string `json:"description,omitempty"`
}

type VirtualService struct {
	Href                  string                         `json:"href,omitempty"`
	CreatedAt             string                         `json:"created_at,omitempty"`
	UpdatedAt             string                         `json:"updated_at,omitempty"`
	DeletedAt             string                         `json:"deleted_at,omitempty"`
	CreatedBy             Href                           `json:"created_by,omitempty"`
	UpdatedBy             Href                           `json:"updated_by,omitempty"`
	DeletedBy             Href                           `json:"deleted_by,omitempty"`
	UpdateType            string                         `json:"update_type,omitempty"`
	Name                  string                         `json:"name,omitempty"`
	Description           string                         `json:"description,omitempty"`
	ExternalDataSet       string                         `json:"external_data_set,omitempty"`
	ExternalDataReference string                         `json:"external_data_reference,omitempty"`
	PCEFqdn               *string                        `json:"pce_fqdn,omitempty"`
	ServicePorts          []VirtualServiceServicePort    `json:"service_ports,omitempty"`
	Labels                []VirtualServiceLabel          `json:"labels,omitempty"`
	IpOverrides           []string                       `json:"ip_overrides,omitempty"`
	ApplyTo               string                         `json:"apply_to,omitempty"`
	Caps                  []string                       `json:"caps,omitempty"`
	ServiceAddresses      []VirtualServiceServiceAddress `json:"service_addresses,omitempty"`
}

func (c *PCE) GetVirtualServicesByParameter(key, value, pversion string) ([]VirtualService, error) {
	var virtualServices []VirtualService

	r := c.newRestyClient()

	resp, err := r.R().
		SetQueryParam(key, value).
		SetResult(&virtualServices).
		Get(fmt.Sprintf("/orgs/1/sec_policy/%s/virtual_services?%s=%s", pversion, key, value))
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		if len(resp.Body()) != 0 {
			return nil, fmt.Errorf("API error: status %d - %s - %s", resp.StatusCode(), resp.Status(), string(resp.Body()))
		} else {
			return nil, fmt.Errorf("API error: status %d - %s", resp.StatusCode(), resp.Status())
		}
	}

	return virtualServices, nil
}
