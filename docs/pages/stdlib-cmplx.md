# Stdlib `cmplx`

The `cmplx` module provides a suite of functions for performing operations on complex numbers. These functions allow you to create, manipulate, and compute advanced mathematical expressions using complex arithmetic.

## Functions

- **`new(real, imag)`**  
  Creates a new complex number from the given real and imaginary parts.

- **`conj(c)`**  
  Returns the complex conjugate of the complex number `c`.

- **`abs(c)`**  
  Returns the modulus (absolute value) of the complex number `c`.

- **`arg(c)`**  
  Returns the phase (argument) of the complex number `c` in radians.

- **`sin(c)`**  
  Returns the sine of the complex number `c`.

- **`cos(c)`**  
  Returns the cosine of the complex number `c`.

- **`acos(c)`**  
  Returns the arc cosine (inverse cosine) of the complex number `c`.

- **`acosh(c)`**  
  Returns the hyperbolic arc cosine of the complex number `c`.

- **`asin(c)`**  
  Returns the arc sine (inverse sine) of the complex number `c`.

- **`asinh(c)`**  
  Returns the hyperbolic arc sine of the complex number `c`.

- **`atan(c)`**  
  Returns the arc tangent (inverse tangent) of the complex number `c`.

- **`atanh(c)`**  
  Returns the hyperbolic arc tangent of the complex number `c`.

- **`cosh(c)`**  
  Returns the hyperbolic cosine of the complex number `c`.

- **`cot(c)`**  
  Returns the cotangent of the complex number `c` (defined as 1/tan(c)).

- **`exp(c)`**  
  Returns the exponential function (e^c) of the complex number `c`.

- **`inf()`**  
  Returns an infinite complex number.

- **`isinf(c)`**  
  Reports whether the complex number `c` is infinite; returns `true` if it is, otherwise `false`.

- **`isnan(c)`**  
  Reports whether the complex number `c` is NaN (not a number); returns `true` if it is, otherwise `false`.

- **`log(c)`**  
  Returns the natural logarithm of the complex number `c`.

- **`log10(c)`**  
  Returns the base-10 logarithm of the complex number `c`.

- **`nan()`**  
  Returns a complex number representing NaN.

- **`phase(c)`**  
  Alias for `arg(c)`; returns the phase (argument) of the complex number `c`.

- **`polar(c)`**  
  Returns the polar coordinates of the complex number `c` as a map with keys:  
  - `r`: The modulus (absolute value)  
  - `theta`: The angle in radians

- **`pow(x, y)`**  
  Returns the result of raising the complex number `x` to the power of complex number `y`.

- **`rect(r, theta)`**  
  Returns the complex number corresponding to the given polar coordinates `r` and `theta`.

- **`sinh(c)`**  
  Returns the hyperbolic sine of the complex number `c`.

- **`sqrt(c)`**  
  Returns the square root of the complex number `c`.

- **`tan(c)`**  
  Returns the tangent of the complex number `c`.

- **`tanh(c)`**  
  Returns the hyperbolic tangent of the complex number `c`.