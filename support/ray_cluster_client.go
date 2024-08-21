/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package support

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	. "github.com/onsi/gomega"
)

type RayJobSetup struct {
	EntryPoint string         `json:"entrypoint"`
	RuntimeEnv map[string]any `json:"runtime_env"`
}

type RayJobResponse struct {
	JobID        string `json:"job_id"`
	SubmissionID string `json:"submission_id"`
}

type RayJobDetailsResponse struct {
	JobID        string `json:"job_id"`
	SubmissionID string `json:"submission_id"`
	Status       string `json:"status"`
}

type RayJobLogsResponse struct {
	Logs string `json:"logs"`
}

type RayClusterClientConfig struct {
	SkipTlsVerification bool
}

var _ RayClusterClient = (*rayClusterClient)(nil)

type rayClusterClient struct {
	endpoint    url.URL
	httpClient  *http.Client
	bearerToken string
}

type RayClusterClient interface {
	CreateJob(job *RayJobSetup) (*RayJobResponse, error)
	GetJobDetails(jobID string) (*RayJobDetailsResponse, error)
	GetJobLogs(jobID string) (string, error)
	GetJobs() ([]map[string]interface{}, error)
	WaitForJobStatus(test Test, jobID string) string
}

func NewRayClusterClient(dashboardEndpoint url.URL, config RayClusterClientConfig, bearerToken string) RayClusterClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: config.SkipTlsVerification},
		Proxy:           http.ProxyFromEnvironment,
	}
	return &rayClusterClient{endpoint: dashboardEndpoint, httpClient: &http.Client{Transport: tr}, bearerToken: bearerToken}
}

func (client *rayClusterClient) CreateJob(job *RayJobSetup) (response *RayJobResponse, err error) {
	marshalled, err := json.Marshal(job)
	if err != nil {
		return
	}

	createJobURL := client.endpoint.String() + "/api/jobs/"

	resp, err := client.httpClient.Post(createJobURL, "application/json", bytes.NewReader(marshalled))
	if err != nil {
		return
	}

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("incorrect response code: %d for creating Ray Job, response body: %s", resp.StatusCode, respData)
	}

	err = json.Unmarshal(respData, &response)
	return
}

func (client *rayClusterClient) GetJobs() ([]map[string]interface{}, error) {
	getAllJobsDetailsURL := client.endpoint.String() + "/api/jobs/"

	req, err := http.NewRequest(http.MethodGet, getAllJobsDetailsURL, nil)
	if err != nil {
		return nil, err
	}
	if client.bearerToken != "" {
		req.Header.Set("Authorization", "Bearer "+client.bearerToken)
	}
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 503 {
		return nil, fmt.Errorf("service unavailable")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *rayClusterClient) GetJobDetails(jobID string) (response *RayJobDetailsResponse, err error) {
	getJobDetailsURL := client.endpoint.String() + "/api/jobs/" + jobID

	req, err := http.NewRequest(http.MethodGet, getJobDetailsURL, nil)
	if err != nil {
		return nil, err
	}
	if client.bearerToken != "" {
		req.Header.Set("Authorization", "Bearer "+client.bearerToken)
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("incorrect response code: %d for retrieving Ray Job details, response body: %s", resp.StatusCode, respData)
	}

	err = json.Unmarshal(respData, &response)
	return
}

func (client *rayClusterClient) GetJobLogs(jobID string) (logs string, err error) {
	getJobLogsURL := client.endpoint.String() + "/api/jobs/" + jobID + "/logs"
	req, err := http.NewRequest(http.MethodGet, getJobLogsURL, nil)
	if err != nil {
		return "", err
	}
	if client.bearerToken != "" {
		req.Header.Set("Authorization", "Bearer "+client.bearerToken)
	}
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("incorrect response code: %d for retrieving Ray Job logs, response body: %s", resp.StatusCode, respData)
	}

	jobLogs := RayJobLogsResponse{}
	err = json.Unmarshal(respData, &jobLogs)
	return jobLogs.Logs, err
}

func (client *rayClusterClient) WaitForJobStatus(test Test, jobID string) string {
	var status string
	fmt.Printf("Waiting for job to be Succeeded...\n")

	test.Eventually(func() string {
		resp, err := client.GetJobDetails(jobID)
		test.Expect(err).ToNot(HaveOccurred())
		statusVal := resp.Status
		if statusVal == "SUCCEEDED" || statusVal == "FAILED" {
			fmt.Printf("JobStatus : %s\n", statusVal)
			status = statusVal
			return status
		}
		if status != statusVal && statusVal != "SUCCEEDED" {
			fmt.Printf("JobStatus : %s...\n", statusVal)
			status = statusVal
		}
		return status
	}, TestTimeoutDouble, 3*time.Second).Should(Or(Equal("SUCCEEDED"), Equal("FAILED")), "Job did not complete within the expected time")

	if status == "SUCCEEDED" {
		fmt.Printf("Job succeeded !\n")
	} else {
		fmt.Printf("Job failed !\n")
	}
	return status
}
