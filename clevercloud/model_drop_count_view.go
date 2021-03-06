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

// DropCountView struct for DropCountView
type DropCountView struct {
	OwnerId   string        `json:"ownerId,omitempty"`
	Count     float32       `json:"count,omitempty"`
	DropPrice DropPriceView `json:"dropPrice,omitempty"`
}
