/*
 * api
 *
 * Cloud Data Streaming Service Open API<br/>https://clouddatastreamingservice.apigw.ntruss.com/api/v1
 *
 * API version: 2021-12-21T11:54:01Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vcdss

type DateTimeVo struct {
	FormattedDate string `json:"formattedDate,omitempty"`
	FormattedDateTime string `json:"formattedDateTime,omitempty"`
	Utc int64 `json:"utc,omitempty"`
}