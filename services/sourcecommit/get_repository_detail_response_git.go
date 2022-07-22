/*
 * api
 *
 * <br/>https://sourcecommit.apigw.ntruss.com/api/v1
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package sourcecommit

type GetRepositoryDetailResponseGit struct {

	// git repository https url
	Https *string `json:"https,omitempty"`

	// git repository ssh url
	Ssh *string `json:"ssh,omitempty"`
}