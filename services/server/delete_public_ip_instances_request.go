/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-07-05T07:55:47Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type DeletePublicIpInstancesRequest struct {

	// 공인IP인스턴스번호리스트
PublicIpInstanceNoList []string `json:"publicIpInstanceNoList"`
}
