/*
 * Clever-Cloud API
 *
 * Public API for managing Clever-Cloud data and products
 *
 * API version: 1.0.1
 * Contact: support@clever-cloud.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package clevercloud

// AddonApplicationSummary struct for AddonApplicationSummary
type AddonApplicationSummary struct {
	ProviderId  string `json:"provider_id,omitempty"`
	HerokuId    string `json:"heroku_id,omitempty"`
	CallbackUrl string `json:"callback_url,omitempty"`
	Plan        string `json:"plan,omitempty"`
	OwnerId     string `json:"owner_id,omitempty"`
}
