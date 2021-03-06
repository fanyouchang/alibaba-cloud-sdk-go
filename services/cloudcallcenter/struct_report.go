package cloudcallcenter

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

// Report is a nested struct in cloudcallcenter response
type Report struct {
	PhoneNumber         string                      `json:"PhoneNumber" xml:"PhoneNumber"`
	Instance            string                      `json:"Instance" xml:"Instance"`
	OverallRealTime     OverallRealTime             `json:"OverallRealTime" xml:"OverallRealTime"`
	Summaries           []KeyValuePair              `json:"Summaries" xml:"Summaries"`
	Items               []BasicStatisticsReportItem `json:"Items" xml:"Items"`
	Indicators          []KeyValuePair              `json:"Indicators" xml:"Indicators"`
	Briefs              []KeyValuePair              `json:"Briefs" xml:"Briefs"`
	OverallHistory      OverallHistory              `json:"OverallHistory" xml:"OverallHistory"`
	SkillGroupsRealTime SkillGroupsRealTime         `json:"SkillGroupsRealTime" xml:"SkillGroupsRealTime"`
}
