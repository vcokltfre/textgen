package main

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
	"github.com/vcokltfre/textgen/textgen"
)

func fileOrURI(path string) ([]byte, error) {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		resp, err := http.Get(path)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		return io.ReadAll(resp.Body)
	}

	return os.ReadFile(path)
}

func main() {
	app := &cli.App{
		Name:  "textgen",
		Usage: "Generate text from a corpus",
		Commands: []*cli.Command{
			{
				Name:  "generate",
				Usage: "Generate text from a weights file",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "length",
						Aliases: []string{"l"},
						Value:   100,
					},
					&cli.StringFlag{
						Name:    "weights",
						Aliases: []string{"w"},
						Value:   "weights.json",
					},
					&cli.StringFlag{
						Name:    "start",
						Aliases: []string{"s"},
						Value:   "",
					},
					&cli.StringFlag{
						Name:    "train",
						Aliases: []string{"t"},
						Value:   "",
					},
				},
				Action: func(c *cli.Context) error {
					length := c.Int("length")
					weightsFile := c.String("weights")
					start := c.String("start")
					train := c.String("train")

					var weights *textgen.Weights

					if train == "" {
						w, err := textgen.LoadWeights(weightsFile)
						if err != nil {
							return err
						}

						weights = w
					} else {
						data, err := fileOrURI(train)
						if err != nil {
							return err
						}

						weights = textgen.NewWeights()
						weights.Train(string(data))
					}

					for i := 0; i < length; i++ {
						start = weights.Predict(start)
						if start == "" {
							continue
						}

						os.Stdout.WriteString(start + " ")
					}

					return nil
				},
			},
			{
				Name:  "train",
				Usage: "Train a weights file from a corpus",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "input",
						Aliases: []string{"i"},
						Value:   "input.txt",
					},
					&cli.StringFlag{
						Name:    "weights",
						Aliases: []string{"w"},
						Value:   "weights.json",
					},
				},
				Action: func(c *cli.Context) error {
					inputFile := c.String("input")
					weightsFile := c.String("weights")

					data, err := fileOrURI(inputFile)
					if err != nil {
						return err
					}

					weights, _ := textgen.LoadWeights(weightsFile)
					if weights == nil {
						weights = textgen.NewWeights()
					}

					weights.Train(string(data))

					return weights.Save(weightsFile)
				},
			},
		},
	}

	app.Run(os.Args)
}
