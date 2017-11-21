package plot

import (
	"github.com/bernardoaraujor/corinda/elementary"
	"github.com/Arafatk/glot"
	//"fmt"
)

func EMLogLog(em elementary.Model){
	tfs := em.TokensNfreqs

	freqs := make([]int, 0)
	is := make([]int, 0)
	for i, tf := range tfs{
		is = append(is, i)
		freqs = append(freqs, tf.Freq)
	}

	dimensions := 2
	// The dimensions supported by the plot
	persist := false
	debug := true
	plot, _ := glot.NewPlot(dimensions, persist, debug)

	pointGroupName := "frequency"
	style := "lines"
	points := [][]int{is, freqs}
	plot.AddPointGroup(pointGroupName, style, points)

	plot.SetTitle(em.Name)

	plot.SetXLabel("Token Rank")
	plot.SetYLabel("Token Freq")

	plot.SetXrange(0, len(is))
	plot.SetYrange(0, freqs[0])

	plot.SetLogscale("x", 10)
	plot.SetLogscale("y", 10)
	//fmt.Println(err)

	plot.SavePlot("plot/plots/" + em.Name + ".png")
}