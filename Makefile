include $(GOROOT)/src/Make.inc
include $(GOROOT)/src/Make.cmd

GOBIN=$(PWD)/bin

DIRS=\
	src/cmd/wc


all.dirs: $(addsuffix .all, $(DIRS))
clean.dirs: $(addsuffix .clean, $(DIRS))
install.dirs: $(addsuffix .install, $(DIRS))
nuke.dirs: $(addsuffix .nuke, $(DIRS))

%.clean:
	+cd $* && gomake clean

%.install:
	+cd $* && gomake install

%.nuke:
	+cd $* && gomake nuke
%.all:
	mkdir $(GOBIN)
	+cd $* && gomake all

clean: clean.dirs

install: install.dirs

all: all.dirs

nuke: nuke.dirs
	rm -rf $(QUOTED_GOBIN)
