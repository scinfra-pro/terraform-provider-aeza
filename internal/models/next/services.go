// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// models/responses/legacy/legacy_models.go
package next

type Service struct {
	ID           int                    `json:"id"`
	Name         string                 `json:"name"`
	IP           string                 `json:"ip"`
	Payload      map[string]interface{} `json:"payload"`
	Price        float64                `json:"price"`
	PaymentTerm  string                 `json:"paymentTerm"`
	AutoProlong  bool                   `json:"autoProlong"`
	CreatedAt    string                 `json:"createdAt"`
	ExpiresAt    string                 `json:"expiresAt"`
	Status       string                 `json:"status"`
	TypeSlug     string                 `json:"typeSlug"`
	ProductName  string                 `json:"productName"`
	LocationCode string                 `json:"locationCode"`
	CurrentTask  *CurrentTask           `json:"currentTask"`
	Capabilities []string               `json:"capabilities"`
}

type CurrentTask struct {
	ID        string `json:"id"`
	Slug      string `json:"slug"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"` // ISO 8601
	Status    string `json:"status"`    // "queued", "running", etc.
}

type ComputedParametersVPS struct {
	CPU      int     `json:"cpu"`
	RAM      int     `json:"ram"`
	ROM      int     `json:"rom"`
	IP       int     `json:"ip"`
	OS       string  `json:"os"`
	Node     string  `json:"node"`
	ISOURL   string  `json:"isoUrl"`
	Recipe   *string `json:"recipe"` // nullable
	Username string  `json:"username"`
}
