# Rotations & Fractals

## Exercises

1. Implement a program that draws a 3D cube that is being rotated through an arbitrary axis. The cube must be projected at the axis $z$.
   Use formulas for rotating a point by a given angle.
   
2. Implement a program that draws a [Sierpiński triangle](https://en.wikipedia.org/wiki/Sierpi%C5%84ski_triangle) fractal using
   the [Chaos game](https://en.wikipedia.org/wiki/Sierpi%C5%84ski_triangle#Chaos_game) method:
   
   1. Take three points in a plane to form a triangle.
   2. Randomly select any point inside the triangle and consider that your current position.
   3. Randomly select any one of the three vertex points.
   4. Move half the distance from your current position to the selected vertex.
   5. Plot the current position.
   6. Repeat from step 3.
