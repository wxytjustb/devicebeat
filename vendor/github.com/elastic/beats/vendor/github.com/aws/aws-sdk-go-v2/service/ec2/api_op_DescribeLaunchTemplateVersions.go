// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeLaunchTemplateVersionsRequest
type DescribeLaunchTemplateVersionsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters.
	//
	//    * create-time - The time the launch template version was created.
	//
	//    * ebs-optimized - A boolean that indicates whether the instance is optimized
	//    for Amazon EBS I/O.
	//
	//    * iam-instance-profile - The ARN of the IAM instance profile.
	//
	//    * image-id - The ID of the AMI.
	//
	//    * instance-type - The instance type.
	//
	//    * is-default-version - A boolean that indicates whether the launch template
	//    version is the default version.
	//
	//    * kernel-id - The kernel ID.
	//
	//    * ram-disk-id - The RAM disk ID.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The ID of the launch template. You must specify either the launch template
	// ID or launch template name in the request.
	LaunchTemplateId *string `type:"string"`

	// The name of the launch template. You must specify either the launch template
	// ID or launch template name in the request.
	LaunchTemplateName *string `min:"3" type:"string"`

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value. This
	// value can be between 1 and 200.
	MaxResults *int64 `type:"integer"`

	// The version number up to which to describe launch template versions.
	MaxVersion *string `type:"string"`

	// The version number after which to describe launch template versions.
	MinVersion *string `type:"string"`

	// The token to request the next page of results.
	NextToken *string `type:"string"`

	// One or more versions of the launch template.
	Versions []string `locationName:"LaunchTemplateVersion" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeLaunchTemplateVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeLaunchTemplateVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeLaunchTemplateVersionsInput"}
	if s.LaunchTemplateName != nil && len(*s.LaunchTemplateName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("LaunchTemplateName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeLaunchTemplateVersionsResult
type DescribeLaunchTemplateVersionsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the launch template versions.
	LaunchTemplateVersions []LaunchTemplateVersion `locationName:"launchTemplateVersionSet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeLaunchTemplateVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeLaunchTemplateVersions = "DescribeLaunchTemplateVersions"

// DescribeLaunchTemplateVersionsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes one or more versions of a specified launch template. You can describe
// all versions, individual versions, or a range of versions.
//
//    // Example sending a request using DescribeLaunchTemplateVersionsRequest.
//    req := client.DescribeLaunchTemplateVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeLaunchTemplateVersions
func (c *Client) DescribeLaunchTemplateVersionsRequest(input *DescribeLaunchTemplateVersionsInput) DescribeLaunchTemplateVersionsRequest {
	op := &aws.Operation{
		Name:       opDescribeLaunchTemplateVersions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeLaunchTemplateVersionsInput{}
	}

	req := c.newRequest(op, input, &DescribeLaunchTemplateVersionsOutput{})
	return DescribeLaunchTemplateVersionsRequest{Request: req, Input: input, Copy: c.DescribeLaunchTemplateVersionsRequest}
}

// DescribeLaunchTemplateVersionsRequest is the request type for the
// DescribeLaunchTemplateVersions API operation.
type DescribeLaunchTemplateVersionsRequest struct {
	*aws.Request
	Input *DescribeLaunchTemplateVersionsInput
	Copy  func(*DescribeLaunchTemplateVersionsInput) DescribeLaunchTemplateVersionsRequest
}

// Send marshals and sends the DescribeLaunchTemplateVersions API request.
func (r DescribeLaunchTemplateVersionsRequest) Send(ctx context.Context) (*DescribeLaunchTemplateVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLaunchTemplateVersionsResponse{
		DescribeLaunchTemplateVersionsOutput: r.Request.Data.(*DescribeLaunchTemplateVersionsOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeLaunchTemplateVersionsRequestPaginator returns a paginator for DescribeLaunchTemplateVersions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeLaunchTemplateVersionsRequest(input)
//   p := ec2.NewDescribeLaunchTemplateVersionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeLaunchTemplateVersionsPaginator(req DescribeLaunchTemplateVersionsRequest) DescribeLaunchTemplateVersionsPaginator {
	return DescribeLaunchTemplateVersionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeLaunchTemplateVersionsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeLaunchTemplateVersionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeLaunchTemplateVersionsPaginator struct {
	aws.Pager
}

func (p *DescribeLaunchTemplateVersionsPaginator) CurrentPage() *DescribeLaunchTemplateVersionsOutput {
	return p.Pager.CurrentPage().(*DescribeLaunchTemplateVersionsOutput)
}

// DescribeLaunchTemplateVersionsResponse is the response type for the
// DescribeLaunchTemplateVersions API operation.
type DescribeLaunchTemplateVersionsResponse struct {
	*DescribeLaunchTemplateVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLaunchTemplateVersions request.
func (r *DescribeLaunchTemplateVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
