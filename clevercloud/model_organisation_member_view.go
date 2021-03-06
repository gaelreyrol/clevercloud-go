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

// OrganisationMemberView struct for OrganisationMemberView
type OrganisationMemberView struct {
	Member OrganisationMemberUserView `json:"member,omitempty"`
	Role   string                     `json:"role,omitempty"`
	Job    string                     `json:"job,omitempty"`
}
