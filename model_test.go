package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestModel(t *testing.T) {
	duration_Milliseconds := 500 * time.Millisecond
	t2 := time.Date(2017, time.February, 16, 0, 0, 0, 0, time.UTC)
	task := Task{"10", USE_CPU, t2, duration_Milliseconds, []string{"toto", "DC1"}, make(map[string]string), false, false, 1}
	res2B, _ := json.Marshal(&task)
	//fmt.Println(string(res2B))
	expected := "{\"id\":\"10\",\"type\":\"USE_CPU\",\"start\":\"2017-02-16T00:00:00Z\",\"duration\":500000000,\"tags\":[\"toto\",\"DC1\"],\"param\":{},\"launched\":false,\"done\":false}"
	if strings.Compare(string(res2B), expected) != 0 {
		t.Errorf("model were not equals got: %v, want: %v.", string(res2B), expected)
	}
}

func TestModel2(t *testing.T) {
	duration_Milliseconds := 60 * time.Second
	t2 := time.Now()
	task := Task{"10", USE_CPU, t2, duration_Milliseconds, []string{"toto", "DC1"}, make(map[string]string), false, false, 1}

	scenario := Scenario{"add cpu", "add cpu", []Task{task}, "1", false}
	res2B, _ := json.Marshal(&scenario)
	fmt.Println(string(res2B))
}

func TestModel3(t *testing.T) {
	duration_Milliseconds := 60 * time.Second
	t2 := time.Now()
	task := Task{"10", USE_CPU, t2, duration_Milliseconds, []string{"toto", "DC1"}, make(map[string]string), false, false, 1}

	scenario := Scenario{"add cpu", "add cpu", []Task{task}, "1", false}
	modifyScenario(&scenario)
	fmt.Println(scenario.Tasks)
	fmt.Println(scenario)
}

func modifyTask(task *Task) {

}

func modifyScenario(scenario *Scenario) {
	scenario.Done = true
	for i := 0; i < len(scenario.Tasks); i++ {
		task := &scenario.Tasks[i]
		task.Done = true
	}
}
