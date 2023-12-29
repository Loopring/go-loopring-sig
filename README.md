# EDdsa Signature Implementation in Go

## Install

[Install Golang](https://go.dev/doc/install)

Install [Xcode Command Line Tools](https://developer.apple.com/downloads/)

```
go install golang.org/x/mobile/cmd/gomobile@latest
```

## wasm
wasm tinygo: v0.30.0
[wasm_exec.js](src%2Fapi%2Fsign%2Fwasm_exec.js) from `$(tinygo env TINYGOROOT)targets/wasm_exec.js`

```
gomobile init
```

## Build

### iOS 

```
gomobile bind -target=ios -o ./ios/loopringGoSign.xcframework
```

This command will generate a binary according to the target platform.

Drag the binary into project and build.

### Android 

```
gomobile bind  -target=android  -o ./android/loopringGoSign.aar
```

You can find the aar file in the android directory.