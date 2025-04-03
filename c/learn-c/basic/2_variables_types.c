// https://www.learn-c.org/en/Variables_and_Types

// stdio is a library that provides input and output functions
#include <stdio.h>

// main is the entry point of the program
int main() {
  int a = 3;
  float b = 4.5;
  double c = 5.25;
  // we can use operators to perform calculations
  float sum = a + b + c;
  
  // %f is used to format the output for a float
  printf("The sum of a, b, and c is %f.", sum);
  return 0;
}