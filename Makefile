GOCMD = go
namespace = github.com/mislav/everyenv/config

PROGRAM_VERSION ?= 0.0.1-g$(shell git rev-parse --short=7 HEAD)

PROGRAM_NAME ?= rbenv
PROGRAM_EXECUTABLE ?= ruby
PROGRAM_FILENAME ?= .$(PROGRAM_EXECUTABLE)-version
PROGRAM_ROOT ?= $$HOME/.$(PROGRAM_NAME)
PROGRAM_ROOT_NAME ?= $(shell tr 'a-z' 'A-Z' <<< $(PROGRAM_NAME))_ROOT
PROGRAM_DIR_NAME ?= $(shell tr 'a-z' 'A-Z' <<< $(PROGRAM_NAME))_DIR

.PHONY: format

all: $(PROGRAM_NAME)

$(PROGRAM_NAME):
	$(GOCMD) build -ldflags ' \
		-X $(namespace).Root $(PROGRAM_ROOT) \
		-X $(namespace).RootEnvName $(PROGRAM_ROOT_NAME) \
		-X $(namespace).VersionFilename $(PROGRAM_FILENAME) \
		-X $(namespace).DirEnvName $(PROGRAM_DIR_NAME) \
		-X $(namespace).MainExecutable $(PROGRAM_EXECUTABLE) \
		-X $(namespace).BuildVersion $(PROGRAM_VERSION) \
		' -o $(PROGRAM_NAME) ./everyenv.go

format:
	$(GOCMD) fmt ./...
