# learning-go
Code written while learning go


## Basic Commands
    - go mod init {{file name}}
    - go env



## Helpful tips
    - new() vs make()
        - new() allocates memory but does not initialize whereas make() allocates memory and intializes
        - new creates zeroed storage( cannot put data) whereas make() gives non zeroed stoage.


## Useful inbuilt Packages
 - bufio package is for taking input (refer 03 and 04)
 - strconv packages helps with converting string
 - strings package helps working with strings
 - time package
 - sync package
 - Mutex



## Things to remember

    - file with the main function needs to start have `package main` at top
    - public variables start in capital letters, private ones start in lower
    - nil is the syntax for null value
    - formatting time : ```go 01-02-2006 15:04:05 Monday```
    - Pointer is a direct reference to memory location. using this ensures original value is being transferred and not a copy
    - array is not a commonly used data type in this. Slice is what is dominantly used
    - ranges in slice, the last range are non-inclusive
    - You need to encode and decode jsons for APIs. If you are creating a JSON you need to encode, if you are recieving a json you need to decode it. Use the json(inbuilt) library
    - Forgot what context is. Need to look into that again



semvar.org(semvar.org)


## Libraries:
    - Routing : Gorilla Mux `go get -u github.com/gorilla/mux`
    - MongoDB Driver: `go get go.mongodb.org/mongo-driver/mongo`