# BigBlueButton Web API

**bbb-web-api** -- is a web server developed with golang. 

## Web Server

The **bbb-client-logger** is a standalone web server that provides an HTTP API for storing logs.

### Compiling

The application can be built per platfrom using `go build` command .If you have goxc installed you should be able to create package by running `./build.sh` file, the output directory is set to `bin`.


* Install Go
```
cd src
sudo ./all.bash
```

* Update your ~/.profile file to contain the following Go environement variables
```
export GOROOT=/home/ubuntu/go
export GOPATH=/home/ubuntu/workspace/go
export GOBIN=/home/ubuntu/go/bin
export PATH=$PATH:$GOROOT/bin:$GOBIN
```

* Refresh your .profile
```
source ~/.profile
```


* Go to the project root and install its dependencies
```
go get github.com/gorilla/mux
go get github.com/spf13/viper
go get github.com/jhoonb/archivex
go get github.com/juju/loggo
go get gopkg.in/natefinch/lumberjack.v2
```

* And finally run the build shell script
```
./build.sh
```

### Usage

Run the server

```sh
./web-api
```

The server is immediately ready for querying.

### Configuration

A default configuration is hardcoded into the application but can be overriden by creating a config.json. The provided config.json contains the settings that can be changed.


### Querying

To test the server you need a HTTP client. The default server port is 8090.


It is possible to choose between two return types, JSON which is by default enabled and XML by adding an HTTP GET parameter named 'marshall'.

### Development

Used third party libraries are:

github.com/gorilla/mux : http://www.gorillatoolkit.org/pkg/mux

github.com/spf13/viper : https://github.com/spf13/viper


