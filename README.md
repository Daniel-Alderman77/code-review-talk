# code-review-talk

Two examples of Go code with deliberate faults added to them I have created for a tech talk.

Server, a simple HTTP server that writes HelloWorld on the /greeting endpoint:
* Default server mux, has no timeouts
* Reads the request body and doesn’t do anything with it
* Panics when it doesn’t need to, should handle the error gracefully and return like bellow
* Bytes shadows bytes package name 

Triangles, a function with a unit test that calculates whether a triangle is Scalene, Equilateral or Isosceles:
* There is no new line character, so it will print on a singular line
* a, b, c are hardcoded, the answer will always be the same
* The triangle types could be hardcoded as consts
* There is no test case for a Scalene triangle

