# palette
Creates a palette based on the colors of a given image.

    palette [flags] image_path [...image_path]
    
### flags

    -d, --downsample float   the factor by which to downsample the image before
                             creating a palette (default 1)
    -h, --help               help
    -k, --k int              size of the palette (default 5)
    -s, --seed int           seed used for initializing cluster means. -1 means
                             no seed is used (default -1)

# building

navigate to root directory and run `go build`

# testing

navigate to root directory and run `go test -v ./...`