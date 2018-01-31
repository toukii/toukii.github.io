package main

import (
	"fmt"
	"os"
	"text/template"
)

type TableInfo struct {
	CnName string
	EnName string
}
type ModuleTable struct {
	ModuleName string
	TableList  []TableInfo
}

func main() {
	t1()
}

func t1() {
	moduleTable := [2]ModuleTable{}
	moduleTable[0].TableList = []TableInfo{TableInfo{CnName: "CN11"}}
	moduleTable[1].TableList = []TableInfo{TableInfo{CnName: "CN22"}}
	tpl, _ := template.ParseFiles("tpl1.tpl")

	err := tpl.Execute(os.Stdout, moduleTable)
	fmt.Println(err)
}
