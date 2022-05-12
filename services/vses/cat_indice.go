/*
 * api
 *
 * Search Engine Service(VPC) Open API<br/>https://vpcsearchengine.apigw.ntruss.com/api/v1
 *
 * API version: 2022-04-21T09:07:51Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vses

type CatIndice struct {
	DocsCount string `json:"docsCount,omitempty"`
	DocsDeleted string `json:"docsDeleted,omitempty"`
	Health string `json:"health,omitempty"`
	IndexName string `json:"indexName,omitempty"`
	PrimaryCount string `json:"primaryCount,omitempty"`
	ReplicaCount string `json:"replicaCount,omitempty"`
	Status string `json:"status,omitempty"`
	StoreSize string `json:"storeSize,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}
