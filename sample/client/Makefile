# variable
binaryName=client
outputPath=_output
ldLibraryPath=.:${LD_LIBRARY_PATH}:/usr/local/lib

build:
	CGO_CXXFLAGS="-g -O2 -D_GLIBCXX_USE_CXX11_ABI=" \
	CGO_LDFLAGS="-lenet" \
	go build -o ${outputPath}/${binaryName}

run: build
	LD_LIBRARY_PATH=${ldLibraryPath} ${outputPath}/${binaryName}

clean:
	rm -rf _output
