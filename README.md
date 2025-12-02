This is my attempt at the WaTor simulation for a college project.
Overall I didn't give the project the amount of time I should have and that will be seem in the lack of working functionalities.

I'm pleased with how my movement system turned out, I didn't handle the wrap around required for the creatures, to avoid memory errors they are allowed to -1 of the width and height specified where they will get stuck & not update at all.

The sharks do eat the fish as intended & as a result live longer, there is no breeding system implemented.

The creature struct idea & the PrintGrid() was inspired by the lazyhacker implementation of the project which falls under the MIT license
Link: https://github.com/lazyhacker/wator

To improve speedup, the intent was to unwrap the  many for loops where possible and attempt to implement a stencil pattern.

Link to the my WaTor github: https://github.com/Ryan-Dunne/ProjectWaTor