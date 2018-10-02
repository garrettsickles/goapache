# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Local Directory
ROOT-DIR=..
CMD-DIR=$(ROOT-DIR)/cmd

# mod-golang-example
TGT_MOD-GOLANG-EXAMPLE=libmod-golang-example.so
DIR_MOD-GOLANG-EXAMPLE=$(CMD-DIR)/mod-golang-example
SRC_MOD-GOLANG-EXAMPLE= \
	$(DIR_MOD-GOLANG-EXAMPLE)/configure.go \
	$(DIR_MOD-GOLANG-EXAMPLE)/hooks.go \
	$(DIR_MOD-GOLANG-EXAMPLE)/module.go \
	$(DIR_MOD-GOLANG-EXAMPLE)/handler.go \
	$(DIR_MOD-GOLANG-EXAMPLE)/initialize.go \

mod-golang-example:
	$(GOBUILD) -v -x -o $(TGT_MOD-GOLANG-EXAMPLE) -buildmode=c-shared $(SRC_MOD-GOLANG-EXAMPLE)