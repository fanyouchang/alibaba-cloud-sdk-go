package imageseg

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

// SegmentLogo invokes the imageseg.SegmentLogo API synchronously
// api document: https://help.aliyun.com/api/imageseg/segmentlogo.html
func (client *Client) SegmentLogo(request *SegmentLogoRequest) (response *SegmentLogoResponse, err error) {
	response = CreateSegmentLogoResponse()
	err = client.DoAction(request, response)
	return
}

// SegmentLogoWithChan invokes the imageseg.SegmentLogo API asynchronously
// api document: https://help.aliyun.com/api/imageseg/segmentlogo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SegmentLogoWithChan(request *SegmentLogoRequest) (<-chan *SegmentLogoResponse, <-chan error) {
	responseChan := make(chan *SegmentLogoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SegmentLogo(request)
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

// SegmentLogoWithCallback invokes the imageseg.SegmentLogo API asynchronously
// api document: https://help.aliyun.com/api/imageseg/segmentlogo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SegmentLogoWithCallback(request *SegmentLogoRequest, callback func(response *SegmentLogoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SegmentLogoResponse
		var err error
		defer close(result)
		response, err = client.SegmentLogo(request)
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

// SegmentLogoRequest is the request struct for api SegmentLogo
type SegmentLogoRequest struct {
	*requests.RpcRequest
	ImageURL string `position:"Query" name:"ImageURL"`
}

// SegmentLogoResponse is the response struct for api SegmentLogo
type SegmentLogoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSegmentLogoRequest creates a request to invoke SegmentLogo API
func CreateSegmentLogoRequest() (request *SegmentLogoRequest) {
	request = &SegmentLogoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imageseg", "2019-12-30", "SegmentLogo", "imageseg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSegmentLogoResponse creates a response to parse from SegmentLogo response
func CreateSegmentLogoResponse() (response *SegmentLogoResponse) {
	response = &SegmentLogoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
