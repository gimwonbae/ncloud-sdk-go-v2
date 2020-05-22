/*
 * monitoring
 *
 * <br/>https://ncloud.apigw.ntruss.com/monitoring/v2
 *
 * API version: 2020-05-22T05:53:06Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package monitoring

type GetMetricStatisticListResponse struct {

RequestId *string `json:"requestId,omitempty"`

ReturnCode *string `json:"returnCode,omitempty"`

ReturnMessage *string `json:"returnMessage,omitempty"`

MetricStatisticList []*MetricStatistic `json:"metricStatisticList,omitempty"`
}