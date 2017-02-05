#include "../strings.h"
#include <stdio.h>

int main() {
  printf("This is a C Application.\n");

  printf("%s\n", C_Hello(0));
  printf("%s\n", C_Hello(1));

  return 0;
}
