package main;

import "github.com/fogleman/gg";
import "fmt";

func main(){
	fmt.Printf("%s\n","test");

	w := 100;
	h := 100;

	ctx := gg.NewContext(w, h);

	ctx.SetRGB(0,0,0);
	ctx.SetPixel(w/2,h/2);

	ctx.SavePNG("mandelbrot.png");
}
