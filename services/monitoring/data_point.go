/*
 * monitoring
 *
 * <br/>https://ncloud.apigw.ntruss.com/monitoring/v2
 *
 * API version: 2018-06-25T02:38:27Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package monitoring

type DataPoint struct {

Timestamp *string `json:"timestamp,omitempty"`

Average *float64 `json:"average,omitempty"`

Unit *string `json:"unit,omitempty"`
}
