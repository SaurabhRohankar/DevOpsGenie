package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/inancgumus/screen"
	"github.com/zcalusic/sysinfo"
)

const (
	// Tools constants
	Docker   = "Docker"
	Ansible  = "Ansible"
	Jenkins  = "Jenkins"
	Terraform = "Terraform"
)

func main() {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("Welcome to the DevOps CLI Tool")

	tools := []string{Docker, Ansible, Jenkins, Terraform}

	printTools(tools)

	choice := getUserChoice(len(tools))

	fmt.Printf("You selected %s\n", tools[choice])

	checkSuperuserPrivilege()

	sysInfo := getSystemInfo()
	osVendor := extractOSVendor(sysInfo)

	fmt.Println("OS Vendor:", osVendor)
}

func printTools(tools []string) {
	for index, tool := range tools {
		fmt.Printf("%d. %s\n", index+1, tool)
	}
}

func getUserChoice(maxChoice int) int {
	var choice int
	fmt.Print("Please select which tool to install: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	if choice < 1 || choice > maxChoice {
		fmt.Println("Invalid choice. Please give a valid number")
		os.Exit(1)
	}

	return choice - 1
}

func checkSuperuserPrivilege() {
	current, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	if current.Uid != "0" {
		log.Fatal("requires superuser privilege")
	}
}

func getSystemInfo() *sysinfo.SysInfo {
	var si sysinfo.SysInfo
	si.GetSysInfo()
	return &si
}

func extractOSVendor(sysInfo *sysinfo.SysInfo) string {
	data, err := json.MarshalIndent(sysInfo, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	var osData map[string]interface{}
	osErr := json.Unmarshal([]byte(data), &osData)
	if osErr != nil {
		fmt.Println("Error:", osErr)
		os.Exit(1)
	}

	osVendor := osData["os"].(map[string]interface{})["vendor"].(string)
	return osVendor
}

