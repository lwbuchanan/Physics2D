CXX = g++

debug:
	cp -r res/** bin/debug/res
	$(CXX) -c src/*.cpp -std=c++14 -m64 -g -Wall -I include
	$(CXX) *.o -o bin/debug/main -lSDL2main -lSDL2 -lSDL2_image
	./bin/debug/main

release:
	cp -r res/** bin/release/res
	$(CXX) -c src/*.cpp -std=c++14 -m64 -O3 -Wall -I include
	$(CXX) *.o -o bin/release/main -s -lSDL2main -lSDL2 -lSDL2_image
	./bin/release/main
