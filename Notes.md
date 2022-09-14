# Go Notes

## Setup

- Go programs are organized into packages
- Go’s standard library, provides different core packages for us to use
- Go is a statically typed language. It’s mean that go compiler need to know the data type when declaring the variable
- Type inference

### Install

- IDE: VS Code
- Go compiler

### Create new module
```shell
$ go mod init <module path>
```

## Variables

- Variables are used to store values
- Like containers for values
- Give the variable a name & reference everywhere in the app

### Benefits

- Update the value only once
- Make our app more flexible

**&** pointer to memory address e.g: 

*var username*

…

&username

### Variable Scopes

Scope is the region of a program, where a define variable can be accessed

**3 Levels of Scopes**

- Local: Declaration within function. Can be use only within that function

    Declaration within block. Can be user only within that block

- Package: Declaration outside all functions. Can be used everywhere in the same package

- Global: Declaration outside all functions and uppercase first letter. Can be used everywhere across all packages

## Data types

- String: For textual data, defined with double quotes
- Booleans => bool
- Arrays
- Integer
- Maps
- Struct

### Array

Declaring array: 

var <array_name> [size]<data type> [e.g](http://e.gL) ***var booking [50]string***

**Slices** : is an abstraction of an Array,  more flexible and powerful

### Integer

Integers: Representing whole number, positive and negative. E.g: 5,120, -20

| Go | Java | Size |
| --- | --- | --- |
| int8 | byte | 0-255 |
| int16 | short | 0-65535 |
| int32 | int | 0-4294967295 |
| int64 | long |  |

### Maps

Declaring maps:

***var** myMap **map**[string]string*

Initializing map list

**var** listMap = ***make([]map[string]string)***

### Struct

Collect different data types of data

**type statement - Custom Types**

the **type** keyword create a new type with the name you specify

Create a type called “UserData” based on a struct of firstName, lastName…

```go
type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}
```

It could also create a type based on every other data type like int, string, etc

## Loops

for {}

for index, <iteration_var>  

## Functions

Syntax:

```go
func functionName (arg1, …) (result type1, result type2) {
    ...
    block of statements
    ...
}

```

### Packages

- Go programs are organised into packages
- A package is a collection of Go files

## Concurrency with Goroutine

**go** keyword starts a new goroutine

A goroutine is a lightweight thread managed by the Go runtime

Synchronize with ***WayGroup*** from package ***sync***