/*
 * api
 *
 * Cloud Data Streaming Service Open API<br/>https://clouddatastreamingservice.apigw.ntruss.com/api/v1
 *
 * API version: 2021-12-21T11:54:01Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vcdss

type GetClusterStatusPerNodeResponseVo struct {
	BrokerNodeTotalNumber int32 `json:"brokerNodeTotalNumber,omitempty"`
	BrokerNumber int32 `json:"brokerNumber,omitempty"`
	BrokerStatus string `json:"brokerStatus,omitempty"`
	BrokerStatusPerNode []ProcessStatusVo `json:"brokerStatusPerNode,omitempty"`
	ClusterName string `json:"clusterName,omitempty"`
	CmakStatus string `json:"cmakStatus,omitempty"`
	ZookeeperNumber int32 `json:"zookeeperNumber,omitempty"`
	ZookeeperStatus string `json:"zookeeperStatus,omitempty"`
	ZookeeperStatusPerNode []ProcessStatusVo `json:"zookeeperStatusPerNode,omitempty"`
	ZookeeperTotalNumber int32 `json:"zookeeperTotalNumber,omitempty"`
}