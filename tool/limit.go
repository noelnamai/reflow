// Copyright 2017 GRAIL, Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package tool

const (
	// The amount of outstanding bytes (send or receive)
	// acceptable from a host repository.
	transferLimit = 20 << 30

	// The number of concurrent stat operations that can
	// be performed against a repository.
	statLimit = 200

	// The number of concurrent cache operations allowed.
	cacheOpConcurrency = 512
)
