package annotatedb

import (
	"context"
	"encoding/csv"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/google/subcommands"
	"github.com/nelhage/taktician/ai"
	"github.com/nelhage/taktician/logs"
	"github.com/nelhage/taktician/ptn"
	"golang.org/x/sync/errgroup"
)

type Command struct {
	minPly, maxPly int
	workers        int
}

func (*Command) Name() string     { return "annotatedb" }
func (*Command) Synopsis() string { return "Annotate the playtak DB with analysis at every move" }
func (*Command) Usage() string {
	return `annotatedb [options] game.db
`
}

func (c *Command) SetFlags(flags *flag.FlagSet) {
	flags.IntVar(&c.minPly, "min-ply", 3, "minimum ply to analyze")
	flags.IntVar(&c.maxPly, "max-ply", 5, "maximum ply to analyze")
	flags.IntVar(&c.workers, "workers", 2, "parallel workers")
}

type batch []annotation

func (c *Command) annotateRows(ctx context.Context, todo <-chan todoRow, out chan<- batch) error {
	for row := range todo {
		var batch batch
		game, err := ptn.ParsePTN(strings.NewReader(row.PTN))
		if err != nil {
			return err
		}
		it := game.Iterator()

		init, err := game.InitialPosition()
		if err != nil {
			return err
		}

		cfg := ai.MinimaxConfig{
			Size: init.Size(),
		}
		engine := ai.NewMinimax(cfg)
		for it.Next() {
			p := it.Position()
			if o, _ := p.GameOver(); o {
				break
			}
			for d := c.minPly; d <= c.maxPly; d++ {
				pv, v, _ := engine.AnalyzeDepth(ctx, d, p)
				batch = append(batch, annotation{
					Game:     row.Id,
					Ply:      p.MoveNumber(),
					Depth:    d,
					Analysis: v,
					Move:     ptn.FormatMove(pv[0]),
				})
			}
		}
		out <- batch
	}
	return nil
}

func (c *Command) Execute(ctx context.Context, flag *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	repo, err := logs.Open(flag.Arg(0))
	if err != nil {
		log.Fatalf("open(%q): %s", flag.Arg(0), err.Error())
	}

	db := repo.DB()

	todo := make(chan todoRow)
	results := make(chan batch)

	go func() {
		defer close(todo)
		rows, err := db.Queryx(selectTODO)
		if err != nil {
			log.Fatalf("query: %s", err.Error())
		}
		for rows.Next() {
			var row todoRow
			err := rows.StructScan(&row)
			if err != nil {
				log.Fatalf("scan: %s", err.Error())
			}
			todo <- row
		}
	}()

	var grp errgroup.Group
	for j := 0; j < c.workers; j++ {
		grp.Go(func() error {
			return c.annotateRows(ctx, todo, results)
		})
	}
	go func() {
		if err := grp.Wait(); err != nil {
			log.Fatalf("parse: %v", err)
		}
		close(results)
	}()

	w := csv.NewWriter(os.Stdout)
	for batch := range results {
		for _, a := range batch {
			row := []string{
				strconv.Itoa(a.Game),
				strconv.Itoa(a.Ply),
				strconv.Itoa(a.Depth),
				strconv.FormatInt(a.Analysis, 10),
				a.Move,
			}
			w.Write(row)
		}
	}

	return subcommands.ExitSuccess
}
