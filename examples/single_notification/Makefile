##
## Makefile for go-notify
## by lenorm_f
##

include $(GOROOT)/src/Make.inc

TARG = example_go

GOFILES =\
	 main.go

CLEANFILES += example_c

include $(GOROOT)/src/Make.cmd

example_c: main.c
	gcc -Wall -Wextra `pkg-config --cflags --libs libnotify` -o example_c main.c
