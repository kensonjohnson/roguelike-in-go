#!/usr/bin/env bash

echoPurple() {
  echo -e "\x1b[95m$1\x1b[0m"
}   

echoYellow() {
  echo -e "\x1b[93m$1\x1b[0m"
}


if [ -d "dist" ]; then
  echoPurple "✅ Dist directory exists"
else
    echoYellow "❌ Dist directory missing"
    echoYellow "🔨 Creating dist directory"
    mkdir dist
    echoPurple "✅ Dist directory created"
fi

echoPurple "✅ Building for WebAssembly"
GOOS=js GOARCH=wasm go build -o dist/rogue-game.wasm
if [ -f "dist/roguelike-demo.wasm" ]; then
  echoPurple "✅ Finished build"
else
    echoYellow "❌ WebAssembly module failed to build. Exiting..."
    exit 1
fi

if [ -f "dist/wasm_exec.js" ]; then
  echoPurple "✅ wasm_exec.js exists in dist directory"
else
    echoYellow "❌ wasm_exec.js not in dist directory"
    echoPurple "🔨Copy the wasm_exec.js file to the dist directory"
    cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" dist
    echoPurple "✅ wasm_exec.js copied to dist directory"
fi

if [ -f "dist/index.html" ]; then
  echoPurple "✅ index.html exists in dist directory"
else
    echoYellow "❌ index.html not in dist directory"
    echoPurple "🔨 Copy the index.html file to the dist directory"
cp index.html dist
    echoPurple "✅ index.html copied to dist directory"
fi

echoPurple "🎉 Build complete"
