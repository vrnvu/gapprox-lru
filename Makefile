.PHONY = help fill run

.DEFAULT_GOAL = help

help:
	@echo "---------------HELP-----------------"
	@echo "fill: fill a /tmp folder with X bytes of data"
	@echo "run: runs one iteration of the approximated lru"
	@echo "------------------------------------"

fill:
	go run cmd/fill/main.go

run:
	go run cmd/run/main.go

clean:
	rm -rf ./tmp
