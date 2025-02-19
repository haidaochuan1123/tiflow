// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"math"

	"github.com/pingcap/tiflow/pkg/errors"
)

// SorterConfig represents sorter config for a changefeed
type SorterConfig struct {
	// the directory used to store the temporary files generated by the sorter
	SortDir string `toml:"sort-dir" json:"sort-dir"`

	// Cache size of sorter in MB.
	CacheSizeInMB uint64 `toml:"cache-size-in-mb" json:"cache-size-in-mb"`

	// the maximum memory use percentage that allows in-memory sorting
	// Deprecated: use CacheSizeInMB instead.
	MaxMemoryPercentage int `toml:"max-memory-percentage" json:"max-memory-percentage"`

	// the maximum memory consumption allowed for in-memory sorting
	// Deprecated: we don't use this field anymore after introducing pull based sink.
	MaxMemoryConsumption uint64 `toml:"max-memory-consumption" json:"max-memory-consumption"`
	// the size of workerpool
	// Deprecated: we don't use this field anymore after introducing pull based sink.
	NumWorkerPoolGoroutine int `toml:"num-workerpool-goroutine" json:"num-workerpool-goroutine"`
	// number of concurrent heap sorts
	// Deprecated: we don't use this field anymore after introducing pull based sink.
	NumConcurrentWorker int `toml:"num-concurrent-worker" json:"num-concurrent-worker"`
	// maximum size for a heap
	// Deprecated: we don't use this field anymore after introducing pull based sink.
	ChunkSizeLimit uint64 `toml:"chunk-size-limit" json:"chunk-size-limit"`
}

// ValidateAndAdjust validates and adjusts the sorter configuration
func (c *SorterConfig) ValidateAndAdjust() error {
	if c.CacheSizeInMB < 8 || c.CacheSizeInMB*uint64(1<<20) > uint64(math.MaxInt64) {
		return errors.ErrIllegalSorterParameter.GenWithStackByArgs("cache-size-in-mb should be greater than 8(MB)")
	}
	return nil
}
