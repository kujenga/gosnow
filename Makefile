
all:
	cd test && gcc -Wall -c ctest.c -I../drafter/src/ -I../drafter/ext/snowcrash/src/
	gcc ./test/ctest.o -L./drafter/build/out/Release/ -ldrafter -o ./test/bin/ctest

install:
	cd drafter && ./configure --shared && make
	ln -s $(CURDIR)/drafter/build/out/Release/libdrafter.dylib /usr/local/lib/libdrafter.dylib
	mkdir -p ./test/bin

travis:
	cd drafter && ./configure --shared && make
	mkdir -p ./test/bin

clean:
	rm /usr/local/lib/libdrafter.dylib
	rm ./test/ctest.o
	rm ./test/bin/*

test: all
	./test/bin/ctest
