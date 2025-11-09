package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/join"
)

const (
	flagField1        = "field1"
	flagField2        = "field2"
	flagOutputFormat  = "output-format"
	flagEmptyString   = "empty"
	flagIgnoreCase    = "ignore-case"
	flagOuterJoin     = "outer-join"
	flagUnpairedFile1 = "unpaired-1"
	flagUnpairedFile2 = "unpaired-2"
	flagCheckOrder    = "check-order"
)

func main() {
	app := &cli.App{
		Name:  "join",
		Usage: "join lines of two files on a common field",
		UsageText: `join [OPTIONS] FILE1 FILE2

   For each pair of input lines with identical join fields, write a line to
   standard output. The default join field is the first, delimited by blanks.`,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    flagField1,
				Aliases: []string{"1"},
				Usage:   "join on this FIELD of file 1",
				Value:   1,
			},
			&cli.IntFlag{
				Name:    flagField2,
				Aliases: []string{"2"},
				Usage:   "join on this FIELD of file 2",
				Value:   1,
			},
			&cli.StringFlag{
				Name:    flagOutputFormat,
				Aliases: []string{"o"},
				Usage:   "obey FORMAT while constructing output line",
			},
			&cli.StringFlag{
				Name:    flagEmptyString,
				Aliases: []string{"e"},
				Usage:   "replace missing input fields with STRING",
			},
			&cli.BoolFlag{
				Name:    flagIgnoreCase,
				Aliases: []string{"i"},
				Usage:   "ignore differences in case when comparing fields",
			},
			&cli.BoolFlag{
				Name:    flagOuterJoin,
				Aliases: []string{"a"},
				Usage:   "also print unpairable lines",
			},
			&cli.BoolFlag{
				Name:    flagUnpairedFile1,
				Aliases: []string{"v1"},
				Usage:   "like -a 1, but suppress joined output lines",
			},
			&cli.BoolFlag{
				Name:    flagUnpairedFile2,
				Aliases: []string{"v2"},
				Usage:   "like -a 2, but suppress joined output lines",
			},
			&cli.BoolFlag{
				Name:  flagCheckOrder,
				Usage: "check that the input is correctly sorted",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "join: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add file arguments (requires exactly 2 files)
	for i := 0; i < c.NArg(); i++ {
		params = append(params, gloo.File(c.Args().Get(i)))
	}

	// Add flags based on CLI options
	if c.IsSet(flagField1) {
		params = append(params, Field1(c.Int(flagField1)))
	}
	if c.IsSet(flagField2) {
		params = append(params, Field2(c.Int(flagField2)))
	}
	if c.IsSet(flagOutputFormat) {
		params = append(params, OutputFormat(c.String(flagOutputFormat)))
	}
	if c.IsSet(flagEmptyString) {
		params = append(params, EmptyString(c.String(flagEmptyString)))
	}
	if c.Bool(flagIgnoreCase) {
		params = append(params, IgnoreCase)
	}
	if c.Bool(flagOuterJoin) {
		params = append(params, OuterJoin)
	}
	if c.Bool(flagUnpairedFile1) {
		params = append(params, UnpairedFile1)
	}
	if c.Bool(flagUnpairedFile2) {
		params = append(params, UnpairedFile2)
	}
	if c.Bool(flagCheckOrder) {
		params = append(params, CheckOrder)
	}

	// Create and execute the join command
	cmd := Join(params...)
	return gloo.Run(cmd)
}
