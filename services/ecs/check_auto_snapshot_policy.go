package ecs

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

func (client *Client) CheckAutoSnapshotPolicy(request *CheckAutoSnapshotPolicyRequest) (response *CheckAutoSnapshotPolicyResponse, err error) {
	response = CreateCheckAutoSnapshotPolicyResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CheckAutoSnapshotPolicyWithChan(request *CheckAutoSnapshotPolicyRequest) (<-chan *CheckAutoSnapshotPolicyResponse, <-chan error) {
	responseChan := make(chan *CheckAutoSnapshotPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckAutoSnapshotPolicy(request)
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

func (client *Client) CheckAutoSnapshotPolicyWithCallback(request *CheckAutoSnapshotPolicyRequest, callback func(response *CheckAutoSnapshotPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckAutoSnapshotPolicyResponse
		var err error
		defer close(result)
		response, err = client.CheckAutoSnapshotPolicy(request)
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

type CheckAutoSnapshotPolicyRequest struct {
	*requests.RpcRequest
	DataDiskPolicyEnabled             requests.Boolean `position:"Query" name:"DataDiskPolicyEnabled"`
	ResourceOwnerId                   requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DataDiskPolicyRetentionDays       requests.Integer `position:"Query" name:"DataDiskPolicyRetentionDays"`
	ResourceOwnerAccount              string           `position:"Query" name:"ResourceOwnerAccount"`
	SystemDiskPolicyRetentionLastWeek requests.Boolean `position:"Query" name:"SystemDiskPolicyRetentionLastWeek"`
	OwnerAccount                      string           `position:"Query" name:"OwnerAccount"`
	SystemDiskPolicyTimePeriod        requests.Integer `position:"Query" name:"SystemDiskPolicyTimePeriod"`
	OwnerId                           requests.Integer `position:"Query" name:"OwnerId"`
	DataDiskPolicyRetentionLastWeek   requests.Boolean `position:"Query" name:"DataDiskPolicyRetentionLastWeek"`
	SystemDiskPolicyRetentionDays     requests.Integer `position:"Query" name:"SystemDiskPolicyRetentionDays"`
	DataDiskPolicyTimePeriod          requests.Integer `position:"Query" name:"DataDiskPolicyTimePeriod"`
	SystemDiskPolicyEnabled           requests.Boolean `position:"Query" name:"SystemDiskPolicyEnabled"`
}

type CheckAutoSnapshotPolicyResponse struct {
	*responses.BaseResponse
	RequestId              string `json:"RequestId" xml:"RequestId"`
	AutoSnapshotOccupation int    `json:"AutoSnapshotOccupation" xml:"AutoSnapshotOccupation"`
	IsPermittedModify      string `json:"IsPermittedModify" xml:"IsPermittedModify"`
}

func CreateCheckAutoSnapshotPolicyRequest() (request *CheckAutoSnapshotPolicyRequest) {
	request = &CheckAutoSnapshotPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CheckAutoSnapshotPolicy", "ecs", "openAPI")
	return
}

func CreateCheckAutoSnapshotPolicyResponse() (response *CheckAutoSnapshotPolicyResponse) {
	response = &CheckAutoSnapshotPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
