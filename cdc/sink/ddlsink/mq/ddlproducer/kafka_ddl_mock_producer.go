// Copyright 2022 PingCAP, Inc.
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

package ddlproducer

import (
	"context"
	"fmt"

	"github.com/pingcap/tiflow/cdc/model"
	"github.com/pingcap/tiflow/pkg/sink/codec/common"
	"github.com/pingcap/tiflow/pkg/sink/kafka"
)

var _ DDLProducer = (*MockDDLProducer)(nil)

// MockDDLProducer is a mock producer for test.
type MockDDLProducer struct {
	events map[string][]*common.Message
}

// NewMockDDLProducer creates a mock producer.
func NewMockDDLProducer(_ context.Context, _ model.ChangeFeedID, _ kafka.Factory) (DDLProducer, error) {
	return &MockDDLProducer{
		events: make(map[string][]*common.Message),
	}, nil
}

// SyncBroadcastMessage stores a message to all partitions of the topic.
func (m *MockDDLProducer) SyncBroadcastMessage(ctx context.Context, topic string,
	totalPartitionsNum int32, message *common.Message,
) error {
	for i := 0; i < int(totalPartitionsNum); i++ {
		key := fmt.Sprintf("%s-%d", topic, i)
		if _, ok := m.events[key]; !ok {
			m.events[key] = make([]*common.Message, 0)
		}
		m.events[key] = append(m.events[key], message)
	}

	return nil
}

// SyncSendMessage stores a message to a partition of the topic.
func (m *MockDDLProducer) SyncSendMessage(_ context.Context, topic string,
	partitionNum int32, message *common.Message,
) error {
	key := fmt.Sprintf("%s-%d", topic, partitionNum)
	if _, ok := m.events[key]; !ok {
		m.events[key] = make([]*common.Message, 0)
	}
	m.events[key] = append(m.events[key], message)

	return nil
}

// Close do nothing.
func (m *MockDDLProducer) Close() {}

// GetAllEvents returns the events received by the mock producer.
func (m *MockDDLProducer) GetAllEvents() []*common.Message {
	var events []*common.Message
	for _, v := range m.events {
		events = append(events, v...)
	}
	return events
}

// GetEvents returns the event filtered by the key.
func (m *MockDDLProducer) GetEvents(topic string,
	partitionNum int32,
) []*common.Message {
	key := fmt.Sprintf("%s-%d", topic, partitionNum)
	return m.events[key]
}
