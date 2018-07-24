/*
 * autoscaling
 *
 * <br/>https://ncloud.apigw.ntruss.com/autoscaling/v2
 *
 * API version: 2018-06-21T02:22:22Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package autoscaling

type CreateAutoScalingGroupRequest struct {

	// 오토스케일링그룹명
AutoScalingGroupName *string `json:"autoScalingGroupName,omitempty"`

	// 론치설정명
LaunchConfigurationName *string `json:"launchConfigurationName"`

	// 기대용량치
DesiredCapacity *int32 `json:"desiredCapacity,omitempty"`

	// 최소사이즈
MinSize *int32 `json:"minSize"`

	// 최대사이즈
MaxSize *int32 `json:"maxSize"`

	// 디폴트쿨다운타임
DefaultCooldown *int32 `json:"defaultCooldown,omitempty"`

	// 로드밸런서명리스트
LoadBalancerNameList []string `json:"loadBalancerNameList,omitempty"`

	// 헬스체크보류기간
HealthCheckGracePeriod *int32 `json:"healthCheckGracePeriod,omitempty"`

	// 헬스체크유형코드
HealthCheckTypeCode *string `json:"healthCheckTypeCode,omitempty"`

	// ZONE번호리스트
ZoneNoList []string `json:"zoneNoList"`
}
