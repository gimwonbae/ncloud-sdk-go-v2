/*
 * api
 *
 * Cloud Data Streaming Service Open API<br/>https://clouddatastreamingservice.apigw.ntruss.com/api/v1
 *
 * API version: 2021-12-21T11:54:01Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vcdss

type MetricInfoVo struct {
	Aggregation string `json:"aggregation,omitempty"`
	ComputeInstanceNo string `json:"computeInstanceNo,omitempty"`
	CwKey string `json:"cw_key,omitempty"`
	Dimensions map[string]string `json:"dimensions,omitempty"`
	Interval string `json:"interval,omitempty"`
	Metric string `json:"metric,omitempty"`
	ProdKey string `json:"prodKey,omitempty"`
	ProductName string `json:"productName,omitempty"`
	QueryAggregation string `json:"queryAggregation,omitempty"`
	ServiceGroupInstanceNo string `json:"serviceGroupInstanceNo,omitempty"`
}