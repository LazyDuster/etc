#include "stdlib.h"
#include "stdio.h"

char a[];

int main()
{
     printf("%s%s;", a, a);
}

char a[] = "#include \"stdlib.h\"\n#include \"stdio.h\"\n\nchar a[];\n\nint main()\n{\n\tprintf(\"%s%s;\", a, a);\n}\n\n char a[] = \"";
