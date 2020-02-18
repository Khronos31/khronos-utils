package main

import (
	"io/ioutil"
	"os"
)

//#cgo LDFLAGS: -framework Foundation -framework UIKit
//#cgo CFLAGS: -ObjC
/*
#import <Foundation/Foundation.h>
#import <UIKit/UIKit.h>

void clipText(char *text);

void clipText(char *text) {
	[[UIPasteboard generalPasteboard] setValue:[NSString stringWithCString:text encoding:NSUTF8StringEncoding] forPasteboardType:@"public.text"];
}
*/
import "C"

func main() {
	if len(os.Args) >= 2 {
		var data = make([]byte, 0, 100*1000)
		for i := 1; i < len(os.Args); i++ {
			data = append(data, ' ')
			data = append(data, os.Args[i]...)
		}
		C.clipText(C.CString(string(data[1:])))
	} else {
		data, _ := ioutil.ReadAll(os.Stdin)
		C.clipText(C.CString(string(data)))
	}
}
