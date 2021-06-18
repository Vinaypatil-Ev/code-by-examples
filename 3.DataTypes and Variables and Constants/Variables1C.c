#include <stdio.h>
#include <stdbool.h>

int main() {
    // Varibles
    // declare the variable as:
    // datatype variable_name = value;
    char str[] = "This is character string in C";
    int i = 1000;
    float f = 10000.77;
    double d = 1010.04443;
    bool b = true;

    // Constant
    // declare the constant using keyword const:
    // const datatype const_name = value;
    const char strC[] = "This is character constant string in C";
    const int ic = 1000;
    const float fc = 10000.77;
    const double dc = 1010.04443;
    const bool bb = false;

    // datatype
    char string[] = "This character datatype"; // 1byte
    int number = 1000;                         // 2 byte
    float div = 23.33;                         // 4 byte
    double value = 345.94949;                  // 8 byte
    // datatype modifires
    short int intShort = 10;                   // 2 byte
    long int intLong = 2147483647;             // 4 byte
    long double dLong = 335.784884;            // 10 byte  
}