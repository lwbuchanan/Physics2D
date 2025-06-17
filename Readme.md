# 2D Physics Engine (go edition)

I am in the process of reworking this project in go w/ raylib.

Raylib has built in physics which I will not be using, but as a rendering engine it is far more convenient than SDL2.

C++ is fun, but I am prioritizing learning the theory behind physics simulation here, not trying to become an expert in memory safe programming. Having a garbage collector is too good to pass up unless I end up trying to optomize it down the line.

The physics math is contained in a submodule that doesn't depend on raylib. I probobly wont ever use this physics engine outside of this project, but its a helpful logical seperation between rendering/physics.

### Some helpful resources I've used
- <https://code.tutsplus.com/how-to-create-a-custom-2d-physics-engine-the-basics-and-impulse-resolution--gamedev-6331t>
- <https://www.chrishecker.com/Rigid_Body_Dynamics>

This guy's explanations are extreamly helpful. A lot of my code was initally translated from the c# in the videos, but I've been working on trying to optomize and refactor it using my own understanding. If you are interested in building an engine yourself, I'd highly recommend you start by following along with his videos.
- <https://www.youtube.com/watch?v=emfGoBgE020&list=PLSlpr6o9vURwq3oxVZSimY8iC-cdd3kIs&index=4>

### Learning Milestones
This section is just a list of things this project has taught me about physics, algorithms, and programming in general. It will be helpful when showing people (job interviewers) how I've been pursuing self-directed learning.

- OO is not the solution to everything
    - Starting this project in C++ was somewhat a misplay. It's not a bad language, but I was unfamiliar with it and dealing with a challenging project in an unfamiliar environment is a recipe for an abandoned project. I switched to go because garbage collection is super convienent if you care more about getting the logic correct than optomizing memory usage. I was worried that it would be almost impossible to write polymorphic code without a solid objected oriented type system, but the more I've used go, the more I value its simplicity. There is definatly a lot more type checking code, but things like collision detection are prone to the [expression problem](https://en.wikipedia.org/wiki/Expression_problem), so OOP doesn't neccessarily make life easier there. In fact, not having an easy way to make abstract types highly encourages using more abstraction in the procedural logic itself, leading to general purpose functionalty with just the occasional ```switch(value.valueType)```.

- The hardest part of physics simulation is not physics
    - Figuring out how to make objects bounce off each other is simple. With a bit of knowledge about dynamics, the equations to calculate the final velocity of an object after a collision is pretty straightforward. The hard part is geometry. In nature, we have the electromagnetic force that conveniently figures out when things are "touching" and push on them. However, unless we want to develop an elemetary particle simulator, we have to rely on shape intersection. This is deceptivly tricky since real objects generally don't occupy the same poition in space-time. When we look at some arbitrary situation like two polygons intersecting along multiple axes, we have to make some assumptions about how we got into that situation and how to react to this abnormality in such a way that convincingly looks like the two objects simply bounced without intersection. We also have to apply an instantaneous impulse instead of the more accurate response which would be to apply a force over the duration of contact. Determining all these details is an exercise in geometry rather than in physics.

- Simple observations can lead to powerful conclusions
    - The separating axis theorem (SAT) is the core of the collision detection algorithm. This theorem can efficiently detect when two shapes have intersected by making note of the fact that when two objects _havn't_ intersected, you can draw a line between them (assuming they aren't concave). That alone is quite obvious, but if we start trying to use this theory in practice, we realize that there are actually only a few lines we have to try before we know that no such line could exist. If two objects are not intersecting, there must be some separting line that is parallel to one of the lines (or tangent to one of the curves). If this weren't the case, then the seperaing line would have to intersect one of the polygons which would mean that the line wasn't, in fact, separating. Moreover, we can check each line by checking the extent of the shapes on the line tangent the tested edge. This means we can check every parallel edge at once. Testing the collision of two parallelograms only requires us to test 4 edges. Convienienly, when we detect an intersection, we also know the normal of the collision since it is parallel to the seperating axis. This algorithm covers all cases in a miminal number of checks and is underpinned by simple understanding of what it means for two shapes to intersect.