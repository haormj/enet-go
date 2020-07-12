build:
	swig -go -intgosize 64 -DENET_CALLBACK= -DENET_API= -I/usr/local/include enet.i

clean:
	rm -rf enet.go enet_wrap.c
