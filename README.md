# caltrain

caltrain provides realtime train timings for various train stations in both
directions. It does so by scraping the Caltrain mobile page.

Please don't abuse this api.

See https://godoc.org/github.com/sent-hil/caltrail-realtime for more api.

## Install

`go get github.com/sent-hil/caltrain-realtime`

## Usage

```go
package main

import (
  "github.com/sent-hil/caltrain-realtime"
)

func main() {
  timings, err := caltrain.GetRealTimings(caltrain.SanFrancisco, caltrail.SouthBound)
}
```
