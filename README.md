
# Blocksize Tech Assignment

**Please do not publicly fork this repository but rather make a private copy, thank you!**
  
This repository contains the boilerplate to a small microservice that is meant for you to show off your coding skills as well as general technical abilities.  
    
In here you will find two directories:
 - `proto` contains all protobuf definitions:  
   - `service.proto` contains the protobuf service & message definitions for the ApikeyService that has to be implemented
 - `assignment` contains a boilerplate Go application with the needed structure for a gRPC server already set up.  

A note on [Go](https://golang.org/):  
At Blocksize we work in a heavily Go-centric environment with almost all of our backend written in Go. If you have not worked with Go before and want to use any language or framework, that's fine. We actively encourage you to use the tools you think you can show your abilities the best and are excited for the solution you're going to come up with in the language of your choice.

The goal of this technical assignment is to implement a service that exposes a gRPC server interface for other services to interact with.  
The service revolves around several actions around apikeys. One of the challenges at Blocksize is to securely store our user's exchange apikeys while still making them accessible by authorized parties and functions.
For the sake of this assignment, an apikey always consists of an "apikey" and a "secret". For both should be taken care to not expose them accidentally, but only when requested to. They should never appear in any logs.
  
The idea is to use this service as a central "repository" for apikeys. They can be added, listed and requested by other services whenever needed.
  
There are three rpc methods to be implemented:
- `AddApikey` can be used to add a new apikey to the system that is assigned to the specified userid  
- `ListApikeys` can be used to list all apikeys that are currently added for the given userid
- `GetApikey` can be used to get the plaintext details of the apikey

A note on the gRPCs/protobufs:  
We've deliberately split the project into two parts, the .proto files and the actual implementation. "Compiling" the .proto files into their language-specific stubs can be a quite tedious task as it requires a lot of dependencies. For Go at least, we've already added the output files to this repository. Should you pick another language that requires compilation, compile them on your machine and add them to your code - there's no need to re-compile them during your Docker build.
  
If there are any changes necessary to any of the already existing files, please do not refrain from changing them in whatever way seems to be necessary to you.  
  
If you have any questions, please ask us. We're happy to help and discuss your ideas. 
  
### What we're expecting  
- An implementation of the ApikeyService server in a language of your choice, listening on port 50051  
- A Dockerfile to (1) build the application if needed and (2) run the application
- The service connects to an external SQL server of your choice for data storage  
- The credentials needed to connect to the SQL server can be passed to the service either using envvars or configuration files during runtime
- Some lightweight documentation how to configure and start the Docker container
  
### Helpful links
- The gRPC quick start guides for all supported languages: https://grpc.io/docs/languages/
- The Protocol Buffers tutorials, especially the language specific sections on "Compiling your Protocol Buffers". E.g. on [Dart](https://developers.google.com/protocol-buffers/docs/darttutorial#compiling-your-protocol-buffers)
- BloomRPC, a Postman-like gRPC GUI client for interacting with gRPC services: https://github.com/uw-labs/bloomrpc

Happy Coding!