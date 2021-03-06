/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dcos

type IamUserPermissionsGroups struct {
	Rid           string      `json:"rid"`
	Gid           string      `json:"gid"`
	Description   string      `json:"description"`
	Aclurl        string      `json:"aclurl"`
	Membershipurl string      `json:"membershipurl"`
	Actions       []IamAction `json:"actions"`
}
