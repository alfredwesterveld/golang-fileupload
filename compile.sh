#!/bin/bash
# ---
echo "Cross compiling for windows, mac, linux\n";
echo "-----------------------------------------";
#----

#---- Compile Windows
echo "[BEGIN] Compiling for windows to bin/fileupload-win32\n";
env GOOS=windows go build -o bin/fileupload-win32; # windows
echo "[DONE] Compiling for windows to bin/fileupload-win32\n";
echo "-----------------------------------------";

# ---- Compile Mac
echo "[BEGIN] Compiling for mac to bin/fileupload-darwin\n"
env GOOS=darwin go build -o bin/fileupload-darwin # darwin / Mac
echo "[DONE] Compiling for mac to bin/fileupload-darwin\n"
echo "-----------------------------------------";

# ---- Start Linux
echo "[BEGIN] Compiling for linux to bin/fileupload-linux\n"
env GOOS=linux go build -o bin/fileupload-linux # linux
echo "[DONE] Compiling for linux to bin/fileupload-linux\n"