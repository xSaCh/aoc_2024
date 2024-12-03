include .env
export $(shell sed 's/=.*//' .env)

VAL := $(shell printf "%02d" $(DAY))

new:
	@ mkdir day$(VAL)
	@ cp template/day_test.go day$(VAL)/day$(VAL)_test.go
	@ cat template/day.go |  sed 's/INPUTFILEPATH/.\/day$(VAL)\/input.txt/g' > day$(VAL)/day$(VAL).go
	@ curl 'https://adventofcode.com/2024/day/$(DAY)/input' -H 'Cookie: session=$(SESSION_ID)' > day$(VAL)/input.txt
