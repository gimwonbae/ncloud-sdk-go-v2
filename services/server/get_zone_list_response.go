/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-07-05T07:55:47Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type GetZoneListResponse struct {

RequestId *string `json:"requestId,omitempty"`

ReturnCode *string `json:"returnCode,omitempty"`

ReturnMessage *string `json:"returnMessage,omitempty"`

ZoneList []Zone `json:"zoneList,omitempty"`
}
