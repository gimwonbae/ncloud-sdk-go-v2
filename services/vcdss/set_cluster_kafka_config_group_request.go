/*
 * api
 *
 * Cloud Data Streaming Service Open API<br/>https://clouddatastreamingservice.apigw.ntruss.com/api/v1
 *
 * API version: 2021-12-21T11:54:01Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vcdss

type SetClusterKafkaConfigGroupRequest struct {
	KafkaVersionCode       string `json:"kafkaVersionCode,omitempty"`
	ServiceGroupInstanceNo int32  `json:"serviceGroupInstanceNo,omitempty"`
}
