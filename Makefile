
all:
	gcc -Wall -c test.c -I./drafter/src/ -I./drafter/ext/snowcrash/src/
	gcc test.o -L./drafter/build/out/Release/ -ldrafter -o test

install:
	cd drafter && ./configure --shared && make	
	ln -s $(CURDIR)/drafter/build/out/Release/libdrafter.dylib /usr/local/lib/libdrafter.dylib

clean:
	rm /usr/local/lib/libdrafter.dylib
	rm test.o test

test: all
	./test
