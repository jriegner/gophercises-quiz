package main

import "flag"

type configuration struct {
	csv string
}

func(c *configuration) init(){
	flag.StringVar(&c.csv, "csv", "problems.csv", "path to problems csv.")
}

func(c *configuration) parse(){
	if !flag.Parsed(){
		flag.Parse()
	}
}
