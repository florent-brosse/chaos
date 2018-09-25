package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type config struct {
	Port           int
	Listen_address string
	Master_address string
	Master_port    int
	Tags           []string
	Cpu            bool
	Cpuusage       string
	Ram            bool
	Ramusage       string
	File           bool
	Filepath       string
	Fileusage      string
	Io             bool
	Iousage        string
}

var (
	conf      *config
	scenarios []Scenario
)

func getConf() *config {
	flag.Int("port", 8080, "default port to listen")

	flag.Bool("cpu", false, "launch cpu task")
	flag.String("cpuusage", "50", "cpu task usage")

	flag.Bool("ram", false, "launch ram task")
	flag.String("ramusage", "50", "ram task usage")

	flag.Bool("io", false, "launch io task")
	flag.String("iousage", "50", "io task usage")

	flag.Bool("file", false, "create file task")
	flag.String("fileusage", "50", "file task usage")
	flag.String("filepath", "/tmp/BIG_FILE", "file task usage")

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.SetEnvPrefix("chaos")
	viper.AutomaticEnv()
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	conf := &config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}
	return conf
}

func init() {
	conf = getConf()
}

func main() {

	if (conf.Cpu != conf.File != conf.Io != conf.Ram) || (!conf.Cpu && !conf.File && !conf.Io && !conf.Ram) {

		switch {
		case conf.Cpu:
			cpu(conf.Cpuusage)
		case conf.Ram:
			ram(conf.Ramusage)
		case conf.Io:
			io(conf.Iousage)
		case conf.File:
			makeFile(conf.Filepath, conf.Fileusage)
		default:
			fmt.Printf("Port: %v\n", conf.Port)
			fmt.Printf("listen_address: %v\n", conf.Listen_address)
			go doScenarios(&scenarios)
			startServer()
		}

	} else {
		var message string = `Error please use
chaos --ram --ramusage 80%
chaos --cpu --cpuusage 80%
chaos --file --fileusage 1% --filepath /tmp/BIGFILE`
		log.Fatal(message)
	}

}
func doScenarios(scenarios *[]Scenario) {
	for {
		for i := 0; i < len(*scenarios); i++ {
			scenario := &((*scenarios)[i])
			if !scenario.Done {
				doScenario(scenario)
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}
func doScenario(scenario *Scenario) {
	allDone := true
	timeNow := time.Now()
	for i := 0; i < len(scenario.Tasks); i++ {
		task := &scenario.Tasks[i]
		if !task.Done && !task.Launched && task.Start.Before(timeNow) && (task.Start.Add(task.Duration)).After(timeNow) {
			go launchTask(task)
		}
		if !task.Done && task.Launched && (task.Start.Add(task.Duration)).Before(timeNow) {
			go stopTask(task)
		}
		if !task.Done && !task.Launched && (task.Start.Add(task.Duration)).Before(timeNow) {
			task.Launched = false
			task.Done = true
		}
		allDone = allDone && task.Done
	}
	scenario.Done = allDone
}

func launchTask(task *Task) {
	task.Launched = true
	var command string
	switch task.Type {
	case KILL_PROCESS:
		command = "service " + task.Param["path"] + " stop"
	case START_PROCESS:
		command = "service " + task.Param["path"] + " start"
	case CREATE_FILE:
		command = "./chaos --file --filepath " + task.Param["path"] + " --fileusage " + task.Param["usage"]
	case USE_RAM:
		command = "./chaos --ram --ramusage " + task.Param["usage"]
	case USE_CPU:
		command = "./chaos --cpu --cpuusage " + task.Param["usage"]
	case USE_IO:
		command = "./chaos --io --iousage " + task.Param["usage"]
	case SHUTDOWN:
		command = "killall " + task.Param["processname"]
	case ADD_LATENCY:
		command = ""
	case CHANGE_TIME:
		command = ""
	}
	task.pid = launchCommand(command)
}
func stopTask(task *Task) {
	task.Done = true
	var command string
	switch task.Type {
	case CREATE_FILE:
		command = "rm -rf " + task.Param["path"]
	case ADD_LATENCY:
		command = ""
	case CHANGE_TIME:
		command = ""
	default:
		command = "kill -TERM -" + strconv.Itoa(task.pid)
	}
	if command != "" {
		launchCommand(command)
	}
}
