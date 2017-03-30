#!/usr/bin/env bash
go install

naval_fate_example ship new aaa bbb ccc
naval_fate_example ship hahaha move xxx yyy --speed=120
naval_fate_example ship shoot xxx yyy
naval_fate_example mine set xxx yyy --moored
naval_fate_example mine remove xxx yyy --drifting

#calculator_example --version
