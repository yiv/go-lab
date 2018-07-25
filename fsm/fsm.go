package main

import (
	"github.com/looplab/fsm"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/log"
	"os"
	"time"
	"sync"
)

const (
	StateIdle      = "idle"
	StateStop      = "stop"
	StateReady     = "ready"
	StatePreparing = "preparing"
	StatePlaying   = "playing"
)

const (
	EnterStateIdle = "enter_idle"
	LeaveStateIdle = "leave_idle"

	EnterStateStop = "enter_stop"
	LeaveStateStop = "leave_stop"

	EnterStateReady = "enter_ready"
	LeaveStateReady = "leave_ready"

	EnterStatePreparing = "enter_preparing"
	LeaveStatePreparing = "leave_preparing"

	EnterStatePlaying = "enter_playing"
	LeaveStatePlaying = "leave_playing"
)

const (
	EventInit   = "init"
	EventTry    = "try"
	EventStart  = "start"
	EventSettle = "settle"
	EventOver   = "over"
	EventReset  = "reset"
)

type Table struct {
	logger    log.Logger
	EventChan chan string
	FSM       *fsm.FSM
}

func main() {
	logger := log.NewLogfmtLogger(os.Stdout)
	logger = level.NewFilter(logger, level.AllowDebug())
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	t := Table{
		EventChan: make(chan string),
		logger:    logger,
	}
	f := fsm.NewFSM(
		StateIdle,
		fsm.Events{
			{Name: EventInit, Src: []string{StateIdle}, Dst: StateReady},
			{Name: EventTry, Src: []string{StateReady}, Dst: StatePreparing},
			{Name: EventStart, Src: []string{StatePreparing}, Dst: StatePlaying},
			{Name: EventSettle, Src: []string{StatePlaying}, Dst: StateStop},
			{Name: EventOver, Src: []string{StateStop}, Dst: StateReady},
			{Name: EventReset, Src: []string{StateReady}, Dst: StateIdle},
			{Name: EventReset, Src: []string{StatePreparing}, Dst: StateIdle},
			{Name: EventReset, Src: []string{StatePlaying}, Dst: StateIdle},
			{Name: EventReset, Src: []string{StateStop}, Dst: StateIdle},
		},
		fsm.Callbacks{
			EnterStateIdle: func(e *fsm.Event) { t.enterStateIdle(e) },
			LeaveStateIdle: func(e *fsm.Event) { t.leaveStateIdle(e) },

			EnterStateReady: func(e *fsm.Event) { t.enterStateReady(e) },
			LeaveStateReady: func(e *fsm.Event) { t.leaveStateReady(e) },

			EnterStatePreparing: func(e *fsm.Event) { t.enterStatePreparing(e) },
			LeaveStatePreparing: func(e *fsm.Event) { t.leaveStatePreparing(e) },

			EnterStatePlaying: func(e *fsm.Event) { t.enterStatePlaying(e) },
			LeaveStatePlaying: func(e *fsm.Event) { t.leaveStatePlaying(e) },

			EnterStateStop: func(e *fsm.Event) { t.enterStateStop(e) },
			LeaveStateStop: func(e *fsm.Event) { t.leaveStateStop(e) },
		},
	)
	wg := sync.WaitGroup{}
	wg.Add(1)
	t.FSM = f
	go t.EventTrigger(EventInit)
	go t.eventListener()
	wg.Wait()
}

func (t *Table) EventTrigger(event string) {
	level.Debug(t.logger).Log("Table", "EventTrigger", "event", event)
	go func() {
		t.EventChan <- event
	}()
}

func (t *Table) enterStateIdle(e *fsm.Event) {
	level.Debug(t.logger).Log("STATE", "enterStateIdle", "state", t.FSM.Current())
	t.EventTrigger(EventInit)
}

func (t *Table) leaveStateIdle(e *fsm.Event) {
	level.Debug(t.logger).Log("STATE", "leaveStateIdle", "state", t.FSM.Current())
}

func (t *Table) enterStateReady(e *fsm.Event) {
	level.Debug(t.logger).Log("STATE", "enterStateReady", "state", t.FSM.Current())
	t.EventTrigger(EventTry)
}

func (t *Table) leaveStateReady(e *fsm.Event) {
	level.Debug(t.logger).Log("STATE", "leaveStateReady", "state", t.FSM.Current())
}

func (t *Table) enterStatePreparing(e *fsm.Event) {
	level.Debug(t.logger).Log("STATE", "enterStatePreparing", "state", t.FSM.Current())
	t.EventTrigger(EventStart)
}

func (t *Table) leaveStatePreparing(e *fsm.Event) {
	level.Debug(t.logger).Log("STATE", "leaveStatePreparing", "state", t.FSM.Current())
}

func (t *Table) enterStatePlaying(e *fsm.Event) {
	level.Debug(t.logger).Log("STATE", "enterStatePlaying", "state", t.FSM.Current())
	time.Sleep(time.Millisecond * 500)
	t.EventTrigger(EventSettle)
}

func (t *Table) leaveStatePlaying(e *fsm.Event) {
	level.Debug(t.logger).Log("STATE", "leaveStatePlaying", "state", t.FSM.Current())
}

func (t *Table) enterStateStop(e *fsm.Event) {
	level.Debug(t.logger).Log("STATE", "enterStateStop", "state", t.FSM.Current())

	t.EventTrigger(EventOver)
}

func (t *Table) leaveStateStop(e *fsm.Event) {
	level.Debug(t.logger).Log("STATE", "leaveStateStop", "state", t.FSM.Current())
}

func (t *Table) eventListener() {
	go func() {
		for {
			select {
			case event := <-t.EventChan:
				if err := t.FSM.Event(event); err != nil {
					level.Error(t.logger).Log("Table", "eventListener", "event", event, "state", t.FSM.Current())
				}
			}
		}
	}()
}
