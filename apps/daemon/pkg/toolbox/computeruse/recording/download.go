// Copyright Daytona Platforms Inc.
// SPDX-License-Identifier: AGPL-3.0

package recording

import (
	"errors"
	"net/http"
	"os"

	common_errors "github.com/daytonaio/common-go/pkg/errors"
	"github.com/daytonaio/daemon/pkg/toolbox/computeruse"
	"github.com/daytonaio/daemon/pkg/recording"
	"github.com/gin-gonic/gin"
)

// DownloadRecording godoc
//
//	@Summary		Download a recording
//	@Description	Download a recording by providing its ID
//	@Tags			computer-use
//	@Produce		octet-stream
//	@Param			id	path		string	true	"Recording ID"
//	@Success		200	{file}		binary
//	@Failure		400	{object}	common.ErrorResponse
//	@Failure		404	{object}	common.ErrorResponse
//	@Failure		500	{object}	common.ErrorResponse
//	@Router			/computeruse/recordings/{id}/download [get]
//
//	@id				DownloadRecording
func (r *RecordingController) DownloadRecording(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		_ = ctx.Error(common_errors.NewBadRequestError(errors.New("id is required")))
		ctx.Abort()
		return
	}

	rec, err := r.recordingService.GetRecording(id)
	if err != nil {
		if errors.Is(err, recording.ErrRecordingNotFound) {
			_ = ctx.Error(newRecordingError(http.StatusNotFound, "recording not found", computeruse.CodeRecordingNotFound))
			ctx.Abort()
			return
		}
		_ = ctx.Error(common_errors.NewInternalServerError(err))
		ctx.Abort()
		return
	}

	if _, err := os.Stat(rec.FilePath); os.IsNotExist(err) {
		_ = ctx.Error(common_errors.NewNotFoundError(errors.New("recording file not found")))
		ctx.Abort()
		return
	} else if err != nil {
		_ = ctx.Error(common_errors.NewInternalServerError(err))
		ctx.Abort()
		return
	}

	ctx.File(rec.FilePath)
}
