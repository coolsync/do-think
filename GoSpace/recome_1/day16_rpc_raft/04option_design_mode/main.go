package main

import "fmt"

// def Options
type Options struct {
	strOption1 string
	strOption2 string
	strOption3 string
	intOption1 int
	intOption2 int
	intOption3 int
}

// def func val, pass options field
type OptionFn func(opts *Options)

// create Options api
func initOptions1(opts ...OptionFn) {
	options := &Options{} // init struct

	// traveser all func
	for _, opt := range opts {
		opt(options) // every func, get fields specific value
	}

	fmt.Printf("options: %#v\n", options)
}

// Assign a specific value to a field method
func WithStrOption1(s string) OptionFn {
	return func(opts *Options) {
		opts.strOption1 = s
	}
}
func WithStrOption2(s string) OptionFn {
	return func(opts *Options) {
		opts.strOption2 = s
	}
}
func WithStrOption3(s string) OptionFn {
	return func(opts *Options) {
		opts.strOption3 = s
	}
}

func WithIntOption1(i int) OptionFn {
	return func(opts *Options) {
		opts.intOption1 = i
	}
}
func WithIntOption2(i int) OptionFn {
	return func(opts *Options) {
		opts.intOption2 = i
	}
}
func WithIntOption3(i int) OptionFn {
	return func(opts *Options) {
		opts.intOption3 = i
	}
}

func main() {
	initOptions1(WithStrOption1("str1"), WithStrOption2("str2"), WithIntOption1(1))
}
