package main

// import (
// 	"fmt"
// 	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
// 	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth"
//     iotda "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotda/v5"
// 	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotda/v5/model"
//     region "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
// )

// func main() {
//     // The AK and SK used for authentication are hard-coded or stored in plaintext, which has great security risks. It is recommended that the AK and SK be stored in ciphertext in configuration files or environment variables and decrypted during use to ensure security.
//     // In this example, AK and SK are stored in environment variables for authentication. Before running this example, set environment variables CLOUD_SDK_AK and CLOUD_SDK_SK in the local environment
//     ak := ""
//     sk := ""
//     // endpoint：请在控制台的"总览"界面的"平台接入地址"中查看"应用侧"的https接入地址
//     endpoint := ""

//     auth := basic.NewCredentialsBuilder().
//         WithAk(ak).
//         WithSk(sk).
// 		WithProjectId("a129d472438743eda97676e5a925f87d").
// 		// 企业版/标准版需要使用衍生算法，基础版请删除该配置"WithDerivedPredicate"
// 		WithDerivedPredicate(auth.GetDefaultDerivedPredicate()).
//         Build()

//     client := iotda.NewIoTDAClient(
//         iotda.IoTDAClientBuilder().
//             // 标准版/企业版需要自行创建region，基础版使用IoTDARegion中的region对象
//             WithRegion(region.NewRegion("cn-north-4", endpoint)).
//             WithCredential(auth).
//             Build())

//     request := &model.CreateCommandRequest{}
// 	request.DeviceId = "6663d853dbfd46fabbf54b9_device_"
// 	var parasDeviceCommandRequest interface{} = "{\"buzzer_switch\":true,\"window_switch\":true}"
// 	request.Body = &model.DeviceCommandRequest{
// 		Paras: &parasDeviceCommandRequest,
// 	}
// 	response, err := client.CreateCommand(request)
// 	if err == nil {
//         fmt.Printf("%+v\n", response)
//     } else {
//         fmt.Println(err)
//     }
// }