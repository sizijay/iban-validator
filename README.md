# iban-validator

###### **INTRODUCTION**
 
The iban validator is a microservice, a REST API written in Golang to facilitate the validation of the _International Bank Account Number_ . 
Reference: [Wikipedia]

The API will accept a POST request and then after the validation of the request, a response will be sent back based on the validity of the IBAN.

The service has the capability to keep different information related to different countries since the IBAN varies from country to country. 


###### **SETUP/INSTALLATION**

The service can be directly downloaded and run or can be run with the makefile or even as a docker container easily.

**Run from terminal:**
The repository can easily be cloned or downloaded from github. Then, following commands can be executed from within the root directory to install dependencies and run.

`$ go mod vendor # go get can also be used` \
`$ go run main.go`

**Run with Makefile:** 
Executed from the root directory.

`$ make  # this will simply build & run the binary`\
`$ make build # build the binary in linux`\
`$ make run # run the built binary`\
`$ make clean # delete built binaries if existed and clean`\
`$ make buildmac # build on macOS`\
`$ make runmac # run built binary on macOS`\
`$ make test # run test cases with coverage`

**Docker Support:**
Dockerfile & Docker-compose.yml are also included. (Docker installation needed.) Executed from the root directory.

`$ docker build . --tag v1.0 # build a new image`\
`$ docker run --publish 8080:8081 v1.0 -d ### this will map 8080(local) port to 8081(container) port and will run the container in detached mode`

Without using above 2 commands, docker-compose.yml can be directly used with a single command.

`$ docker compose up -d ### this will build and run the binary in detached mode. docker commands can be used thereafter to manage`


###### **REQUEST-RESPONSE STRUCTURE**
The binary's default port is set to _8081_. (Can be changed from within the code or compose file if needed.) Request will consist of the IBAN. Attached herewith is the curl.

`curl --location --request POST '127.0.0.1:8081/validate/iban' \
 --header 'Content-Type: application/json' \
 --data-raw '{
     "IBAN": "LU28 0019 4006 4475 0000"
 }'` 

If the IBAN is valid, the server will respond with a _200_ status having the following response body.

`{
     "Data": {
         "IsValidIBAN": true
     }
 }`
 
 
The value will be _false_ otherwise and error responses will be sent accordingly.

[Postman Collection]

Configurations are saved in the yaml directory and can be updated with new records easily. Recommended Go version >1.16.

©Sisira Jayawardena




[Wikipedia]: https://en.wikipedia.org/wiki/International_Bank_Account_Number
[Postman Collection]: https://www.getpostman.com/collections/f6b5a110c328282d58f0