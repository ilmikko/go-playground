package main;

import "github.com/fogleman/gg";
import "fmt";
import "math";

// Color is three floats r,g,b between 0..1
type color [3]float64;

// Palette is a list of colors
type palette []color;

// Get the fractional component of a float (3.14159 -> 0.14159)
func frac(f float64) (float64){
	return f-math.Floor(f);
}

func getColor(p float64, palette palette) (float64, float64, float64){
	// Get a color from the palette using p which is 0..1

	p *= float64(len(palette)); // Scale to the size of our palette
	index := int(math.Floor(p)); // Floor to an index

	var r float64;
	var g float64;
	var b float64;

	// Because we Floor, there should be always a color at index+1.
	// In the (rare) case palette[index+1] does not exist, we return palette[index].
	if (index+1<len(palette)) {
		//r,g,b = palette[index+1]; //???
		// Calculate a value between these two colors
		r1 := palette[index][0];
		g1 := palette[index][1];
		b1 := palette[index][2];

		r2 := palette[index+1][0];
		g2 := palette[index+1][1];
		b2 := palette[index+1][2];

		// Get the fractional part of p
		f := frac(p);
		// Use the fractional part to lerp between these two values
		r = r1 + (r2-r1)*f;
		g = g1 + (g2-g1)*f;
		b = b1 + (b2-b1)*f;
	} else {
		//r,g,b = palette[index]; //????
		r = palette[index][0];
		g = palette[index][1];
		b = palette[index][2];
	}

	return r,g,b;
}

// Canvas definitions
const pixelW = {{RESX}};
const pixelH = {{RESY}};

// Math definitions
var mathX = {{VIEWX}};
var mathY = {{VIEWY}};
var mathW = {{VIEWW}};
var mathH = {{VIEWH}};

// Palettes
var paletteRainbow = palette{color{0,0,0},color{1,0,0},color{1,1,0},color{0,1,0},color{0,1,1},color{0,0,1},color{1,0,1},color{1,1,1}};
var paletteBNW = palette{color{0,0,0},color{1,1,1}};

func renderZoomCrosshair(ctx *gg.Context, x int, y int){
	ctx.SetLineWidth(3);
	ctx.SetRGB(0,0,0);

	ctx.DrawLine(float64(x),0,float64(x),float64(pixelH));
	ctx.Stroke();

	ctx.DrawLine(0,float64(y),float64(pixelW),float64(y));
	ctx.Stroke();

	ctx.DrawLine(0,0,float64(x),float64(y));
	ctx.Stroke();
}

func getMathCoordinates(pixelX int, pixelY int) (float64, float64){
	// Convert the pixelX and pixelY to x and y in the definition.
	// We know that x and y are the center coordinates, w and h are the width and height of our "math viewport".
	// This means that the left side of the canvas, where pixelX = 0, will be mathX-mathW/2.
	// The right side where pixelX = pixelW is mathX+mathW/2.

	px := float64(pixelX)/float64(pixelW); // 0..1
	py := float64(pixelY)/float64(pixelH); // 0..1

	x := float64(mathX) - float64(mathW)/2 + px*float64(mathW);
	y := float64(mathY) - float64(mathH)/2 + py*float64(mathH);

	return x, y;
}

func renderImage(ctx *gg.Context, filename string){
	ctx.SetRGB(1,1,1);
	ctx.Clear();

	iterations := {{ITERATIONS}};

	ctx.SetRGB(0,0,0);
	// Loop through every pixel
	// TODO: This loop could be threaded, as a pixel need not know its neighbor to calculate its value.
	for pixelX := 0; pixelX < pixelW; pixelX++ {
		for pixelY := 0; pixelY < pixelH; pixelY++ {
			x, y := getMathCoordinates(pixelX, pixelY);

			// Save the initial values
			x0 := x;
			y0 := y;

			for iteration := 0; iteration < iterations; iteration++ {
				if {{EXPRESSION}} {
					// The value is (still) inside the set
					// Perform transformations on both x and y
					newx := {{TRANSFORMX}};
					newy := {{TRANSFORMY}};
					x = newx;
					y = newy;
				} else {
					// The value escaped the set
					// Calculate the color
					r,g,b := getColor(float64(iteration)/float64(iterations),paletteRainbow);
					ctx.SetRGB(r,g,b);
					ctx.SetPixel(pixelX,pixelY);
					break;
				}
			}
		}
	}

	// Get a zoom in crosshair, calculate the pixel values
	//zoomX := mathX+0.1;
	//zoomY := mathY+0.1;

	pixelZoomX := 500;
	pixelZoomY := 600;

	// Render the zoom in crosshair for debug
	renderZoomCrosshair(ctx,pixelZoomX,pixelZoomY);

	// Get the coordinates for this pixel
	x, y := getMathCoordinates(pixelZoomX, pixelZoomY);

	// Set zoom to x and y
	fmt.Printf("Math coords: %f,%f -> %f,%f\n",mathX,mathY,x,y);
	mathX = x;
	mathY = y;

	ctx.SavePNG(filename);
}

func main(){
	fmt.Printf("%s\n","Running `{{EXPRESSION}}` for {{ITERATIONS}} iterations");

	ctx := gg.NewContext(pixelW, pixelH);

	renderImage(ctx, "{{EXPRESSION}}.png");
}
