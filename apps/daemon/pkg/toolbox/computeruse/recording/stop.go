// Copyright Daytona Platforms Inc.
// SPDX-License-Identifier: AGPL-3.0

package recording

import (
	"errors"
	"net/http"

	common_errors "github.com/daytonaio/common-go/pkg/errors"
	"github.com/daytonaio/daemon/pkg/toolbox/computeruse"
	"github.com/gin-gonic/gin"

	recordingservice "github.com/daytonaio/daemon/pkg/recording"
)

// StopRecording godoc
//
//	@Summary		Stop a recording
//	@Description	Stop an active screen recording session
//	@Tags			computer-use
//	@Accept			json
//	@Produce		json
//	@Param			request	body		StopRecordingRequest	true	"Recording ID to stop"
//	@Success		200		{object}	RecordingDTO
//	@Failure		400		{object}	common.ErrorResponse
//	@Failure		404		{object}	common.ErrorResponse
//	@Failure		500		{object}	common.ErrorResponse
//	@Router			/computeruse/recordings/stop [post]
//
//	@id				StopRecording
func (r *RecordingController) StopRecording(ctx *gin.Context) {
	var request StopRecordingRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		_ = ctx.Error(common_errors.NewInvalidBodyRequestError(err))
		ctx.Abort()
		return
	}

	if request.ID == "" {
		_ = ctx.Error(common_errors.NewBadRequestError(errors.New("id is required")))
		ctx.Abort()
		return
	}

	recording, err := r.recordingService.StopRecording(request.ID)
	if err != nil {
		if errors.Is(err, recordingservice.ErrRecordingNotFound) {
			_ = ctx.Error(newRecordingError(http.StatusNotFound, "recording not found", computeruse.CodeRecordingNotFound))
			ctx.Abort()
			return
		}
		_ = ctx.Error(common_errors.NewInternalServerError(err))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, *RecordingToDTO(recording))
}
