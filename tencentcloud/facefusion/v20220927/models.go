// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20220927

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type DescribeMaterialListRequestParams struct {
	// 活动Id
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 素材Id
	MaterialId *string `json:"MaterialId,omitnil,omitempty" name:"MaterialId"`

	// 每次拉取条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeMaterialListRequest struct {
	*tchttp.BaseRequest
	
	// 活动Id
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 素材Id
	MaterialId *string `json:"MaterialId,omitnil,omitempty" name:"MaterialId"`

	// 每次拉取条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeMaterialListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaterialListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ActivityId")
	delete(f, "MaterialId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMaterialListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaterialListResponseParams struct {
	// 素材列表数据
	MaterialInfos []*PublicMaterialInfos `json:"MaterialInfos,omitnil,omitempty" name:"MaterialInfos"`

	// 素材条数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMaterialListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMaterialListResponseParams `json:"Response"`
}

func (r *DescribeMaterialListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaterialListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FaceInfo struct {
	// 人脸框的横坐标
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// 人脸框的纵坐标
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// 人脸框的宽度
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 人脸框的高度
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`
}

type FaceRect struct {
	// 人脸框左上角横坐标。
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// 人脸框左上角纵坐标。
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// 人脸框宽度。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 人脸框高度。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`
}

// Predefined struct for user
type FuseFaceRequestParams struct {
	// 活动 ID，请在<a href="https://console.cloud.tencent.com/facefusion" target="_blank">人脸融合控制台</a>查看。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 素材 ID，请在<a href="https://console.cloud.tencent.com/facefusion" target="_blank">人脸融合控制台</a>查看。
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 返回图像方式（url 或 base64) ，二选一。url有效期为7天。
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`

	// 用户人脸图片、素材模板图的人脸位置信息。
	MergeInfos []*MergeInfo `json:"MergeInfos,omitnil,omitempty" name:"MergeInfos"`

	// 脸型融合比例，数值越高，融合后的脸型越像素材人物。取值范围[0,100] 
	// 若此参数不填写，则使用人脸融合控制台中脸型参数数值。（换脸版算法暂不支持此参数调整）
	FuseProfileDegree *int64 `json:"FuseProfileDegree,omitnil,omitempty" name:"FuseProfileDegree"`

	// 五官融合比例，数值越高，融合后的五官越像素材人物。取值范围[0,100] 
	// 若此参数不填写，则使用人脸融合控制台中五官参数数值。（换脸版算法暂不支持此参数调整）
	FuseFaceDegree *int64 `json:"FuseFaceDegree,omitnil,omitempty" name:"FuseFaceDegree"`

	// 为融合结果图添加合成标识的开关，默认为1。
	// 1：添加标识。
	// 0：不添加标识。
	// 其他数值：默认按1处理。
	// 建议您使用显著标识来提示结果图使用了人脸融合技术，是AI合成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在融合结果图右下角添加“本图片为AI合成图片”字样，您可根据自身需要替换为其他的Logo图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// 融合参数。
	FuseParam *FuseParam `json:"FuseParam,omitnil,omitempty" name:"FuseParam"`
}

type FuseFaceRequest struct {
	*tchttp.BaseRequest
	
	// 活动 ID，请在<a href="https://console.cloud.tencent.com/facefusion" target="_blank">人脸融合控制台</a>查看。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 素材 ID，请在<a href="https://console.cloud.tencent.com/facefusion" target="_blank">人脸融合控制台</a>查看。
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 返回图像方式（url 或 base64) ，二选一。url有效期为7天。
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`

	// 用户人脸图片、素材模板图的人脸位置信息。
	MergeInfos []*MergeInfo `json:"MergeInfos,omitnil,omitempty" name:"MergeInfos"`

	// 脸型融合比例，数值越高，融合后的脸型越像素材人物。取值范围[0,100] 
	// 若此参数不填写，则使用人脸融合控制台中脸型参数数值。（换脸版算法暂不支持此参数调整）
	FuseProfileDegree *int64 `json:"FuseProfileDegree,omitnil,omitempty" name:"FuseProfileDegree"`

	// 五官融合比例，数值越高，融合后的五官越像素材人物。取值范围[0,100] 
	// 若此参数不填写，则使用人脸融合控制台中五官参数数值。（换脸版算法暂不支持此参数调整）
	FuseFaceDegree *int64 `json:"FuseFaceDegree,omitnil,omitempty" name:"FuseFaceDegree"`

	// 为融合结果图添加合成标识的开关，默认为1。
	// 1：添加标识。
	// 0：不添加标识。
	// 其他数值：默认按1处理。
	// 建议您使用显著标识来提示结果图使用了人脸融合技术，是AI合成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在融合结果图右下角添加“本图片为AI合成图片”字样，您可根据自身需要替换为其他的Logo图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// 融合参数。
	FuseParam *FuseParam `json:"FuseParam,omitnil,omitempty" name:"FuseParam"`
}

func (r *FuseFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FuseFaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ModelId")
	delete(f, "RspImgType")
	delete(f, "MergeInfos")
	delete(f, "FuseProfileDegree")
	delete(f, "FuseFaceDegree")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	delete(f, "FuseParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FuseFaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FuseFaceResponseParams struct {
	// RspImgType 为 url 时，返回结果的 url， RspImgType 为 base64 时返回 base64 数据。
	FusedImage *string `json:"FusedImage,omitnil,omitempty" name:"FusedImage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FuseFaceResponse struct {
	*tchttp.BaseResponse
	Response *FuseFaceResponseParams `json:"Response"`
}

func (r *FuseFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FuseFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FuseParam struct {
	// 图片编码参数
	ImageCodecParam *ImageCodecParam `json:"ImageCodecParam,omitnil,omitempty" name:"ImageCodecParam"`
}

type ImageCodecParam struct {
	// 元数据
	MetaData []*MetaData `json:"MetaData,omitnil,omitempty" name:"MetaData"`
}

type LogoParam struct {
	// 标识图片位于融合结果图中的坐标，将按照坐标对标识图片进行位置和大小的拉伸匹配
	LogoRect *FaceRect `json:"LogoRect,omitnil,omitempty" name:"LogoRect"`

	// 标识图片Url地址
	LogoUrl *string `json:"LogoUrl,omitnil,omitempty" name:"LogoUrl"`

	// 标识图片base64
	LogoImage *string `json:"LogoImage,omitnil,omitempty" name:"LogoImage"`
}

type MaterialFaces struct {
	// 人脸序号
	FaceId *string `json:"FaceId,omitnil,omitempty" name:"FaceId"`

	// 人脸框信息
	FaceInfo *FaceInfo `json:"FaceInfo,omitnil,omitempty" name:"FaceInfo"`
}

type MergeInfo struct {
	// 输入图片base64
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 输入图片url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 上传的图片人脸位置信息（人脸框）
	InputImageFaceRect *FaceRect `json:"InputImageFaceRect,omitnil,omitempty" name:"InputImageFaceRect"`

	// 素材人脸ID，不填默认取最大人脸。
	TemplateFaceID *string `json:"TemplateFaceID,omitnil,omitempty" name:"TemplateFaceID"`

	// 模版中人脸位置信息(人脸框)，不填默认取最大人脸。此字段仅适用于图片融合自定义模版素材场景。
	TemplateFaceRect *FaceRect `json:"TemplateFaceRect,omitnil,omitempty" name:"TemplateFaceRect"`
}

type MetaData struct {
	// MetaData的Key
	MetaKey *string `json:"MetaKey,omitnil,omitempty" name:"MetaKey"`

	// MetaData的Value
	MetaValue *string `json:"MetaValue,omitnil,omitempty" name:"MetaValue"`
}

type PublicMaterialInfos struct {
	// 素材Id
	MaterialId *string `json:"MaterialId,omitnil,omitempty" name:"MaterialId"`

	// 素材状态
	MaterialStatus *int64 `json:"MaterialStatus,omitnil,omitempty" name:"MaterialStatus"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 人脸信息
	MaterialFaceList []*MaterialFaces `json:"MaterialFaceList,omitnil,omitempty" name:"MaterialFaceList"`

	// 素材名
	MaterialName *string `json:"MaterialName,omitnil,omitempty" name:"MaterialName"`

	// 审核原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuditResult *string `json:"AuditResult,omitnil,omitempty" name:"AuditResult"`
}