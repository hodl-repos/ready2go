package r2o

import (
	"context"
	"fmt"
	"net/http"
)

// JobStatusService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type JobStatusService service

type JobResponse struct {
	JobStatusFinishedAt string `json:"jobStatus_finishedAt"`
	JobStatusID         int64  `json:"jobStatus_id"`
	JobStatusStartedAt  string `json:"jobStatus_startedAt"`
	JobStatusStatus     string `json:"jobStatus_status"`
}

func (as *JobStatusService) GetJobStatus(ctx context.Context, id *int) (*JobResponse, error) {
	responseData := JobResponse{}

	u := fmt.Sprintf("job_status/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
