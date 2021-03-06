package nlp_automl

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

// RunSmartCallService invokes the nlp_automl.RunSmartCallService API synchronously
func (client *Client) RunSmartCallService(request *RunSmartCallServiceRequest) (response *RunSmartCallServiceResponse, err error) {
	response = CreateRunSmartCallServiceResponse()
	err = client.DoAction(request, response)
	return
}

// RunSmartCallServiceWithChan invokes the nlp_automl.RunSmartCallService API asynchronously
func (client *Client) RunSmartCallServiceWithChan(request *RunSmartCallServiceRequest) (<-chan *RunSmartCallServiceResponse, <-chan error) {
	responseChan := make(chan *RunSmartCallServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RunSmartCallService(request)
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

// RunSmartCallServiceWithCallback invokes the nlp_automl.RunSmartCallService API asynchronously
func (client *Client) RunSmartCallServiceWithCallback(request *RunSmartCallServiceRequest, callback func(response *RunSmartCallServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RunSmartCallServiceResponse
		var err error
		defer close(result)
		response, err = client.RunSmartCallService(request)
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

// RunSmartCallServiceRequest is the request struct for api RunSmartCallService
type RunSmartCallServiceRequest struct {
	*requests.RpcRequest
	Product      string           `position:"Body" name:"Product"`
	SessionId    string           `position:"Body" name:"SessionId"`
	RobotId      requests.Integer `position:"Body" name:"RobotId"`
	ParamContent string           `position:"Body" name:"ParamContent"`
	ServiceName  string           `position:"Body" name:"ServiceName"`
}

// RunSmartCallServiceResponse is the response struct for api RunSmartCallService
type RunSmartCallServiceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateRunSmartCallServiceRequest creates a request to invoke RunSmartCallService API
func CreateRunSmartCallServiceRequest() (request *RunSmartCallServiceRequest) {
	request = &RunSmartCallServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("nlp-automl", "2019-11-11", "RunSmartCallService", "nlpautoml", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRunSmartCallServiceResponse creates a response to parse from RunSmartCallService response
func CreateRunSmartCallServiceResponse() (response *RunSmartCallServiceResponse) {
	response = &RunSmartCallServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
