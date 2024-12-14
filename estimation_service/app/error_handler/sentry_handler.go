package error_handler

import (
	"github.com/getsentry/sentry-go"
)

func CaptureServerException(errMessage error) {
	sentry.CaptureException(errMessage)
}
