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

// OAuth1AccessTokenView struct for OAuth1AccessTokenView
type OAuth1AccessTokenView struct {
	Token           string             `json:"token,omitempty"`
	Consumer        OAuth1ConsumerView `json:"consumer,omitempty"`
	CreationDate    int64              `json:"creationDate,omitempty"`
	LastUtilisation int64              `json:"lastUtilisation,omitempty"`
	Rights          OAuthRightsView    `json:"rights,omitempty"`
}
