/*
 * vsourcedeploy
 *
 * <br/>https://vpcsourcedeploy.apigw.ntruss.com/api/v1
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vsourcedeploy

type GetScenarioDetailResponse struct {

Project *GetIdNameResponse `json:"project,omitempty"`

Stage *GetIdNameResponse `json:"stage,omitempty"`

Id *int32 `json:"id,omitempty"`

Name *string `json:"name,omitempty"`

Type_ *string `json:"type,omitempty"`

Description *string `json:"description,omitempty"`

Config *GetScenarioConfig `json:"config,omitempty"`
}