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
	Iopath         string
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
	flag.String("iopath", "50", "io task path")

	flag.Bool("file", false, "create file task")
	flag.String("fileusage", "50", "file task usage")
	flag.String("filepath", "/tmp/BIG_FILE", "file task path")

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
			io(conf.Iousage, conf.Iopath)
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
		durationSec := time.Second * time.Duration(task.Duration)
		endDate := task.Start.Add(durationSec)
		if !task.Done && !task.Launched && task.Start.Before(timeNow) && (endDate).After(timeNow) {
			task.Launched = true
			launchTask(task)
		}
		if !task.Done && task.Launched && (endDate).Before(timeNow) {
			task.Done = true
			stopTask(task)
		}
		if !task.Done && !task.Launched && (endDate).Before(timeNow) {
			task.Launched = false
			task.Done = true
		}
		allDone = allDone && task.Done
	}
	scenario.Done = allDone
}

func launchTask(task *Task) {
	var command string
	switch task.Type {
	case STOP_SERVICE:
		command = "service " + task.Param["servicename"] + " stop"
	case START_SERVICE:
		command = "service " + task.Param["servicename"] + " start"
	case CREATE_FILE:
		command = "./chaos --file --filepath " + task.Param["path"] + " --fileusage " + task.Param["usage"]
	case USE_RAM:
		command = "./chaos --ram --ramusage " + task.Param["usage"]
	case USE_CPU:
		command = "./chaos --cpu --cpuusage " + task.Param["usage"]
	case USE_IO:
		command = "./chaos --io --iopath " + task.Param["path"] + " --iousage " + task.Param["usage"]
	case KILL_PROCESS:
		command = "killall " + task.Param["processname"]
	case SHUTDOWN:
		command = "shutdown now"
	case RUN_COMMAND:
		command = task.Param["command"]
	case ADD_LATENCY:
		command = "tc qdisc add dev " + task.Param["interface"] + " root netem delay " + task.Param["delay"] + "ms"
	case CHANGE_TIME:
		command = "" //not yet how to do it a lot of code use time.Now()
	case BLOCK_RANGE_INPUT_PORT:
		command = "/sbin/iptables -A INPUT -p tcp --destination-port " + task.Param["rangeport"] + " -j DROP -m comment --comment \"TASK_ID='" + task.Id + "'\""
	}
	task.pid = launchCommand(command)
}
func stopTask(task *Task) {
	var command string
	switch task.Type {
	case CREATE_FILE:
		command = "rm -rf " + task.Param["path"]
	case ADD_LATENCY:
		command = "tc qdisc del dev " + task.Param["interface"] + " root netem"
	case CHANGE_TIME:
		command = ""
	case BLOCK_RANGE_INPUT_PORT:
		command = "/sbin/iptables -D INPUT -p tcp --destination-port " + task.Param["rangeport"] + " -j DROP"
	default:
		command = "kill " + strconv.Itoa(task.pid) + ";pkill -P " + strconv.Itoa(task.pid)
	}
	if command != "" {
		launchCommand(command)
	}
}
