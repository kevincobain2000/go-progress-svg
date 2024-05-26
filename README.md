<h1 align="center">
  SVG - Circle & Bar Progress generator
  <br>
  in Golang.
  <br>
  <br>
</h1>

<p align="center">
  <img alt="svg circle progress sample" src="https://imgur.com/UOfAB33.png">
</p>
<p align="center">
  <img alt="svg circle progress sample" src="https://imgur.com/LToiOr4.png" width="300">
</p>

---

**Circle Progress:** Generate pure SVG circle progress bar.

**Bar Progress:** Generate pure SVG bar progress bar.

**Supports Captions:** Add captions horizontally or vertically.

**Customizable:** Customize with various color, width, height, background and caption options.

**Lightweight:** No dependencies, just a single file.

**Beautiful:** Customizable to rounded corners, different colors, and caption options.

---
## Reports from [coveritup](https://coveritup.app/readme?org=kevincobain2000&repo=go-progress-svg&branch=master)

![go-build-time](https://coveritup.app/badge?org=kevincobain2000&repo=go-progress-svg&branch=master&type=go-build-time)
![coverage](https://coveritup.app/badge?org=kevincobain2000&repo=go-progress-svg&branch=master&type=coverage)

![go-lint-errors](https://coveritup.app/badge?org=kevincobain2000&repo=go-progress-svg&branch=master&type=go-lint-errors)
![go-test-run-time](https://coveritup.app/badge?org=kevincobain2000&repo=go-progress-svg&branch=master&type=go-test-run-time)

![go-mod-dependencies](https://coveritup.app/badge?org=kevincobain2000&repo=go-progress-svg&branch=master&type=go-mod-dependencies)

![go-lint-errors](https://coveritup.app/chart?org=kevincobain2000&repo=go-progress-svg&branch=master&type=go-lint-errors&theme=light&line=fill&width=150&height=150&output=svg&line=fill)
![go-test-run-time](https://coveritup.app/chart?org=kevincobain2000&repo=go-progress-svg&branch=master&type=go-test-run-time&theme=light&line=fill&width=150&height=150&output=svg&line=fill)
![go-build-time](https://coveritup.app/chart?org=kevincobain2000&repo=go-progress-svg&branch=master&type=go-build-time&theme=light&line=fill&width=150&height=150&output=svg&line=fill)
![coverage](https://coveritup.app/chart?org=kevincobain2000&repo=go-progress-svg&branch=master&type=coverage&theme=light&line=fill&width=150&height=150&output=svg&line=fill)
![go-mod-dependencies](https://coveritup.app/chart?org=kevincobain2000&repo=go-progress-svg&branch=master&type=go-mod-dependencies&theme=light&line=fill&width=150&height=150&output=svg&line=fill)



## Usage

### Circle Progress

```go
import (
	"fmt"
	"os"

	gps "github.com/kevincobain2000/go-progress-svg"
)

func main() {
	circular, _ := gps.NewCircular(func(o *gps.CircularOptions) error {
		o.Size = 200
		o.CircleWidth = 16
		o.ProgressWidth = 16
		o.CircleColor = "#e0e0e0"
		o.ProgressColor = "#76e5b1"
		o.TextColor = "#6bdba7"
		o.TextSize = 52
		o.ShowPercentage = true
		o.BackgroundColor = ""
		o.Caption = ""
		o.CaptionPos = "bottom"
		o.CaptionSize = 20
		o.CaptionColor = "#000000"
		return nil
	})

	circular.SVG()
}
```

### Bar Progress

```go
import (
    "fmt"
    "os"

    gps "github.com/kevincobain2000/go-progress-svg"
)

func main() {
    bar, _ := gps.NewBar(func(o *gps.BarOptions) error {
		o.Progress = 97
		o.Width = 200
		o.Height = 50
		o.ProgressColor = "#76e5b1"
		o.TextColor = "#6bdba7"
		o.TextSize = 20
		o.ShowPercentage = true
		o.Caption = ""
		o.CaptionSize = 16
		o.CaptionColor = "#000000"
		o.BackgroundColor = "#e0e0e0"
		o.CornerRadius = 10
        return nil
    })

    bar.SVG()
}
```


## CHANGE LOG

### v1.0.0 - initial release