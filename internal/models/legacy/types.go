// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// models/responses/legacy/types.go
package legacy

type ServiceType struct {
	Slug                 string                 `json:"slug"`
	Order                int                    `json:"order"`
	Names                map[string]string      `json:"names"`
	Payload              map[string]interface{} `json:"payload"`
	LocaledPayload       map[string]interface{} `json:"localedPayload"`
	Name                 string                 `json:"name"`
	PrettyLocaledPayload map[string]string      `json:"prettyLocaledPayload"`
}
