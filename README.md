# Textgen

Text generation from existing source text using weighted predictions.

## Installation

```bash
go get github.com/vcokltfre/textgen@latest
```

## Usage

To train from a corpus:

```bash
textgen train -i corpus.txt
```

To generate text:

```bash
textgen generate
```

### Additional options

For the `train` command you can specify:

- `-i` or `--input` to specify the input file
- `-w` or `--weights` to specify the weights output file

For the `generate` command you can specify:

- `-w` or `--weights` to specify the weights input file
- `-l` or `--length` to specify the length of the generated text
- `-s` or `--start` to specify the seed word for the generator

The `train` command accepts either a file path or an HTTP/HTTPS URL as input.

## Pre-trained Weights

In the `weights/` directory you can find pre-trained weights for the following corpora:

- [The complete works of Shakespeare](https://ocw.mit.edu/ans7870/6/6.006/s08/lecturenotes/files/t8.shakespeare.txt)
- [The Bee Movie script](http://www.script-o-rama.com/movie_scripts/a1/bee-movie-script-transcript-seinfeld.html)
