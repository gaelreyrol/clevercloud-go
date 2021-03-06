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

// WannabeNamespace struct for WannabeNamespace
type WannabeNamespace struct {
	Namespace string `json:"namespace,omitempty"`
	MinPort   int64  `json:"minPort,omitempty"`
	MaxPort   int64  `json:"maxPort,omitempty"`
}
