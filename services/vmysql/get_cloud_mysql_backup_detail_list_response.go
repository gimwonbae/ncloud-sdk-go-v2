/*
 * vmysql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmysql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmysql

type GetCloudMysqlBackupDetailListResponse struct {
	RequestId *string `json:"requestId,omitempty"`

	ReturnCode *string `json:"returnCode,omitempty"`

	ReturnMessage *string `json:"returnMessage,omitempty"`

	TotalRows *int32 `json:"totalRows,omitempty"`

	// CloudMysql백업상세리스트
	CloudMysqlBackupDetailList []*CloudMysqlBackupDetail `json:"cloudMysqlBackupDetailList,omitempty"`
}
