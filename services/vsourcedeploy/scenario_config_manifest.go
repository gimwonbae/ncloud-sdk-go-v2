/*
 * vsourcedeploy
 *
 * <br/>https://vpcsourcedeploy.apigw.ntruss.com/api/v1
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vsourcedeploy

type ScenarioConfigManifest struct {

Type_ *string `json:"type,omitempty"`

Repository *string `json:"repository,omitempty"`

Branch *string `json:"branch,omitempty"`

Path []*string `json:"path,omitempty"`
}