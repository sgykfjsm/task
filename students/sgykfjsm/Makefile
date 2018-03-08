BIN_DIR = bin
APP_NAME = task
LD_FLAGS = -w -s

BATS_CMD := $(PWD)/bats/libexec/bats

init:
	go get -u github.com/golang/dep/cmd/dep
	dep init

deps:
	dep ensure
	@if [ "$(ARGS)" = "update" ]; then \
		dep ensure --update; \
	fi

test: build
	$(RM) -r -v task.db
	@if [ ! -d "bats" ]; then \
		git clone https://github.com/sstephenson/bats.git; \
	fi
	$(BATS_CMD) main.bats

build:
	mkdir -pv $(BIN_DIR)
	go build -ldflags '$(LD_FLAGS)' -o $(BIN_DIR)/$(APP_NAME) .

clean:
	$(RM) -r task.db $(BIN_DIR) bats
