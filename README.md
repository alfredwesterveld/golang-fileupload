# Warning

This service has no security at all so it should only run in trusted environment (behind firewall)

# Fileupload

this is a real simple static binary to have upload/download service in one.

## installation

Most easily just download cross compiled binaries from [releases](https://github.com/alfredwesterveld/golang-fileupload/releases)

## Usage

Standard it will run at `0.0.0.0:1323`. The <file> you upload will be available in the folder `./public/` which is also accessible via http://0.0.0.0:1323/<file>.
So if file is `google.png` it is accessible via `http://0.0.0.0:1323/google.png`.

you can change address bind to by using environment variable `ADDRESS`.
