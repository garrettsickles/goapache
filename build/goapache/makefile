# Go parameters
GOCMD   = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST  = $(GOCMD) test
GOGET   = $(GOCMD) get

# Local Directory
ROOT_DIR = ../..

# goapache
DIR_GOAPACHE = $(ROOT_DIR)
SRC_GOAPACHE =                        \
	$(DIR_GOAPACHE)/apache_read.go    \
	$(DIR_GOAPACHE)/apache_request.go \

# You must also run this command to allow all linker flags...
#   $ export CGO_LDFLAGS_ALLOW=".*"
build:
	$(GOBUILD) -x $(SRC_GOAPACHE)