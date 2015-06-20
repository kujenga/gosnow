#include <stdio.h>
#include <stdlib.h>
#include "cdrafter.h"

int main(int argc, char const *argv[])
{
	const char* source = "# My API\n## GET /message\n + Response 200 (text/plain)\n\n        Hello World\n";
	char *result = NULL;
	int ret = drafter_c_parse(source, 0, &result);

	printf("Result: %s\n", ret == 0 ? "OK" : "ERROR");
	printf("Serialized JSON result:\n%s\n", result);

	free(result); /* we MUST release allocted memory for result */
	return 0;
}
