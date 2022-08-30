/*
 * api
 *
 * Cloud Data Streaming Service Open API<br/>https://clouddatastreamingservice.apigw.ntruss.com/api/v1
 *
 * API version: 2021-12-21T11:54:01Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vcdss

type GetKafkaConfigGroupListResponseVo struct {
	CurrentPage          int32              `json:"currentPage,omitempty"`
	IsFirst              bool               `json:"isFirst,omitempty"`
	IsLast               bool               `json:"isLast,omitempty"`
	KafkaConfigGroupList []KafkaConfigGroup `json:"kafkaConfigGroupList,omitempty"`
	PageSize             int32              `json:"pageSize,omitempty"`
	TotalCount           int64              `json:"totalCount,omitempty"`
	TotalPage            int32              `json:"totalPage,omitempty"`
}