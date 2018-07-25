/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-07-05T07:55:47Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type BlockStorageSnapshotInstance struct {

	// 블록스토리지스냅샷인스턴스번호
BlockStorageSnapshotInstanceNo *string `json:"blockStorageSnapshotInstanceNo,omitempty"`

	// 블록스토리지스냅샷명
BlockStorageSnapshotName *string `json:"blockStorageSnapshotName,omitempty"`

	// 블록스토지리볼륨사이즈
BlockStorageSnapshotVolumeSize *int64 `json:"blockStorageSnapshotVolumeSize,omitempty"`

	// 원본블록스토리지인스턴스번호
OriginalBlockStorageInstanceNo *string `json:"originalBlockStorageInstanceNo,omitempty"`

	// 원본블록스토리지명
OriginalBlockStorageName *string `json:"originalBlockStorageName,omitempty"`

	// 블록스토리지스냅샷인스턴스상태
BlockStorageSnapshotInstanceStatus **CommonCode `json:"blockStorageSnapshotInstanceStatus,omitempty"`

	// 블록스토리지스냅샷인스턴스OP
BlockStorageSnapshotInstanceOperation **CommonCode `json:"blockStorageSnapshotInstanceOperation,omitempty"`

BlockStorageSnapshotInstanceStatusName *string `json:"blockStorageSnapshotInstanceStatusName,omitempty"`

	// 생성일시
CreateDate *string `json:"createDate,omitempty"`

	// 블록스토리지스냅샷인스턴스설명
BlockStorageSnapshotInstanceDescription *string `json:"blockStorageSnapshotInstanceDescription,omitempty"`

	// 서버이미지상품코드
ServerImageProductCode *string `json:"serverImageProductCode,omitempty"`

	// OS정보
OsInformation *string `json:"osInformation,omitempty"`
}
