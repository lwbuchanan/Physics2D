# 2D Physics Engine (go edition)

I am in the process of reworking this project in go w/ raylib.

Raylib has built in physics which I will not be using, but as a rendering engine it is far more convenient than SDL2.

C++ is fun, but I am prioritizing learning the theory behind physics simulation here, not trying to become an expert in memory safe programming. Having a garbage collector is too good to pass up unless I end up trying to optomize it down the line.

The physics math is contained in a submodule that doesn't depend on raylib. I probobly wont ever use this physics engine outside of this project, but its a helpful logical seperation between rendering/physics.

### Some helpful resources I've used
- <https://code.tutsplus.com/how-to-create-a-custom-2d-physics-engine-the-basics-and-impulse-resolution--gamedev-6331t>

This guy's explanations are extreamly helpful. A lot of my code was initally translated from the c# in the videos, but I've been working on trying to optomize and refactor it using my own understanding. If you are interested in building an engine yourself, I'd highly recommend you start by following along with his videos.
- <https://www.youtube.com/watch?v=emfGoBgE020&list=PLSlpr6o9vURwq3oxVZSimY8iC-cdd3kIs&index=4>

### Learning Milestones
This section is just a list of things this project has taught me about physics, algorithms, and programming in general. It will be helpful when showing people (job interviewers) how I've been pursuing self-directed learning.

- OO is not the solution to everything
    - Starting this project in C++ was somewhat a misplay. It's not a bad language, but I was unfamiliar with it and dealing with a challenging project in an unfamiliar environment is a recipe for an abandoned project. I switched to go because garbage collection is super convienent if you care more about getting the logic correct than optomizing memory usage. I was worried that it would be almost impossible to write polymorphic code without a solid objected oriented type system, but the more I've used go, the more I value its simplicity. There is definatly a lot more type checking code, but things like collision detection are prone to the [expression problem](https://en.wikipedia.org/wiki/Expression_problem), so OOP doesn't neccessarily make life easier there. In fact, not having an easy way to make abstract types highly encourages using more abstraction in the procedural logic itself, leading to general purpose functionalty with just the occasional ```switch(value.valueType)```.