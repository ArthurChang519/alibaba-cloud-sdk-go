package emas_appmonitor

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

// Point is a nested struct in emas_appmonitor response
type Point struct {
	Time             int64          `json:"Time" xml:"Time"`
	Data             float64        `json:"Data" xml:"Data"`
	ErrorInfo        ErrorInfo      `json:"ErrorInfo" xml:"ErrorInfo"`
	PointData        PointData      `json:"PointData" xml:"PointData"`
	CentroidItemList []CentroidItem `json:"CentroidItemList" xml:"CentroidItemList"`
}
