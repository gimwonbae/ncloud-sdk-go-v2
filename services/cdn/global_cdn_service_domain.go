/*
 * cdn
 *
 * <br/>https://ncloud.apigw.ntruss.com/cdn/v2
 *
 * API version: 2018-07-04T02:02:10Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cdn

type GlobalCdnServiceDomain struct {

	// 서비스도메인구분코드
ServiceDomainTypeCode *string `json:"serviceDomainTypeCode,omitempty"`

	// 프로토콜구분코드
ProtocolTypeCode *string `json:"protocolTypeCode,omitempty"`

	// 디폴트도메인이름
DefaultDomainName *string `json:"defaultDomainName,omitempty"`

	// 유저도메인이름
UserDomainName *string `json:"userDomainName,omitempty"`
}
