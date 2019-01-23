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

// DescribeApis invokes the cloudapi.DescribeApis API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapis.html
func (client *Client) DescribeApis(request *DescribeApisRequest) (response *DescribeApisResponse, err error) {
	response = CreateDescribeApisResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeApisWithChan invokes the cloudapi.DescribeApis API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapis.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApisWithChan(request *DescribeApisRequest) (<-chan *DescribeApisResponse, <-chan error) {
	responseChan := make(chan *DescribeApisResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApis(request)
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

// DescribeApisWithCallback invokes the cloudapi.DescribeApis API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapis.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApisWithCallback(request *DescribeApisRequest, callback func(response *DescribeApisResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApisResponse
		var err error
		defer close(result)
		response, err = client.DescribeApis(request)
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

// DescribeApisRequest is the request struct for api DescribeApis
type DescribeApisRequest struct {
	*requests.RpcRequest
	ApiName       string           `position:"Query" name:"ApiName"`
	CatalogId     string           `position:"Query" name:"CatalogId"`
	Visibility    string           `position:"Query" name:"Visibility"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	GroupId       string           `position:"Query" name:"GroupId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	ApiId         string           `position:"Query" name:"ApiId"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribeApisResponse is the response struct for api DescribeApis
type DescribeApisResponse struct {
	*responses.BaseResponse
	RequestId   string                    `json:"RequestId" xml:"RequestId"`
	TotalCount  int                       `json:"TotalCount" xml:"TotalCount"`
	PageSize    int                       `json:"PageSize" xml:"PageSize"`
	PageNumber  int                       `json:"PageNumber" xml:"PageNumber"`
	ApiSummarys ApiSummarysInDescribeApis `json:"ApiSummarys" xml:"ApiSummarys"`
}

// CreateDescribeApisRequest creates a request to invoke DescribeApis API
func CreateDescribeApisRequest() (request *DescribeApisRequest) {
	request = &DescribeApisRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApis", "apigateway", "openAPI")
	return
}

// CreateDescribeApisResponse creates a response to parse from DescribeApis response
func CreateDescribeApisResponse() (response *DescribeApisResponse) {
	response = &DescribeApisResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
