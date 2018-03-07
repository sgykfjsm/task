BIN_DIR = bin
APP_NAME = task
LD_FLAGS = -w -s

init:
	go get -u github.com/golang/dep/cmd/dep
	dep init

deps:
	dep ensure
	@if [ "$(ARGS)" = "update" ]; then \
		dep ensure --update; \
	fi

build:
	mkdir -pv $(BIN_DIR)
	go build -ldflags '$(LD_FLAGS)' -o $(BIN_DIR)/$(APP_NAME) .

clean:
	$(RM) -rv task $(BIN_DIR)
