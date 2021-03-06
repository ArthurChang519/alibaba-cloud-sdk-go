package retailcloud

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

// WorkLoad is a nested struct in retailcloud response
type WorkLoad struct {
	AllNodeCount        int    `json:"AllNodeCount" xml:"AllNodeCount"`
	CpuUse              string `json:"CpuUse" xml:"CpuUse"`
	CpuUsePercent       string `json:"CpuUsePercent" xml:"CpuUsePercent"`
	MemRequestPercent   string `json:"MemRequestPercent" xml:"MemRequestPercent"`
	MemUsePercent       string `json:"MemUsePercent" xml:"MemUsePercent"`
	AllocateAllPodCount int    `json:"AllocateAllPodCount" xml:"AllocateAllPodCount"`
	MemTotal            string `json:"MemTotal" xml:"MemTotal"`
	CpuLevel            string `json:"CpuLevel" xml:"CpuLevel"`
	CpuCapacityTotal    string `json:"CpuCapacityTotal" xml:"CpuCapacityTotal"`
	MemUse              string `json:"MemUse" xml:"MemUse"`
	MemCapacityTotal    string `json:"MemCapacityTotal" xml:"MemCapacityTotal"`
	MemLevel            string `json:"MemLevel" xml:"MemLevel"`
	CpuRequest          string `json:"CpuRequest" xml:"CpuRequest"`
	CpuTotal            string `json:"CpuTotal" xml:"CpuTotal"`
	CpuRequestPercent   string `json:"CpuRequestPercent" xml:"CpuRequestPercent"`
	AllocateAppPodCount int    `json:"AllocateAppPodCount" xml:"AllocateAppPodCount"`
	MemRequest          string `json:"MemRequest" xml:"MemRequest"`
}
