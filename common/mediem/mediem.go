/**
* @Author: TheLife
* @Date: 2021/5/10 上午11:02
 */
package mediem

import "math"

const abortIndex int8 = math.MaxInt8 / 2

// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(*Context)

// HandlersChain defines a HandlerFunc array.
type HandlersChain []HandlerFunc

type Context struct {
	Result Data

	index    int8
	handlers HandlersChain
}
type Data struct {
	Err  error
	Data interface{}
}

func (c *Context) Use(middleware ...HandlerFunc) *Context {
	c.handlers = append(c.handlers, middleware...)

	return c
}

func (c *Context) Run(middleware ...HandlerFunc) *Context {
	c.Next()

	return c
}

// Next should be used only inside middleware.
// It executes the pending handlers in the chain inside the calling handler.
// See example in GitHub.
func (c *Context) Next(middleware ...HandlerFunc) *Context {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}

	return c
}

// IsAborted returns true if the current context was aborted.
func (c *Context) IsAborted() bool {
	return c.index >= abortIndex
}

// Abort prevents pending handlers from being called. Note that this will not stop the current handler.
// Let's say you have an authorization middleware that validates that the current request is authorized.
// If the authorization fails (ex: the password does not match), call Abort to ensure the remaining handlers
// for this request are not called.
func (c *Context) Abort() {
	c.index = abortIndex
}

// Error attaches an error to the current context. The error is pushed to a list of errors.
// It's a good idea to call Error for each error that occurred during the resolution of a request.
// A middleware can be used to collect all the errors and push them to a database together,
// print a log, or append it in the HTTP response.
// Error will panic if err is nil.
func (c *Context) Error(err error) {
	c.Result.Err = err
}
