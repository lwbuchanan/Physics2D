# 2D Physics Engine (go edition)

I am in the process of reworking this project in go w/ raylib.

Raylib has built in physics which I will not be using, but as a rendering engine it is far more convenient than SDL2.

C++ is fun, but I am prioritizing learning the theory behind physics simulation here, not trying to become an expert in memory safe programming. Having a garbage collector is too good to pass up unless I end up trying to optomize it down the line.

The physics math is contained in a submodule that doesn't depend on raylib. I probobly wont ever use this physics engine outside of this project, but its a helpful logical seperation between rendering/physics.

### Some helpful resources I've used
- <https://code.tutsplus.com/how-to-create-a-custom-2d-physics-engine-the-basics-and-impulse-resolution--gamedev-6331t>
- <https://www.youtube.com/watch?v=emfGoBgE020&list=PLSlpr6o9vURwq3oxVZSimY8iC-cdd3kIs&index=4>
