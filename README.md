# A sample Golang project 

## References

- [Language reference](https://golang.org/doc)
- [Language spec](https://golang.org/ref/spec)
- [Go by examples](https://gobyexample.com/)
- [Web application](https://astaxie.gitbooks.io/build-web-application-with-golang/en/)
- [A web application tutorial](https://golang.org/doc/articles/wiki/)

## Overview

_Golang_ is a procedure language with its syntax similar to _C_ language. It also has many syntax similar to _JavaScript_ (or _TypeScript_) and _Java_. Below is a list of syntax highlights:

| Subject | Explanation |
| --- | ------ |
| scope | _package_ level scope, defined by _<b>package</b> pkg_name_; upper case of the first character of a variable or function indicates it can be accessible from other packages |
|  | _block_ level scope, variables inside a _block_ surrounded by '{}' |
| data declaration | explicitly declare data type: <code>var foo Foo</code>, or implicitly <code>foo := Foo{ ...}</code> |
|  | a variable can be a <i>pointer</i> of a data type, <code>*foo</code>; relation between a data and its pointer is similar to _C_'s, but there is no passing by reference concept |
| function declaration | it is declared either implicitly as <code>func( ... ) { ...}</code>, or as a named function: <code> func func_name( ... ) (string, error) { ... } </code> |
| derieved data type | <code> <em>type</em> derivedType <em>struct</em> { ... }</code>|
| interface | <code> <em>type</em> interfaceType <em>interface</em> { func1( ) }</code>; derived data can implement functions declared in an interface |
| collection | array: []fooType (size fixed); slice: []fooType (size not fixed); map[fooType] barType;  One may view slice as a special case of map: map[int] fooType
| modulation | data and function of other packages can be included using <code>import ( pkgname ... )</code>|
| coroutine | by using <em>go</em>, a separate thread is launched; e.g: <code> go funcName()</code> |
| | data communication between parent and child thread uses <em>chan</em> variable: <code>chanVar := make(chan string)</code> with data assignment defined by channel operator, '<-': <code>chanVar <- "Hello, World"</code> |
| _defer_ | statements can have <em>defer</em> before the statements; the deferred statements are guranteed to be executed, if even _panic()_, is called, on the FILO order, like a stack. |
| function pointer, closure | The language supports functional pointer as well as closure |
| OOP discussion | It implements data abstraction and encapsulation. However, it does not support method overload nor extend struct and/or interface. |

## Project Structure

| file name | features |
| --- | ---- |
| note.go | _init_ and _main_ functions are defined |
| | illustration of defining a data type _config_, an interface _cmdLine_ |
|         | _config_ implements _cmdLine_ interface and _flag_ package is used to parse command line options; <note>Note</note> default flag.Usage() is over-written |
| server.go | uses _http_ package to create a web server, which handles '/', '/new', and '/(view|edit|save)/\<title\>' calls, _http/template_ package to display web pages, and _regexp_ package is used in validating and extracting substring |
|         | also demonstrates with using functional pointer and closure  |
| persist.go | uses _io/ioutil_ and _path/filepath_ packages to perform directory walk and file IO; _os_ package is used to create data directory |

## Development Tools

- VC Code is the IDE used in developing this project

