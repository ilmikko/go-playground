package main;

import "github.com/fogleman/gg";
import "fmt";

func main(){
	fmt.Printf("%s\n","Running {{EXPRESSION}}");

	// Canvas definitions
	pixelW := 100;
	pixelH := 100;

	// Math definitions
	mathX := 0;
	mathY := 0;
	mathW := 4;
	mathH := 4;

	ctx := gg.NewContext(pixelW, pixelH);

	ctx.SetRGB(1,1,1);
	ctx.Clear();

	ctx.SetRGB(0,0,0);
	// Loop through every pixel
	// TODO: This loop could be threaded, as a pixel need not know its neighbor to calculate its value.
	for pixelX := 0; pixelX < pixelW; pixelX++ {
		for pixelY := 0; pixelY < pixelH; pixelY++ {
			// Convert the pixelX and pixelY to x and y in the definition.
			// We know that x and y are the center coordinates, w and h are the width and height of our "math viewport".
			// This means that the left side of the canvas, where pixelX = 0, will be mathX-mathW/2.
			// The right side where pixelX = pixelW is mathX+mathW/2.

			px := float64(pixelX)/float64(pixelW); // 0..1
			py := float64(pixelY)/float64(pixelH); // 0..1

			x := float64(mathX) - float64(mathW)/2 + px*float64(mathW);
			y := float64(mathY) - float64(mathH)/2 + py*float64(mathH);

			if {{EXPRESSION}} {
				ctx.SetPixel(pixelX,pixelY);
			}
		}
	}

	ctx.SavePNG("{{EXPRESSION}}.png");
}
