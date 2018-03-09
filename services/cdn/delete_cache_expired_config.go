package cdn

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// invoke DeleteCacheExpiredConfig api with *DeleteCacheExpiredConfigRequest synchronously
// api document: https://help.aliyun.com/api/cdn/deletecacheexpiredconfig.html
func (client *Client) DeleteCacheExpiredConfig(request *DeleteCacheExpiredConfigRequest) (response *DeleteCacheExpiredConfigResponse, err error) {
	response = CreateDeleteCacheExpiredConfigResponse()
	err = client.DoAction(request, response)
	return
}

// invoke DeleteCacheExpiredConfig api with *DeleteCacheExpiredConfigRequest asynchronously
// api document: https://help.aliyun.com/api/cdn/deletecacheexpiredconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCacheExpiredConfigWithChan(request *DeleteCacheExpiredConfigRequest) (<-chan *DeleteCacheExpiredConfigResponse, <-chan error) {
	responseChan := make(chan *DeleteCacheExpiredConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCacheExpiredConfig(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// invoke DeleteCacheExpiredConfig api with *DeleteCacheExpiredConfigRequest asynchronously
// api document: https://help.aliyun.com/api/cdn/deletecacheexpiredconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCacheExpiredConfigWithCallback(request *DeleteCacheExpiredConfigRequest, callback func(response *DeleteCacheExpiredConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCacheExpiredConfigResponse
		var err error
		defer close(result)
		response, err = client.DeleteCacheExpiredConfig(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type DeleteCacheExpiredConfigRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	CacheType     string           `position:"Query" name:"CacheType"`
	ConfigID      string           `position:"Query" name:"ConfigID"`
}

type DeleteCacheExpiredConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// create a request to invoke DeleteCacheExpiredConfig API
func CreateDeleteCacheExpiredConfigRequest() (request *DeleteCacheExpiredConfigRequest) {
	request = &DeleteCacheExpiredConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DeleteCacheExpiredConfig", "", "")
	return
}

// create a response to parse from DeleteCacheExpiredConfig response
func CreateDeleteCacheExpiredConfigResponse() (response *DeleteCacheExpiredConfigResponse) {
	response = &DeleteCacheExpiredConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}