#!/usr/bin/env bash
go install

quick_example tcp 127.0.0.1 1080 --timeout=110
quick_example serial 80 --baud=5800 --timeout=120

quick_example --version
