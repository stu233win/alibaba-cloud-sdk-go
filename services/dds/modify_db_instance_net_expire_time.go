package dds

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

// ModifyDBInstanceNetExpireTime invokes the dds.ModifyDBInstanceNetExpireTime API synchronously
// api document: https://help.aliyun.com/api/dds/modifydbinstancenetexpiretime.html
func (client *Client) ModifyDBInstanceNetExpireTime(request *ModifyDBInstanceNetExpireTimeRequest) (response *ModifyDBInstanceNetExpireTimeResponse, err error) {
	response = CreateModifyDBInstanceNetExpireTimeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBInstanceNetExpireTimeWithChan invokes the dds.ModifyDBInstanceNetExpireTime API asynchronously
// api document: https://help.aliyun.com/api/dds/modifydbinstancenetexpiretime.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBInstanceNetExpireTimeWithChan(request *ModifyDBInstanceNetExpireTimeRequest) (<-chan *ModifyDBInstanceNetExpireTimeResponse, <-chan error) {
	responseChan := make(chan *ModifyDBInstanceNetExpireTimeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBInstanceNetExpireTime(request)
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

// ModifyDBInstanceNetExpireTimeWithCallback invokes the dds.ModifyDBInstanceNetExpireTime API asynchronously
// api document: https://help.aliyun.com/api/dds/modifydbinstancenetexpiretime.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBInstanceNetExpireTimeWithCallback(request *ModifyDBInstanceNetExpireTimeRequest, callback func(response *ModifyDBInstanceNetExpireTimeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBInstanceNetExpireTimeResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBInstanceNetExpireTime(request)
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

// ModifyDBInstanceNetExpireTimeRequest is the request struct for api ModifyDBInstanceNetExpireTime
type ModifyDBInstanceNetExpireTimeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId          requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ConnectionString         string           `position:"Query" name:"ConnectionString"`
	ClassicExpendExpiredDays requests.Integer `position:"Query" name:"ClassicExpendExpiredDays"`
	SecurityToken            string           `position:"Query" name:"SecurityToken"`
	DBInstanceId             string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount     string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount             string           `position:"Query" name:"OwnerAccount"`
	OwnerId                  requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyDBInstanceNetExpireTimeResponse is the response struct for api ModifyDBInstanceNetExpireTime
type ModifyDBInstanceNetExpireTimeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBInstanceNetExpireTimeRequest creates a request to invoke ModifyDBInstanceNetExpireTime API
func CreateModifyDBInstanceNetExpireTimeRequest() (request *ModifyDBInstanceNetExpireTimeRequest) {
	request = &ModifyDBInstanceNetExpireTimeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "ModifyDBInstanceNetExpireTime", "", "")
	return
}

// CreateModifyDBInstanceNetExpireTimeResponse creates a response to parse from ModifyDBInstanceNetExpireTime response
func CreateModifyDBInstanceNetExpireTimeResponse() (response *ModifyDBInstanceNetExpireTimeResponse) {
	response = &ModifyDBInstanceNetExpireTimeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
