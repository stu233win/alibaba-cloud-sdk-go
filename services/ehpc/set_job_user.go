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

func (client *Client) SetJobUser(request *SetJobUserRequest) (response *SetJobUserResponse, err error) {
	response = CreateSetJobUserResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SetJobUserWithChan(request *SetJobUserRequest) (<-chan *SetJobUserResponse, <-chan error) {
	responseChan := make(chan *SetJobUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetJobUser(request)
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

func (client *Client) SetJobUserWithCallback(request *SetJobUserRequest, callback func(response *SetJobUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetJobUserResponse
		var err error
		defer close(result)
		response, err = client.SetJobUser(request)
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

type SetJobUserRequest struct {
	*requests.RpcRequest
	RunasUserPassword string `position:"Query" name:"RunasUserPassword"`
	RunasUser         string `position:"Query" name:"RunasUser"`
	ClusterId         string `position:"Query" name:"ClusterId"`
}

type SetJobUserResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateSetJobUserRequest() (request *SetJobUserRequest) {
	request = &SetJobUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2017-07-14", "SetJobUser", "ehs", "openAPI")
	return
}

func CreateSetJobUserResponse() (response *SetJobUserResponse) {
	response = &SetJobUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
