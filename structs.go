package winlog

import (
	"sync"
	"time"
)

// Stores the common fields from a log event
type WinLogEvent struct {
	// From EvtRender
	ProviderName string
	EventId      uint64
	Qualifiers   uint64
	Level        uint64
	Task         uint64
	Opcode       uint64
	Created      time.Time
	RecordId     uint64
	ProcessId    uint64
	ThreadId     uint64
	Channel      string
	ComputerName string
	Version      uint64

	// From EvtFormatMessage
	Msg          string
	LevelText    string
	TaskText     string
	OpcodeText   string
	Keywords     []string
	ChannelText  string
	ProviderText string
	IdText       string

	// Serialied XML bookmark to
	// restart at this event
	Bookmark string
}

type channelWatcher struct {
	subscription ListenerHandle
	callback     *LogEventCallbackWrapper
	bookmark     BookmarkHandle
}

// Watches one or more event log channels
// and publishes events and errors to Go
// channels
type WinLogWatcher struct {
	errChan   chan error
	eventChan chan *WinLogEvent

	renderContext SysRenderContext
	watches       map[string]*channelWatcher
	watchMutex    sync.Mutex
	shutdown      chan interface{}
}

type SysRenderContext uint64
type ListenerHandle uint64
type PublisherHandle uint64
type EventHandle uint64
type BookmarkHandle uint64

type LogEventCallback interface {
	PublishError(error)
	PublishEvent(EventHandle)
}

type LogEventCallbackWrapper struct {
	callback LogEventCallback
}

