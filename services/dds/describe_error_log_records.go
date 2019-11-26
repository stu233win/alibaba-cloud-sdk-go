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

// DescribeErrorLogRecords invokes the dds.DescribeErrorLogRecords API synchronously
// api document: https://help.aliyun.com/api/dds/describeerrorlogrecords.html
func (client *Client) DescribeErrorLogRecords(request *DescribeErrorLogRecordsRequest) (response *DescribeErrorLogRecordsResponse, err error) {
	response = CreateDescribeErrorLogRecordsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeErrorLogRecordsWithChan invokes the dds.DescribeErrorLogRecords API asynchronously
// api document: https://help.aliyun.com/api/dds/describeerrorlogrecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeErrorLogRecordsWithChan(request *DescribeErrorLogRecordsRequest) (<-chan *DescribeErrorLogRecordsResponse, <-chan error) {
	responseChan := make(chan *DescribeErrorLogRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeErrorLogRecords(request)
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

// DescribeErrorLogRecordsWithCallback invokes the dds.DescribeErrorLogRecords API asynchronously
// api document: https://help.aliyun.com/api/dds/describeerrorlogrecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeErrorLogRecordsWithCallback(request *DescribeErrorLogRecordsRequest, callback func(response *DescribeErrorLogRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeErrorLogRecordsResponse
		var err error
		defer close(result)
		response, err = client.DescribeErrorLogRecords(request)
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

// DescribeErrorLogRecordsRequest is the request struct for api DescribeErrorLogRecords
type DescribeErrorLogRecordsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	RoleType             string           `position:"Query" name:"RoleType"`
	NodeId               string           `position:"Query" name:"NodeId"`
	SQLId                requests.Integer `position:"Query" name:"SQLId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBName               string           `position:"Query" name:"DBName"`
}

// DescribeErrorLogRecordsResponse is the response struct for api DescribeErrorLogRecords
type DescribeErrorLogRecordsResponse struct {
	*responses.BaseResponse
	RequestId        string                         `json:"RequestId" xml:"RequestId"`
	Engine           string                         `json:"Engine" xml:"Engine"`
	TotalRecordCount int                            `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       int                            `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  int                            `json:"PageRecordCount" xml:"PageRecordCount"`
	Items            ItemsInDescribeErrorLogRecords `json:"Items" xml:"Items"`
}

// CreateDescribeErrorLogRecordsRequest creates a request to invoke DescribeErrorLogRecords API
func CreateDescribeErrorLogRecordsRequest() (request *DescribeErrorLogRecordsRequest) {
	request = &DescribeErrorLogRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "DescribeErrorLogRecords", "", "")
	return
}

// CreateDescribeErrorLogRecordsResponse creates a response to parse from DescribeErrorLogRecords response
func CreateDescribeErrorLogRecordsResponse() (response *DescribeErrorLogRecordsResponse) {
	response = &DescribeErrorLogRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
