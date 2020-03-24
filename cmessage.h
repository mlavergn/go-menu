#ifndef CMESSAGE
#define CMESSAGE

typedef enum COperation {
	None = 0,
	Reset = 10,
	Demo = 20,
	Quit = 99
} COperation;

typedef struct CMessage {
	COperation op;
	const char *text;
	const char *err;
	const char *dataString;
	int dataInt;
	bool dataBool;
} CMessage;

#endif