package memorydll

//#include<stdlib.h> //for C.free
//import "C" it cannot test, so it's commented
import (
	"testing"
	"encoding/base64"
	"fmt"
)
/*
//file example.c
gcc -c -fpic example.c

gcc -shared example.o -o example.dll

#include <stdio.h>
/ * Compute the greatest common divisor of positive integers * /

int gcd(int x, int y) {
    int g;
    g = y;
    while (x > 0) {
        g = x;
        x = y % x;
        y = g;
    }
    return g;
}

void print(char * hello){
    printf(hello);
}
 */

//the following dll is example.dll
var testdllbase64=`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
9Ou+jbYAAAAAjbwnAAAAAFWhQEAQcInlhcB0BF3DZpBduAEAAACjQEAQcOuDkJCQVbloMBBwieXr
FI22AAAAAItRBIsBg8EIAYIAABBwgfloMBBwcupdw5CQkJCQkJCQ/yWAYBBwkJD/JYhgEHCQkP8l
jGAQcJCQ/yWEYBBwkJD/JZBgEHCQkP8llGAQcJCQ/yVwYBBwkJD/JXRgEHCQkFWJ5YPsGOil/f//
xwQk/BEQcOg5/f//ycOQkJD/////wBMQcAAAAAD/////AAAAAAAAAAAAAAAAAAAAAAAAAADsExBw
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGxpYmdj
Y19zX2R3Mi0xLmRsbABfX3JlZ2lzdGVyX2ZyYW1lX2luZm8AbGliZ2NqX3MuZGxsAF9Kdl9SZWdp
c3RlckNsYXNzZXMAX19kZXJlZ2lzdGVyX2ZyYW1lX2luZm8AAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMFF
FFcAAAAAPFAAAAEAAAACAAAAAgAAAChQAAAwUAAAOFAAAEASAABvEgAASFAAAExQAAAAAAEAZXhh
bXBsZS5kbGwAZ2NkAHByaW50AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAYAAAAAAA
AAAAAAAIYQAAcGAAAFBgAAAAAAAAAAAAADBhAACAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
nGAAALBgAAAAAAAAAAAAAMJgAADQYAAA2mAAAORgAADsYAAA9mAAAAAAAAAAAAAAnGAAALBgAAAA
AAAAAAAAAMJgAADQYAAA2mAAAORgAADsYAAA9mAAAAAAAABRAUdldE1vZHVsZUhhbmRsZUEAAGwB
R2V0UHJvY0FkZHJlc3MAADQAX19kbGxvbmV4aXQAtgBfZXJybm8AAGICZmZsdXNoAABxAmZyZWUA
AKQCbWFsbG9jAACxAnByaW50ZgAAAGAAAABgAABLRVJORUwzMi5kbGwAAAAAFGAAABRgAAAUYAAA
FGAAABRgAAAUYAAAbXN2Y3J0LmRsbAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAABkAAAA
CTATMDEwTDCiMMQw0zASMSAxQjFQMXkxizGiMakxsTG8Mc4x5DEFMhcyLTKnMrQyvzLYMvMy/TIO
MxkzMjNHM1IzajNwM4IzijOSM5ozojOqM7IzujPOM+AzAAAAIAAADAAAAAAwAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMgAAAAIAAAAA
AKQCAABpAgAAX19DVE9SX0xJU1RfXwCGAgAAX19EVE9SX0xJU1RfXwAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACgAgAAAgAAAAAA
BAFHTlUgQyA0LjQuMAABLi4vLi4vLi4vZ2NjLTQuNC4wL2xpYmdjYy8uLi9nY2MvbGliZ2NjMi5j
AEM6XE1pbkdXXHNyY1xnY2NiZlxtaW5ndzMyXGxpYmdjYwCAExBwgBMQcAAAAAACBAVpbnQAAgQH
dW5zaWduZWQgaW50AAICB3Nob3J0IHVuc2lnbmVkIGludAACAQZjaGFyAANfaW9idWYAIAGCPgEA
AARfcHRyAAGDPgEAAAIjAARfY250AAGEcgAAAAIjBARfYmFzZQABhT4BAAACIwgEX2ZsYWcAAYZy
AAAAAiMMBF9maWxlAAGHcgAAAAIjEARfY2hhcmJ1ZgABiHIAAAACIxQEX2J1ZnNpegABiXIAAAAC
IxgEX3RtcGZuYW1lAAGKPgEAAAIjHAAFBJ8AAAAGRklMRQABi6cAAAACCAVsb25nIGxvbmcgaW50
AAIEBWxvbmcgaW50AAICBXNob3J0IGludAACBAdsb25nIHVuc2lnbmVkIGludAAHBAcCAQZzaWdu
ZWQgY2hhcgACAQh1bnNpZ25lZCBjaGFyAAIIB2xvbmcgbG9uZyB1bnNpZ25lZCBpbnQAAgQEZmxv
YXQAAggEZG91YmxlAAIIA2NvbXBsZXggZmxvYXQAAhADY29tcGxleCBkb3VibGUAAgwEbG9uZyBk
b3VibGUAAhgDY29tcGxleCBsb25nIGRvdWJsZQAGZnVuY19wdHIAAig4AgAABQQ+AgAACAEJRAEA
AEsCAAAKAAtfaW9iAAGaQAIAAAEBCSgCAABpAgAADI8BAAABAA1fX0NUT1JfTElTVF9fAAOqCFkC
AAABBQPcExBwDV9fRFRPUl9MSVNUX18AA6sIWQIAAAEFA+gTEHAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAREBJQgTCwMIGwgR
ARIBEAYAAAIkAAsLPgsDCAAAAxMBAwgLCzoLOwsBEwAABA0AAwg6CzsLSRM4CgAABQ8ACwtJEwAA
BhYAAwg6CzsLSRMAAAckAAsLPgsAAAgVACcMAAAJAQFJEwETAAAKIQAAAAs0AAMIOgs7C0kTPww8
DAAADCEASRMvCwAADTQAAwg6CzsFSRM/DAIKAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACJAAAAAgCDAAAAAQH7
Dg0AAQEBAQAAAAEAAAEvbWluZ3cvbGliL2djYy9taW5ndzMyLy4uLy4uLy4uL2luY2x1ZGUALi4v
Li4vLi4vZ2NjLTQuNC4wL2xpYmdjYy8uLi9nY2MAAHN0ZGlvLmgAAQAAZ2JsLWN0b3JzLmgAAgAA
bGliZ2NjMi5jAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAC5maWxlAAAADwAAAP7/
AABnAWRsbGNydDEuYwAAAAAAAAAAAAAAAAA6AAAAAAAAAAEAIAADAQAAAAAAAAAAAAAAAAAAAAAA
AAAAAABGAAAAAAAAAAQAAAADAAAAAABUAAAAEAAAAAQAAAADAAAAAABhAAAAYAAAAAEAIAACAF9h
dGV4aXQAEAEAAAEAIAACAF9fb25leGl0QAEAAAEAIAACAC50ZXh0AAAAAAAAAAEAAAADAWUBAAAU
AAAAAAAAAAAAAAAAAC5kYXRhAAAAAAAAAAIAAAADAQAAAAAAAAAAAAAAAAAAAAAAAC5ic3MAAAAA
AAAAAAQAAAADASAAAAAAAAAAAAAAAAAAAAAAAC5maWxlAAAAIQAAAP7/AABnAWN5Z21pbmctY3J0
YmVnaW4uYwAAAAB3AAAAZAAAAAMAAAADAF9vYmoAAAAAIAAAAAQAAAADAAAAAACLAAAABAAAAAIA
AAADAAAAAACZAAAAcAEAAAEAIAACAQAAAAAAAAAAAAAAAAAAAAAAAAAAAACvAAAA/AEAAAEAIAAC
AC50ZXh0AAAAcAEAAAEAAAADAc8AAAAUAAAAAAAAAAAAAAAAAC5kYXRhAAAAAAAAAAIAAAADAQAA
AAAAAAAAAAAAAAAAAAAAAC5ic3MAAAAAIAAAAAQAAAADASAAAAAAAAAAAAAAAAAAAAAAAC5yZGF0
YQAAAAAAAAMAAAADAWIAAAAAAAAAAAAAAAAAAAAAAAAAAADHAAAAZAAAAAMAAAADAC5qY3IAAAAA
BAAAAAIAAAADAC5maWxlAAAALAAAAP7/AABnAWV4YW1wbGUuYwAAAAAAAAAAAF9nY2QAAAAAQAIA
AAEAIAACAQAAAAAAAAAAAAAAAAAAAAAAAF9wcmludAAAbwIAAAEAIAACAC50ZXh0AAAAQAIAAAEA
AAADAUIAAAABAAAAAAAAAAAAAAAAAC5kYXRhAAAAAAAAAAIAAAADAQAAAAAAAAAAAAAAAAAAAAAA
AC5ic3MAAAAAQAAAAAQAAAADAQAAAAAAAAAAAAAAAAAAAAAAAC5maWxlAAAANgAAAP7/AABnAWRs
bG1haW4uYwAAAAAAAAAAAAAAAADRAAAAkAIAAAEAIAACAQAAAAAAAAAAAAAAAAAAAAAAAC50ZXh0
AAAAkAIAAAEAAAADAQwAAAAAAAAAAAAAAAAAAAAAAC5kYXRhAAAAAAAAAAIAAAADAQAAAAAAAAAA
AAAAAAAAAAAAAC5ic3MAAAAAQAAAAAQAAAADAQAAAAAAAAAAAAAAAAAAAAAAAC5maWxlAAAARAAA
AP7/AABnAWdjY21haW4uYwAAAAAAAAAAAAAAAADdAAAAQAAAAAQAAAADAHAuMAAAAAAAAAAAAAIA
AAADAAAAAADqAAAAoAIAAAEAIAACAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAD9AAAA0AIAAAEAIAAC
AF9fX21haW4AMAMAAAEAIAACAC50ZXh0AAAAoAIAAAEAAAADAa0AAAALAAAAAAAAAAAAAAAAAC5k
YXRhAAAAAAAAAAIAAAADAQQAAAABAAAAAAAAAAAAAAAAAC5ic3MAAAAAQAAAAAQAAAADARAAAAAA
AAAAAAAAAAAAAAAAAC5maWxlAAAATgAAAP7/AABnAXBzZXVkby1yZWxvYy5jAAAAAAAAAAAQAQAA
UAMAAAEAIAACAQAAAAAAAAAAAAAAAAAAAAAAAC50ZXh0AAAAUAMAAAEAAAADASgAAAADAAAAAAAA
AAAAAAAAAC5kYXRhAAAABAAAAAIAAAADAQAAAAAAAAAAAAAAAAAAAAAAAC5ic3MAAAAAUAAAAAQA
AAADAQAAAAAAAAAAAAAAAAAAAAAAAC5maWxlAAAAVgAAAP7/AABnAQAAAAArAQAAAAAAAAAAAAAA
AC50ZXh0AAAAgAMAAAEAAAADAQAAAAAAAAAAAAAAAAAAAAAAAC5kYXRhAAAABAAAAAIAAAADAQAA
AAAAAAAAAAAAAAAAAAAAAC5ic3MAAAAAUAAAAAQAAAADAQIAAAAAAAAAAAAAAAAAAAAAAC5maWxl
AAAAkAAAAP7/AABnAWxpYmdjYzIuYwAAAAAAAAAAAC50ZXh0AAAAgAMAAAEAAAADAQAAAAAAAAAA
AAAAAAAAAAAAAC5kYXRhAAAABAAAAAIAAAADAQAAAAAAAAAAAAAAAAAAAAAAAC5ic3MAAAAAVAAA
AAQAAAADAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAA/AQAAAAAAAAoAAAADAZwAAAAAAAAAAAAAAAAA
AAAAAAAAAABNAQAAAAAAAAkAAAADAaQCAAAGAAAAAAAAAAAAAAAAAAAAAABZAQAAAAAAAAsAAAAD
AY0AAAAAAAAAAAAAAAAAAAAAAAAAAABlAQAAAAAAAAgAAAADATYAAAABAAAAAAAAAAAAAAAAAC50
ZXh0AAAAgAMAAAEAAAADAC5kYXRhAAAABAAAAAIAAAADAC5ic3MAAAAAVAAAAAQAAAADAC5pZGF0
YSQ3GAEAAAYAAAADAC5pZGF0YSQ1gAAAAAYAAAADAC5pZGF0YSQ0UAAAAAYAAAADAC5pZGF0YSQ2
wgAAAAYAAAADAC50ZXh0AAAAiAMAAAEAAAADAC5kYXRhAAAABAAAAAIAAAADAC5ic3MAAAAAVAAA
AAQAAAADAC5pZGF0YSQ3IAEAAAYAAAADAC5pZGF0YSQ1iAAAAAYAAAADAC5pZGF0YSQ0WAAAAAYA
AAADAC5pZGF0YSQ22gAAAAYAAAADAC50ZXh0AAAAkAMAAAEAAAADAC5kYXRhAAAABAAAAAIAAAAD
AC5ic3MAAAAAVAAAAAQAAAADAC5pZGF0YSQ3JAEAAAYAAAADAC5pZGF0YSQ1jAAAAAYAAAADAC5p
ZGF0YSQ0XAAAAAYAAAADAC5pZGF0YSQ25AAAAAYAAAADAC50ZXh0AAAAmAMAAAEAAAADAC5kYXRh
AAAABAAAAAIAAAADAC5ic3MAAAAAVAAAAAQAAAADAC5pZGF0YSQ3HAEAAAYAAAADAC5pZGF0YSQ1
hAAAAAYAAAADAC5pZGF0YSQ0VAAAAAYAAAADAC5pZGF0YSQ20AAAAAYAAAADAC50ZXh0AAAAoAMA
AAEAAAADAC5kYXRhAAAABAAAAAIAAAADAC5ic3MAAAAAVAAAAAQAAAADAC5pZGF0YSQ3KAEAAAYA
AAADAC5pZGF0YSQ1kAAAAAYAAAADAC5pZGF0YSQ0YAAAAAYAAAADAC5pZGF0YSQ27AAAAAYAAAAD
AC50ZXh0AAAAqAMAAAEAAAADAC5kYXRhAAAABAAAAAIAAAADAC5ic3MAAAAAVAAAAAQAAAADAC5p
ZGF0YSQ3LAEAAAYAAAADAC5pZGF0YSQ1lAAAAAYAAAADAC5pZGF0YSQ0ZAAAAAYAAAADAC5pZGF0
YSQ29gAAAAYAAAADAC5maWxlAAAAoAAAAP7/AABnAWZha2UAAAAAAAAAAAAAAAAAAGhuYW1lAAAA
UAAAAAYAAAADAGZ0aHVuawAAgAAAAAYAAAADAC50ZXh0AAAAsAMAAAEAAAADAQAAAAAAAAAAAAAA
AAAAAAAAAC5kYXRhAAAABAAAAAIAAAADAQAAAAAAAAAAAAAAAAAAAAAAAC5ic3MAAAAAVAAAAAQA
AAADAQAAAAAAAAAAAAAAAAAAAAAAAC5pZGF0YSQyFAAAAAYAAAADARQAAAADAAAAAAAAAAAAAAAA
AC5pZGF0YSQ1fAAAAAYAAAADAQQAAAAAAAAAAAAAAAAAAAAAAC5pZGF0YSQ0TAAAAAYAAAADAQQA
AAAAAAAAAAAAAAAAAAAAAC5maWxlAAAAvAAAAP7/AABnAWZha2UAAAAAAAAAAAAAAAAAAC50ZXh0
AAAAsAMAAAEAAAADAQAAAAAAAAAAAAAAAAAAAAAAAC5kYXRhAAAABAAAAAIAAAADAQAAAAAAAAAA
AAAAAAAAAAAAAC5ic3MAAAAAVAAAAAQAAAADAQAAAAAAAAAAAAAAAAAAAAAAAC5pZGF0YSQ0aAAA
AAYAAAADAQQAAAAAAAAAAAAAAAAAAAAAAC5pZGF0YSQ1mAAAAAYAAAADAQQAAAAAAAAAAAAAAAAA
AAAAAC5pZGF0YSQ3MAEAAAYAAAADAQsAAAAAAAAAAAAAAAAAAAAAAC50ZXh0AAAAsAMAAAEAAAAD
AC5kYXRhAAAABAAAAAIAAAADAC5ic3MAAAAAVAAAAAQAAAADAC5pZGF0YSQ3AAEAAAYAAAADAC5p
ZGF0YSQ1cAAAAAYAAAADAC5pZGF0YSQ0QAAAAAYAAAADAC5pZGF0YSQ2nAAAAAYAAAADAC50ZXh0
AAAAuAMAAAEAAAADAC5kYXRhAAAABAAAAAIAAAADAC5ic3MAAAAAVAAAAAQAAAADAC5pZGF0YSQ3
BAEAAAYAAAADAC5pZGF0YSQ1dAAAAAYAAAADAC5pZGF0YSQ0RAAAAAYAAAADAC5pZGF0YSQ2sAAA
AAYAAAADAC5maWxlAAAAzAAAAP7/AABnAWZha2UAAAAAAAAAAAAAAAAAAGhuYW1lAAAAQAAAAAYA
AAADAGZ0aHVuawAAcAAAAAYAAAADAC50ZXh0AAAAwAMAAAEAAAADAQAAAAAAAAAAAAAAAAAAAAAA
AC5kYXRhAAAABAAAAAIAAAADAQAAAAAAAAAAAAAAAAAAAAAAAC5ic3MAAAAAVAAAAAQAAAADAQAA
AAAAAAAAAAAAAAAAAAAAAC5pZGF0YSQyAAAAAAYAAAADARQAAAADAAAAAAAAAAAAAAAAAC5pZGF0
YSQ1bAAAAAYAAAADAQQAAAAAAAAAAAAAAAAAAAAAAC5pZGF0YSQ0PAAAAAYAAAADAQQAAAAAAAAA
AAAAAAAAAAAAAC5maWxlAAAA2gAAAP7/AABnAWZha2UAAAAAAAAAAAAAAAAAAC50ZXh0AAAAwAMA
AAEAAAADAQAAAAAAAAAAAAAAAAAAAAAAAC5kYXRhAAAABAAAAAIAAAADAQAAAAAAAAAAAAAAAAAA
AAAAAC5ic3MAAAAAVAAAAAQAAAADAQAAAAAAAAAAAAAAAAAAAAAAAC5pZGF0YSQ0SAAAAAYAAAAD
AQQAAAAAAAAAAAAAAAAAAAAAAC5pZGF0YSQ1eAAAAAYAAAADAQQAAAAAAAAAAAAAAAAAAAAAAC5p
ZGF0YSQ3CAEAAAYAAAADAQ0AAAAAAAAAAAAAAAAAAAAAAC5maWxlAAAA7AAAAP7/AABnAWN5Z21p
bmctY3J0ZW5kLmMAAAAAAAB1AQAAZAAAAAMAAAADAAAAAACEAQAABAAAAAIAAAADAAAAAACRAQAA
wAMAAAEAIAADAQAAAAAAAAAAAAAAAAAAAAAAAC50ZXh0AAAAwAMAAAEAAAADARkAAAADAAAAAAAA
AAAAAAAAAC5kYXRhAAAABAAAAAIAAAADAQAAAAAAAAAAAAAAAAAAAAAAAC5ic3MAAAAAVAAAAAQA
AAADAQAAAAAAAAAAAAAAAAAAAAAAAAAAAADHAAAAZAAAAAMAAAADAQQAAAAAAAAAAAAAAAAAAAAA
AC5qY3IAAAAABAAAAAIAAAADAQQAAAAAAAAAAAAAAAAAAAAAAAAAAACmAQAA4AMAAAEAAAADAQQA
AAABAAAAAAAAAAAAAAAAAAAAAACzAQAAaAAAAAMAAAACAAAAAADSAQAAAAAAAAIAAAACAAAAAADh
AQAA6AMAAAEAAAACAF9mcmVlAAAAkAMAAAEAIAACAAAAAADwAQAAAAAAAP//AAACAAAAAAAhAgAA
AAAAAAAAAAACAAAAAAAwAgAAMAEAAAYAAAACAF9fZXJybm8AmAMAAAEAIAACAAAAAABEAgAAAAAA
AP//AAACAAAAAAB3AgAAABAAAP//AAACAAAAAACQAgAAAAAgAP//AAACAAAAAACqAgAABAAAAP//
AAACAAAAAADGAgAAAAAAAAAAAAACAAAAAADYAgAAAAAAAAAAAAACAAAAAADqAgAAAAAAAAAAAAAC
AAAAAAD6AgAAsAMAAAEAIAACAAAAAAAOAwAAAAAAAAAAIABpARQAAAABAAAAAAAAAAAAAAAAAAAA
AAAlAwAAAAAAAAQAAAACAAAAAAAzAwAAaAAAAAMAAAACAAAAAABWAwAAABAAAP//AAACAAAAAABu
AwAAhAAAAAYAAAACAAAAAAB8AwAAdAAAAAYAAAACAAAAAACUAwAAuAMAAAEAIAACAAAAAACmAwAA
AAAAAAAAAAACAAAAAAC4AwAAAAAAAAAAAAACAF9fZGxsX18AAAAAAP//AAACAAAAAADIAwAAAAAA
AP//AAACAAAAAADdAwAAFAAAAAYAAAACAAAAAADwAwAAAAAQcP//AAACAAAAAAD/AwAAABAAAP//
AAACAAAAAAAVBAAAaAAAAAMAAAACAAAAAAAzBAAACAAAAAIAAAACAAAAAABABAAA3AMAAAEAAAAC
AF9mZmx1c2gAiAMAAAEAIAACAAAAAABOBAAAdAAAAAQAAAACAAAAAABaBAAAAAAAAAAAAAACAAAA
AABqBAAAAAAAAAAAAAACAAAAAAB8BAAA3AMAAAEAAAACAAAAAACLBAAAgAAAAAYAAAACAAAAAACe
BAAAAAIAAP//AAACAAAAAACxBAAAkAAAAAYAAAACAAAAAAC/BAAABAAAAP//AAACAF9fZW5kX18A
AAAAAAAAAAACAAAAAADUBAAAcAAAAAYAAAACAAAAAADuBAAAgAMAAAEAIAACAF9tYWxsb2MAoAMA
AAEAIAACAAAAAAD7BAAA6AMAAAEAAAACAAAAAAAJBQAAAAAAAP//AAACAAAAAAA+BQAAAAAQAP//
AAACAAAAAABXBQAAAAAAAAAAAAACAAAAAABpBQAAAAAQcP//AAACAAAAAAB2BQAAAwAAAP//AAAC
AAAAAACEBQAAiAAAAAYAAAACAAAAAACSBQAAAAAAAAAAIABpARMAAAABAAAAAAAAAAAAAAAAAAAA
AACnBQAAAAAAAAAAAAACAAAAAAC0BQAAjAAAAAYAAAACAAAAAADABQAAAAAAAAAAIABpARIAAAAB
AAAAAAAAAAAAAAAAAAAAAADZBQAAAQAAAP//AAACAAAAAADxBQAAAAAAAP//AAACAAAAAAACBgAA
lAAAAAYAAAACAAAAAAAQBgAAAAAAAAYAAAACAAAAAAAlBgAAAAAAAP//AAACAAAAAABBBgAAAAAA
AP//AAACAF9wcmludGYAqAMAAAEAIAACAAAAAABZBgAAaAAAAAMAAAACAAAAAAB7BgAACAEAAAYA
AAACAAAAAACRBgAAAAAAAAAAAAACAKEGAAAuZGVidWdfcHVibmFtZXMALmRlYnVnX2luZm8ALmRl
YnVnX2FiYnJldgAuZGVidWdfbGluZQBfX19kbGxfZXhpdABfZmlyc3RfYXRleGl0AF9uZXh0X2F0
ZXhpdABfRGxsTWFpbkNSVFN0YXJ0dXBAMTIAX19fRUhfRlJBTUVfQkVHSU5fXwBfX19KQ1JfTElT
VF9fAF9fX2djY19yZWdpc3Rlcl9mcmFtZQBfX19nY2NfZGVyZWdpc3Rlcl9mcmFtZQAuZWhfZnJh
bWUAX0RsbE1haW5AMTIAX2luaXRpYWxpemVkAF9fX2RvX2dsb2JhbF9kdG9ycwBfX19kb19nbG9i
YWxfY3RvcnMAX19wZWkzODZfcnVudGltZV9yZWxvY2F0b3IAcHNldWRvLXJlbG9jLWxpc3QuYwAu
ZGVidWdfYWJicmV2AC5kZWJ1Z19pbmZvAC5kZWJ1Z19saW5lAC5kZWJ1Z19wdWJuYW1lcwBfX19G
UkFNRV9FTkRfXwBfX19KQ1JfRU5EX18AX3JlZ2lzdGVyX2ZyYW1lX2N0b3IALmN0b3JzLjY1NTM1
AF9fX1JVTlRJTUVfUFNFVURPX1JFTE9DX0xJU1RfXwBfX2RhdGFfc3RhcnRfXwBfX19EVE9SX0xJ
U1RfXwAud2Vhay5fX0p2X1JlZ2lzdGVyQ2xhc3Nlcy5fX19nY2NfcmVnaXN0ZXJfZnJhbWUAX19f
dGxzX3N0YXJ0X18AX19saWJtc3ZjcnRfYV9pbmFtZQAud2Vhay5fX19yZWdpc3Rlcl9mcmFtZV9p
bmZvLl9fX2djY19yZWdpc3Rlcl9mcmFtZQBfX3NpemVfb2Zfc3RhY2tfY29tbWl0X18AX19zaXpl
X29mX3N0YWNrX3Jlc2VydmVfXwBfX21ham9yX3N1YnN5c3RlbV92ZXJzaW9uX18AX19fY3J0X3hs
X3N0YXJ0X18AX19fY3J0X3hpX3N0YXJ0X18AX19fY3J0X3hpX2VuZF9fAF9HZXRNb2R1bGVIYW5k
bGVBQDQAX19fcmVnaXN0ZXJfZnJhbWVfaW5mbwBfX2Jzc19zdGFydF9fAF9fX1JVTlRJTUVfUFNF
VURPX1JFTE9DX0xJU1RfRU5EX18AX19zaXplX29mX2hlYXBfY29tbWl0X18AX19pbXBfX19lcnJu
bwBfX2ltcF9fR2V0UHJvY0FkZHJlc3NAOABfR2V0UHJvY0FkZHJlc3NAOABfX19jcnRfeHBfc3Rh
cnRfXwBfX19jcnRfeHBfZW5kX18AX19taW5vcl9vc192ZXJzaW9uX18AX19oZWFkX2xpYm1zdmNy
dF9hAF9faW1hZ2VfYmFzZV9fAF9fc2VjdGlvbl9hbGlnbm1lbnRfXwBfX1JVTlRJTUVfUFNFVURP
X1JFTE9DX0xJU1RfXwBfX2RhdGFfZW5kX18AX19DVE9SX0xJU1RfXwBfX2Jzc19lbmRfXwBfX19j
cnRfeGNfZW5kX18AX19fY3J0X3hjX3N0YXJ0X18AX19fQ1RPUl9MSVNUX18AX19pbXBfX19fZGxs
b25leGl0AF9fZmlsZV9hbGlnbm1lbnRfXwBfX2ltcF9fbWFsbG9jAF9fbWFqb3Jfb3NfdmVyc2lv
bl9fAF9faW1wX19HZXRNb2R1bGVIYW5kbGVBQDQAX19fZGxsb25leGl0AF9fRFRPUl9MSVNUX18A
LndlYWsuX19fZGVyZWdpc3Rlcl9mcmFtZV9pbmZvLl9fX2djY19yZWdpc3Rlcl9mcmFtZQBfX3Np
emVfb2ZfaGVhcF9yZXNlcnZlX18AX19fY3J0X3h0X3N0YXJ0X18AX19fSW1hZ2VCYXNlAF9fc3Vi
c3lzdGVtX18AX19pbXBfX2ZmbHVzaABfX0p2X1JlZ2lzdGVyQ2xhc3NlcwBfX190bHNfZW5kX18A
X19pbXBfX2ZyZWUAX19fZGVyZWdpc3Rlcl9mcmFtZV9pbmZvAF9fbWFqb3JfaW1hZ2VfdmVyc2lv
bl9fAF9fbG9hZGVyX2ZsYWdzX18AX19pbXBfX3ByaW50ZgBfX2hlYWRfbGlia2VybmVsMzJfYQBf
X21pbm9yX3N1YnN5c3RlbV92ZXJzaW9uX18AX19taW5vcl9pbWFnZV92ZXJzaW9uX18AX19SVU5U
SU1FX1BTRVVET19SRUxPQ19MSVNUX0VORF9fAF9fbGlia2VybmVsMzJfYV9pbmFtZQBfX19jcnRf
eHRfZW5kX18A`
func TestMemoryLoadLibrary(t *testing.T) {
	testdll,_:=base64.StdEncoding.DecodeString(testdllbase64)
	dll,err:=NewDLL(testdll,"example.dll");
	if err!=nil{
		t.Error(err)
		return
	}
	proc,err:=dll.FindProc("gcd")
	if err!=nil{
		t.Error(err)
		return
	}
	result,_,_:=proc.Call(uintptr(4),uintptr(8))
	//fmt.Println("rsult=",result)
	if int(result)!=4{
		t.Error("gcd calc error")
	}
	//go test complains import "C"
//	proc,err=dll.FindProc("print")
//	if err!=nil{
//		t.Error(err)
//		return
//	}
//	cstr:=C.CString("hello,world")
//	defer C.free(unsafe.Pointer(cstr))
//	proc.Call(cstr)
}

