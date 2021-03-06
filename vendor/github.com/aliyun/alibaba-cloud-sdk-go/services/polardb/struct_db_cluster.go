package polardb

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

// DBCluster is a nested struct in polardb response
type DBCluster struct {
	DBNodeClass          string                      `json:"DBNodeClass" xml:"DBNodeClass"`
	DeletionLock         int                         `json:"DeletionLock" xml:"DeletionLock"`
	DBType               string                      `json:"DBType" xml:"DBType"`
	DBVersion            string                      `json:"DBVersion" xml:"DBVersion"`
	Engine               string                      `json:"Engine" xml:"Engine"`
	PayType              string                      `json:"PayType" xml:"PayType"`
	CreateTime           string                      `json:"CreateTime" xml:"CreateTime"`
	DBClusterNetworkType string                      `json:"DBClusterNetworkType" xml:"DBClusterNetworkType"`
	DBClusterId          string                      `json:"DBClusterId" xml:"DBClusterId"`
	DBClusterStatus      string                      `json:"DBClusterStatus" xml:"DBClusterStatus"`
	RegionId             string                      `json:"RegionId" xml:"RegionId"`
	DeletedTime          string                      `json:"DeletedTime" xml:"DeletedTime"`
	ZoneId               string                      `json:"ZoneId" xml:"ZoneId"`
	DBNodeNumber         int                         `json:"DBNodeNumber" xml:"DBNodeNumber"`
	ResourceGroupId      string                      `json:"ResourceGroupId" xml:"ResourceGroupId"`
	StorageUsed          int64                       `json:"StorageUsed" xml:"StorageUsed"`
	IsDeleted            int                         `json:"IsDeleted" xml:"IsDeleted"`
	DBClusterDescription string                      `json:"DBClusterDescription" xml:"DBClusterDescription"`
	ExpireTime           string                      `json:"ExpireTime" xml:"ExpireTime"`
	VpcId                string                      `json:"VpcId" xml:"VpcId"`
	Expired              string                      `json:"Expired" xml:"Expired"`
	LockMode             string                      `json:"LockMode" xml:"LockMode"`
	DBNodes              DBNodesInDescribeDBClusters `json:"DBNodes" xml:"DBNodes"`
	Tags                 TagsInDescribeDBClusters    `json:"Tags" xml:"Tags"`
}
