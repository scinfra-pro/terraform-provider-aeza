// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/utils/format_utils.go
package utils

import (
	"fmt"
	"strings"
	"time"
)

func FormatPrice(price float64) string {
		euros := float64(price) / 100.0
		formatted := fmt.Sprintf("%.2f", euros)
	formatted = strings.Replace(formatted, ".", ",", 1)
	return fmt.Sprintf("%s €", formatted)
}

func FormatDate(isoDate string) string {
		t, err := time.Parse(time.RFC3339, isoDate)
	if err != nil {
		return isoDate
	}

		return t.Format("02.01.2006 T15:04:05.000Z")
}

func FormatDateFromUnix(unixTimestamp int64) string {
	if unixTimestamp == 0 {
		return ""
	}

	t := time.Unix(unixTimestamp, 0)
	return t.Format("02.01.2006 T15:04:05.000Z")
}

func stringToPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
