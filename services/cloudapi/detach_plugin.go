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

// DetachPlugin invokes the cloudapi.DetachPlugin API synchronously
// api document: https://help.aliyun.com/api/cloudapi/detachplugin.html
func (client *Client) DetachPlugin(request *DetachPluginRequest) (response *DetachPluginResponse, err error) {
	response = CreateDetachPluginResponse()
	err = client.DoAction(request, response)
	return
}

// DetachPluginWithChan invokes the cloudapi.DetachPlugin API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/detachplugin.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DetachPluginWithChan(request *DetachPluginRequest) (<-chan *DetachPluginResponse, <-chan error) {
	responseChan := make(chan *DetachPluginResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachPlugin(request)
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

// DetachPluginWithCallback invokes the cloudapi.DetachPlugin API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/detachplugin.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DetachPluginWithCallback(request *DetachPluginRequest, callback func(response *DetachPluginResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachPluginResponse
		var err error
		defer close(result)
		response, err = client.DetachPlugin(request)
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

// DetachPluginRequest is the request struct for api DetachPlugin
type DetachPluginRequest struct {
	*requests.RpcRequest
	StageName     string `position:"Query" name:"StageName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	PluginId      string `position:"Query" name:"PluginId"`
	GroupId       string `position:"Query" name:"GroupId"`
	ApiId         string `position:"Query" name:"ApiId"`
	ApiIds        string `position:"Query" name:"ApiIds"`
}

// DetachPluginResponse is the response struct for api DetachPlugin
type DetachPluginResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDetachPluginRequest creates a request to invoke DetachPlugin API
func CreateDetachPluginRequest() (request *DetachPluginRequest) {
	request = &DetachPluginRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DetachPlugin", "apigateway", "openAPI")
	return
}

// CreateDetachPluginResponse creates a response to parse from DetachPlugin response
func CreateDetachPluginResponse() (response *DetachPluginResponse) {
	response = &DetachPluginResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
