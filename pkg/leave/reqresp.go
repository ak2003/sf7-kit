package leave

import (
	"context"
	"encoding/json"
	"net/http"
	"fmt"
	"os"
	"time"
	"io"
	"path/filepath"

	"gitlab.dataon.com/gophers/sf7-kit/pkg/leave/model"

	"github.com/gorilla/schema"
)

func decodeGetLeaveRequestListingReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetLeaveRequestListingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetLeaveRequestFilterListingReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetLeaveRequestListingFilterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetDataTypeOfLeaveReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetDataTypeOfLeaveReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetDataRequestForReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetDataRequestForReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetDataRemainingLeaveReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetDataRemainingLeaveReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeCreateLeaveRequestReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.CreateLeaveRequestReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeCreateLeaveRequestFormReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.CreateLeaveRequestFormReq

	// For request body type application/x-www-form-urlencoded (URLSearchParams) / multipart/form-data (FormData)
	err := r.ParseForm()

	if err := r.ParseMultipartForm(1024); err != nil {
		return nil, err
	}

	uploadedFile, handler, err := r.FormFile("refdoc")
	if err != nil {
		return nil, err
	}
	defer uploadedFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	t := time.Now()
	filename := fmt.Sprintf("%s%s", t.Format("20060102150405"), handler.Filename)

	fileLocation := filepath.Join(dir, "upload", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		return nil, err
	}

	if err != nil {
		// Handle error
		return nil, err
	}

	err = schema.NewDecoder().Decode(&req, r.PostForm)
	if err != nil {
		// Handle error
		return nil, err
	}

	req.Refdoc = filename

	return req, nil
}
