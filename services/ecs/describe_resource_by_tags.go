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

func (client *Client) DescribeResourceByTags(request *DescribeResourceByTagsRequest) (response *DescribeResourceByTagsResponse, err error) {
	response = CreateDescribeResourceByTagsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeResourceByTagsWithChan(request *DescribeResourceByTagsRequest) (<-chan *DescribeResourceByTagsResponse, <-chan error) {
	responseChan := make(chan *DescribeResourceByTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeResourceByTags(request)
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

func (client *Client) DescribeResourceByTagsWithCallback(request *DescribeResourceByTagsRequest, callback func(response *DescribeResourceByTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeResourceByTagsResponse
		var err error
		defer close(result)
		response, err = client.DescribeResourceByTags(request)
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

type DescribeResourceByTagsRequest struct {
	*requests.RpcRequest
	Tag4Value            string           `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Tag2Key              string           `position:"Query" name:"Tag.2.Key"`
	Tag5Key              string           `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Tag3Key              string           `position:"Query" name:"Tag.3.Key"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceType         string           `position:"Query" name:"ResourceType"`
	Tag5Value            string           `position:"Query" name:"Tag.5.Value"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	Tag1Key              string           `position:"Query" name:"Tag.1.Key"`
	Tag1Value            string           `position:"Query" name:"Tag.1.Value"`
	Tag2Value            string           `position:"Query" name:"Tag.2.Value"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	Tag4Key              string           `position:"Query" name:"Tag.4.Key"`
	Tag3Value            string           `position:"Query" name:"Tag.3.Value"`
}

type DescribeResourceByTagsResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	Resources  struct {
		Resource []struct {
			ResourceId   string `json:"ResourceId" xml:"ResourceId"`
			ResourceType string `json:"ResourceType" xml:"ResourceType"`
			RegionId     string `json:"RegionId" xml:"RegionId"`
		} `json:"Resource" xml:"Resource"`
	} `json:"Resources" xml:"Resources"`
}

func CreateDescribeResourceByTagsRequest() (request *DescribeResourceByTagsRequest) {
	request = &DescribeResourceByTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeResourceByTags", "ecs", "openAPI")
	return
}

func CreateDescribeResourceByTagsResponse() (response *DescribeResourceByTagsResponse) {
	response = &DescribeResourceByTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
