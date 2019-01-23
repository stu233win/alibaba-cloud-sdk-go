package cloudapi

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

// DescribePlugins invokes the cloudapi.DescribePlugins API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describeplugins.html
func (client *Client) DescribePlugins(request *DescribePluginsRequest) (response *DescribePluginsResponse, err error) {
	response = CreateDescribePluginsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePluginsWithChan invokes the cloudapi.DescribePlugins API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeplugins.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePluginsWithChan(request *DescribePluginsRequest) (<-chan *DescribePluginsResponse, <-chan error) {
	responseChan := make(chan *DescribePluginsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePlugins(request)
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

// DescribePluginsWithCallback invokes the cloudapi.DescribePlugins API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeplugins.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePluginsWithCallback(request *DescribePluginsRequest, callback func(response *DescribePluginsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePluginsResponse
		var err error
		defer close(result)
		response, err = client.DescribePlugins(request)
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

// DescribePluginsRequest is the request struct for api DescribePlugins
type DescribePluginsRequest struct {
	*requests.RpcRequest
	PluginType    string           `position:"Query" name:"PluginType"`
	PluginName    string           `position:"Query" name:"PluginName"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	PluginId      string           `position:"Query" name:"PluginId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribePluginsResponse is the response struct for api DescribePlugins
type DescribePluginsResponse struct {
	*responses.BaseResponse
	RequestId  string                   `json:"RequestId" xml:"RequestId"`
	TotalCount int                      `json:"TotalCount" xml:"TotalCount"`
	PageSize   int                      `json:"PageSize" xml:"PageSize"`
	PageNumber int                      `json:"PageNumber" xml:"PageNumber"`
	Plugins    PluginsInDescribePlugins `json:"Plugins" xml:"Plugins"`
}

// CreateDescribePluginsRequest creates a request to invoke DescribePlugins API
func CreateDescribePluginsRequest() (request *DescribePluginsRequest) {
	request = &DescribePluginsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribePlugins", "apigateway", "openAPI")
	return
}

// CreateDescribePluginsResponse creates a response to parse from DescribePlugins response
func CreateDescribePluginsResponse() (response *DescribePluginsResponse) {
	response = &DescribePluginsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
