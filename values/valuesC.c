#include<stdio.h>
#include<stdbool.h>

int main() {
    printf("%s\n", "This is string");
    printf("%d\n", 1101);
    printf("%f\n", 1101.01);
    printf("%0.2f\n", 1101.088);
    printf("%s\n", true ? "true": "false"); // need stdbool.h
    printf("%s\n", false ? "true": "false"); // need stdbool.h
}