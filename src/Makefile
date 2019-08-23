BINARIES = forftanks upload.cgi

CFLAGS = -Wall

all: $(BINARIES)

install: all
	install -d $(DESTDIR)/usr/bin
	install forftanks $(DESTDIR)/usr/bin

	install -d $(DESTDIR)/usr/lib/tanks
	cp -r examples $(DESTDIR)/usr/lib/tanks/examples

forftanks: forftanks.o ctanks.o forf.o
forftanks: LDLIBS = -lm

forftanks.o: forf.h ctanks.h
forf.o: forf.c forf.h
ctanks.o: ctanks.h

%.html: %.html.m4 nav.html.inc
	m4 $< > $@

clean:
	rm -f *.o next-round round-*.html
	rm -f $(BINARIES)
