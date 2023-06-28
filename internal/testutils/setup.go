// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package testutils

import (
	"io"
	"testing"

	"github.com/lasthyphen/dijets-cli/pkg/ux"
	"github.com/lasthyphen/dijetsnode/utils/logging"
	"github.com/stretchr/testify/require"
)

func SetupTest(t *testing.T) *require.Assertions {
	// use io.Discard to not print anything
	ux.NewUserLog(logging.NoLog{}, io.Discard)
	return require.New(t)
}
