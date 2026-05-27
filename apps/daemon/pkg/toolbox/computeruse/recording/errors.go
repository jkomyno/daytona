// Copyright Daytona Platforms Inc.
// SPDX-License-Identifier: AGPL-3.0

package recording

import (
	common_errors "github.com/daytonaio/common-go/pkg/errors"
	"github.com/daytonaio/daemon/pkg/toolbox/computeruse"
)

func newRecordingError(statusCode int, message string, code computeruse.ComputerUseErrorCode) error {
	return common_errors.NewCustomError(statusCode, message, string(code))
}
