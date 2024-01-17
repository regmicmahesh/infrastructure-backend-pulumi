SHELL := /bin/bash

export APP_ROOT := $(shell pwd)

-include $(APP_ROOT)/env/override.mk

include $(APP_ROOT)/base/Makefile

include $(APP_ROOT)/targets/pulumi.mk
include $(APP_ROOT)/targets/docker.mk
include $(APP_ROOT)/targets/go.mk



