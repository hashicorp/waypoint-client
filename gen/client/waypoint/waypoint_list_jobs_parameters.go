// Code generated by go-swagger; DO NOT EDIT.

package waypoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewWaypointListJobsParams creates a new WaypointListJobsParams object
// with the default values initialized.
func NewWaypointListJobsParams() *WaypointListJobsParams {
	var ()
	return &WaypointListJobsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointListJobsParamsWithTimeout creates a new WaypointListJobsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointListJobsParamsWithTimeout(timeout time.Duration) *WaypointListJobsParams {
	var ()
	return &WaypointListJobsParams{

		timeout: timeout,
	}
}

// NewWaypointListJobsParamsWithContext creates a new WaypointListJobsParams object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointListJobsParamsWithContext(ctx context.Context) *WaypointListJobsParams {
	var ()
	return &WaypointListJobsParams{

		Context: ctx,
	}
}

// NewWaypointListJobsParamsWithHTTPClient creates a new WaypointListJobsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointListJobsParamsWithHTTPClient(client *http.Client) *WaypointListJobsParams {
	var ()
	return &WaypointListJobsParams{
		HTTPClient: client,
	}
}

/*WaypointListJobsParams contains all the parameters to send to the API endpoint
for the waypoint list jobs operation typically these are written to a http.Request
*/
type WaypointListJobsParams struct {

	/*ApplicationApplication*/
	ApplicationApplication *string
	/*ApplicationProject*/
	ApplicationProject *string
	/*JobState*/
	JobState []string
	/*NamespaceID
	  namespace_id is the id of the namespace.

	*/
	NamespaceID string
	/*PaginationNextPageToken
	  Specifies a page token to use to retrieve the next page. Set this to the
	`next_page_token` returned by previous list requests to get the next page of
	results. If set, `previous_page_token` must not be set.

	*/
	PaginationNextPageToken *string
	/*PaginationPageSize
	  The max number of results per page that should be returned. If the number
	of available results is larger than `page_size`, a `next_page_token` is
	returned which can be used to get the next page of results in subsequent
	requests. A value of zero will cause `page_size` to be defaulted.

	*/
	PaginationPageSize *int64
	/*PaginationPreviousPageToken
	  Specifies a page token to use to retrieve the previous page. Set this to
	the `previous_page_token` returned by previous list requests to get the
	previous page of results. If set, `next_page_token` must not be set.

	*/
	PaginationPreviousPageToken *string
	/*PipelinePipelineID
	  ID of the pipeline.

	*/
	PipelinePipelineID *string
	/*PipelinePipelineName
	  Name of the pipeline.

	*/
	PipelinePipelineName *string
	/*PipelineRunSequence
	  Pipeline run sequence.

	*/
	PipelineRunSequence *string
	/*PipelineStep
	  Step name within the pipeline.

	*/
	PipelineStep *string
	/*ProjectProject*/
	ProjectProject *string
	/*TargetRunnerIDID*/
	TargetRunnerIDID *string
	/*WorkspaceWorkspace*/
	WorkspaceWorkspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint list jobs params
func (o *WaypointListJobsParams) WithTimeout(timeout time.Duration) *WaypointListJobsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint list jobs params
func (o *WaypointListJobsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint list jobs params
func (o *WaypointListJobsParams) WithContext(ctx context.Context) *WaypointListJobsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint list jobs params
func (o *WaypointListJobsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint list jobs params
func (o *WaypointListJobsParams) WithHTTPClient(client *http.Client) *WaypointListJobsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint list jobs params
func (o *WaypointListJobsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationApplication adds the applicationApplication to the waypoint list jobs params
func (o *WaypointListJobsParams) WithApplicationApplication(applicationApplication *string) *WaypointListJobsParams {
	o.SetApplicationApplication(applicationApplication)
	return o
}

// SetApplicationApplication adds the applicationApplication to the waypoint list jobs params
func (o *WaypointListJobsParams) SetApplicationApplication(applicationApplication *string) {
	o.ApplicationApplication = applicationApplication
}

// WithApplicationProject adds the applicationProject to the waypoint list jobs params
func (o *WaypointListJobsParams) WithApplicationProject(applicationProject *string) *WaypointListJobsParams {
	o.SetApplicationProject(applicationProject)
	return o
}

// SetApplicationProject adds the applicationProject to the waypoint list jobs params
func (o *WaypointListJobsParams) SetApplicationProject(applicationProject *string) {
	o.ApplicationProject = applicationProject
}

// WithJobState adds the jobState to the waypoint list jobs params
func (o *WaypointListJobsParams) WithJobState(jobState []string) *WaypointListJobsParams {
	o.SetJobState(jobState)
	return o
}

// SetJobState adds the jobState to the waypoint list jobs params
func (o *WaypointListJobsParams) SetJobState(jobState []string) {
	o.JobState = jobState
}

// WithNamespaceID adds the namespaceID to the waypoint list jobs params
func (o *WaypointListJobsParams) WithNamespaceID(namespaceID string) *WaypointListJobsParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint list jobs params
func (o *WaypointListJobsParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the waypoint list jobs params
func (o *WaypointListJobsParams) WithPaginationNextPageToken(paginationNextPageToken *string) *WaypointListJobsParams {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the waypoint list jobs params
func (o *WaypointListJobsParams) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the waypoint list jobs params
func (o *WaypointListJobsParams) WithPaginationPageSize(paginationPageSize *int64) *WaypointListJobsParams {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the waypoint list jobs params
func (o *WaypointListJobsParams) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the waypoint list jobs params
func (o *WaypointListJobsParams) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *WaypointListJobsParams {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the waypoint list jobs params
func (o *WaypointListJobsParams) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WithPipelinePipelineID adds the pipelinePipelineID to the waypoint list jobs params
func (o *WaypointListJobsParams) WithPipelinePipelineID(pipelinePipelineID *string) *WaypointListJobsParams {
	o.SetPipelinePipelineID(pipelinePipelineID)
	return o
}

// SetPipelinePipelineID adds the pipelinePipelineId to the waypoint list jobs params
func (o *WaypointListJobsParams) SetPipelinePipelineID(pipelinePipelineID *string) {
	o.PipelinePipelineID = pipelinePipelineID
}

// WithPipelinePipelineName adds the pipelinePipelineName to the waypoint list jobs params
func (o *WaypointListJobsParams) WithPipelinePipelineName(pipelinePipelineName *string) *WaypointListJobsParams {
	o.SetPipelinePipelineName(pipelinePipelineName)
	return o
}

// SetPipelinePipelineName adds the pipelinePipelineName to the waypoint list jobs params
func (o *WaypointListJobsParams) SetPipelinePipelineName(pipelinePipelineName *string) {
	o.PipelinePipelineName = pipelinePipelineName
}

// WithPipelineRunSequence adds the pipelineRunSequence to the waypoint list jobs params
func (o *WaypointListJobsParams) WithPipelineRunSequence(pipelineRunSequence *string) *WaypointListJobsParams {
	o.SetPipelineRunSequence(pipelineRunSequence)
	return o
}

// SetPipelineRunSequence adds the pipelineRunSequence to the waypoint list jobs params
func (o *WaypointListJobsParams) SetPipelineRunSequence(pipelineRunSequence *string) {
	o.PipelineRunSequence = pipelineRunSequence
}

// WithPipelineStep adds the pipelineStep to the waypoint list jobs params
func (o *WaypointListJobsParams) WithPipelineStep(pipelineStep *string) *WaypointListJobsParams {
	o.SetPipelineStep(pipelineStep)
	return o
}

// SetPipelineStep adds the pipelineStep to the waypoint list jobs params
func (o *WaypointListJobsParams) SetPipelineStep(pipelineStep *string) {
	o.PipelineStep = pipelineStep
}

// WithProjectProject adds the projectProject to the waypoint list jobs params
func (o *WaypointListJobsParams) WithProjectProject(projectProject *string) *WaypointListJobsParams {
	o.SetProjectProject(projectProject)
	return o
}

// SetProjectProject adds the projectProject to the waypoint list jobs params
func (o *WaypointListJobsParams) SetProjectProject(projectProject *string) {
	o.ProjectProject = projectProject
}

// WithTargetRunnerIDID adds the targetRunnerIDID to the waypoint list jobs params
func (o *WaypointListJobsParams) WithTargetRunnerIDID(targetRunnerIDID *string) *WaypointListJobsParams {
	o.SetTargetRunnerIDID(targetRunnerIDID)
	return o
}

// SetTargetRunnerIDID adds the targetRunnerIdId to the waypoint list jobs params
func (o *WaypointListJobsParams) SetTargetRunnerIDID(targetRunnerIDID *string) {
	o.TargetRunnerIDID = targetRunnerIDID
}

// WithWorkspaceWorkspace adds the workspaceWorkspace to the waypoint list jobs params
func (o *WaypointListJobsParams) WithWorkspaceWorkspace(workspaceWorkspace string) *WaypointListJobsParams {
	o.SetWorkspaceWorkspace(workspaceWorkspace)
	return o
}

// SetWorkspaceWorkspace adds the workspaceWorkspace to the waypoint list jobs params
func (o *WaypointListJobsParams) SetWorkspaceWorkspace(workspaceWorkspace string) {
	o.WorkspaceWorkspace = workspaceWorkspace
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointListJobsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApplicationApplication != nil {

		// query param application.application
		var qrApplicationApplication string
		if o.ApplicationApplication != nil {
			qrApplicationApplication = *o.ApplicationApplication
		}
		qApplicationApplication := qrApplicationApplication
		if qApplicationApplication != "" {
			if err := r.SetQueryParam("application.application", qApplicationApplication); err != nil {
				return err
			}
		}

	}

	if o.ApplicationProject != nil {

		// query param application.project
		var qrApplicationProject string
		if o.ApplicationProject != nil {
			qrApplicationProject = *o.ApplicationProject
		}
		qApplicationProject := qrApplicationProject
		if qApplicationProject != "" {
			if err := r.SetQueryParam("application.project", qApplicationProject); err != nil {
				return err
			}
		}

	}

	valuesJobState := o.JobState

	joinedJobState := swag.JoinByFormat(valuesJobState, "multi")
	// query array param jobState
	if err := r.SetQueryParam("jobState", joinedJobState...); err != nil {
		return err
	}

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
	}

	if o.PaginationNextPageToken != nil {

		// query param pagination.next_page_token
		var qrPaginationNextPageToken string
		if o.PaginationNextPageToken != nil {
			qrPaginationNextPageToken = *o.PaginationNextPageToken
		}
		qPaginationNextPageToken := qrPaginationNextPageToken
		if qPaginationNextPageToken != "" {
			if err := r.SetQueryParam("pagination.next_page_token", qPaginationNextPageToken); err != nil {
				return err
			}
		}

	}

	if o.PaginationPageSize != nil {

		// query param pagination.page_size
		var qrPaginationPageSize int64
		if o.PaginationPageSize != nil {
			qrPaginationPageSize = *o.PaginationPageSize
		}
		qPaginationPageSize := swag.FormatInt64(qrPaginationPageSize)
		if qPaginationPageSize != "" {
			if err := r.SetQueryParam("pagination.page_size", qPaginationPageSize); err != nil {
				return err
			}
		}

	}

	if o.PaginationPreviousPageToken != nil {

		// query param pagination.previous_page_token
		var qrPaginationPreviousPageToken string
		if o.PaginationPreviousPageToken != nil {
			qrPaginationPreviousPageToken = *o.PaginationPreviousPageToken
		}
		qPaginationPreviousPageToken := qrPaginationPreviousPageToken
		if qPaginationPreviousPageToken != "" {
			if err := r.SetQueryParam("pagination.previous_page_token", qPaginationPreviousPageToken); err != nil {
				return err
			}
		}

	}

	if o.PipelinePipelineID != nil {

		// query param pipeline.pipeline_id
		var qrPipelinePipelineID string
		if o.PipelinePipelineID != nil {
			qrPipelinePipelineID = *o.PipelinePipelineID
		}
		qPipelinePipelineID := qrPipelinePipelineID
		if qPipelinePipelineID != "" {
			if err := r.SetQueryParam("pipeline.pipeline_id", qPipelinePipelineID); err != nil {
				return err
			}
		}

	}

	if o.PipelinePipelineName != nil {

		// query param pipeline.pipeline_name
		var qrPipelinePipelineName string
		if o.PipelinePipelineName != nil {
			qrPipelinePipelineName = *o.PipelinePipelineName
		}
		qPipelinePipelineName := qrPipelinePipelineName
		if qPipelinePipelineName != "" {
			if err := r.SetQueryParam("pipeline.pipeline_name", qPipelinePipelineName); err != nil {
				return err
			}
		}

	}

	if o.PipelineRunSequence != nil {

		// query param pipeline.run_sequence
		var qrPipelineRunSequence string
		if o.PipelineRunSequence != nil {
			qrPipelineRunSequence = *o.PipelineRunSequence
		}
		qPipelineRunSequence := qrPipelineRunSequence
		if qPipelineRunSequence != "" {
			if err := r.SetQueryParam("pipeline.run_sequence", qPipelineRunSequence); err != nil {
				return err
			}
		}

	}

	if o.PipelineStep != nil {

		// query param pipeline.step
		var qrPipelineStep string
		if o.PipelineStep != nil {
			qrPipelineStep = *o.PipelineStep
		}
		qPipelineStep := qrPipelineStep
		if qPipelineStep != "" {
			if err := r.SetQueryParam("pipeline.step", qPipelineStep); err != nil {
				return err
			}
		}

	}

	if o.ProjectProject != nil {

		// query param project.project
		var qrProjectProject string
		if o.ProjectProject != nil {
			qrProjectProject = *o.ProjectProject
		}
		qProjectProject := qrProjectProject
		if qProjectProject != "" {
			if err := r.SetQueryParam("project.project", qProjectProject); err != nil {
				return err
			}
		}

	}

	if o.TargetRunnerIDID != nil {

		// query param targetRunner.id.id
		var qrTargetRunnerIDID string
		if o.TargetRunnerIDID != nil {
			qrTargetRunnerIDID = *o.TargetRunnerIDID
		}
		qTargetRunnerIDID := qrTargetRunnerIDID
		if qTargetRunnerIDID != "" {
			if err := r.SetQueryParam("targetRunner.id.id", qTargetRunnerIDID); err != nil {
				return err
			}
		}

	}

	// path param workspace.workspace
	if err := r.SetPathParam("workspace.workspace", o.WorkspaceWorkspace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
