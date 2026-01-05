// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// models/legacy/recipe.go
package legacy

type Recipe struct {
	ID              int            `json:"id"`
	Slug            string         `json:"slug"`
	Name            string         `json:"name"`
	Icon            string         `json:"icon"`
	Description     string         `json:"description"`
	Tags            []string       `json:"tags"`
	Enabled         bool           `json:"enabled"`
	OS              int            `json:"os"`
	OperatingSystem string         `json:"operatingSystem"`
	Targets         map[string]int `json:"targets"`
}
