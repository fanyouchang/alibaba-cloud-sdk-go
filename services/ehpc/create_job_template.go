package ehpc

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

// invoke CreateJobTemplate api with *CreateJobTemplateRequest synchronously
// api document: https://help.aliyun.com/api/ehpc/createjobtemplate.html
func (client *Client) CreateJobTemplate(request *CreateJobTemplateRequest) (response *CreateJobTemplateResponse, err error) {
	response = CreateCreateJobTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// invoke CreateJobTemplate api with *CreateJobTemplateRequest asynchronously
// api document: https://help.aliyun.com/api/ehpc/createjobtemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateJobTemplateWithChan(request *CreateJobTemplateRequest) (<-chan *CreateJobTemplateResponse, <-chan error) {
	responseChan := make(chan *CreateJobTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateJobTemplate(request)
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

// invoke CreateJobTemplate api with *CreateJobTemplateRequest asynchronously
// api document: https://help.aliyun.com/api/ehpc/createjobtemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateJobTemplateWithCallback(request *CreateJobTemplateRequest, callback func(response *CreateJobTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateJobTemplateResponse
		var err error
		defer close(result)
		response, err = client.CreateJobTemplate(request)
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

type CreateJobTemplateRequest struct {
	*requests.RpcRequest
	CommandLine        string           `position:"Query" name:"CommandLine"`
	Name               string           `position:"Query" name:"Name"`
	RunasUser          string           `position:"Query" name:"RunasUser"`
	Priority           requests.Integer `position:"Query" name:"Priority"`
	PackagePath        string           `position:"Query" name:"PackagePath"`
	StdoutRedirectPath string           `position:"Query" name:"StdoutRedirectPath"`
	StderrRedirectPath string           `position:"Query" name:"StderrRedirectPath"`
	ReRunable          requests.Boolean `position:"Query" name:"ReRunable"`
	ArrayRequest       string           `position:"Query" name:"ArrayRequest"`
	Variables          string           `position:"Query" name:"Variables"`
}

type CreateJobTemplateResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TemplateId string `json:"TemplateId" xml:"TemplateId"`
}

// create a request to invoke CreateJobTemplate API
func CreateCreateJobTemplateRequest() (request *CreateJobTemplateRequest) {
	request = &CreateJobTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2017-07-14", "CreateJobTemplate", "ehs", "openAPI")
	return
}

// create a response to parse from CreateJobTemplate response
func CreateCreateJobTemplateResponse() (response *CreateJobTemplateResponse) {
	response = &CreateJobTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}