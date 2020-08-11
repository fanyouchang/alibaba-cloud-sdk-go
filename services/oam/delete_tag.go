package oam

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

// DeleteTag invokes the oam.DeleteTag API synchronously
// api document: https://help.aliyun.com/api/oam/deletetag.html
func (client *Client) DeleteTag(request *DeleteTagRequest) (response *DeleteTagResponse, err error) {
	response = CreateDeleteTagResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteTagWithChan invokes the oam.DeleteTag API asynchronously
// api document: https://help.aliyun.com/api/oam/deletetag.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteTagWithChan(request *DeleteTagRequest) (<-chan *DeleteTagResponse, <-chan error) {
	responseChan := make(chan *DeleteTagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteTag(request)
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

// DeleteTagWithCallback invokes the oam.DeleteTag API asynchronously
// api document: https://help.aliyun.com/api/oam/deletetag.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteTagWithCallback(request *DeleteTagRequest, callback func(response *DeleteTagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteTagResponse
		var err error
		defer close(result)
		response, err = client.DeleteTag(request)
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

// DeleteTagRequest is the request struct for api DeleteTag
type DeleteTagRequest struct {
	*requests.RpcRequest
	TagId *[]string `position:"Query" name:"TagId"  type:"Repeated"`
}

// DeleteTagResponse is the response struct for api DeleteTag
type DeleteTagResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteTagRequest creates a request to invoke DeleteTag API
func CreateDeleteTagRequest() (request *DeleteTagRequest) {
	request = &DeleteTagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oam", "2017-01-01", "DeleteTag", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteTagResponse creates a response to parse from DeleteTag response
func CreateDeleteTagResponse() (response *DeleteTagResponse) {
	response = &DeleteTagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
