/*
 * api
 *
 * Cloud Data Streaming Service Open API<br/>https://clouddatastreamingservice.apigw.ntruss.com/api/v1
 *
 * API version: 2021-12-21T11:54:01Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vcdss

type OpenApiCreateClusterRequestVo struct {
	// 브로커 노드 갯수 | min:3 | only 3 to up
	BrokerNodeCount int32 `json:"brokerNodeCount"`
	// 브로커 노드 상품(Spec) 코드
	BrokerNodeProductCode string `json:"brokerNodeProductCode"`
	// 데이터 노드 스토리지 사이즈
	BrokerNodeStorageSize int64 `json:"brokerNodeStorageSize"`
	// 데이터 노드 스토리지 타입 | Default: SSD
	BrokerNodeStorageType2Code string `json:"brokerNodeStorageType2Code"`
	// 브로커 서브넷 이름
	BrokerNodeSubnetName string `json:"brokerNodeSubnetName"`
	// 브로커 서브넷 번호
	BrokerNodeSubnetNo int32 `json:"brokerNodeSubnetNo"`
	// 클러스터 이름
	ClusterName string `json:"clusterName"`
	// 카프카 매니저 접근 유저 네임
	KafkaManagerUserName string `json:"kafkaManagerUserName"`
	// 카프카 매니저 접근 유저 비밀번호
	KafkaManagerUserPassword string `json:"kafkaManagerUserPassword"`
	// Application 버전 코드
	KafkaVersionCode string `json:"kafkaVersionCode"`
	// 매니저 노드 갯수 | Default:1 | only 1
	ManagerNodeCount int32 `json:"managerNodeCount"`
	// 매니저 노드 상품(Spec) 코드
	ManagerNodeProductCode string `json:"managerNodeProductCode"`
	// 매니저 서브넷 이름
	ManagerNodeSubnetName string `json:"managerNodeSubnetName,omitempty"`
	// 매니저 서브넷 번호
	ManagerNodeSubnetNo int32 `json:"managerNodeSubnetNo,omitempty"`
	// 소프트웨어 상품(OS) 코드
	SoftwareProductCode string `json:"softwareProductCode"`
	// VPC 이름
	VpcName string `json:"vpcName,omitempty"`
	// VPC 번호
	VpcNo int32 `json:"vpcNo,omitempty"`
}
