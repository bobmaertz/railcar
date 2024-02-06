# Railcar

Railcar (working name) is an extendable oauth server written in Go. Its still under very early development and is not quite ready for use. 

This is a project to learn more about the fundamentals of the OAuth RFC from a bare bones implementation. This project has minimal external dependencies.

## Goals 
- No Dependencies
- Support for (RFC6749)[https://datatracker.ietf.org/doc/html/rfc6749)
## Build
```sh 
make build 
```


## Run 

```sh 
./bin/railcar-darwin -config_file ./path/to/config.yml
```

or 

```sh 
make all 

```

## TODO
[] Documentation / Examples 
[] Integration Tests 
[] Github Actions CI/CD
[] Database support 
