package main

/*
#cgo arm LDFLAGS: -framework Foundation -framework UIKit -isysroot /usr/share/SDKs/iPhoneOS.sdk -miphoneos-version-min=7.0 -arch armv7
#cgo arm CFLAGS: -ObjC -isysroot /usr/share/SDKs/iPhoneOS.sdk -miphoneos-version-min=7.0 -arch armv7
#cgo arm64 LDFLAGS: -framework Foundation -framework UIKit -isysroot /usr/share/SDKs/iPhoneOS.sdk -miphoneos-version-min=7.0 -arch arm64
#cgo arm64 CFLAGS: -ObjC -isysroot /usr/share/SDKs/iPhoneOS.sdk -miphoneos-version-min=7.0 -arch arm64

#import <Foundation/Foundation.h>
#import <UIKit/UIKit.h>

void clipText(char *data);

void clipText(char *data) {
	NSString *text = [NSString stringWithCString:data encoding:NSUTF8StringEncoding];
	UIPasteboard *pasteboard = [UIPasteboard generalPasteboard];
	if ([UIDevice currentDevice].systemVersion.doubleValue >= 10) {
		pasteboard.string = text;
	} else {
		[pasteboard setValue:text forPasteboardType:@"public.text"];
	}
}
*/
import "C"

import (
	"io/ioutil"
	"os"
)

func main() {
	var text string
	if len(os.Args) >= 2 {
		var data = make([]byte, 0, 100*1000)
		for i := 1; i < len(os.Args); i++ {
			data = append(data, ' ')
			data = append(data, os.Args[i]...)
		}
		text = string(data[1:])
	} else {
		data, _ := ioutil.ReadAll(os.Stdin)
		text = string(data)
	}
	C.clipText(C.CString(text))
}
