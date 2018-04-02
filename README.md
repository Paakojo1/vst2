# Golang VST package

## Dependencies 
For Go dependencies **dep** is used.

VST2 SDK is also required and since it is under commercial lincense, it cannot be a part of public repository. 

To use this package:
1. Go to [Steinberg web site](https://www.steinberg.net/en/company/developers.html) to get the SDK
2. Extract downloaded archive
3. Set environment variable CGO_CFLAGS="-I<PATH TO /VST2_SDK/public.sdk/source/vst2.x>"
* aeffect.h
* aeffectx.h
* vstfxstore.h