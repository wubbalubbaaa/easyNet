// Copyright 2020 wubbalubbaaa. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package easyNet

import (
	"errors"
)

var (
	errClosed             = errors.New("conn closed")
	errReadTimeout        = errors.New("read timeout")
	errWriteTimeout       = errors.New("write timeout")
	errInvalidFixedLength = errors.New("invalid fixed length")
	errUnexpectedEOF      = errors.New("unexpected EOF error")
	errTooLessLength      = errors.New("too less length")
	errUnsupportedLength  = errors.New("unsupported length")
)
