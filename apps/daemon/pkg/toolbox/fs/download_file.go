// Copyright 2025 Daytona Platforms Inc.
// SPDX-License-Identifier: AGPL-3.0

package fs

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	common_errors "github.com/daytonaio/common-go/pkg/errors"
	"github.com/daytonaio/daemon/pkg/common"
	"github.com/gin-gonic/gin"
)

// DownloadFile godoc
//
//	@Summary		Download a file
//	@Description	Download a file by providing its path
//	@Tags			file-system
//	@Produce		octet-stream
//	@Param			path	query		string	true	"File path to download"
//	@Success		200		{file}		binary
//	@Failure		400		{object}	common.ErrorResponse
//	@Failure		403		{object}	common.ErrorResponse
//	@Failure		404		{object}	common.ErrorResponse
//	@Router			/files/download [get]
//
//	@id				DownloadFile
func DownloadFile(c *gin.Context) {
	requestedPath := c.Query("path")
	if requestedPath == "" {
		_ = c.Error(common_errors.NewBadRequestError(errors.New("path is required")))
		return
	}

	absPath, err := filepath.Abs(requestedPath)
	if err != nil {
		_ = c.Error(common_errors.NewCustomError(http.StatusBadRequest, fmt.Sprintf("invalid path: %s", err.Error()), string(common.CodeInvalidFilePath)))
		return
	}

	fileInfo, err := os.Stat(absPath)
	if err != nil {
		if os.IsNotExist(err) {
			_ = c.Error(common_errors.NewCustomError(http.StatusNotFound, err.Error(), string(common.CodeFileNotFound)))
			return
		}
		if os.IsPermission(err) {
			_ = c.Error(common_errors.NewCustomError(http.StatusForbidden, err.Error(), string(common.CodeFileAccessDenied)))
			return
		}
		_ = c.Error(common_errors.NewBadRequestError(err))
		return
	}

	if fileInfo.IsDir() {
		_ = c.Error(common_errors.NewBadRequestError(errors.New("path must be a file")))
		return
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Type", "application/octet-stream")
	filename := filepath.Base(absPath)
	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"; filename*=utf-8''%s`,
		toLatin1(filename), encodeRFC5987(filename)))
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Expires", "0")
	c.Header("Cache-Control", "must-revalidate")
	c.Header("Pragma", "public")

	c.File(absPath)
}
