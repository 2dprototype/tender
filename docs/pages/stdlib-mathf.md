# Stdlib `mathf`

The `mathf` module provides mathematical utilities for common operations including interpolation, angle manipulation, power-of-two calculations, and noise generation. It's inspired by Unity's Mathf library and optimized for game development and real-time applications.

## Constants

- `deg2rad`: Conversion factor from degrees to radians (π/180).
- `rad2deg`: Conversion factor from radians to degrees (180/π).
- `eps`: Epsilon value (2^-52) used for floating-point comparisons.
- `neg_inf`: Negative infinity value.

## Functions

### Core Utilities

- `abs(x)`: Returns the absolute value of x.

- `sign(x)`: Returns 1 if x is positive or zero, -1 if x is negative.

- `clamp(value, min, max)`: Clamps a value between a minimum and maximum float value.

- `clamp01(value)`: Clamps a value between 0 and 1.

- `approx(a, b)`: Compares two floating-point values and returns true if they are similar within epsilon tolerance.

- `closest_pow2(value)`: Returns the closest power of two value (32-bit).

- `closest_pow2l(value)`: Returns the closest power of two value (64-bit).

- `is_pow2(value)`: Returns true if the value is a power of two.

- `next_pow2(value)`: Returns the next power of two value.

- `round(x)`: Returns x rounded to the nearest integer (ties to even).

### Angle Utilities

- `delta_angle(current, target)`: Calculates the shortest difference between two given angles in degrees.

- `lerp_angle(a, b, t)`: Same as lerp but ensures values interpolate correctly when they wrap around 360 degrees.

- `move_angle(current, target, max_delta)`: Moves an angle towards a target angle with a maximum delta step.

### Interpolation & Smoothing

- `lerp(a, b, t)`: Linearly interpolates between a and b by t (clamped between 0 and 1).

- `lerp_uncl(a, b, t)`: Linearly interpolates between a and b by t (unclamped).

- `inverse_lerp(a, b, value)`: Calculates the linear parameter t that produces the interpolant value within the range [a, b].

- `move_towards(current, target, max_delta)`: Moves a value towards a target with a maximum delta step.

- `smooth_step(from, to, t)`: Interpolates between from and to with smoothing at the limits (Hermite interpolation).

### Loops & Waves

- `repeat(t, length)`: Loops the value t so that it is never larger than length and never smaller than 0.

- `pingpong(t, length)`: Ping-pongs the value t so that it oscillates between 0 and length.

### Color Space

- `gamma2linear(value)`: Converts the given value from gamma (sRGB) to linear color space.

- `linear2gamma(value)`: Converts the given value from linear to gamma (sRGB) color space.

### Noise

- `perlin(x, y)`: Generates 2D Perlin noise value in the range [-1, 1].

### Optimization

- `inv_sqrt(x)`: Returns the inverse square root of x using the Quake III fast algorithm.

## Examples

```go
import "mathf"

// Core utilities
mathf.abs(-5)           // 5
mathf.sign(-3.5)        // -1
mathf.clamp(10, 0, 5)   // 5
mathf.clamp01(0.8)      // 0.8
mathf.approx(0.1, 0.1000001) // true
mathf.is_pow2(8)        // true
mathf.next_pow2(5)      // 8
mathf.round(2.5)        // 2 (ties to even)

// Interpolation
mathf.lerp(0, 10, 0.5)  // 5
mathf.inverse_lerp(0, 10, 5) // 0.5
mathf.smooth_step(0, 10, 0.5) // 5 (with smooth easing)

// Angles
mathf.delta_angle(350, 10) // 20 (shortest path)
mathf.lerp_angle(350, 10, 0.5) // 0 (correctly wraps)

// Loops
mathf.repeat(5, 3)     // 2
mathf.pingpong(5, 3)   // 1

// Color
mathf.gamma2linear(0.5)  // 0.217637640824031
mathf.linear2gamma(0.5)  // 0.7297400528400191

// Noise
mathf.perlin(0.5, 0.3)  // Perlin noise value between -1 and 1

// Optimization
mathf.inv_sqrt(4)       // 0.5 (fast inverse square root)
```

## Constants Example

```go
import "mathf"

mathf.deg2rad          // 0.017453292519943295
mathf.rad2deg          // 57.29577951308232
mathf.eps              // 2.220446049250313e-16
mathf.neg_inf          // -Infinity
```