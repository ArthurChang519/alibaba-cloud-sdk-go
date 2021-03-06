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

// BindIntervenePackageAndModel invokes the nlp_automl.BindIntervenePackageAndModel API synchronously
func (client *Client) BindIntervenePackageAndModel(request *BindIntervenePackageAndModelRequest) (response *BindIntervenePackageAndModelResponse, err error) {
	response = CreateBindIntervenePackageAndModelResponse()
	err = client.DoAction(request, response)
	return
}

// BindIntervenePackageAndModelWithChan invokes the nlp_automl.BindIntervenePackageAndModel API asynchronously
func (client *Client) BindIntervenePackageAndModelWithChan(request *BindIntervenePackageAndModelRequest) (<-chan *BindIntervenePackageAndModelResponse, <-chan error) {
	responseChan := make(chan *BindIntervenePackageAndModelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BindIntervenePackageAndModel(request)
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

// BindIntervenePackageAndModelWithCallback invokes the nlp_automl.BindIntervenePackageAndModel API asynchronously
func (client *Client) BindIntervenePackageAndModelWithCallback(request *BindIntervenePackageAndModelRequest, callback func(response *BindIntervenePackageAndModelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BindIntervenePackageAndModelResponse
		var err error
		defer close(result)
		response, err = client.BindIntervenePackageAndModel(request)
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

// BindIntervenePackageAndModelRequest is the request struct for api BindIntervenePackageAndModel
type BindIntervenePackageAndModelRequest struct {
	*requests.RpcRequest
	Product      string           `position:"Query" name:"Product"`
	ModelId      requests.Integer `position:"Query" name:"ModelId"`
	PackageId    requests.Integer `position:"Query" name:"PackageId"`
	TenantId     requests.Integer `position:"Query" name:"TenantId"`
	ProjectId    requests.Integer `position:"Query" name:"ProjectId"`
	ModelVersion string           `position:"Query" name:"ModelVersion"`
}

// BindIntervenePackageAndModelResponse is the response struct for api BindIntervenePackageAndModel
type BindIntervenePackageAndModelResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   int    `json:"Message" xml:"Message"`
	Success   string `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateBindIntervenePackageAndModelRequest creates a request to invoke BindIntervenePackageAndModel API
func CreateBindIntervenePackageAndModelRequest() (request *BindIntervenePackageAndModelRequest) {
	request = &BindIntervenePackageAndModelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("nlp-automl", "2019-07-01", "BindIntervenePackageAndModel", "nlpautoml", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBindIntervenePackageAndModelResponse creates a response to parse from BindIntervenePackageAndModel response
func CreateBindIntervenePackageAndModelResponse() (response *BindIntervenePackageAndModelResponse) {
	response = &BindIntervenePackageAndModelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
