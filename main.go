package main

/*
#cgo CFLAGS: -x objective-c -g -Wall
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>
#import "cmessage.h"

extern void onMessageGo(struct CMessage *msg);

@interface GoMenu : NSObject
	-(void) addItemWithTitle:(NSString *)title key:(NSString *)key tag:(COperation)tag;
	-(void) sendMessage:(id)sender;
	+(void) recvMessage:(struct CMessage *)msg;
@end

void onMessageObjc(struct CMessage *msg) {
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

	-(void) addItemWithTitle:(NSString *)title key:(NSString *)key tag:(enum COperation)tag {
		if (!statusItem.menu) {
			[statusItem setMenu: [[NSMenu new] autorelease]];
		}
		NSMenuItem* statusMenuItem = [[[NSMenuItem alloc] initWithTitle:title action:@selector(sendMessage:) keyEquivalent:key] autorelease];
		// [statusMenuItem setTarget: self];
		statusMenuItem.target = self;
		statusMenuItem.tag = tag;
		[statusItem.menu addItem: statusMenuItem];
	}

	-(void) sendMessage:(id)sender {
		NSString *str = @"world";
		NSLog(@"objc sendMessage %@ via %@", str, sender);
		struct CMessage msg;
		msg.op = ((NSMenuItem*)sender).tag;
		msg.text = str.UTF8String;
		onMessageGo(&msg);
  }

	+(void) recvMessage:(struct CMessage *)msg {
		switch(msg->op) {
		case None:
			NSLog(@"objc::recvMessage op->None");
			break;
		case Demo:
			NSLog(@"objc::recvMessage op->Demo");
			break;
		case Reset:
			NSLog(@"objc::recvMessage op->Reset");
			break;
		case Quit:
			NSLog(@"objc::recvMessage op->Quit");
			break;
		default:
			NSLog(@"objc::recvMessage unknown %@", @(msg->op));
		}
	}
@end

int StartApp(void) {
	[NSAutoreleasePool new];
	[NSApplication sharedApplication];
	GoMenu *menu = [[[GoMenu alloc] initWithTitle:@"Demo"] autorelease];
	[menu addItemWithTitle:@"Reset" key:@"r" tag:Reset];
	[menu addItemWithTitle:@"Quit" key:@"q" tag:Quit];
	[NSApp run];
	return 0;
}
*/
import "C"

import (
	"fmt"
	"os"
	"time"
	"unsafe"
)

func sendMessage(str string, op COperation) {
	fmt.Println("golang sendMessage", str)
	time.Sleep(1 * time.Second)

	msg := NewCMessage()
	msg.op = op
	msg.text = str
	cmsg := msg.Write()
	defer msg.Free()

	C.onMessageObjc(cmsg)
}

func recvMessage(ptr unsafe.Pointer) {
	msg := NewCMessage()
	gmsg := msg.Read(ptr)

	switch gmsg.op {
	case none:
		fmt.Println("none")
	case demo:
		fmt.Println("demo")
	case reset:
		fmt.Println("reset")
	case quit:
		fmt.Println("quit")
		os.Exit(1)
	}
}

func main() {
	fmt.Println("GoMenu Demo")
	sendMessage("hello", demo)
	C.StartApp()
}
