package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrMissingBucketName = errors.New("require bucket name missing")
)

// CreateR2Bucket Creates a new R2 bucket.
//
// API reference: https://api.cloudflare.com/#r2-bucket-create-bucket
func (api *API) CreateR2Bucket(ctx context.Context, rc *ResourceContainer, bucketName string) error {
	if rc.Identifier == "" {
		return ErrMissingAccountID
	}

	if bucketName == "" {
		return ErrMissingBucketName
	}

	uri := fmt.Sprintf("/accounts/%s/r2/buckets/%s", rc.Identifier, bucketName)
	_, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)

	return err
}

// DeleteR2Bucket Deletes an existing R2 bucket.
//
// API reference: https://api.cloudflare.com/#r2-bucket-delete-bucket
func (api *API) DeleteR2Bucket(ctx context.Context, rc *ResourceContainer, bucketName string) error {
	if rc.Identifier == "" {
		return ErrMissingAccountID
	}

	if bucketName == "" {
		return ErrMissingBucketName
	}

	uri := fmt.Sprintf("/accounts/%s/r2/buckets/%s", rc.Identifier, bucketName)
	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)

	return err
}
