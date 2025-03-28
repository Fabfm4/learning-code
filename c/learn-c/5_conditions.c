// https://www.learn-c.org/en/Conditions

// stdio is a library that provides input and output functions
#include <stdio.h>

void guessNumber(int guess) {
    if ( guess == 555) {
        printf("Correct. You guessed it!\n");
    } else if ( guess > 555) {
        printf("Your guess is too high.\n");
    } else {
        printf("Your guess is too low.\n");
    }
}

// main is the entry point of the program
int main() {
    guessNumber(500);
    guessNumber(600);
    guessNumber(555);
}