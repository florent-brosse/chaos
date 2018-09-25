package main

import (
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
	Duration time.Duration     `json:"duration"`
	Tags     []string          `json:"tags"`
	Param    map[string]string `json:"param"`
	Launched bool              `json:"launched"`
	Done     bool              `json:"done"`
	pid      int
}

func FinishTime(scenario *Scenario) time.Time {
	longestTime := time.Time{}
	for _, v := range scenario.Tasks {
		currentFinishTaskTime := v.Start.Add(v.Duration)
		if currentFinishTaskTime.After(longestTime) {
			longestTime = currentFinishTaskTime
		}
	}
	return longestTime
}

type TaskType int

const (
	KILL_PROCESS TaskType = iota
	START_PROCESS
	CREATE_FILE
	USE_RAM
	USE_CPU
	USE_IO
	SHUTDOWN
	ADD_LATENCY
	CHANGE_TIME
)
