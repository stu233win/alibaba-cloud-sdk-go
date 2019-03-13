package edas

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

// BindServerlessSlb invokes the edas.BindServerlessSlb API synchronously
// api document: https://help.aliyun.com/api/edas/bindserverlessslb.html
func (client *Client) BindServerlessSlb(request *BindServerlessSlbRequest) (response *BindServerlessSlbResponse, err error) {
	response = CreateBindServerlessSlbResponse()
	err = client.DoAction(request, response)
	return
}

// BindServerlessSlbWithChan invokes the edas.BindServerlessSlb API asynchronously
// api document: https://help.aliyun.com/api/edas/bindserverlessslb.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BindServerlessSlbWithChan(request *BindServerlessSlbRequest) (<-chan *BindServerlessSlbResponse, <-chan error) {
	responseChan := make(chan *BindServerlessSlbResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BindServerlessSlb(request)
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

// BindServerlessSlbWithCallback invokes the edas.BindServerlessSlb API asynchronously
// api document: https://help.aliyun.com/api/edas/bindserverlessslb.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BindServerlessSlbWithCallback(request *BindServerlessSlbRequest, callback func(response *BindServerlessSlbResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BindServerlessSlbResponse
		var err error
		defer close(result)
		response, err = client.BindServerlessSlb(request)
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

// BindServerlessSlbRequest is the request struct for api BindServerlessSlb
type BindServerlessSlbRequest struct {
	*requests.RoaRequest
	Intranet string `position:"Query" name:"Intranet"`
	AppId    string `position:"Query" name:"AppId"`
	Internet string `position:"Query" name:"Internet"`
}

// BindServerlessSlbResponse is the response struct for api BindServerlessSlb
type BindServerlessSlbResponse struct {
	*responses.BaseResponse
	Code          int    `json:"Code" xml:"Code"`
	Message       string `json:"Message" xml:"Message"`
	ChangeOrderId string `json:"ChangeOrderId" xml:"ChangeOrderId"`
}

// CreateBindServerlessSlbRequest creates a request to invoke BindServerlessSlb API
func CreateBindServerlessSlbRequest() (request *BindServerlessSlbRequest) {
	request = &BindServerlessSlbRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "BindServerlessSlb", "/pop/v5/k8s/acs/serverless_slb_binding", "edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBindServerlessSlbResponse creates a response to parse from BindServerlessSlb response
func CreateBindServerlessSlbResponse() (response *BindServerlessSlbResponse) {
	response = &BindServerlessSlbResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}