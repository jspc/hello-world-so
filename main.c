#include "strings.h"
#include <stdio.h>

int main() {
  GoString helloTrue = Hello(1);
  GoString helloFalse = Hello(0);

  printf("This is a C Application.\n");

  printf("%.*s\n", (int)helloTrue.n, helloTrue.p);
  printf("%.*s\n", (int)helloFalse.n, helloFalse.p);

  return 0;
}
