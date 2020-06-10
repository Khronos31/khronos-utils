package main

/*
#cgo arm LDFLAGS: -framework Foundation -framework AVFoundation -isysroot /usr/share/SDKs/iPhoneOS.sdk -miphoneos-version-min=7.0 -arch armv7
#cgo arm CFLAGS: -ObjC -isysroot /usr/share/SDKs/iPhoneOS.sdk -miphoneos-version-min=7.0 -arch armv7
#cgo arm64 LDFLAGS: -framework Foundation -framework AVFoundation -isysroot /usr/share/SDKs/iPhoneOS.sdk -miphoneos-version-min=7.0 -arch arm64
#cgo arm64 CFLAGS: -ObjC -isysroot /usr/share/SDKs/iPhoneOS.sdk -miphoneos-version-min=7.0 -arch arm64

#import <Foundation/Foundation.h>
#import <AVFoundation/AVFoundation.h>

void initSpeechSynthesizer();
void speechText(char *text);

void initSpeechSynthesizer() {
	[[AVAudioSession sharedInstance] setCategory:AVAudioSessionCategoryPlayback error:nil];
	[[AVAudioSession sharedInstance] setActive:YES error:nil];
}

void speechText(char *text) {
	AVSpeechSynthesizer *speechSynthesizer = [[AVSpeechSynthesizer alloc] init];
	NSString *speakingText = [NSString stringWithCString:text encoding:NSUTF8StringEncoding];
	AVSpeechUtterance *utterance = [AVSpeechUtterance speechUtteranceWithString:speakingText];
	[speechSynthesizer speakUtterance:utterance];
	do {
		sleep(1);
	} while(speechSynthesizer.speaking);
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
	C.initSpeechSynthesizer()
	C.speechText(C.CString(text))
}
