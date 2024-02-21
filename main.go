package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

func main() {
	// Define curve parameters for a small example curve, y^2 = x^3 + Ax + B over Fp
	A := big.NewInt(1)  // Example value for A
	B := big.NewInt(1)  // Example value for B
	P := big.NewInt(17) // Example prime number for P

	// Initialize the curve
	curve := NewEllipticCurve(A, B, P)
	fmt.Println("Elliptic curve created with parameters:")
	fmt.Printf("A: %s, B: %s, P: %s\n", curve.A, curve.B, curve.P)

	// Generate a random point on the curve
	fmt.Println("Generating a random point on the curve...")
	point, err := GenerateRandomPoint(curve)
	if err != nil {
		fmt.Println("Error generating random point:", err)
		os.Exit(1)
	}
	fmt.Printf("Random point on curve: (%s, %s)\n", point.X.Text(10), point.Y.Text(10))

	// Perform point addition with the generated point and a known point on the curve
	fmt.Println("Performing point addition...")
	knownPoint := &Point{
		X:     big.NewInt(5),
		Y:     big.NewInt(1),
		Curve: curve,
	}
	sum := knownPoint.Add(point)
	fmt.Printf("Sum of points: (%s, %s)\n", sum.X.Text(10), sum.Y.Text(10))

	// Perform scalar multiplication
	fmt.Println("Performing scalar multiplication...")
	scalar := big.NewInt(3) // Example scalar
	multipliedPoint := knownPoint.ScalarMult(scalar)
	fmt.Printf("Scalar multiplication result: (%s, %s)\n", multipliedPoint.X.Text(10), multipliedPoint.Y.Text(10))
}

func GenerateRandomPoint(curve *EllipticCurve) (*Point, error) {
	for {
		x, err := rand.Int(rand.Reader, curve.P)
		if err != nil {
			return nil, err
		}

		// Calculate y^2 = x^3 + ax + b
		ySquared := new(big.Int).Exp(x, big.NewInt(3), curve.P) // y^2 = x^3 mod P
		ySquared.Add(ySquared, new(big.Int).Mul(curve.A, x))    // y^2 = (x^3 + Ax) mod P
		ySquared.Add(ySquared, curve.B)                         // y^2 = (x^3 + Ax + B) mod P
		ySquared.Mod(ySquared, curve.P)                         // Ensure it's mod P

		// Check if ySquared is a quadratic residue modulo P, meaning it has a square root mod P
		if EulerCriterion(ySquared, curve.P) {
			y := TonelliShanks(ySquared, curve.P) // Find a square root of ySquared
			if y != nil {
				// Return the point (x, y)
				return NewPoint(x, y, curve), nil
			}
			// If Tonelli-Shanks fails, try another x
		}
	}
}
