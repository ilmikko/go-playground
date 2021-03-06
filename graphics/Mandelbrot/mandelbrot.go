package main;

import "github.com/fogleman/gg";
import "fmt";
import "os";
import "time";
import "math";
import "math/rand";

// Color is three floats r,g,b between 0..1
type color [3]float64;

// A pixel is always two integers, like 30,30
type pixel [2]int;

// A saved value contains the calculated value to a pixel.
type savedValue [2]float64;

// Palette is a list of colors
type palette []color;

// Get the fractional component of a float (3.14159 -> 0.14159)
func frac(f float64) (float64) {
	return f-math.Floor(f);
}

func abs(i int) (int) {
	if (i < 0) {
		return -i;
	} else {
		return i;
	}
}

func getColor(p float64, palette palette) (float64, float64, float64){
	// Get a color from the palette using p which is 0..1

	p = frac(p);
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
var paletteRainbow = palette{color{0,0,0},color{1,0,0},color{1,1,0},color{0,1,0},color{0,1,1},color{0,0,1},color{1,0,1},color{1,1,1},color{0,0,0}};
var paletteBNW = palette{color{0,0,0},color{1,1,1}};

func renderZoomCrosshair(ctx *gg.Context, x int, y int, s float64){
	ctx.SetLineWidth(3);
	ctx.SetRGB(0,0,0);

	// Vertical line
	ctx.DrawLine(float64(x),0,float64(x),float64(pixelH));
	ctx.Stroke();

	// Horizontal line
	ctx.DrawLine(0,float64(y),float64(pixelW),float64(y));
	ctx.Stroke();

	// Topleft helper
	ctx.DrawLine(0,0,float64(x),float64(y));
	ctx.Stroke();

	// Scale helper rectangle
	rw, rh := float64(pixelW)/s, float64(pixelH)/s;
	ctx.DrawRectangle(float64(x)-rw/2,float64(y)-rh/2,rw,rh);
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

func renderHistogram(ctx *gg.Context, escapeHistogram []int){
	// Some sweet histogram data!
	for i := 0; i < len(escapeHistogram); i++ {
		fmt.Printf("%d ", escapeHistogram[i]);
	}

	ctx.SetRGB(0,0,0);

	for i := 0; i < len(escapeHistogram); i++ {
		for h := 0; h < escapeHistogram[i]; h++ {
			ctx.SetPixel(i,h);
		}
	}
}

func renderImage(ctx *gg.Context, filename string){
	ctx.SetRGB(1,1,1);
	ctx.Clear();

	iterations := {{ITERATIONS}};

	interestingPixels := []pixel{};

	escapeHistogram := make(map[int]int);
	escapedPixels := 0;
	skippedPixels := 0;
	unescapedPixels := make(map[int]map[int]savedValue);

	// Loop through every pixel
	// TODO: This loop could be threaded, as a pixel need not know its neighbor to calculate its value.

	iterationChunk := 50;
	currentIteration := 0;

	for {
		for pixelX := 0; pixelX < pixelW; pixelX++ {
			for pixelY := 0; pixelY < pixelH; pixelY++ {
				var x, y float64;

				x, y = getMathCoordinates(pixelX, pixelY);

				// Save the initial values
				x0 := x;
				y0 := y;

				// Why are you like this?
				_, ok := unescapedPixels[pixelX];
				if (ok==true) {
					_, ok2 := unescapedPixels[pixelX][pixelY];
					if (ok2==true) {
						previousValue := unescapedPixels[pixelX][pixelY];
						x = previousValue[0];
						y = previousValue[1];
						if (pixelX == 0 && pixelY == 0){
							fmt.Printf("%d, %d = %f %f", pixelX, pixelY, x, y);
						}
					}
					ok = ok && ok2;
				}

				// Skip pixels that are already escaped
				if (currentIteration > 0 && !ok) {
					skippedPixels++;
					continue;
				}

				// Track the min iteration to crack the pixels
				minIteration := currentIteration+iterationChunk;

				for iteration := currentIteration; iteration < currentIteration+iterationChunk; iteration++ {
					if ({{EXPRESSION}}) {
						// The value is (still) inside the set
						// Perform transformations on both x and y
						newx := {{TRANSFORMX}};
						newy := {{TRANSFORMY}};
						x = newx;
						y = newy;

						if (iteration == currentIteration+iterationChunk-1) {
							// Put this pixel into the map in case we do a re-run
							if (unescapedPixels[pixelX] == nil) {
								unescapedPixels[pixelX] = map[int]savedValue{};
							}
							unescapedPixels[pixelX][pixelY] = savedValue{x,y};
						}
					} else {
						// The value escaped the set
						// Calculate the color
						r,g,b := getColor(float64(iteration)/float64(iterations),paletteRainbow);
						ctx.SetRGB(r,g,b);
						ctx.SetPixel(pixelX,pixelY);

						_, ok := escapeHistogram[iteration];
						if (!ok) {
							escapeHistogram[iteration] = 0;
						}
						escapeHistogram[iteration]++;

						escapedPixels++;

						_, ok = unescapedPixels[pixelX];
						if (ok) {
							delete(unescapedPixels[pixelX],pixelY);
						}

						// Update the min iteration
						if (iteration < minIteration) {
							minIteration = iteration;
						}

						// Because we are zooming in, we want to see how interesting this pixel is. The most interesting pixels will be chosen to act as the next bunch of centers for the zoom-in.
						if (iteration==(currentIteration+iterationChunk+minIteration)/2){
							// Interesting enough to be considered
							interestingPixels = append(interestingPixels, pixel{pixelX, pixelY});
							// DEBUG: Color the interesting pixels above the treshold as black
							//ctx.SetRGB(0,0,0);
							//ctx.SetPixel(pixelX,pixelY);
						}
						break;
					}
				}
			}
		}

		currentIteration += iterationChunk;
		// Check if it's time to stop.
		// Get the last 10 values and determine how much the values have changed
		sampleSize := 10;

		// Make sure we have enough iterations.
		if (currentIteration > sampleSize) {
			changeDelta := 0;
			for g := 0; g < sampleSize; g++ {
				changeDelta += abs(escapeHistogram[currentIteration-1-g]-escapeHistogram[currentIteration-2-g]);
			}

			fmt.Printf("%d - %d Change delta: %d\n", escapedPixels, skippedPixels, changeDelta);
			if (changeDelta < 256) {
				break;
			}
		}
	}

	// Get a zoom in crosshair, calculate the pixel values
	// Which pixels were the most interesting? Zoom there!
	spread := 10.0; // How many pixels from the center we can spread in case there are no interesting pixels
	length := len(interestingPixels);
	interestingPixel := pixel{int(float64(pixelW)/2.0+(rand.Float64()-0.5)*spread),int(float64(pixelH)/2.0+(rand.Float64()-0.5)*spread)};
	if (length > 0) {
		interestingPixel = interestingPixels[int(rand.Float64()*float64(length))];
	}
	pixelZoomX := interestingPixel[0];
	pixelZoomY := interestingPixel[1];
	pixelZoomS := 1.5;

	// Render the zoom in crosshair for debug
	// renderZoomCrosshair(ctx,pixelZoomX,pixelZoomY,pixelZoomS);

	// Get the coordinates for this pixel
	x, y := getMathCoordinates(pixelZoomX, pixelZoomY);

	// Set zoom to x and y
	fmt.Printf("Math coords: %f,%f -> %f,%f\n",mathX,mathY,x,y);
	mathX = x;
	mathY = y;

	// Get new scale from zoom level
	mathW /= pixelZoomS;
	mathH /= pixelZoomS;

	//renderHistogram(ctx,escapeHistogram);

	ctx.SavePNG(filename);
}

func main(){
	rand.Seed(time.Now().UTC().UnixNano());

	fmt.Printf("%s\n","Running `{{EXPRESSION}}` for {{ITERATIONS}} iterations");

	ctx := gg.NewContext(pixelW, pixelH);

	dir := "./"+fmt.Sprintf("%d",time.Now().UTC().UnixNano())+"/";

	os.Mkdir(dir,0755);

	for i := 0; i < {{IMAGES}}; i++ {
		renderImage(ctx, fmt.Sprintf(dir+"%s-%f-%f,%fx%f.png","{{EXPRESSION}}", float64(i), float64(mathX), float64(mathY), float64(mathW)));
	}
}
