/*
 * vmssql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmssql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmssql

type GetCloudMssqlBackupListResponse struct {

RequestId *string `json:"requestId,omitempty"`

ReturnCode *string `json:"returnCode,omitempty"`

ReturnMessage *string `json:"returnMessage,omitempty"`

TotalRows *int32 `json:"totalRows,omitempty"`

	// CloudMssql백업리스트
CloudMssqlBackupList *CloudMssqlBackupList `json:"cloudMssqlBackupList,omitempty"`
}
