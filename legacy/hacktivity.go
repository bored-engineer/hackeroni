// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package legacy

// HacktivityService handles communication with the report related methods of the H1 API.
type HacktivityService service

// Define some convinience functions
const (
	HacktivityFilterDisclosed                     string = "type:public"
	HacktivityFilterBugBounty                     string = "type:bounty-awarded"
	HacktivityFilterAll                           string = "type:all"
	HacktivitySortTypeLatestDisclosableActivityAt string = "latest_disclosable_activity_at"
)

// HacktivityListOptions
type HacktivityListOptions struct {
	SortType string `url:"sort_type"`
	Page     uint64 `url:"page"`
	Filter   string `url:"filter"`
}

// List hacktivity matching given criteria
func (s *HacktivityService) List(opts HacktivityListOptions) ([]*Report, *Response, error) {
	u, _ := addOptions("hacktivity", &opts)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")

	hacktivityResponse := struct {
		Reports []*Report `json:"reports"`
	}{}
	resp, err := s.client.Do(req, &hacktivityResponse)
	if err != nil {
		return nil, resp, err
	}

	return hacktivityResponse.Reports, resp, err
}
