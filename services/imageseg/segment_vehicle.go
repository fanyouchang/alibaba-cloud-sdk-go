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

// SegmentVehicle invokes the imageseg.SegmentVehicle API synchronously
// api document: https://help.aliyun.com/api/imageseg/segmentvehicle.html
func (client *Client) SegmentVehicle(request *SegmentVehicleRequest) (response *SegmentVehicleResponse, err error) {
	response = CreateSegmentVehicleResponse()
	err = client.DoAction(request, response)
	return
}

// SegmentVehicleWithChan invokes the imageseg.SegmentVehicle API asynchronously
// api document: https://help.aliyun.com/api/imageseg/segmentvehicle.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SegmentVehicleWithChan(request *SegmentVehicleRequest) (<-chan *SegmentVehicleResponse, <-chan error) {
	responseChan := make(chan *SegmentVehicleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SegmentVehicle(request)
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

// SegmentVehicleWithCallback invokes the imageseg.SegmentVehicle API asynchronously
// api document: https://help.aliyun.com/api/imageseg/segmentvehicle.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SegmentVehicleWithCallback(request *SegmentVehicleRequest, callback func(response *SegmentVehicleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SegmentVehicleResponse
		var err error
		defer close(result)
		response, err = client.SegmentVehicle(request)
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

// SegmentVehicleRequest is the request struct for api SegmentVehicle
type SegmentVehicleRequest struct {
	*requests.RpcRequest
	ImageURL string `position:"Body" name:"ImageURL"`
}

// SegmentVehicleResponse is the response struct for api SegmentVehicle
type SegmentVehicleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSegmentVehicleRequest creates a request to invoke SegmentVehicle API
func CreateSegmentVehicleRequest() (request *SegmentVehicleRequest) {
	request = &SegmentVehicleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imageseg", "2019-12-30", "SegmentVehicle", "imageseg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSegmentVehicleResponse creates a response to parse from SegmentVehicle response
func CreateSegmentVehicleResponse() (response *SegmentVehicleResponse) {
	response = &SegmentVehicleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
