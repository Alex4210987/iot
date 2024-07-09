package util

import (
	iotda "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotda/v5"
	hwmodel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotda/v5/model"

)

func SendIoTCommand(client *iotda.IoTDAClient, commandParams string, commandName string, serviceId string) (*hwmodel.CreateCommandResponse, error) {
	var paras interface{} = commandParams
	
	request := &hwmodel.CreateCommandRequest{}
	request.Body = &hwmodel.DeviceCommandRequest{
		Paras:       &paras,
		CommandName: &commandName,
		ServiceId:   &serviceId,
	}

	response, err := client.CreateCommand(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}