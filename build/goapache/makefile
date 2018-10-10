# Go parameters
GOCMD   = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST  = $(GOCMD) test
GOGET   = $(GOCMD) get

# Local Directory
ROOT_DIR = ../..

# goapache
LDFLAGS_GOAPACHE = "-Wl,-z,relro,-z,now"
DIR_GOAPACHE = $(ROOT_DIR)
SRC_GOAPACHE =              \
	$(DIR_GOAPACHE)/read.go \

# You must also run this command to allow all linker flags...
#   $ export CGO_LDFLAGS_ALLOW=".*"
build:
	$(GOBUILD) -x $(SRC_GOAPACHE)