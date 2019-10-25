package smartag

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

// AddSnatEntry invokes the smartag.AddSnatEntry API synchronously
// api document: https://help.aliyun.com/api/smartag/addsnatentry.html
func (client *Client) AddSnatEntry(request *AddSnatEntryRequest) (response *AddSnatEntryResponse, err error) {
	response = CreateAddSnatEntryResponse()
	err = client.DoAction(request, response)
	return
}

// AddSnatEntryWithChan invokes the smartag.AddSnatEntry API asynchronously
// api document: https://help.aliyun.com/api/smartag/addsnatentry.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddSnatEntryWithChan(request *AddSnatEntryRequest) (<-chan *AddSnatEntryResponse, <-chan error) {
	responseChan := make(chan *AddSnatEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddSnatEntry(request)
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

// AddSnatEntryWithCallback invokes the smartag.AddSnatEntry API asynchronously
// api document: https://help.aliyun.com/api/smartag/addsnatentry.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddSnatEntryWithCallback(request *AddSnatEntryRequest, callback func(response *AddSnatEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddSnatEntryResponse
		var err error
		defer close(result)
		response, err = client.AddSnatEntry(request)
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

// AddSnatEntryRequest is the request struct for api AddSnatEntry
type AddSnatEntryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
	CidrBlock            string           `position:"Query" name:"CidrBlock"`
	SnatIp               string           `position:"Query" name:"SnatIp"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// AddSnatEntryResponse is the response struct for api AddSnatEntry
type AddSnatEntryResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	InstanceId   string `json:"InstanceId" xml:"InstanceId"`
}

// CreateAddSnatEntryRequest creates a request to invoke AddSnatEntry API
func CreateAddSnatEntryRequest() (request *AddSnatEntryRequest) {
	request = &AddSnatEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "AddSnatEntry", "smartag", "openAPI")
	return
}

// CreateAddSnatEntryResponse creates a response to parse from AddSnatEntry response
func CreateAddSnatEntryResponse() (response *AddSnatEntryResponse) {
	response = &AddSnatEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}