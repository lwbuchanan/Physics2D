# 2D Physics Engine

This is my custom physics simulation tool which simulates rigidbody dynamics on circles and convex polygons. It supports gravity and normal forces, allowing for stacking. It uses the separating axis theorem to detect collisions and resolves them using the conservation of linear and angular momentum. Objects currently have a static restituion to allow for inelastic collisions, and I plan to eventually add support for static and dynamic forces. Objects with zero mass are unaffected by forces but act as collision obejcts. 

The simulation runs on a dynamic tick rate which standardizes physics speed regardless of frame rate. Each physics tick divides the delta-time into a fixed number of steps to more accuratly integrate the changes in velocity.

I originally started this project in c++ using sld2, but later moved to go with raylib for rendering. Raylib is incredibly easy to work with, and moving away from manual memory management allowed me to focus on understanding the math without worrying about performance and memory issues. In the future, I plan to optomize this engine using several steps of broad phase collision detections.

### Tools used
- go (language)
- raylib (for rendering)

### Some helpful resources I've used
- <https://code.tutsplus.com/how-to-create-a-custom-2d-physics-engine-the-basics-and-impulse-resolution--gamedev-6331t>
- <https://www.chrishecker.com/Rigid_Body_Dynamics>

This guy's explanations are extreamly helpful. If you are interested in building an engine yourself, I'd highly recommend you start by following along with his videos.
- <https://www.youtube.com/watch?v=emfGoBgE020&list=PLSlpr6o9vURwq3oxVZSimY8iC-cdd3kIs&index=4>
