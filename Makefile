include $(GOROOT)/src/Make.inc

GOBIN=$(PWD)/bin

DIRS=\
	src/lib\
	src/cmd/cat\
	src/cmd/wc\


all.dirs: $(addsuffix .all, $(DIRS))
clean.dirs: $(addsuffix .clean, $(DIRS))
install.dirs: $(addsuffix .install, $(DIRS))
nuke.dirs: $(addsuffix .nuke, $(DIRS))

%.clean:
	+cd $* && gomake clean

%.install:
	test -d $(GOBIN) || mkdir $(GOBIN)
	+cd $* && gomake install

%.nuke:
	+cd $* && gomake nuke
%.all:
	+cd $* && gomake all

clean: clean.dirs

install: install.dirs

all: all.dirs

nuke: nuke.dirs
	rm -rf $(QUOTED_GOBIN)
