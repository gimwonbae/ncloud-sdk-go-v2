/*
 * clouddb
 *
 * Cloud DB<br/>https://ncloud.apigw.ntruss.com/clouddb/v2
 *
 * API version: 2018-07-04T02:46:49Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package clouddb

type FlushCloudDbInstanceResponse struct {

RequestId *string `json:"requestId,omitempty"`

ReturnCode *string `json:"returnCode,omitempty"`

ReturnMessage *string `json:"returnMessage,omitempty"`

TotalRows *int32 `json:"totalRows,omitempty"`

CloudDBInstanceList []CloudDbInstance `json:"cloudDBInstanceList,omitempty"`
}
