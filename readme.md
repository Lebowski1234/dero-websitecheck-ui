# User Interface for Website Validation Smart Contract

This is the user interface (UI) for the Website Validation Smart Contract written for the Dero Stargate Smart Contract competition, located in this repository: [https://github.com/lebowski1234/dero-websitecheck](https://github.com/lebowski1234/dero-websitecheck). This document contains basic usage instructions.

Binaries for Windows and Linux (both 64 bit) are located here:

* Linux: [https://github.com/Lebowski1234/dero-websitecheck-ui/raw/master/binaries/website-check-linux64.tar.gz](https://github.com/Lebowski1234/dero-websitecheck-ui/raw/master/binaries/website-check-linux64.tar.gz)

* Windows: [https://github.com/Lebowski1234/dero-websitecheck-ui/raw/master/binaries/website-check-windows64.rar](https://github.com/Lebowski1234/dero-websitecheck-ui/raw/master/binaries/website-check-windows64.rar)

Or follow the instructions below to compile. 


## Compiling

All development was done in Ubuntu using Go version 1.11.4.

First, download dependencies:

```
$ go get -u github.com/tidwall/gjson
$ go get -u github.com/dixonwille/wmenu
$ go get -u github.com/atotto/clipboard
```

Then build:

```
$ go build website-check.go
```


## Running

The Dero Stargate daemon must be running first, with the standard RPC port open. 

Get the Dero Stargate binaries here:

[https://git.dero.io/DeroProject/Dero_Stargate_testnet_binaries](https://git.dero.io/DeroProject/Dero_Stargate_testnet_binaries)


To run the Dero Stargate daemon (in Linux):

```
./derod-linux-amd64 --testnet
```

To run the website validation user interface:

```
$ ./website-check
```

The instructions are the same for Windows, without the './'


## Usage

Refer to the website validation smart contract [readme](https://github.com/lebowski1234/dero-websitecheck) for an explanation of how the contract works. All options in the user interface are self explanatory and intuitive. Run the user interface and choose from Options 1 to 4 (e.g. type '1' then enter):


### Option 1 - Check Website Using SCID From Clipboard

This is the default option, and can also be called just by pressing enter. Copy the SCID to the clipboard, then press 1 or hit enter. The program fetches the domain name, IP address, and description from the smart contract, then carries out a lookup on the domain name to check the live IP address. The live IP address is then compared to the IP address stored in the smart contract. If the two match, then the domain name from the smart contract is copied to the clipboard, ready for pasting into your web browser. If the IP addresses do not match, then a warning is displayed, and the domain is not copied to the clipboard. 

Note that to use this option in Linux, a suitable clipboard application must be installed first. For example in Ubuntu:

```
$ sudo apt-get update
$ sudo apt-get install xclip
```

### Option 2 - Enter Smart Contract ID (SCID)

A manual alternative to Option 1. Enter the smart contract ID (SCID). This must be done before Option 3 is selected. 

### Option 3 - Check Website Using Entered SCID

As per Option 1, but the domain name is not copied to the clipboard. This is useful for Linux distributions which do not have a clipboard application installed. 

### Option 4 - Exit

Exit the user interface. 


## Contact Details

I plan to update the smart contract and user interface for the Dero main network, when smart contracts become live. To report a bug, please open an issue in github. 

My contact details are: thedudelebowski1234@gmail.com

Finally, if you found this useful, any Dero donations are most welcome! dERoSME4c5GNUvPo27NsRFeJPR1FKiYt87g8Gknbm6JU9eL3xRPDs6JijHuVNxVzyFZXg1wxjbh52Hu9gUfWd3Lx5QRNTXvJWZ



