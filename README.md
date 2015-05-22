# mst
Morten's Statistics Thing

    go install github.com/msiebuhr/mst/cmd/mst

Tired of doing ad-hoc AWK/shell/... scripts to do basic statistics on a set of numbers.

Inspired by [st](https://github.com/nferraz/st), which had just one small
annoyance too many for me...

## Usage

    > echo 1\\n2\\n2\\n4 | mst # Right now it's line-based; probably should split on words
    min 1
    q1 2
    mean 2
    q3 4
    max 4
    sum 9
    count 4
    average 2.25


## License
MIT
