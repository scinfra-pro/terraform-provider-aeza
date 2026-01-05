// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// models/legacy/responses.go
package legacy

type ListServicesResponse struct {
	Data struct {
		SelectorMode string        `json:"selectorMode"`
		Filter       *string       `json:"filter"`
		Items        []ServiceBase `json:"items"`
		Total        int           `json:"total"`
		Edit         bool          `json:"edit"`
	} `json:"data"`
}

type ListServicesVPSResponse struct {
	Data struct {
		SelectorMode string       `json:"selectorMode"`
		Filter       *string      `json:"filter"`
		Items        []ServiceVPS `json:"items"`
		Total        int          `json:"total"`
		Edit         bool         `json:"edit"`
	} `json:"data"`
}

type ListServiceTypesResponse struct {
	Data struct {
		Items []ServiceType `json:"items"`
		Total int           `json:"total"`
	} `json:"data"`
}

type ServiceGroupsResponse struct {
	Data struct {
		Items []ServiceGroup `json:"items"`
		Total int            `json:"total"`
	} `json:"data"`
}

type ListProductsResponse struct {
	Data struct {
		Items []Product `json:"items"`
		Total int       `json:"total"`
	} `json:"data"`
}

type OSResponse struct {
	Data struct {
		Items []OperatingSystem `json:"items"`
	} `json:"data"`
}

type ListRecipesResponse struct {
	Data struct {
		Items []Recipe `json:"items"`
	} `json:"data"`
}

type ServiceDeleteResponse struct {
	Data string `json:"data"`
}
