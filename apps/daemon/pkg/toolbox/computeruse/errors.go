// Copyright Daytona Platforms Inc.
// SPDX-License-Identifier: AGPL-3.0

package computeruse

import (
	"errors"
	"net/http"

	common_errors "github.com/daytonaio/common-go/pkg/errors"
)

type ComputerUseErrorCode string

const (
	// Computer use operations
	CodeComputerUseOperationFailed ComputerUseErrorCode = "COMPUTER_USE_OPERATION_FAILED"

	// Accessibility
	CodeA11yUnavailable  ComputerUseErrorCode = "A11Y_UNAVAILABLE"
	CodeA11yNodeNotFound ComputerUseErrorCode = "A11Y_NODE_NOT_FOUND"
	CodeA11yInternal     ComputerUseErrorCode = "A11Y_INTERNAL"

	// Recording
	CodeRecordingNotFound    ComputerUseErrorCode = "RECORDING_NOT_FOUND"
	CodeRecordingStillActive  ComputerUseErrorCode = "RECORDING_STILL_ACTIVE"
	CodeRecordingFfmpegNotFound ComputerUseErrorCode = "RECORDING_FFMPEG_NOT_FOUND"

	// Plugin lifecycle
	CodeComputerUseUnavailable ComputerUseErrorCode = "COMPUTER_USE_UNAVAILABLE"
)

func newComputerUseError(statusCode int, message string, code ComputerUseErrorCode) error {
	return common_errors.NewCustomError(statusCode, message, string(code))
}

func classifyA11yError(err error) error {
	msg := err.Error()
	switch {
	case hasA11ySentinel(msg, a11yMsgUnavailable):
		return newComputerUseError(http.StatusServiceUnavailable, msg, CodeA11yUnavailable)
	case hasA11ySentinel(msg, a11yMsgNodeNotFound):
		return newComputerUseError(http.StatusNotFound, msg, CodeA11yNodeNotFound)
	case hasA11ySentinel(msg, a11yMsgNoAccessibleRoot):
		return common_errors.NewNotFoundError(errors.New(msg))
	case hasA11ySentinel(msg, a11yMsgActionNotSupported):
		return common_errors.NewBadRequestError(errors.New(msg))
	case hasA11ySentinel(msg, a11yMsgInvalidScope):
		return common_errors.NewBadRequestError(errors.New(msg))
	case hasA11ySentinel(msg, a11yMsgInvalidRequest):
		return common_errors.NewBadRequestError(errors.New(msg))
	default:
		return newComputerUseError(http.StatusInternalServerError, msg, CodeA11yInternal)
	}
}
