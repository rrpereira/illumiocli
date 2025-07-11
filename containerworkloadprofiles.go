package illumiocli

import "fmt"

type Href struct {
	Href string `json:"href,omitempty"`
}

type ContainerWorkloadProfileLabelAssignment struct {
	Href  string `json:"href,omitempty"`
	Value string `json:"value,omitempty"`
}

type ContainerWorkloadProfileLabel struct {
	Key        string                                  `json:"key,omitempty"`
	Assignment ContainerWorkloadProfileLabelAssignment `json:"assignment,omitempty"`
}

type ContainerWorkloadProfile struct {
	Href            string                          `json:"href,omitempty"`
	Namespace       string                          `json:"namespace,omitempty"`
	Description     string                          `json:"description,omitempty"`
	EnforcementMode string                          `json:"enforcement_mode,omitempty"`
	VisibilityLevel string                          `json:"visibility_level,omitempty"`
	Managed         bool                            `json:"managed,omitempty"`
	AssignLabels    []Href                          `json:"assign_labels,omitempty"`
	Labels          []ContainerWorkloadProfileLabel `json:"labels,omitempty"`
	Linked          bool                            `json:"linked,omitempty"`
	CreatedAt       string                          `json:"created_at,omitempty"`
	CreatedBy       *Href                           `json:"created_by,omitempty"`
	UpdatedAt       string                          `json:"updated_at,omitempty"`
	UpdatedBy       *Href                           `json:"updated_by,omitempty"`
}

func (c *PCE) GetContainerWorkloadProfilesByContainerClusterID(id string) ([]ContainerWorkloadProfile, error) {
	var profiles []ContainerWorkloadProfile

	r := c.newRestyClient()

	resp, err := r.R().
		SetResult(&profiles).
		Get(fmt.Sprintf("/orgs/1/container_clusters/%s/container_workload_profiles", id))
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

	return profiles, nil
}

func (c *PCE) PutContainerWorkloadProfileByHref(href string, cwp *ContainerWorkloadProfile) (*ContainerWorkloadProfile, error) {
	var profile ContainerWorkloadProfile

	r := c.newRestyClient()

	resp, err := r.R().
		SetBody(cwp).
		SetResult(&profile).
		Put(href)
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

	return &profile, nil
}
