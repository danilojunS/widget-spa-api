package utils_test

import (
	utils "github.com/danilojunS/widgets-spa-api/utils"
	"testing"
)

func TestCheckError(t *testing.T) {
	// do not panics if error is nil
	utils.CheckError(nil)
}
