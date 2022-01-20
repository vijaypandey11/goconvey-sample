# goconvey-test-example
goconvey testing example for mathematical operators

**Get Started:**
To get started with goconvey, we need to install goconvey first. Follow below steps to get going with goconvey


Installation
------------

	$ go get github.com/smartystreets/goconvey
  
 
[Quick start]
-----------

Create a test go file, for example:

```go
package testing

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

//
func TestAdd(t *testing.T) {

	Convey("The result of 5 + 11 = 16 ?", t, func() {
		x, y := 5, 11
		So(Add(x, y), ShouldEqual, 16)
	})
}
```

Get started with goconvey with your code
----------------------------------------

Start up the GoConvey web server at your project's path:

	$ $GOPATH/bin/goconvey

Make sure the bin folder has the goconvey exe within and your code has go module, go sum and  _*_test.go_ files etc

This process would look like below in git bash

![image](https://user-images.githubusercontent.com/45930971/150318836-a54b1149-9e0a-443a-b931-55552cfbc141.png)


#### [In the browser]

Check the test results displayed in your browser at:

	http://localhost:8080

If the browser doesn't open automatically, please click [http://localhost:8080](http://localhost:8080) to open manually.

The result and UI would look like:

![image](https://user-images.githubusercontent.com/45930971/150319296-974cd70e-8b3c-4e2e-a566-cbc7b820482a.png)



