package zstack

import (
	"fmt"
	"github.com/gzxgogh/zstack-sdk-go19-v4/errcode"
	"github.com/gzxgogh/zstack-sdk-go19-v4/model"
	"github.com/gzxgogh/zstack-sdk-go19-v4/request"
	"github.com/maczh/mgin/logs"
	"github.com/maczh/mgin/models"
	"github.com/maczh/mgin/utils"
)

// 创建云主机规格
func CreateInstanceOffering(params model.CreateInstanceOfferingRequest) models.Result[any] {
	//POST zstack/v1/instance-offerings
	url := fmt.Sprintf("zstack/v1/instance-offerings")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreateInstanceOfferingFailed, err.Error())
	}
	var result model.CreateInstanceOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success[any](result)
}

// 删除云主机规格
func DeleteInstanceOffering(params model.DeleteInstanceOfferingRequest) models.Result[any] {
	//DELETE zstack/v1/instance-offerings/{uuid}
	url := fmt.Sprintf("zstack/v1/instance-offerings/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeleteInstanceOfferingFailed, err.Error())
	}
	var result model.DeleteInstanceOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success[any](result)
}

// 查询云主机规格
func QueryInstanceOffering(params model.QueryInstanceOfferingRequest) models.Result[any] {
	//GET zstack/v1/instance-offerings
	//GET zstack/v1/instance-offerings/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/instance-offerings")
	} else {
		url = fmt.Sprintf("zstack/v1/instance-offerings/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryInstanceOfferingFailed, err.Error())
	}
	var result model.QueryInstanceOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success[any](result)
}

// 更改云主机规格
func ChangeInstanceOffering(params model.ChangeInstanceOfferingRequest) models.Result[any] {
	//PUT zstack/v1/vm-instances/{vmInstanceUuid}/actions
	url := fmt.Sprintf("zstack/v1/vm-instances/%s/actions", params.VmInstanceUuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangeInstanceOfferingFailed, err.Error())
	}
	var result model.ChangeInstanceOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success[any](result)
}

// 更新云主机规格
func UpdateInstanceOffering(params model.UpdateInstanceOfferingRequest) models.Result[any] {
	//PUT zstack/v1/instance-offerings/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/instance-offerings/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdateInstanceOfferingFailed, err.Error())
	}
	var result model.UpdateInstanceOfferingResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success[any](result)
}

// 更改云主机规格的启用状态
func ChangeInstanceOfferingState(params model.ChangeInstanceOfferingStateRequest) models.Result[any] {
	//PUT zstack/v1/instance-offerings/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/instance-offerings/%s/actions", params.Uuid)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangeInstanceOfferingStateFailed, err.Error())
	}
	var result model.ChangeInstanceOfferingStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success[any](result)
}
