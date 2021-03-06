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

// InlineObject10 struct for InlineObject10
type InlineObject10 struct {
	InvitationKey          string `json:"invitationKey,omitempty"`
	AddonBetaInvitationKey string `json:"addonBetaInvitationKey,omitempty"`
	Email                  string `json:"email,omitempty"`
	Pass                   string `json:"pass,omitempty"`
	UrlNext                string `json:"url_next,omitempty"`
	Terms                  string `json:"terms,omitempty"`
	SubscriptionSource     string `json:"subscription_source,omitempty"`
	CleverFlavor           string `json:"clever_flavor,omitempty"`
	OauthToken             string `json:"oauth_token,omitempty"`
}
