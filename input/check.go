/*

=======================
Scilla - Information Gathering Tool
=======================

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see http://www.gnu.org/licenses/.

	@Repository:  https://github.com/edoardottt/scilla

	@Author:      edoardottt, https://www.edoardoottavianelli.it

*/

package input

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/edoardottt/scilla/output"
	"github.com/edoardottt/scilla/utils"
)

//ReportSubcommandCheckFlags >
func ReportSubcommandCheckFlags(reportCommand flag.FlagSet, reportTargetPtr *string,
	reportOutputPtr *string, reportPortsPtr *string, reportCommonPtr *bool,
	reportSpysePtr *bool, reportVirusTotalPtr *bool, reportSubdomainDBPtr *bool,
	StartPort int, EndPort int, reportIgnoreDirPtr *string,
	reportIgnoreSubPtr *string, reportTimeoutPort *int) (int, int, []int, bool, []string, []string) {
	// Required Flags
	if *reportTargetPtr == "" {
		reportCommand.PrintDefaults()
		os.Exit(1)
	}
	//Verify good inputs
	if !utils.IsURL(*reportTargetPtr) {
		fmt.Println("The inputted target is not valid.")
		os.Exit(1)
	}
	if !output.OutputFormatIsOk(*reportOutputPtr) {
		fmt.Println("The output format is not valid.")
		os.Exit(1)
	}
	//common and p not together
	if *reportPortsPtr != "" && *reportCommonPtr {
		fmt.Println("You can't specify a port range and common option together.")
		os.Exit(1)
	}

	if *reportSpysePtr && !*reportSubdomainDBPtr {
		fmt.Println("You can't specify Spyse and not the Open Database option.")
		fmt.Println("If you want to use Spyse Api, set also -db option.")
		os.Exit(1)
	}

	if *reportVirusTotalPtr && !*reportSubdomainDBPtr {
		fmt.Println("You can't specify VirusTotal and not the Open Database option.")
		fmt.Println("If you want to use VirusTotal Api, set also -db option.")
		os.Exit(1)
	}

	var portsArray []int
	var portArrayBool bool
	if *reportPortsPtr != "" {
		if strings.Contains(*reportPortsPtr, "-") && strings.Contains(*reportPortsPtr, ",") {
			fmt.Println("You can specify a ports range or an array, not both.")
			os.Exit(1)
		}
		if strings.Contains(*reportPortsPtr, "-") {
			portsRange := string(*reportPortsPtr)
			StartPort, EndPort = utils.CheckPortsRange(portsRange, StartPort, EndPort)
			portArrayBool = false
		} else if strings.Contains(*reportPortsPtr, ",") {
			portsArray = utils.CheckPortsArray(*reportPortsPtr)
			portArrayBool = true
		} else {
			portsRange := string(*reportPortsPtr)
			StartPort, EndPort = utils.CheckPortsRange(portsRange, StartPort, EndPort)
			portArrayBool = false
		}
	}
	var reportIgnoreDir []string
	var reportIgnoreSub []string
	if *reportIgnoreDirPtr != "" {
		toBeIgnored := string(*reportIgnoreDirPtr)
		reportIgnoreDir = utils.CheckIgnore(toBeIgnored)
	}
	if *reportIgnoreSubPtr != "" {
		toBeIgnored := string(*reportIgnoreSubPtr)
		reportIgnoreSub = utils.CheckIgnore(toBeIgnored)
	}
	if *reportTimeoutPort < 1 || *reportTimeoutPort > 100 {
		fmt.Println("Port Scan timeout must be an integer between 1 and 100.")
		os.Exit(1)
	}

	return StartPort, EndPort, portsArray, portArrayBool, reportIgnoreDir, reportIgnoreSub
}

//DNSSubcommandCheckFlags >
func DNSSubcommandCheckFlags(dnsCommand flag.FlagSet, dnsTargetPtr *string, dnsOutputPtr *string) {
	// Required Flags
	if *dnsTargetPtr == "" {
		dnsCommand.PrintDefaults()
		os.Exit(1)
	}
	//Verify good inputs
	if !utils.IsURL(*dnsTargetPtr) {
		fmt.Println("The inputted target is not valid.")
		os.Exit(1)
	}
	if !output.OutputFormatIsOk(*dnsOutputPtr) {
		fmt.Println("The output format is not valid.")
		os.Exit(1)
	}
}

