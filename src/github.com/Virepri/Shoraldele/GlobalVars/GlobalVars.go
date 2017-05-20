package GlobalVars

import "sync"

var ConfigLocs map[string]string
var SetupFuncs map[string]func(string) //runs on program startup
var CmdFuncs map[string]func(string,string)
var ModuleRoutines map[string]func()
var WaitGroup sync.WaitGroup