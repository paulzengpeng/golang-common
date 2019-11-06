package base_server_octopus

import (
	"encoding/json"
	"github.com/becent/golang-common/base-server-sdk"
	"strconv"
)

//orgId 组织id
// formFile 需要上传的文件列表：key是field的name，value是文件路径
// formFile := make(map[string]string)
// formFile["file1"] = "path/to/test.log"
// formFile["file2"] = "path/to/test.2.log"
// 返回
//	result = {
//		"file1": "https://xxx.com/path/to/file1",
//		"file2": "https://xxx.com/path/to/file2"
//	}
func Upload(orgId int, formFile map[string]string) (map[string]string, *base_server_sdk.Error) {
	if orgId <= 0 || len(formFile) == 0 {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.OctopusServerHost, "resource", "upload", params, formFile)
	if err != nil {
		return nil, err
	}
	var resp map[string]string
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, base_server_sdk.ErrServiceBusy
	}

	return resp, nil
}
