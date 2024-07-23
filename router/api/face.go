package api

import (
	"backend/util"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	frs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/frs/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/frs/v2/model"
	frsRegion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/frs/v2/region"
)

var FaceClient *frs.FrsClient

func InitFaceClient() {
	ak := os.Getenv("FACE_CLOUD_SDK_AK")
	sk := os.Getenv("FACE_CLOUD_SDK_SK")
	if ak == "" || sk == "" {
		panic("FACE_CLOUD_SDK_AK or FACE_CLOUD_SDK_SK environment variables are not set")
	}
	auth, _ := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		SafeBuild()

	region, _ := frsRegion.SafeValueOf("cn-east-3")

	builder, _ := frs.FrsClientBuilder().
		WithRegion(region).
		WithCredential(auth).
		SafeBuild()

	FaceClient = frs.NewFrsClient(builder)
}

func AddFaceHandler(c *gin.Context) {
	if FaceClient == nil {
		InitFaceClient()
	}
	//
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the uploaded file to a temporary location
	tempFilePath := filepath.Join(os.TempDir(), file.Filename)

	if err := c.SaveUploadedFile(file, tempFilePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
		return
	}
	// Create a FilePart from the saved file path
	fileContent, err := os.Open(tempFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open file"})
		return
	}
	defer fileContent.Close()

	// Add face to face set
	addRequest := &model.AddFacesByFileRequest{
		FaceSetName: "test", // Replace with your face set name
		Body: &model.AddFacesByFileRequestBody{
			ImageFile: def.NewFilePart(fileContent),
		},
	}

	addResponse, err := FaceClient.AddFacesByFile(addRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Add faces response: %+v\n", addResponse)
	c.JSON(http.StatusOK, gin.H{"message": "人脸添加成功"})
}

func SearchFaceHandler(c *gin.Context) {
	if FaceClient == nil {
		InitFaceClient()
	}

	contentType := c.GetHeader("Content-Type")
	if contentType != "image/jpeg" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Content-Type. Expected image/jpeg"})
		return
	}

	fileContent, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	// Save the uploaded file to a temporary location
	tempFilePath := filepath.Join(os.TempDir(), "uploaded.jpg")
	err = os.WriteFile(tempFilePath, fileContent, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	fileContentFile, err := os.Open(tempFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}

	// Search face in face set
	searchRequest := &model.SearchFaceByFileRequest{
		FaceSetName: "test", // Replace with your face set name
		Body: &model.SearchFaceByFileRequestBody{
			ImageFile: def.NewFilePart(fileContentFile),
		},
	}

	searchResponse, err := FaceClient.SearchFaceByFile(searchRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	flag := false
	for _, face := range *searchResponse.Faces {
		fmt.Printf("Face similarity: %f\n", *face.Similarity)
		if *face.Similarity > 0.85 {
			flag = true
		}
	}

	commandParams := map[string]interface{}{
		"access_control_switch": flag,
	}

	if HWClient == nil {
		SettingUpEnvironment()
		InitHuaweiCloudClient()
	}
	util.SendIoTCommand(HWClient, DeviceId, commandParams, "access_control_commands", "personnal_access")
	c.JSON(http.StatusOK, searchResponse)
}
