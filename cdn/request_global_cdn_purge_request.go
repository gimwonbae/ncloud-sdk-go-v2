/*
 * cdn
 *
 * <br/>https://ncloud.apigw.ntruss.com/cdn/v2
 *
 * API version: 2018-07-04T02:02:10Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cdn

type RequestGlobalCdnPurgeRequest struct {

	// CDN인스턴스번호
CdnInstanceNo *string `json:"cdnInstanceNo"`

	// 서비스도메인이름리스트
ServiceDomainNameList []string `json:"serviceDomainNameList,omitempty"`

	// 전체퍼지여부
IsWholePurge *bool `json:"isWholePurge"`

	// 전체도메인여부
IsWholeDomain *bool `json:"isWholeDomain"`

	// 대상파일리스트
TargetFileList []string `json:"targetFileList,omitempty"`
}
