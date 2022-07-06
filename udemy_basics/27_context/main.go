package main

import (
	"context"
	"fmt"
)

// context is a tool used with gorutine patterns to avoid leaking gorutines.
// leaking go rutine is where we have a process that lunches multiple goruttines and we shit the process down. The go rutines will still run using the resources
// context can help us kill gorutines after the process ends
// context is more of a advance topic

// often requests on go serwers are handled in its own go rutine each and they should be canceled in some situation (i.t timneout). We can also use context for this, canceling a gorutine after some time
func main() {
	parentContext := context.Background()                                // creating context
	derivedContext, contextCancelFn := context.WithCancel(parentContext) // create derived context and cancel function for this context

	fmt.Println("context: \t", derivedContext)
	fmt.Println("context error: \t", derivedContext.Err())
	fmt.Printf("context type: %T \t", derivedContext)
	fmt.Println("contextcancel function: \t", contextCancelFn)

	fmt.Println("----------------")

	contextCancelFn()

	fmt.Println("context: \t", derivedContext)
	fmt.Println("context error: \t", derivedContext.Err())
	fmt.Printf("context type: %T \t", derivedContext)
	fmt.Println("contextcancel function: \t", contextCancelFn)
}
