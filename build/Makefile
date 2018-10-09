# Go parameters
GOCMD   = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST  = $(GOCMD) test
GOGET   = $(GOCMD) get

# Local Directory
ROOT_DIR = ..
CMD_DIR  = $(ROOT_DIR)/cmd

# mod_example
MOD_EXAMPLE_NAME   = mod_goapache_example
TGT_MOD_EXAMPLE_SO = $(MOD_EXAMPLE_NAME).so
DIR_MOD_EXAMPLE    = $(CMD_DIR)/$(MOD_EXAMPLE_NAME)
SRC_MOD_EXAMPLE =                    \
	$(DIR_MOD_EXAMPLE)/configure.go  \
	$(DIR_MOD_EXAMPLE)/directives.go \
	$(DIR_MOD_EXAMPLE)/handler.go    \
	$(DIR_MOD_EXAMPLE)/hooks.go      \
	$(DIR_MOD_EXAMPLE)/initialize.go \
	$(DIR_MOD_EXAMPLE)/module.go

# You must also run this command to allow all linker flags...
#   $ export CGO_LDFLAGS_ALLOW=".*"
mod_goapache_example:
	export CGO_LDFLAGS_ALLOW=".*" && \
	$(GOBUILD) -x -o $(TGT_MOD_EXAMPLE_SO) -buildmode=c-shared $(SRC_MOD_EXAMPLE)