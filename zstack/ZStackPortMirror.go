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

// 创建端口镜像
func CreatePortMirror(params model.CreatePortMirrorRequest) models.Result[any] {
	//POST zstack/v1/port-mirrors
	url := fmt.Sprintf("zstack/v1/port-mirrors")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreatePortMirrorFailed, err.Error())
	}
	var result model.CreatePortMirrorResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success[any](result)
}

// 查询端口镜像
func QueryPortMirror(params model.QueryPortMirrorRequest) models.Result[any] {
	//GET zstack/v1/port-mirrors
	//GET zstack/v1/portMirrors/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/port-mirrors")
	} else {
		url = fmt.Sprintf("zstack/v1/port-mirrors/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryPortMirrorFailed, err.Error())
	}
	var result model.QueryPortMirrorResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success[any](result)
}

// 更新端口镜像
func UpdatePortMirror(params model.UpdatePortMirrorRequest) models.Result[any] {
	//PUT zstack/v1/port-mirrors/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/port-mirrors/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.UpdatePortMirrorFailed, err.Error())
	}
	var result model.UpdatePortMirrorResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success[any](result)
}

// 删除端口镜像
func DeletePortMirror(params model.DeletePortMirrorRequest) models.Result[any] {
	//DELETE zstack/v1/port-mirrors/{uuid}
	url := fmt.Sprintf("zstack/v1/port-mirrors/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeletePortMirrorFailed, err.Error())
	}
	var result model.DeletePortMirrorResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success[any](result)
}

// 创建端口镜像会话
func CreatePortMirrorSession(params model.CreatePortMirrorSessionRequest) models.Result[any] {
	//POST zstack/v1/port-mirrors/sessions
	url := fmt.Sprintf("zstack/v1/port-mirrors/sessions")

	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(errcode.CreatePortMirrorSessionFailed, err.Error())
	}
	var result model.CreatePortMirrorSessionResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success[any](result)
}

// 查询端口镜像会话
func QueryPortMirrorSession(params model.QueryPortMirrorSessionRequest) models.Result[any] {
	//GET zstack/v1/port-mirrors/sessions
	//GET zstack/v1/port-mirrors/sessions/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/port-mirrors/sessions")
	} else {
		url = fmt.Sprintf("zstack/v1/port-mirrors/sessions/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryPortMirrorSessionFailed, err.Error())
	}
	var result model.QueryPortMirrorSessionResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success[any](result)
}

// 更新端口镜像服务状态
func ChangePortMirrorState(params model.ChangePortMirrorStateRequest) models.Result[any] {
	//PUT zstack/v1/port-mirrors/{uuid}/actions
	url := fmt.Sprintf("zstack/v1/port-mirrors/%s/actions", params.UUID)

	dataStr, err := request.Put(url, params)
	if err != nil {
		return models.Error(errcode.ChangePortMirrorStateFailed, err.Error())
	}
	var result model.ChangePortMirrorStateResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success[any](result)
}

// 删除端口镜像会话
func DeletePortMirrorSession(params model.DeletePortMirrorSessionRequest) models.Result[any] {
	//DELETE zstack/v1/port-mirrors/sessons/{uuid}
	url := fmt.Sprintf("zstack/v1/port-mirrors/sessons/%s", params.UUID)

	dataStr, err := request.Delete(url, params)
	if err != nil {
		return models.Error(errcode.DeletePortMirrorSessionFailed, err.Error())
	}
	var result model.DeletePortMirrorSessionResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success[any](result)
}

// 获取系统中可以使用端口镜像服务的网卡
func GetCandidateVmNicsForPortMirror(params model.GetCandidateVmNicsForPortMirrorRequest) models.Result[any] {
	//GET zstack/v1/port-mirrors/{portMirrorUuid}/vm-instances/candidate-nics/{type}
	url := fmt.Sprintf("zstack/v1/port-mirrors/%s/vm-instances/candidate-nics/%s", params.PortMirrorUuid, params.Type)

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.GetCandidateVmNicsForPortMirrorFailed, err.Error())
	}
	var result model.GetCandidateVmNicsForPortMirrorResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success[any](result)
}

// 查询端口镜像网络分配的IP
func QueryPortMirrorNetworkUsedIp(params model.QueryPortMirrorNetworkUsedIpRequest) models.Result[any] {
	//GET zstack/v1/port-mirrors/networks/usedIps
	//GET zstack/v1/port-mirrors/networks/usedIps/{uuid}
	var url string
	if params.UUID == "" {
		url = fmt.Sprintf("zstack/v1/port-mirrors/networks/usedIps")
	} else {
		url = fmt.Sprintf("zstack/v1/port-mirrors/networks/usedIps/%s", params.UUID)
	}

	dataStr, err := request.Get(url, params)
	if err != nil {
		return models.Error(errcode.QueryPortMirrorNetworkUsedIpFailed, err.Error())
	}
	var result model.QueryPortMirrorNetworkUsedIpResponse
	utils.FromJSON(dataStr, &result)
	logs.Debug("最终结果:{}", result)

	return models.Success[any](result)
}
