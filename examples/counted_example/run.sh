#!/usr/bin/env bash
go install

counted_example -vvvvvvvvvv
counted_example go go
counted_example --path ./here --path ./there
counted_example this.txt that.txt

#counted_example --version
