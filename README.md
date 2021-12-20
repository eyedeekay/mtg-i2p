# mtg-i2p

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

This modifcation is created by changing the `internet/config/type_hostport.go`
so that the `config.toml` file make include an I2P value. The modification is:

```diff
diff --git a/internal/config/type_hostport.go b/internal/config/type_hostport.go
index 937c637..8dcb254 100644
--- a/internal/config/type_hostport.go
+++ b/internal/config/type_hostport.go
@@ -4,7 +4,6 @@ import (
 	"fmt"
 	"net"
 	"strconv"
-	"strings"
 )
 
 type TypeHostPort struct {
@@ -14,12 +13,6 @@ type TypeHostPort struct {
 }
 
 func (t *TypeHostPort) Set(value string) error {
-	if strings.HasSuffix(value, ".i2p") {
-		t.Port = uint(0)
-		t.Host = value
-		t.Value = value
-		return nil
-	}
 	host, port, err := net.SplitHostPort(value)
 	if err != nil {
 		return fmt.Errorf("incorrect host:port value (%v): %w", value, err)
```

So you configure it by setting the "`bind-to`" value in `config.toml` to
an I2P hostname, such as "telegram.i2p" like so:

```toml
bind-to = "telegram.i2p"
```

The second modification we do is to change `internal/utils/net_listener.go`
file to produce an I2P listener when it encounters an I2P hostname. This is
the required modification:

```diff
diff --git a/internal/utils/net_listener.go b/internal/utils/net_listener.go
index 9486940..496f51b 100644
--- a/internal/utils/net_listener.go
+++ b/internal/utils/net_listener.go
@@ -3,10 +3,8 @@ package utils
 import (
 	"fmt"
 	"net"
-	"strings"
 
 	"github.com/9seconds/mtg/v2/network"
-	sam "github.com/eyedeekay/sam3/helper"
 )
 
 type Listener struct {
@@ -29,15 +27,6 @@ func (l Listener) Accept() (net.Conn, error) {
 }
 
 func NewListener(bindTo string, bufferSize int) (net.Listener, error) {
-	if strings.HasSuffix(bindTo, ".i2p") {
-		base, err := sam.I2PListener(bindTo, "127.0.0.1:7656", bindTo)
-		if err != nil {
-			return nil, fmt.Errorf("cannot build a base I2P listener: %w", err)
-		}
-		return Listener{
-			Listener: base,
-		}, nil
-	}
 	base, err := net.Listen("tcp", bindTo)
 	if err != nil {
 		return nil, fmt.Errorf("cannot build a base listener: %w", err)
```

Otherwise, mtg is unchanged from it's original stable, efficient self,
effectively turning mtg into a stable, efficient proxy for Telegram
clients to reach services with.