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

// SetupIntentView PaymentIntent id from Stripe
type SetupIntentView struct {
	OwnerId      string `json:"ownerId,omitempty"`
	Id           string `json:"id,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
	Customer     string `json:"customer,omitempty"`
}
