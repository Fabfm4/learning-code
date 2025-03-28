// https://www.learn-c.org/en/Pointers
// I need more context to understand the code below
// I will try to run it and see what happens


// Copilot, can you explain what is happening in this code?
// Sure! This code declares an integer variable n and initializes it with the value 10.
// It then declares a pointer variable pointer_to_n and assigns the address of n to it.
// The code then changes the value of n by dereferencing the pointer and assigning the value 11 to it.
// Finally, it checks if the pointer is pointing to the correct address and if the value of n has been updated correctly.
// If everything is correct, it prints "Done!" to the console.


#include <stdio.h>

int main() {
  int n = 10;

  int * pointer_to_n = &n;
  *pointer_to_n = 11;


  if (pointer_to_n != &n) return 1;
  if (*pointer_to_n != 11) return 1;

  printf("Done!\n");
  return 0;
}