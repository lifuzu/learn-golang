## import declaration
```
  // Call the methods without explicitly stating the package - NOT RECOMMENDED
  import . "fmt"
  // Doesn’t require to use the following package `io` in importing file but init function(s) from imported package will be executed anyway (package and it dependencies will be initialized)
  import _ "io"
  // Rename the following package to `log`
  import log "github.com/sirupsen/logrus"
  // Rename the standard package `math` to `m`
  import m "math"
```
## init function

  #### init functions are defined in package block and are used for:

  variables initialization if cannot be done in initialization expression,
  checking / fixing program’s state,
  registering,
  running one-time computations,
  and many more.

  #### Package initialization
  To use a imported package it needs to be initialized first. It’s done by Golang’s runtime system and consists of (order matters):

  1. initialization of imported packages (recursive definition)
  2. computing and assigning initial values for variables declared in a package block
  3. executing init functions inside the package
  | Package initialization is done only once even if package is imported many times.

  Package in Go can contain many files; The order of variables initialization and init functions calls if they’re spread across many package’s files depends on the order of files presented to the compiler.

## golang web framework - echo: https://echo.labstack.com/



### References
[medium - golangspec](https://medium.com/golangspec)
[import declaration](https://medium.com/golangspec/import-declarations-in-go-8de0fd3ae8ff)
[3 - init function](https://medium.com/golangspec/init-functions-in-go-eac191b3860a)
[4 - initialize dependencies](https://medium.com/golangspec/initialization-dependencies-in-go-51ae7b53f24c)