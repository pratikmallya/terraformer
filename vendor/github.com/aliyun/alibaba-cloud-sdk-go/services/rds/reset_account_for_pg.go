package rds

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

// ResetAccountForPG invokes the rds.ResetAccountForPG API synchronously
// api document: https://help.aliyun.com/api/rds/resetaccountforpg.html
func (client *Client) ResetAccountForPG(request *ResetAccountForPGRequest) (response *ResetAccountForPGResponse, err error) {
	response = CreateResetAccountForPGResponse()
	err = client.DoAction(request, response)
	return
}

// ResetAccountForPGWithChan invokes the rds.ResetAccountForPG API asynchronously
// api document: https://help.aliyun.com/api/rds/resetaccountforpg.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResetAccountForPGWithChan(request *ResetAccountForPGRequest) (<-chan *ResetAccountForPGResponse, <-chan error) {
	responseChan := make(chan *ResetAccountForPGResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResetAccountForPG(request)
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

// ResetAccountForPGWithCallback invokes the rds.ResetAccountForPG API asynchronously
// api document: https://help.aliyun.com/api/rds/resetaccountforpg.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResetAccountForPGWithCallback(request *ResetAccountForPGRequest, callback func(response *ResetAccountForPGResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResetAccountForPGResponse
		var err error
		defer close(result)
		response, err = client.ResetAccountForPG(request)
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

// ResetAccountForPGRequest is the request struct for api ResetAccountForPG
type ResetAccountForPGRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	AccountName          string           `position:"Query" name:"AccountName"`
	AccountPassword      string           `position:"Query" name:"AccountPassword"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

// ResetAccountForPGResponse is the response struct for api ResetAccountForPG
type ResetAccountForPGResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateResetAccountForPGRequest creates a request to invoke ResetAccountForPG API
func CreateResetAccountForPGRequest() (request *ResetAccountForPGRequest) {
	request = &ResetAccountForPGRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ResetAccountForPG", "rds", "openAPI")
	return
}

// CreateResetAccountForPGResponse creates a response to parse from ResetAccountForPG response
func CreateResetAccountForPGResponse() (response *ResetAccountForPGResponse) {
	response = &ResetAccountForPGResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}