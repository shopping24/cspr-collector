package csprcollector

import "time"

type CSPReport struct {
	DocumentUri        string    `json:"document-uri"`
	Referrer           string    `json:"referrer"`
	ViolatedDirective  string    `json:"violated-directive"`
	EffectiveDirective string    `json:"effective-directive"`
	OriginalPolicy     string    `json:"original-policy"`
	Disposition        string    `json:"disposition"`
	BlockedUri         string    `json:"blocked-uri"`
	StatusCode         int       `json:"status-code"`
	ScriptSample       string    `json:"script-sample"`
	UserAgent          string    `json:"user-agent"`
	Occurred           time.Time `json:"occurred"`
}

func NewCSPRequest() CSPRequest {
	report := CSPRequest{
		Report: CSPReport{
			Occurred: time.Now(),
		},
	}

	return report
}

type CSPRequest struct {
	Report CSPReport `json:"csp-report"`
}
