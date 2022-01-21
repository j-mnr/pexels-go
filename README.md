# pexels-go

The unofficial Pexels API client written in Go (Golang).

## Supported Endpoints & Features

### Photo Endpoints

- [x] Get Photo by ID
- [x] Get Curated Photos
- [x] Search Photos

### Video Endpoints

- [x] Get Video by ID
- [x] Get Popular Videos
- [x] Search Videos

### Collection Endpoints

- [x] Get User's Collection
- [x] Get User's Collections

## Get Started

You will need to request an API key from
[Pexels](https://www.pexels.com/api/new/)

Set an environment variable `PEXELS_API_KEY` to your received API key.
Copy and paste this code snippet into terminal for a quick example.

```sh
mkdir temp && cd temp
go mod init example
go get -u github.com/JayMonari/pexels-go
cat << EOF > main.go
package main

import (
  "encoding/json"
  "fmt"
  "log"
  "os"

  "github.com/JayMonari/pexels-go"
)

func main() {
  myAPIKey := os.Getenv("PEXELS_API_KEY")
  client, err := pexels.NewClient(pexels.Options{APIKey: myAPIKey})
  if err != nil {
    log.Fatal(err)
  }
  params := pexels.PhotoSearchParams{
    Query:       "Ocean",
    Locale:      pexels.EN_US,
    Orientation: pexels.Landscape,
    Size:        pexels.Medium,
    Color:       pexels.Red,
    Page:        3,
    PerPage:     3,
  }
  photos, err := client.SearchPhotos(&params)
  p, _ := json.MarshalIndent(photos.Payload, "", "  ")
  fmt.Println(string(p))
}
EOF
go run main.go
```

### Example Output

```json
{
  "photos": [
    {
      "id": 1136456,
      "width": 5108,
      "height": 2874,
      "url": "https://www.pexels.com/photo/silhouette-of-clouds-during-golden-hour-photograph-1136456/",
      "photographer": "Johannes Plenio",
      "photographer_url": "https://www.pexels.com/@jplenio",
      "photographer_id": 424445,
      "avg_color": "#552313",
      "src": {
        "original": "https://images.pexels.com/photos/1136456/pexels-photo-1136456.jpeg",
        "large2x": "https://images.pexels.com/photos/1136456/pexels-photo-1136456.jpeg?auto=compress\u0026cs=tinysrgb\u0026dpr=2\u0026h=650\u0026w=940",
        "large": "https://images.pexels.com/photos/1136456/pexels-photo-1136456.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=650\u0026w=940",
        "medium": "https://images.pexels.com/photos/1136456/pexels-photo-1136456.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=350",
        "small": "https://images.pexels.com/photos/1136456/pexels-photo-1136456.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=130",
        "portrait": "https://images.pexels.com/photos/1136456/pexels-photo-1136456.jpeg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=1200\u0026w=800",
        "landscape": "https://images.pexels.com/photos/1136456/pexels-photo-1136456.jpeg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=627\u0026w=1200",
        "tiny": "https://images.pexels.com/photos/1136456/pexels-photo-1136456.jpeg?auto=compress\u0026cs=tinysrgb\u0026dpr=1\u0026fit=crop\u0026h=200\u0026w=280"
      },
      "liked": false
    },
    {
      "id": 884549,
      "width": 4107,
      "height": 3286,
      "url": "https://www.pexels.com/photo/sunset-view-on-sea-884549/",
      "photographer": "Rifqi Ramadhan",
      "photographer_url": "https://www.pexels.com/@rifkyilhamrd",
      "photographer_id": 257470,
      "avg_color": "#401603",
      "src": {
        "original": "https://images.pexels.com/photos/884549/pexels-photo-884549.jpeg",
        "large2x": "https://images.pexels.com/photos/884549/pexels-photo-884549.jpeg?auto=compress\u0026cs=tinysrgb\u0026dpr=2\u0026h=650\u0026w=940",
        "large": "https://images.pexels.com/photos/884549/pexels-photo-884549.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=650\u0026w=940",
        "medium": "https://images.pexels.com/photos/884549/pexels-photo-884549.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=350",
        "small": "https://images.pexels.com/photos/884549/pexels-photo-884549.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=130",
        "portrait": "https://images.pexels.com/photos/884549/pexels-photo-884549.jpeg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=1200\u0026w=800",
        "landscape": "https://images.pexels.com/photos/884549/pexels-photo-884549.jpeg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=627\u0026w=1200",
        "tiny": "https://images.pexels.com/photos/884549/pexels-photo-884549.jpeg?auto=compress\u0026cs=tinysrgb\u0026dpr=1\u0026fit=crop\u0026h=200\u0026w=280"
      },
      "liked": false
    },
    {
      "id": 2749928,
      "width": 4668,
      "height": 3112,
      "url": "https://www.pexels.com/photo/photo-of-ocean-during-sunset-2749928/",
      "photographer": "David Frampton",
      "photographer_url": "https://www.pexels.com/@david-frampton-1235333",
      "photographer_id": 1235333,
      "avg_color": "#97694F",
      "src": {
        "original": "https://images.pexels.com/photos/2749928/pexels-photo-2749928.jpeg",
        "large2x": "https://images.pexels.com/photos/2749928/pexels-photo-2749928.jpeg?auto=compress\u0026cs=tinysrgb\u0026dpr=2\u0026h=650\u0026w=940",
        "large": "https://images.pexels.com/photos/2749928/pexels-photo-2749928.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=650\u0026w=940",
        "medium": "https://images.pexels.com/photos/2749928/pexels-photo-2749928.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=350",
        "small": "https://images.pexels.com/photos/2749928/pexels-photo-2749928.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=130",
        "portrait": "https://images.pexels.com/photos/2749928/pexels-photo-2749928.jpeg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=1200\u0026w=800",
        "landscape": "https://images.pexels.com/photos/2749928/pexels-photo-2749928.jpeg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=627\u0026w=1200",
        "tiny": "https://images.pexels.com/photos/2749928/pexels-photo-2749928.jpeg?auto=compress\u0026cs=tinysrgb\u0026dpr=1\u0026fit=crop\u0026h=200\u0026w=280"
      },
      "liked": false
    }
  ],
  "total_results": 8000,
  "page": 3,
  "per_page": 3,
  "prev_page": "https://api.pexels.com/v1/search/?color=red\u0026locale=en-US\u0026orientation=landscape\u0026page=2\u0026per_page=3\u0026query=Ocean\u0026size=medium",
  "next_page": "https://api.pexels.com/v1/search/?color=red\u0026locale=en-US\u0026orientation=landscape\u0026page=4\u0026per_page=3\u0026query=Ocean\u0026size=medium"
}
```

## License

This package is distributed under the terms of the [MIT](LICENSE) License

## Helpful Links

- [Pexels API Docs](https://www.pexels.com/api/documentation/)
