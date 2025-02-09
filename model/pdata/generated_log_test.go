// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "model/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "go run model/internal/cmd/pdatagen/main.go".

package pdata

import (
	"testing"

	"github.com/stretchr/testify/assert"

	otlplogs "go.opentelemetry.io/collector/model/internal/data/protogen/logs/v1"
)

func TestResourceLogsSlice(t *testing.T) {
	es := NewResourceLogsSlice()
	assert.EqualValues(t, 0, es.Len())
	es = newResourceLogsSlice(&[]*otlplogs.ResourceLogs{})
	assert.EqualValues(t, 0, es.Len())

	es.EnsureCapacity(7)
	emptyVal := newResourceLogs(&otlplogs.ResourceLogs{})
	testVal := generateTestResourceLogs()
	assert.EqualValues(t, 7, cap(*es.orig))
	for i := 0; i < es.Len(); i++ {
		el := es.AppendEmpty()
		assert.EqualValues(t, emptyVal, el)
		fillTestResourceLogs(el)
		assert.EqualValues(t, testVal, el)
	}
}

func TestResourceLogsSlice_CopyTo(t *testing.T) {
	dest := NewResourceLogsSlice()
	// Test CopyTo to empty
	NewResourceLogsSlice().CopyTo(dest)
	assert.EqualValues(t, NewResourceLogsSlice(), dest)

	// Test CopyTo larger slice
	generateTestResourceLogsSlice().CopyTo(dest)
	assert.EqualValues(t, generateTestResourceLogsSlice(), dest)

	// Test CopyTo same size slice
	generateTestResourceLogsSlice().CopyTo(dest)
	assert.EqualValues(t, generateTestResourceLogsSlice(), dest)
}

func TestResourceLogsSlice_EnsureCapacity(t *testing.T) {
	es := generateTestResourceLogsSlice()
	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	expectedEs := make(map[*otlplogs.ResourceLogs]bool)
	for i := 0; i < es.Len(); i++ {
		expectedEs[es.At(i).orig] = true
	}
	assert.Equal(t, es.Len(), len(expectedEs))
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	foundEs := make(map[*otlplogs.ResourceLogs]bool, es.Len())
	for i := 0; i < es.Len(); i++ {
		foundEs[es.At(i).orig] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	oldLen := es.Len()
	expectedEs = make(map[*otlplogs.ResourceLogs]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[es.At(i).orig] = true
	}
	assert.Equal(t, oldLen, len(expectedEs))
	es.EnsureCapacity(ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.orig))
	foundEs = make(map[*otlplogs.ResourceLogs]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[es.At(i).orig] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)
}

func TestResourceLogsSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestResourceLogsSlice()
	dest := NewResourceLogsSlice()
	src := generateTestResourceLogsSlice()
	src.MoveAndAppendTo(dest)
	assert.EqualValues(t, generateTestResourceLogsSlice(), dest)
	assert.EqualValues(t, 0, src.Len())
	assert.EqualValues(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.EqualValues(t, generateTestResourceLogsSlice(), dest)
	assert.EqualValues(t, 0, src.Len())
	assert.EqualValues(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestResourceLogsSlice().MoveAndAppendTo(dest)
	assert.EqualValues(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.EqualValues(t, expectedSlice.At(i), dest.At(i))
		assert.EqualValues(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestResourceLogsSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewResourceLogsSlice()
	emptySlice.RemoveIf(func(el ResourceLogs) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := generateTestResourceLogsSlice()
	pos := 0
	filtered.RemoveIf(func(el ResourceLogs) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestResourceLogs_MoveTo(t *testing.T) {
	ms := generateTestResourceLogs()
	dest := NewResourceLogs()
	ms.MoveTo(dest)
	assert.EqualValues(t, NewResourceLogs(), ms)
	assert.EqualValues(t, generateTestResourceLogs(), dest)
}

func TestResourceLogs_CopyTo(t *testing.T) {
	ms := NewResourceLogs()
	generateTestResourceLogs().CopyTo(ms)
	assert.EqualValues(t, generateTestResourceLogs(), ms)
}

func TestResourceLogs_Resource(t *testing.T) {
	ms := NewResourceLogs()
	fillTestResource(ms.Resource())
	assert.EqualValues(t, generateTestResource(), ms.Resource())
}

func TestResourceLogs_SchemaUrl(t *testing.T) {
	ms := NewResourceLogs()
	assert.EqualValues(t, "", ms.SchemaUrl())
	testValSchemaUrl := "https://opentelemetry.io/schemas/1.5.0"
	ms.SetSchemaUrl(testValSchemaUrl)
	assert.EqualValues(t, testValSchemaUrl, ms.SchemaUrl())
}

func TestResourceLogs_InstrumentationLibraryLogs(t *testing.T) {
	ms := NewResourceLogs()
	assert.EqualValues(t, NewInstrumentationLibraryLogsSlice(), ms.InstrumentationLibraryLogs())
	fillTestInstrumentationLibraryLogsSlice(ms.InstrumentationLibraryLogs())
	testValInstrumentationLibraryLogs := generateTestInstrumentationLibraryLogsSlice()
	assert.EqualValues(t, testValInstrumentationLibraryLogs, ms.InstrumentationLibraryLogs())
}

func TestInstrumentationLibraryLogsSlice(t *testing.T) {
	es := NewInstrumentationLibraryLogsSlice()
	assert.EqualValues(t, 0, es.Len())
	es = newInstrumentationLibraryLogsSlice(&[]*otlplogs.InstrumentationLibraryLogs{})
	assert.EqualValues(t, 0, es.Len())

	es.EnsureCapacity(7)
	emptyVal := newInstrumentationLibraryLogs(&otlplogs.InstrumentationLibraryLogs{})
	testVal := generateTestInstrumentationLibraryLogs()
	assert.EqualValues(t, 7, cap(*es.orig))
	for i := 0; i < es.Len(); i++ {
		el := es.AppendEmpty()
		assert.EqualValues(t, emptyVal, el)
		fillTestInstrumentationLibraryLogs(el)
		assert.EqualValues(t, testVal, el)
	}
}

func TestInstrumentationLibraryLogsSlice_CopyTo(t *testing.T) {
	dest := NewInstrumentationLibraryLogsSlice()
	// Test CopyTo to empty
	NewInstrumentationLibraryLogsSlice().CopyTo(dest)
	assert.EqualValues(t, NewInstrumentationLibraryLogsSlice(), dest)

	// Test CopyTo larger slice
	generateTestInstrumentationLibraryLogsSlice().CopyTo(dest)
	assert.EqualValues(t, generateTestInstrumentationLibraryLogsSlice(), dest)

	// Test CopyTo same size slice
	generateTestInstrumentationLibraryLogsSlice().CopyTo(dest)
	assert.EqualValues(t, generateTestInstrumentationLibraryLogsSlice(), dest)
}

func TestInstrumentationLibraryLogsSlice_EnsureCapacity(t *testing.T) {
	es := generateTestInstrumentationLibraryLogsSlice()
	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	expectedEs := make(map[*otlplogs.InstrumentationLibraryLogs]bool)
	for i := 0; i < es.Len(); i++ {
		expectedEs[es.At(i).orig] = true
	}
	assert.Equal(t, es.Len(), len(expectedEs))
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	foundEs := make(map[*otlplogs.InstrumentationLibraryLogs]bool, es.Len())
	for i := 0; i < es.Len(); i++ {
		foundEs[es.At(i).orig] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	oldLen := es.Len()
	expectedEs = make(map[*otlplogs.InstrumentationLibraryLogs]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[es.At(i).orig] = true
	}
	assert.Equal(t, oldLen, len(expectedEs))
	es.EnsureCapacity(ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.orig))
	foundEs = make(map[*otlplogs.InstrumentationLibraryLogs]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[es.At(i).orig] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)
}

func TestInstrumentationLibraryLogsSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestInstrumentationLibraryLogsSlice()
	dest := NewInstrumentationLibraryLogsSlice()
	src := generateTestInstrumentationLibraryLogsSlice()
	src.MoveAndAppendTo(dest)
	assert.EqualValues(t, generateTestInstrumentationLibraryLogsSlice(), dest)
	assert.EqualValues(t, 0, src.Len())
	assert.EqualValues(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.EqualValues(t, generateTestInstrumentationLibraryLogsSlice(), dest)
	assert.EqualValues(t, 0, src.Len())
	assert.EqualValues(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestInstrumentationLibraryLogsSlice().MoveAndAppendTo(dest)
	assert.EqualValues(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.EqualValues(t, expectedSlice.At(i), dest.At(i))
		assert.EqualValues(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestInstrumentationLibraryLogsSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewInstrumentationLibraryLogsSlice()
	emptySlice.RemoveIf(func(el InstrumentationLibraryLogs) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := generateTestInstrumentationLibraryLogsSlice()
	pos := 0
	filtered.RemoveIf(func(el InstrumentationLibraryLogs) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestInstrumentationLibraryLogs_MoveTo(t *testing.T) {
	ms := generateTestInstrumentationLibraryLogs()
	dest := NewInstrumentationLibraryLogs()
	ms.MoveTo(dest)
	assert.EqualValues(t, NewInstrumentationLibraryLogs(), ms)
	assert.EqualValues(t, generateTestInstrumentationLibraryLogs(), dest)
}

func TestInstrumentationLibraryLogs_CopyTo(t *testing.T) {
	ms := NewInstrumentationLibraryLogs()
	generateTestInstrumentationLibraryLogs().CopyTo(ms)
	assert.EqualValues(t, generateTestInstrumentationLibraryLogs(), ms)
}

func TestInstrumentationLibraryLogs_InstrumentationLibrary(t *testing.T) {
	ms := NewInstrumentationLibraryLogs()
	fillTestInstrumentationLibrary(ms.InstrumentationLibrary())
	assert.EqualValues(t, generateTestInstrumentationLibrary(), ms.InstrumentationLibrary())
}

func TestInstrumentationLibraryLogs_SchemaUrl(t *testing.T) {
	ms := NewInstrumentationLibraryLogs()
	assert.EqualValues(t, "", ms.SchemaUrl())
	testValSchemaUrl := "https://opentelemetry.io/schemas/1.5.0"
	ms.SetSchemaUrl(testValSchemaUrl)
	assert.EqualValues(t, testValSchemaUrl, ms.SchemaUrl())
}

func TestInstrumentationLibraryLogs_LogRecords(t *testing.T) {
	ms := NewInstrumentationLibraryLogs()
	assert.EqualValues(t, NewLogRecordSlice(), ms.LogRecords())
	fillTestLogRecordSlice(ms.LogRecords())
	testValLogRecords := generateTestLogRecordSlice()
	assert.EqualValues(t, testValLogRecords, ms.LogRecords())
}

func TestLogRecordSlice(t *testing.T) {
	es := NewLogRecordSlice()
	assert.EqualValues(t, 0, es.Len())
	es = newLogRecordSlice(&[]*otlplogs.LogRecord{})
	assert.EqualValues(t, 0, es.Len())

	es.EnsureCapacity(7)
	emptyVal := newLogRecord(&otlplogs.LogRecord{})
	testVal := generateTestLogRecord()
	assert.EqualValues(t, 7, cap(*es.orig))
	for i := 0; i < es.Len(); i++ {
		el := es.AppendEmpty()
		assert.EqualValues(t, emptyVal, el)
		fillTestLogRecord(el)
		assert.EqualValues(t, testVal, el)
	}
}

func TestLogRecordSlice_CopyTo(t *testing.T) {
	dest := NewLogRecordSlice()
	// Test CopyTo to empty
	NewLogRecordSlice().CopyTo(dest)
	assert.EqualValues(t, NewLogRecordSlice(), dest)

	// Test CopyTo larger slice
	generateTestLogRecordSlice().CopyTo(dest)
	assert.EqualValues(t, generateTestLogRecordSlice(), dest)

	// Test CopyTo same size slice
	generateTestLogRecordSlice().CopyTo(dest)
	assert.EqualValues(t, generateTestLogRecordSlice(), dest)
}

func TestLogRecordSlice_EnsureCapacity(t *testing.T) {
	es := generateTestLogRecordSlice()
	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	expectedEs := make(map[*otlplogs.LogRecord]bool)
	for i := 0; i < es.Len(); i++ {
		expectedEs[es.At(i).orig] = true
	}
	assert.Equal(t, es.Len(), len(expectedEs))
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	foundEs := make(map[*otlplogs.LogRecord]bool, es.Len())
	for i := 0; i < es.Len(); i++ {
		foundEs[es.At(i).orig] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	oldLen := es.Len()
	expectedEs = make(map[*otlplogs.LogRecord]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[es.At(i).orig] = true
	}
	assert.Equal(t, oldLen, len(expectedEs))
	es.EnsureCapacity(ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.orig))
	foundEs = make(map[*otlplogs.LogRecord]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[es.At(i).orig] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)
}

func TestLogRecordSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestLogRecordSlice()
	dest := NewLogRecordSlice()
	src := generateTestLogRecordSlice()
	src.MoveAndAppendTo(dest)
	assert.EqualValues(t, generateTestLogRecordSlice(), dest)
	assert.EqualValues(t, 0, src.Len())
	assert.EqualValues(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.EqualValues(t, generateTestLogRecordSlice(), dest)
	assert.EqualValues(t, 0, src.Len())
	assert.EqualValues(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestLogRecordSlice().MoveAndAppendTo(dest)
	assert.EqualValues(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.EqualValues(t, expectedSlice.At(i), dest.At(i))
		assert.EqualValues(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestLogRecordSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewLogRecordSlice()
	emptySlice.RemoveIf(func(el LogRecord) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := generateTestLogRecordSlice()
	pos := 0
	filtered.RemoveIf(func(el LogRecord) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestLogRecord_MoveTo(t *testing.T) {
	ms := generateTestLogRecord()
	dest := NewLogRecord()
	ms.MoveTo(dest)
	assert.EqualValues(t, NewLogRecord(), ms)
	assert.EqualValues(t, generateTestLogRecord(), dest)
}

func TestLogRecord_CopyTo(t *testing.T) {
	ms := NewLogRecord()
	generateTestLogRecord().CopyTo(ms)
	assert.EqualValues(t, generateTestLogRecord(), ms)
}

func TestLogRecord_Timestamp(t *testing.T) {
	ms := NewLogRecord()
	assert.EqualValues(t, Timestamp(0), ms.Timestamp())
	testValTimestamp := Timestamp(1234567890)
	ms.SetTimestamp(testValTimestamp)
	assert.EqualValues(t, testValTimestamp, ms.Timestamp())
}

func TestLogRecord_TraceID(t *testing.T) {
	ms := NewLogRecord()
	assert.EqualValues(t, NewTraceID([16]byte{}), ms.TraceID())
	testValTraceID := NewTraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1})
	ms.SetTraceID(testValTraceID)
	assert.EqualValues(t, testValTraceID, ms.TraceID())
}

func TestLogRecord_SpanID(t *testing.T) {
	ms := NewLogRecord()
	assert.EqualValues(t, NewSpanID([8]byte{}), ms.SpanID())
	testValSpanID := NewSpanID([8]byte{1, 2, 3, 4, 5, 6, 7, 8})
	ms.SetSpanID(testValSpanID)
	assert.EqualValues(t, testValSpanID, ms.SpanID())
}

func TestLogRecord_Flags(t *testing.T) {
	ms := NewLogRecord()
	assert.EqualValues(t, uint32(0), ms.Flags())
	testValFlags := uint32(0x01)
	ms.SetFlags(testValFlags)
	assert.EqualValues(t, testValFlags, ms.Flags())
}

func TestLogRecord_SeverityText(t *testing.T) {
	ms := NewLogRecord()
	assert.EqualValues(t, "", ms.SeverityText())
	testValSeverityText := "INFO"
	ms.SetSeverityText(testValSeverityText)
	assert.EqualValues(t, testValSeverityText, ms.SeverityText())
}

func TestLogRecord_SeverityNumber(t *testing.T) {
	ms := NewLogRecord()
	assert.EqualValues(t, SeverityNumberUNDEFINED, ms.SeverityNumber())
	testValSeverityNumber := SeverityNumberINFO
	ms.SetSeverityNumber(testValSeverityNumber)
	assert.EqualValues(t, testValSeverityNumber, ms.SeverityNumber())
}

func TestLogRecord_Name(t *testing.T) {
	ms := NewLogRecord()
	assert.EqualValues(t, "", ms.Name())
	testValName := "test_name"
	ms.SetName(testValName)
	assert.EqualValues(t, testValName, ms.Name())
}

func TestLogRecord_Body(t *testing.T) {
	ms := NewLogRecord()
	fillTestAttributeValue(ms.Body())
	assert.EqualValues(t, generateTestAttributeValue(), ms.Body())
}

func TestLogRecord_Attributes(t *testing.T) {
	ms := NewLogRecord()
	assert.EqualValues(t, NewAttributeMap(), ms.Attributes())
	fillTestAttributeMap(ms.Attributes())
	testValAttributes := generateTestAttributeMap()
	assert.EqualValues(t, testValAttributes, ms.Attributes())
}

func TestLogRecord_DroppedAttributesCount(t *testing.T) {
	ms := NewLogRecord()
	assert.EqualValues(t, uint32(0), ms.DroppedAttributesCount())
	testValDroppedAttributesCount := uint32(17)
	ms.SetDroppedAttributesCount(testValDroppedAttributesCount)
	assert.EqualValues(t, testValDroppedAttributesCount, ms.DroppedAttributesCount())
}

func generateTestResourceLogsSlice() ResourceLogsSlice {
	tv := NewResourceLogsSlice()
	fillTestResourceLogsSlice(tv)
	return tv
}

func fillTestResourceLogsSlice(tv ResourceLogsSlice) {
	l := 7
	tv.EnsureCapacity(l)
	for i := 0; i < l; i++ {
		fillTestResourceLogs(tv.AppendEmpty())
	}
}

func generateTestResourceLogs() ResourceLogs {
	tv := NewResourceLogs()
	fillTestResourceLogs(tv)
	return tv
}

func fillTestResourceLogs(tv ResourceLogs) {
	fillTestResource(tv.Resource())
	tv.SetSchemaUrl("https://opentelemetry.io/schemas/1.5.0")
	fillTestInstrumentationLibraryLogsSlice(tv.InstrumentationLibraryLogs())
}

func generateTestInstrumentationLibraryLogsSlice() InstrumentationLibraryLogsSlice {
	tv := NewInstrumentationLibraryLogsSlice()
	fillTestInstrumentationLibraryLogsSlice(tv)
	return tv
}

func fillTestInstrumentationLibraryLogsSlice(tv InstrumentationLibraryLogsSlice) {
	l := 7
	tv.EnsureCapacity(l)
	for i := 0; i < l; i++ {
		fillTestInstrumentationLibraryLogs(tv.AppendEmpty())
	}
}

func generateTestInstrumentationLibraryLogs() InstrumentationLibraryLogs {
	tv := NewInstrumentationLibraryLogs()
	fillTestInstrumentationLibraryLogs(tv)
	return tv
}

func fillTestInstrumentationLibraryLogs(tv InstrumentationLibraryLogs) {
	fillTestInstrumentationLibrary(tv.InstrumentationLibrary())
	tv.SetSchemaUrl("https://opentelemetry.io/schemas/1.5.0")
	fillTestLogRecordSlice(tv.LogRecords())
}

func generateTestLogRecordSlice() LogRecordSlice {
	tv := NewLogRecordSlice()
	fillTestLogRecordSlice(tv)
	return tv
}

func fillTestLogRecordSlice(tv LogRecordSlice) {
	l := 7
	tv.EnsureCapacity(l)
	for i := 0; i < l; i++ {
		fillTestLogRecord(tv.AppendEmpty())
	}
}

func generateTestLogRecord() LogRecord {
	tv := NewLogRecord()
	fillTestLogRecord(tv)
	return tv
}

func fillTestLogRecord(tv LogRecord) {
	tv.SetTimestamp(Timestamp(1234567890))
	tv.SetTraceID(NewTraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1}))
	tv.SetSpanID(NewSpanID([8]byte{1, 2, 3, 4, 5, 6, 7, 8}))
	tv.SetFlags(uint32(0x01))
	tv.SetSeverityText("INFO")
	tv.SetSeverityNumber(SeverityNumberINFO)
	tv.SetName("test_name")
	fillTestAttributeValue(tv.Body())
	fillTestAttributeMap(tv.Attributes())
	tv.SetDroppedAttributesCount(uint32(17))
}
