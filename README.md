# glonag-first
## Setup
 Download from officially website.
## if you are using ubuntu then follow listed steps.
 1. sudo wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
 2. sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
 3. nano ~/.bashrc
 ## copy paste listed code in above file
       export PATH=$PATH:/usr/local/go/bin
       export GOPATH=$HOME/go 
       export PATH=$PATH:$GOPATH/bin
## Save and exit, then run below commond
       source ~/.bashrc
## check version 
       go version
## run file
       go mod init student-service   
       go run main.go

## Initialize Directory
## If your project is not in /home/go/src, then you can do any where and need to execute below commond in project.
       go mod init student-service
       go mod tidy
