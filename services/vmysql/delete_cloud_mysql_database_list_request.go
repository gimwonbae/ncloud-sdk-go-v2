/*
 * vmysql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmysql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmysql

type DeleteCloudMysqlDatabaseListRequest struct {

	// REGION코드
	RegionCode *string `json:"regionCode,omitempty"`

	// CloudMysql인스턴스번호
	CloudMysqlInstanceNo *string `json:"cloudMysqlInstanceNo"`

	// CloudMysqlDatabase리스트
	CloudMysqlDatabaseNameList []*string `json:"cloudMysqlDatabaseNameList"`
}
