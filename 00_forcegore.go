package main

// Try and force our init function to run before any of the eggos
// code, in the hopes that this allows us to set a log level nice
// and early.
//
// This is pretty gory. Thus: forcegore
import _ "github.com/jspc/primes/forcegore"
