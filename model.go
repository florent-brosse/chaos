package main

import (
	"bytes"
	"encoding/json"
	"time"
)

type Scenario struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Tasks       []Task `json:"tasks"`
	Id          string `json:"id"`
	Done        bool   `json:"done"`
}

type Task struct {
	Id       string            `json:"id"`
	Type     TaskType          `json:"type"`
	Start    time.Time         `json:"start"`
	Duration int               `json:"duration"`
	Tags     []string          `json:"tags"`
	Param    map[string]string `json:"param"`
	Launched bool              `json:"launched"`
	Done     bool              `json:"done"`
	pid      int
}

func FinishTime(scenario *Scenario) time.Time {
	longestTime := time.Time{}
	for _, task := range scenario.Tasks {
		currentFinishTaskTime := task.Start.Add(time.Second * time.Duration(task.Duration))
		if currentFinishTaskTime.After(longestTime) {
			longestTime = currentFinishTaskTime
		}
	}
	return longestTime
}

type TaskType int

const (
	KILL_PROCESS TaskType = iota
	START_SERVICE
	STOP_SERVICE
	CREATE_FILE
	USE_RAM
	USE_CPU
	USE_IO
	SHUTDOWN
	ADD_LATENCY
	CHANGE_TIME
	RUN_COMMAND
	BLOCK_RANGE_INPUT_PORT
)

func (t TaskType) String() string {
	return taskTypesId[t]
}

var taskTypesId = map[TaskType]string{
	KILL_PROCESS:  "KILL_PROCESS",
	START_SERVICE: "START_SERVICE",
	STOP_SERVICE:  "STOP_SERVICE",
	CREATE_FILE:   "CREATE_FILE",
	USE_RAM:       "USE_RAM",
	USE_CPU:       "USE_CPU",
	USE_IO:        "USE_IO",
	SHUTDOWN:      "SHUTDOWN",
	ADD_LATENCY:   "ADD_LATENCY",
	CHANGE_TIME:   "CHANGE_TIME",
	RUN_COMMAND:   "RUN_COMMAND",
}

var taskTypesName = map[string]TaskType{
	"KILL_PROCESS":  KILL_PROCESS,
	"START_SERVICE": START_SERVICE,
	"STOP_SERVICE":  STOP_SERVICE,
	"CREATE_FILE":   CREATE_FILE,
	"USE_RAM":       USE_RAM,
	"USE_CPU":       USE_CPU,
	"USE_IO":        USE_IO,
	"SHUTDOWN":      SHUTDOWN,
	"ADD_LATENCY":   ADD_LATENCY,
	"CHANGE_TIME":   CHANGE_TIME,
	"RUN_COMMAND":   RUN_COMMAND,
}

func (d *TaskType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(taskTypesId[*d])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (d *TaskType) UnmarshalJSON(b []byte) error {
	// unmarshal as string
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	// lookup value
	*d = taskTypesName[s]
	return nil
}
