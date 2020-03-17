package main

/*
#cgo LDFLAGS: -framework Foundation -framework UIKit -isysroot /usr/share/SDKs/iPhoneOS.sdk
#cgo CFLAGS: -ObjC -isysroot /usr/share/SDKs/iPhoneOS.sdk

#import <Foundation/Foundation.h>
#import <UIKit/UIKit.h>

void clipText(char *text);

void clipText(char *text) {
	[[UIPasteboard generalPasteboard] setValue:[NSString stringWithCString:text encoding:NSUTF8StringEncoding] forPasteboardType:@"public.text"];
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
