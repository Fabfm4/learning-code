// https://www.learn-c.org/en/Arrays

// stdio is a library that provides input and output functions
#include <stdio.h>

// main is the entry point of the program
int main() {
  int grades[3];

  int average;

  grades[0] = 80;
  grades[1] = 85;
  grades[2] = 90;

  average = (grades[0] + grades[1] + grades[2]) / 3;
  printf("The average of the 3 grades is: %d", average);

  return 0;
}