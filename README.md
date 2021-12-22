# mtg-i2p

Current version has known bugs. Testers/developers only.

I2P Service modification of the highly opinionated mtg Telegram proxy.
[Telegram](https://telegram.org/). This repository contains mostly code which
is copied from the original mtg source code, specifically, the `internal/cli`,
`internal/utils` and `internal/config` directories. It uses the underlying
`mtglib` that is available at 
[https://github.com/9seconds/mtg](https://github.com/9seconds/mtg). This approach
was taken in order to retain the maxiumum number of features from mtg. If features
are found to be irrelevant to I2P and also obstructive to our use cases, they may
be removed as well. Otherwise, the modifications are intentionally minimal in order
to make it easy to keep our distribution up-to-date.

[Look at MODIFICATIONS.md to see a summary of the modifications made](MODIFICATIONS.md)

Plugin install URL's
--------------------

- **Mac OSX** - [http://idk.i2p/mtg-i2p/mtg-i2p-darwin-amd64.su3](http://idk.i2p/mtg-i2p/mtg-i2p-darwin-amd64.su3)
- **Mac OSX ARM** - [http://idk.i2p/mtg-i2p/mtg-i2p-darwin-arm64.su3](http://idk.i2p/mtg-i2p/mtg-i2p-darwin-arm64.su3)
- **FreeBSD AMD64** - [http://idk.i2p/mtg-i2p/mtg-i2p-freebsd-amd64.su3](http://idk.i2p/mtg-i2p/mtg-i2p-freebsd-amd64.su3)
- **Linux x86** -[http://idk.i2p/mtg-i2p/mtg-i2p-linux-386.su3](http://idk.i2p/mtg-i2p/mtg-i2p-linux-386.su3)
- **Linux AMD64** - [http://idk.i2p/mtg-i2p/mtg-i2p-linux-amd64.su3](http://idk.i2p/mtg-i2p/mtg-i2p-linux-amd64.su3)
- **Linux ARM64** - [http://idk.i2p/mtg-i2p/mtg-i2p-linux-arm64.su3](http://idk.i2p/mtg-i2p/mtg-i2p-linux-arm64.su3)
- **OpenBSD AMD64** - [http://idk.i2p/mtg-i2p/mtg-i2p-openbsd-amd64.su3](http://idk.i2p/mtg-i2p/mtg-i2p-openbsd-amd64.su3)
- **Windows x86** - [http://idk.i2p/mtg-i2p/mtg-i2p-windows-386.su3](http://idk.i2p/mtg-i2p/mtg-i2p-windows-386.su3)
- **Windows AMD64** - [http://idk.i2p/mtg-i2p/mtg-i2p-windows-amd64.su3](http://idk.i2p/mtg-i2p/mtg-i2p-windows-amd64.su3)

Once you install the plugin, your MTProto proxy's base32 address is visible in the file:
`$PLUGIN/telegram.i2p.i2p.public.txt`. You can connect to it by creating an I2P tunnel and
configuring it as your MTProto proxy.
