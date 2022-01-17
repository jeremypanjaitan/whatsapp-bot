# About
Whatsapp-bot was created for sending a message to many whatsapp contact. You will provide a message and a list of contact and than the program will send it to all of the contact that you have provide. This project using https://github.com/Rhymen/go-whatsapp. This package implements the Whatsapp Web API.

# How to run ?
- Create .env file consist of these variables
```
WHATSAPP_CONN_TIMEOUT= //e.g 30
QR_FILE_NAME= //e.g qr
REMOTE_JID_SUFFIX=s.whatsapp.net
PHONE_NUMBER_FILE_NAME= //e.g phone_number_list.csv
MESSAGE_FILE_NAME= //e.g message.txt
```
- Create message.txt, this file must be in the root directory of this project. The file consist of the message you want to send
- Create phone_number_list.csv, this file also in the root directory  of this project. This file consist of the users target number for example `6287834234`. The first `62` is the country code
- Run `go mod download` to download dependencies
- Run  `go run .`
- qr.png file generated, then you need to scan that qr using your phone whatsapp. 
- Message successfully delivered.
# Requirements
- go version >= 1.17.3 