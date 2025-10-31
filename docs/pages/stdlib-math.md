# Stdlib `math`

The `math` module provides mathematical functions and constants for performing various calculations.

## Constants

- `e`: Euler's number (approximately 2.71828).
- `pi`: The mathematical constant pi (approximately 3.14159).
- `phi`: The golden ratio (approximately 1.61803).
- `sqrt2`: The square root of 2 (approximately 1.41421).
- `sqrtE`: The square root of Euler's number.
- `sqrtPi`: The square root of pi.
- `sqrtPhi`: The square root of the golden ratio.
- `ln2`: The natural logarithm of 2.
- `log2E`: The base-2 logarithm of Euler's number.
- `ln10`: The natural logarithm of 10.
- `log10E`: The base-10 logarithm of Euler's number.

## Functions

The following functions are available for performing mathematical operations:

- `abs(x)`: Returns the absolute value of x.
- `acos(x)`: Returns the arccosine of x in radians.
- `acosh(x)`: Returns the inverse hyperbolic cosine of x.
- `asin(x)`: Returns the arcsine of x in radians.
- `asinh(x)`: Returns the inverse hyperbolic sine of x.
- `atan(x)`: Returns the arctangent of x in radians.
- `atan2(y, x)`: Returns the arctangent of y/x in radians, using the signs of both parameters to determine the quadrant of the result.
- `atanh(x)`: Returns the inverse hyperbolic tangent of x.
- `cbrt(x)`: Returns the cube root of x.
- `ceil(x)`: Returns the smallest integer value greater than or equal to x.
- `copysign(x, y)`: Returns x with the sign of y.
- `cos(x)`: Returns the cosine of x (x is in radians).
- `cosh(x)`: Returns the hyperbolic cosine of x.
- `dim(x, y)`: Returns the maximum of x and y.
- `erf(x)`: Returns the error function of x.
- `erfc(x)`: Returns the complementary error function of x.
- `exp(x)`: Returns e^x, where e is Euler's number.
- `exp2(x)`: Returns 2 raised to the power of x.
- `expm1(x)`: Returns e^x - 1.
- `floor(x)`: Returns the largest integer value less than or equal to x.
- `gamma(x)`: Returns the gamma function of x.
- `hypot(x, y)`: Returns sqrt(x^2 + y^2) without intermediate overflow or underflow.
- `ilogb(x)`: Returns the exponent of the radix representation of x.
- `inf(sign)`: Returns positive infinity if sign is positive, negative infinity if sign is negative.
- `is_inf(x, sign)`: Reports whether x is positive infinity or negative infinity.
- `is_nan(x)`: Reports whether x is NaN (not a number).
- `j0(x)`: Returns the order-zero Bessel function of the first kind.
- `j1(x)`: Returns the order-one Bessel function of the first kind.
- `jn(n, x)`: Returns the nth order Bessel function of the first kind.
- `ldexp(frac, exp)`: Returns frac × 2**exp.
- `log(x)`: Returns the natural logarithm of x.
- `log10(x)`: Returns the base-10 logarithm of x.
- `log1p(x)`: Returns the natural logarithm of 1 plus x.
- `log2(x)`: Returns the base-2 logarithm of x.
- `logb(x)`: Returns the unbiased exponent of x in the IEEE 754 floating-point representation.
- `max(x, y)`: Returns the larger of x or y.
- `min(x, y)`: Returns the smaller of x or y.
- `mod(x, y)`: Returns the floating-point remainder of x/y.
- `nan(s)`: Returns an IEEE 754 "not-a-number" value.
- `nextafter(x, y)`: Returns the next representable float value after x towards y.
- `pow(x, y)`: Returns x**y, the base-x exponential of y.
- `pow10(n)`: Returns 10**n.
- `remainder(x, y)`: Returns the IEEE 754 floating-point remainder of x/y.
- `signbit(x)`: Reports whether x is negative or negative zero.
- `sin(x)`: Returns the sine of x (x is in radians).
- `sinh(x)`: Returns the hyperbolic sine of x.
- `sqrt(x)`: Returns the square root of x.
- `tan(x)`: Returns the tangent of x (x is in radians).
- `tanh(x)`: Returns the hyperbolic tangent of x.
- `trunc(x)`: Returns the integer value of x truncated towards zero.
- `y0(x)`: Returns the order-zero Bessel function of the second kind.
- `y1(x)`: Returns the order-one Bessel function of the second kind.
- `yn(n, x)`: Returns the nth order Bessel function of the second kind.