//SubdomainSubcommandCheckFlags >
func SubdomainSubcommandCheckFlags(subdomainCommand flag.FlagSet, subdomainTargetPtr *string, subdomainOutputPtr *string,
	subdomainNoCheckPtr *bool, subdomainDBPtr *bool, subdomainWordlistPtr *string, subdomainIgnorePtr *string,
	subdomainCrawlerPtr *bool, subdomainSpysePtr *bool, subdomainVirusTotalPtr *bool) []string {
	// Required Flags
	if *subdomainTargetPtr == "" {
		subdomainCommand.PrintDefaults()
		os.Exit(1)
	}
	//Verify good inputs
	if !utils.IsURL(*subdomainTargetPtr) {
		fmt.Println("The inputted target is not valid.")
		os.Exit(1)
	}
	if !output.OutputFormatIsOk(*subdomainOutputPtr) {
		fmt.Println("The output format is not valid.")
		os.Exit(1)
	}

	//no-check checks
	if *subdomainNoCheckPtr && !*subdomainDBPtr {
		fmt.Println("You can use no-check only with db option.")
		os.Exit(1)
	}
	if *subdomainNoCheckPtr && *subdomainWordlistPtr != "" {
		fmt.Println("You can't use no-check with wordlist option.")
		os.Exit(1)
	}
	if *subdomainNoCheckPtr && *subdomainIgnorePtr != "" {
		fmt.Println("You can't use no-check with ignore option.")
		os.Exit(1)
	}
	if *subdomainNoCheckPtr && *subdomainCrawlerPtr {
		fmt.Println("You can't use no-check with crawler option.")
		os.Exit(1)
	}

	if *subdomainSpysePtr && !*subdomainDBPtr {
		fmt.Println("You can't specify Spyse and not the Open Database option.")
		fmt.Println("If you want to use Spyse Api, set also -db option.")
		os.Exit(1)
	}

	if *subdomainVirusTotalPtr && !*subdomainDBPtr {
		fmt.Println("You can't specify VirusTotal and not the Open Database option.")
		fmt.Println("If you want to use VirusTotal Api, set also -db option.")
		os.Exit(1)
	}

	var subdomainIgnore []string
	if *subdomainIgnorePtr != "" {
		toBeIgnored := string(*subdomainIgnorePtr)
		subdomainIgnore = utils.CheckIgnore(toBeIgnored)
	}
	return subdomainIgnore
}

//PortSubcommandCheckFlags >
func PortSubcommandCheckFlags(portCommand flag.FlagSet, portTargetPtr *string, portsPtr *string,
	portCommonPtr *bool, StartPort int, EndPort int, portOutputPtr *string, portTimeout *int) (int, int, []int, bool) {
	// Required Flags
	if *portTargetPtr == "" {
		portCommand.PrintDefaults()
		os.Exit(1)
	}
	//common and p not together
	if *portsPtr != "" && *portCommonPtr {
		fmt.Println("You can't specify a port range and common option together.")
		os.Exit(1)
	}
	var portArrayBool bool
	var portsArray []int
	if *portsPtr != "" {
		if strings.Contains(*portsPtr, "-") && strings.Contains(*portsPtr, ",") {
			fmt.Println("You can specify a ports range or an array, not both.")
			os.Exit(1)
		}
		if strings.Contains(*portsPtr, "-") {
			portsRange := string(*portsPtr)
			StartPort, EndPort = utils.CheckPortsRange(portsRange, StartPort, EndPort)
			portArrayBool = false
		} else if strings.Contains(*portsPtr, ",") {
			portsArray = utils.CheckPortsArray(*portsPtr)
			portArrayBool = true
		} else {
			portsRange := string(*portsPtr)
			StartPort, EndPort = utils.CheckPortsRange(portsRange, StartPort, EndPort)
			portArrayBool = false
		}
	}
	//Verify good inputs
	if !utils.IsURL(*portTargetPtr) {
		fmt.Println("The inputted target is not valid.")
		os.Exit(1)
	}
	if !output.OutputFormatIsOk(*portOutputPtr) {
		fmt.Println("The output format is not valid.")
		os.Exit(1)
	}
	if *portTimeout < 1 || *portTimeout > 100 {
		fmt.Println("Port Scan timeout must be an integer between 1 and 100.")
		os.Exit(1)
	}

	return StartPort, EndPort, portsArray, portArrayBool
}

//DirSubcommandCheckFlags >
func DirSubcommandCheckFlags(dirCommand flag.FlagSet, dirTargetPtr *string, dirOutputPtr *string,
	dirIgnorePtr *string) []string {
	// Required Flags
	if *dirTargetPtr == "" {
		dirCommand.PrintDefaults()
		os.Exit(1)
	}
	//Verify good inputs
	if !utils.IsURL(*dirTargetPtr) {
		fmt.Println("The inputted target is not valid.")
		os.Exit(1)
	}
	if !output.OutputFormatIsOk(*dirOutputPtr) {
		fmt.Println("The output format is not valid.")
		os.Exit(1)
	}
	var dirIgnore []string
	if *dirIgnorePtr != "" {
		toBeIgnored := string(*dirIgnorePtr)
		dirIgnore = utils.CheckIgnore(toBeIgnored)
	}
	return dirIgnore
}
