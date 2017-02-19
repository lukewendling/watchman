package main

import (
	"fmt"
)

// Between return loopback querystring for between queries
func Between(field, lower, upper string) string {
    return fmt.Sprintf("?filter[where][%s][between][0]=%s&filter[where][%s][between][1]=%s", field, lower, field, upper)
}

// Eq returns loopback querystring for equality queries
func Eq(field, val string) string {
    return fmt.Sprintf("?filter[where][%s]=%s", field, val)
}