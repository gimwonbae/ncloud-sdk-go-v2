/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-07-05T07:55:47Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type AddNasVolumeAccessControlRequest struct {

	// NAS볼륨인스턴스번호
NasVolumeInstanceNo *string `json:"nasVolumeInstanceNo"`

	// 서버인스턴스번호리스트
ServerInstanceNoList []string `json:"serverInstanceNoList,omitempty"`

	// 커스텀IP리스트
CustomIpList []string `json:"customIpList,omitempty"`
}
