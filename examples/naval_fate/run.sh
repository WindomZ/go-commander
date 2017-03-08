#!/usr/bin/env bash
go install

naval_fate ship new aaa bbb ccc
naval_fate ship hahaha move xxx yyy --speed=120
naval_fate ship shoot xxx yyy
naval_fate mine set xxx yyy --moored
naval_fate mine remove xxx yyy --drifting

#calculator_example --version
