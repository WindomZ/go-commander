#!/usr/bin/env bash
go install

quick_example tcp 127.0.0.1 1080
quick_example serial 80 --timeout=120

#quick_example --version
