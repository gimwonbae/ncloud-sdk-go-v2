/*
 * clouddb
 *
 * Cloud DB<br/>https://ncloud.apigw.ntruss.com/clouddb/v2
 *
 * API version: 2018-07-04T02:46:49Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package clouddb

type CloudDbConfig struct {

	// 설정명
ConfigName *string `json:"configName,omitempty"`

	// 설정값
ConfigValue *string `json:"configValue,omitempty"`
}
