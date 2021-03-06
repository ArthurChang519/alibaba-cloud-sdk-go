package emr

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

// ListAdviceAction invokes the emr.ListAdviceAction API synchronously
func (client *Client) ListAdviceAction(request *ListAdviceActionRequest) (response *ListAdviceActionResponse, err error) {
	response = CreateListAdviceActionResponse()
	err = client.DoAction(request, response)
	return
}

// ListAdviceActionWithChan invokes the emr.ListAdviceAction API asynchronously
func (client *Client) ListAdviceActionWithChan(request *ListAdviceActionRequest) (<-chan *ListAdviceActionResponse, <-chan error) {
	responseChan := make(chan *ListAdviceActionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAdviceAction(request)
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

// ListAdviceActionWithCallback invokes the emr.ListAdviceAction API asynchronously
func (client *Client) ListAdviceActionWithCallback(request *ListAdviceActionRequest, callback func(response *ListAdviceActionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAdviceActionResponse
		var err error
		defer close(result)
		response, err = client.ListAdviceAction(request)
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

// ListAdviceActionRequest is the request struct for api ListAdviceAction
type ListAdviceActionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	Component       string           `position:"Query" name:"Component"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	ServiceName     string           `position:"Query" name:"ServiceName"`
}

// ListAdviceActionResponse is the response struct for api ListAdviceAction
type ListAdviceActionResponse struct {
	*responses.BaseResponse
	RequestId  string                  `json:"RequestId" xml:"RequestId"`
	PageNumber int                     `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                     `json:"PageSize" xml:"PageSize"`
	TotalCount int                     `json:"TotalCount" xml:"TotalCount"`
	Items      ItemsInListAdviceAction `json:"Items" xml:"Items"`
}

// CreateListAdviceActionRequest creates a request to invoke ListAdviceAction API
func CreateListAdviceActionRequest() (request *ListAdviceActionRequest) {
	request = &ListAdviceActionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListAdviceAction", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAdviceActionResponse creates a response to parse from ListAdviceAction response
func CreateListAdviceActionResponse() (response *ListAdviceActionResponse) {
	response = &ListAdviceActionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
