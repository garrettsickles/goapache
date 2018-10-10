# Go parameters
GOCMD   = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST  = $(GOCMD) test
GOGET   = $(GOCMD) get

# Local Directory
ROOT_DIR = ../..

# libgoapache
DIR_GOAPACHE = $(ROOT_DIR)
SRC_GOAPACHE =              \
	$(DIR_GOAPACHE)/read.go \

goapache:
	$(GOBUILD) -x $(SRC_GOAPACHE)