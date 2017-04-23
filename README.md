CERROJO
=======
Go package to work with [TREZOR](http://bitcointrezor.com/) and [KEEPKEY](http://keepkey.com/) devices. This repository is an unofficial package written in Go (golang), for any official package, code or tool, please go to the official [TREZOR repository](https://github.com/trezor) or [KEEPKEY repository](https://github.com/keepkey).

## What is TREZOR?

TREZOR is a single purpose device which allows you to make secure Bitcoin transactions. With TREZOR, transactions are completely safe even when initiated on a compromised or vulnerable computer.


## What is KEEPKEY?

KeepKey is a hardware wallet that secures bitcoin, ethereum, litecoin, dogecoin, dash, and namecoin. Your assets are protected from hackers and thieves..


## Installation
```bash
$ go get github.com/conejoninja/cerrojo
```

## Documentation
None yet, it's a work in progress

## Supported methods
Almost everything is supported except *debuglink* related stuff. Transactions methods are done but not tested.

## Tests
Go to the *tests* folder and run them with
```bash
// Put your device in bootloader mode
go test -v cerrojo_bootloader_test.go
// Disconnect and connect your device in normal mode
go test -v cerrojo_test.go
```

Running tests the *traditional* Go way (*go test*) will not work, as for cerrojo_bootloader_test.go to run you need to put your device in *bootloader* mode, the rest of the tests are run in normal mode.

## Contributing to this project:

If you find any improvement or issue you want to fix, feel free to send me a pull request.

## License

This is distributed under the Apache License v2.0

Copyright 2016-2017 Daniel Esteban  -  conejo@conejo.me

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.


## Note


![](https://raw.githubusercontent.com/conejoninja/cerrojo/master/assets/ribbon.png)

If you would like to donate via Bitcoin, please send your donation to this wallet:

   ![](https://raw.githubusercontent.com/conejoninja/cerrojo/master/assets/qr.png)

Bitcoin: 1G9d7uVvioNt8Emsv6fVmCdAPc41nX1c8J