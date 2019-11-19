package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>

struct Message {
  const char *data;
  int id;
};

extern void onMessageGo(struct Message *msg);

@interface GoMenu : NSObject
	-(void) addItemWithTitle:(NSString *)title key:(NSString *)key;
  -(void) sendMessage:(id)sender;
  +(void) recvMessage:(struct Message *)msg;
@end

void onMessageObjc(struct Message *msg) {
  [GoMenu recvMessage: msg];
}

@implementation GoMenu {
		NSStatusItem *statusItem;
	}

	-(instancetype) initWithTitle:(NSString *)title {
		self = [super init];
		if (self) {
			id statusBar = [NSStatusBar systemStatusBar];
			statusItem = [statusBar statusItemWithLength:NSVariableStatusItemLength];
			statusItem.button.title = title;
		}
		return self;
	}
	
	-(void) addItemWithTitle:(NSString *)title key:(NSString *)key {
    if (!statusItem.menu) {
      [statusItem setMenu: [[NSMenu new] autorelease]];
    }
    id statusMenuItem = [[[NSMenuItem alloc] initWithTitle:title action:@selector(sendMessage:) keyEquivalent:key] autorelease];
    [statusMenuItem setTarget: self];
		[statusItem.menu addItem: statusMenuItem];
	}

	-(void) sendMessage:(id)sender {
    NSString *str = @"world";
		NSLog(@"objc sendMessage %@ via %@", str, sender);
		struct Message msg;
		msg.id = 100;
		msg.data = str.UTF8String;
		onMessageGo(&msg);
  }
  
  +(void) recvMessage:(struct Message *)msg {
    NSLog(@"objc revcMessage %@", [NSString stringWithUTF8String: msg->data]);
  }
@end

int StartApp(void) {
	[NSAutoreleasePool new];
	[NSApplication sharedApplication];
	GoMenu *menu = [[[GoMenu alloc] initWithTitle:@"Demo"] autorelease];
	[menu addItemWithTitle:@"Demo Item" key:@"d"];
  [NSApp run];
  return 0;
}
*/
import "C"

import (
  "fmt"
	"time"
	"unsafe"
)

func sendMessage(str string, id int) {
  fmt.Println("golang sendMessage", str)
  time.Sleep(1 * time.Second)
  
	strC := C.CString(str)
  defer C.free(unsafe.Pointer(strC))

  idC := C.int(id)

	msg := C.struct_Message{
		data: strC,
		id:   idC,
	}

	// ptr := C.malloc(C.sizeof_char * 1024)
	// defer C.free(unsafe.Pointer(ptr))
	// C.update(&msg, (*C.char)(ptr))

	C.onMessageObjc(&msg)
}

func revcMessage(msg *C.struct_Message) {
  str := C.GoString(msg.data)
  fmt.Println("golang revcMessage", str)
}

func main() {
  fmt.Println("GoMenu Demo")
	go sendMessage("hello", -100)
	C.StartApp()
}
