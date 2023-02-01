/*
 * api
 *
 * Search Engine Service(VPC) Open API<br/>https://vpcsearchengine.apigw.ntruss.com/api/v2)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vses2

type GetOpenApiAccessControlGroupRuleListResponseVo struct {
	InboundAcgRules  []AcgRuleResponseDetailVo `json:"inboundAcgRules,omitempty"`
	OutboundAcgRules []AcgRuleResponseDetailVo `json:"outboundAcgRules,omitempty"`
}