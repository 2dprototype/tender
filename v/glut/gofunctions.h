/*
 * This is free and unencumbered software released into the public domain.
 * 
 * Anyone is free to copy, modify, publish, use, compile, sell, or
 * distribute this software, either in source code form or as a compiled
 * binary, for any purpose, commercial or non-commercial, and by any
 * means.
 * 
 * In jurisdictions that recognize copyright laws, the author or authors
 * of this software dedicate any and all copyright interest in the
 * software to the public domain. We make this dedication for the benefit
 * of the public at large and to the detriment of our heirs and
 * successors. We intend this dedication to be an overt act of
 * relinquishment in perpetuity of all present and future rights to this
 * software under copyright law.
 * 
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
 * IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
 * OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
 * ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
 * OTHER DEALINGS IN THE SOFTWARE.
 * 
 * For more information, please refer to <http://unlicense.org>
 */


#ifndef GOFUNCTIONS_H
#define GOFUNCTIONS_H


void glutSolidSphereCompatibilityWrapper( double radius, GLint slices, GLint stacks );
void glutWireSphereCompatibilityWrapper( double radius, GLint slices, GLint stacks );
void glutSolidCubeCompatibilityWrapper( double size );
void glutWireCubeCompatibilityWrapper( double size );
void glutSolidConeCompatibilityWrapper( double base, double height, GLint slices, GLint stacks );
void glutWireConeCompatibilityWrapper( double base, double height, GLint slices, GLint stacks );
void glutSolidTorusCompatibilityWrapper( double innerRadius, double outerRadius, GLint nsides, GLint rings );
void glutWireTorusCompatibilityWrapper( double innerRadius, double outerRadius, GLint nsides, GLint rings );
void glutSolidTeapotCompatibilityWrapper( double size );
void glutWireTeapotCompatibilityWrapper( double size );

void register_display();
void register_overlayDisplay();
int create_menu();
int create_menuNullCallback();
void register_reshape();
void register_keyboard();
void register_mouse();
void register_motion();
void register_passiveMotion();
void register_visibility();
void register_entry();
void register_special();
void register_spaceballMotion();
void register_spaceballRotate();
void register_spaceballButton();
void register_buttonBox();
void register_dials();
void register_tabletMotion();
void register_tabletButton();
void register_menuStatus();
void register_menuState();
void register_idle();
void register_timer(unsigned int msecs, int timerId);
void register_windowStatus();
void register_keyboardUp();
void register_specialUp();
void register_joystick(int pollInterval);

void unregister_display();
void unregister_overlayDisplay();
void unregister_reshape();
void unregister_keyboard();
void unregister_mouse();
void unregister_motion();
void unregister_passiveMotion();
void unregister_visibility();
void unregister_entry();
void unregister_special();
void unregister_spaceballMotion();
void unregister_spaceballRotate();
void unregister_spaceballButton();
void unregister_buttonBox();
void unregister_dials();
void unregister_tabletMotion();
void unregister_tabletButton();
void unregister_menuStatus();
void unregister_menuState();
void unregister_idle();
void unregister_windowStatus();
void unregister_keyboardUp();
void unregister_specialUp();
void unregister_joystick(int pollInterval);

void *const_GLUT_STROKE_ROMAN();
void *const_GLUT_STROKE_MONO_ROMAN();
void *const_GLUT_BITMAP_9_BY_15();
void *const_GLUT_BITMAP_8_BY_13();
void *const_GLUT_BITMAP_TIMES_ROMAN_10();
void *const_GLUT_BITMAP_TIMES_ROMAN_24();
void *const_GLUT_BITMAP_HELVETICA_10();
void *const_GLUT_BITMAP_HELVETICA_12();
void *const_GLUT_BITMAP_HELVETICA_18();

#endif /* GOFUNCTIONS_H */






