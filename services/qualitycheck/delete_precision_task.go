package qualitycheck

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

// DeletePrecisionTask invokes the qualitycheck.DeletePrecisionTask API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/deleteprecisiontask.html
func (client *Client) DeletePrecisionTask(request *DeletePrecisionTaskRequest) (response *DeletePrecisionTaskResponse, err error) {
	response = CreateDeletePrecisionTaskResponse()
	err = client.DoAction(request, response)
	return
}

// DeletePrecisionTaskWithChan invokes the qualitycheck.DeletePrecisionTask API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/deleteprecisiontask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeletePrecisionTaskWithChan(request *DeletePrecisionTaskRequest) (<-chan *DeletePrecisionTaskResponse, <-chan error) {
	responseChan := make(chan *DeletePrecisionTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePrecisionTask(request)
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

// DeletePrecisionTaskWithCallback invokes the qualitycheck.DeletePrecisionTask API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/deleteprecisiontask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeletePrecisionTaskWithCallback(request *DeletePrecisionTaskRequest, callback func(response *DeletePrecisionTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePrecisionTaskResponse
		var err error
		defer close(result)
		response, err = client.DeletePrecisionTask(request)
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

// DeletePrecisionTaskRequest is the request struct for api DeletePrecisionTask
type DeletePrecisionTaskRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// DeletePrecisionTaskResponse is the response struct for api DeletePrecisionTask
type DeletePrecisionTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateDeletePrecisionTaskRequest creates a request to invoke DeletePrecisionTask API
func CreateDeletePrecisionTaskRequest() (request *DeletePrecisionTaskRequest) {
	request = &DeletePrecisionTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "DeletePrecisionTask", "", "")
	return
}

// CreateDeletePrecisionTaskResponse creates a response to parse from DeletePrecisionTask response
func CreateDeletePrecisionTaskResponse() (response *DeletePrecisionTaskResponse) {
	response = &DeletePrecisionTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}