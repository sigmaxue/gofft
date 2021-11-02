package main

import (
        "os"
        "fmt"
        "math"
        "math/cmplx"
	chart "github.com/wcharczuk/go-chart/v2" //exposes "chart"
        
        "github.com/mjibson/go-dsp/fft"
)

const NUM_POINTS int = 64 


func generate_signal(points int) []complex128 {
	r := make([]complex128, points)
	for i:=0; i<points; i++ {
		r[i] = complex(math.Cos(math.Pi/float64(points)*float64(2*i)) + math.Cos(math.Pi/float64(points)*float64(6*i)),
			math.Sin(math.Pi/float64(points)*float64(2*i)))
	}
	return r
}

func main() {
	r := generate_signal(NUM_POINTS)
	amplitude := make([]float64, NUM_POINTS)
	for k,v := range fft.FFT(r) {
		amplitude[k] = cmplx.Abs(v)
	}

	fmt.Println(amplitude)
	// graph := chart.Chart{
	// 	Series: []chart.Series{
	// 		chart.ContinuousSeries{
	// 			Name: "Frequency Amplitude",
	// 			XValues: chart.Seq{Sequence: chart.NewLinearSequence().WithStart(1.0).WithEnd(float64(NUM_POINTS))}.Values(),
	// 			YValues: amplitude,
	// 		},
	// 	},
	// }
	//buffer := bytes.NewBuffer([]byte{})
	//err := graph.Render(chart.PNG, buffer)

	bars := make([]chart.Value, NUM_POINTS)
	for k,v := range amplitude {
		bars[k].Value = v
		bars[k].Label = fmt.Sprintf("%d", k)
	}

	graph := chart.BarChart{
		Title: "Frequency Amplitude",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 1,
		Bars: bars,
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}

