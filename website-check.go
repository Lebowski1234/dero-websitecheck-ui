/*
User interface for Website Validation Smart Contract by thedudelebowski. 
Version 1.0

Please note: This UI was written for the Dero Stargate testnet. Use at your own risk! The code could be simplified in some areas, and error handling may not be 100% complete. 

Github link: https://github.com/lebowski1234/dero-websitecheck-ui 
*/

package main

import (
	"bufio"	
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"net"
	"net/http"	
	"github.com/dixonwille/wmenu"
	"github.com/tidwall/gjson"
	"github.com/atotto/clipboard"

)

//For menu
type menuItem int


//For daemon call
type PayloadKeys struct { 
	TxsHashes []string `json:"txs_hashes"` 
	ScKeys []string `json:"sc_keys"` 
}



//For menu
const (
	autoCheck menuItem = iota	
	enterSCID 
	viewWebsite
	exit
)

//For menu
var menuItemStrings = map[menuItem]string{
	autoCheck:	"Check website using SCID from clipboard, then copy website URL to clipboard",
	enterSCID:	"Enter smart contract ID (SCID)",
	viewWebsite:	"Check website using entered SCID",
	exit:		"Exit",		
}


var SCID string
var websiteURL string
var passed bool

func main() {
	mm := mainMenu()
	err := mm.Run()
	if err != nil {
		wmenu.Clear()		
		fmt.Println(err)
		mm.Run()		
		
	}


	
}


/*-----------------------------------------------------------Menu-----------------------------------------------------------------*/


func mainMenu() *wmenu.Menu {
	menu := wmenu.NewMenu("Website check: choose an option.")
	
	menu.Option(menuItemStrings[autoCheck], autoCheck, true, func(opt wmenu.Opt) error { //change false to true to make default option
		wmenu.Clear()
		autoWebsiteCheck() 	
		mm := mainMenu()
		return mm.Run()
	})

	menu.Option(menuItemStrings[enterSCID], enterSCID, false, func(opt wmenu.Opt) error { 
		wmenu.Clear()		
		getSCID() 	
		mm := mainMenu()
		return mm.Run()
	})

	menu.Option(menuItemStrings[viewWebsite], viewWebsite, false, func(opt wmenu.Opt) error {
		wmenu.Clear()
		if SCID != "" {		
			getWebsiteDetails(SCID)		
		} else {
			fmt.Println("Please enter a SCID (Menu Option 1)\n")
		}			
		mm := mainMenu()
		return mm.Run()
	})
	

	menu.Option(menuItemStrings[exit], exit, false, func(opt wmenu.Opt) error {
		wmenu.Clear()		
		return nil //Exit		
		
	})
	
	menu.Action(func(opts []wmenu.Opt) error {
		if len(opts) != 1 {
			return errors.New("wrong number of options chosen")
		}
		wmenu.Clear()
		mm := mainMenu()
		return mm.Run()
		
	})
	return menu
}



//Carry out check using SCID from clipboard, then copy website URL from SC to clipboard
func autoWebsiteCheck() {
	SCID, err:= clipboard.ReadAll()
	if err != nil {
		fmt.Println(err) //if xclip not installed in Linux, will print error then return
		return
	}
	
	if SCID != "" {		
		getWebsiteDetails(SCID)		
	} else {
		fmt.Println("Please enter a SCID (Menu Option 1)\n")
	}
	
	//only copy SC website url to clipboard if security check has passed
	if passed == true {
		clipboard.WriteAll(websiteURL)	
	}

}



//Get SCID, save to memory
func getSCID() {
	scanner := bufio.NewScanner(os.Stdin)
	var text string
	fmt.Print("Enter SCID: ")
	scanner.Scan()
	text = scanner.Text()
	wmenu.Clear()	
	SCID = text
	fmt.Println("SCID entered: ", text)
	fmt.Print("Press 'Enter' to continue...")
  	bufio.NewReader(os.Stdin).ReadBytes('\n')
      
}



func pressToContinue() {
	fmt.Print("Press 'Enter' to continue...")
  	bufio.NewReader(os.Stdin).ReadBytes('\n')
	//wmenu.Clear()
      
}




/*-----------------------------------------------------------RPC Functions-----------------------------------------------------------------*/


//getKeysFromDaemon: send RPC call with list of keys, do error checking, return raw data in string form for JSON extraction
func getKeysFromDaemon(scid string, keys []string) string {
	
	deamonURL:= "http://127.0.0.1:30306/gettransactions"
	txHashes:= []string{scid}
	
	data := PayloadKeys{
		TxsHashes: txHashes,
		ScKeys: keys, 
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		println("Error in function getKeysFromDaemon:")
		fmt.Println(err)
		return ""
	}
	body := bytes.NewReader(payloadBytes)
	
	result, err:=rpcPost(body, deamonURL)
	if err != nil {
		println("Error in function getKeysFromDaemon:")
		fmt.Println(err)
		return ""
	}
	
	//Check to see if we have got the expected response from the daemon:
	if !gjson.Valid(result) {
		println("Error, result not in JSON format")
		return ""
	}
	
	validResponse := gjson.Get(result, "txs_as_hex") 	
	//fmt.Printf("Array value is: %s", validResponse.Array()[0])
	
	validResponse0:= validResponse.Array()[0]

	if validResponse0.String() == "" { //array position 0 value will be [""] if SCID not found or invalid
		println("Error, SCID not found")
		return ""
	}
	
	
	return result
	
			
}



//getWebsiteDetails: displays website security details
func getWebsiteDetails(scid string) {
	
	//var websiteURL string
	var websiteIP string
	var websiteDescription string
			
	scKeys:= []string{"websiteURL", "websiteIP", "websiteDescription"}
	result:= getKeysFromDaemon(scid, scKeys)
	if result == "" {return}


	//Response ok, extract keys from JSON
	
	
	raw := gjson.Get(result, "txs.#.sc_keys.websiteURL")
	resultString:=raw.String() 
	if resultString != "[]" { //weird way of checking value exists, raw.Exists() doesn't seem to work in this context
		websiteURL=raw.Array()[0].String()
				
	}

	raw = gjson.Get(result, "txs.#.sc_keys.websiteIP")
	resultString=raw.String() 
	if resultString != "[]" {
		websiteIP=raw.Array()[0].String()
				
	}

	
	raw = gjson.Get(result, "txs.#.sc_keys.websiteDescription")
	resultString=raw.String() 
	if resultString != "[]" { 
		websiteDescription=raw.Array()[0].String()
				
	}

	checkIP, err := net.LookupIP(websiteURL)	
	if err != nil {
		fmt.Println(err)
		
	}

		
	fmt.Printf("Website URl from smart contract is: %s \n", websiteURL)
	fmt.Printf("Website Description from smart contract is: %s \n", websiteDescription)
	fmt.Printf("Website IP address from smart contract is: %s \n", websiteIP)
	fmt.Printf("Website IP address from net lookup is: %s \n", checkIP[0].String())

	if checkIP[0].String() == websiteIP {
		fmt.Printf("Website IP address matches, website is safe!")
		passed = true
	} else {
		fmt.Printf("Website IP address does not match smart contract, security breach!")
		passed = false
	}	
	
	fmt.Printf("\n") 

	pressToContinue()



}


//rpcPost: Send RPC request, return response body as string 
func rpcPost(body *bytes.Reader, url string) (string, error) {

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		println("Error in function rpcPost:")
		fmt.Println(err)		
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")


	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		println("Error in function rpcPost:")
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	response, err:= ioutil.ReadAll(resp.Body)
	result:=string(response)

	return result, err
}



