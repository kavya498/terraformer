// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatch

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type TagResourceInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the CloudWatch resource that you're adding tags to.
	//
	// The ARN format of an alarm is arn:aws:cloudwatch:Region:account-id:alarm:alarm-name
	//
	// The ARN format of a Contributor Insights rule is arn:aws:cloudwatch:Region:account-id:insight-rule:insight-rule-name
	//
	// For more information on ARN format, see Resource Types Defined by Amazon
	// CloudWatch (https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazoncloudwatch.html#amazoncloudwatch-resources-for-iam-policies)
	// in the Amazon Web Services General Reference.
	//
	// ResourceARN is a required field
	ResourceARN *string `min:"1" type:"string" required:"true"`

	// The list of key-value pairs to associate with the alarm.
	//
	// Tags is a required field
	Tags []Tag `type:"list" required:"true"`
}

// String returns the string representation
func (s TagResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TagResourceInput"}

	if s.ResourceARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceARN"))
	}
	if s.ResourceARN != nil && len(*s.ResourceARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceARN", 1))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type TagResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s TagResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opTagResource = "TagResource"

// TagResourceRequest returns a request value for making API operation for
// Amazon CloudWatch.
//
// Assigns one or more tags (key-value pairs) to the specified CloudWatch resource.
// Currently, the only CloudWatch resources that can be tagged are alarms and
// Contributor Insights rules.
//
// Tags can help you organize and categorize your resources. You can also use
// them to scope user permissions, by granting a user permission to access or
// change only resources with certain tag values.
//
// Tags don't have any semantic meaning to AWS and are interpreted strictly
// as strings of characters.
//
// You can use the TagResource action with an alarm that already has tags. If
// you specify a new tag key for the alarm, this tag is appended to the list
// of tags associated with the alarm. If you specify a tag key that is already
// associated with the alarm, the new tag value that you specify replaces the
// previous value for that tag.
//
// You can associate as many as 50 tags with a CloudWatch resource.
//
//    // Example sending a request using TagResourceRequest.
//    req := client.TagResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/TagResource
func (c *Client) TagResourceRequest(input *TagResourceInput) TagResourceRequest {
	op := &aws.Operation{
		Name:       opTagResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TagResourceInput{}
	}

	req := c.newRequest(op, input, &TagResourceOutput{})
	return TagResourceRequest{Request: req, Input: input, Copy: c.TagResourceRequest}
}

// TagResourceRequest is the request type for the
// TagResource API operation.
type TagResourceRequest struct {
	*aws.Request
	Input *TagResourceInput
	Copy  func(*TagResourceInput) TagResourceRequest
}

// Send marshals and sends the TagResource API request.
func (r TagResourceRequest) Send(ctx context.Context) (*TagResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TagResourceResponse{
		TagResourceOutput: r.Request.Data.(*TagResourceOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TagResourceResponse is the response type for the
// TagResource API operation.
type TagResourceResponse struct {
	*TagResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TagResource request.
func (r *TagResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}