CXX = g++

debug:
	cp -r res/** bin/debug/res
	$(CXX) -c src/*.cpp -std=c++14 -m64 -g -Wall -I include
	$(CXX) *.o -o bin/debug/physics2d -lSDL2main -lSDL2 -lSDL2_image
	./bin/debug/physics2d

release:
	cp -r res/** bin/release/res
	$(CXX) -c src/*.cpp -std=c++14 -m64 -O3 -Wall -I include
	$(CXX) *.o -o bin/release/physics2d -s -lSDL2main -lSDL2 -lSDL2_image
	./bin/release/physics2d
