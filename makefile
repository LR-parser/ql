# Makefile for a go project
# 	
# Targets:
# 	all: Builds the code
# 	build: Builds the code
# 	fmt: Formats the source files
# 	clean: cleans the code
# 	install: Installs the code to the GOPATH
# 	iref: Installs referenced projects
#	test: Runs the tests
#

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test -o ql.test
GODEP=$(GOTEST) -i
GOFMT=gofmt -s -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")

# Package lists
TOPLEVEL_PKG := ql
INT_LIST := interfaces	#<-- Interface directories
IMPL_LIST := . symbols gui ast typechecker ast/expr ast/stmt ast/vari #<-- Implementation directories
CMD_LIST := 	#<-- Command directories
MAIN_LIST := main.go

# List building
ALL_LIST = $(INT_LIST) $(IMPL_LIST) $(CMD_LIST)

BUILD_LIST = $(foreach int, $(ALL_LIST), $(int)_build)
CLEAN_LIST = $(foreach int, $(ALL_LIST), $(int)_clean)
INSTALL_LIST = $(foreach int, $(ALL_LIST), $(int)_install)
IREF_LIST = $(foreach int, $(ALL_LIST), $(int)_iref)
TEST_LIST = $(foreach int, $(ALL_LIST), $(int)_test)
FMT_TEST = $(foreach int, $(ALL_LIST), $(int)_fmt)
RUN_LIST = $(foreach int, $(MAIN_LIST), $(int)_run)

# All are .PHONY for now because dependencyness is hard
.PHONY: $(CLEAN_LIST) $(TEST_LIST) $(FMT_LIST) $(INSTALL_LIST) $(BUILD_LIST) $(IREF_LIST) $(RUN_LIST)

all: build
build: $(BUILD_LIST)
clean: $(CLEAN_LIST)
install: $(INSTALL_LIST)
test: $(TEST_LIST)
iref: $(IREF_LIST)
fmt: $(FMT_TEST)
run: $(RUN_LIST)
doc:
	godoc -http=:8080 &
	open 'http://localhost:8080/pkg/ql/'

$(BUILD_LIST): %_build: %_fmt %_iref
	$(GOBUILD) $(TOPLEVEL_PKG)/$*
$(RUN_LIST): %_run:
	$(GOBUILD) main.go && ./main
$(CLEAN_LIST): %_clean:
	$(GOCLEAN) $(TOPLEVEL_PKG)/$*
$(INSTALL_LIST): %_install:
	$(GOINSTALL) $(TOPLEVEL_PKG)/$*
$(IREF_LIST): %_iref:
	$(GODEP) $(TOPLEVEL_PKG)/$*
$(TEST_LIST): %_test:
	$(GOTEST) $(TOPLEVEL_PKG)/$*
$(FMT_TEST): %_fmt:
	$(GOFMT)


