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


#include <stdio.h>

/* Go functions can not be passed to c directly.            */
/* They can only be called from c.                          */
/* This code is an indirection to call Go callbacks.        */
/* _cgo_export.h is generated automatically by cgo.         */
#include "_cgo_export.h"

/* Exported functions from Go are:                          */
/* goDisplay                                                */
/* goOverlayDisplay                                         */
/* goReshape                                                */
/* goKeyboard                                               */
/* goMouse                                                  */
/* goMotion                                                 */
/* goPassiveMotion                                          */
/* goVisibility                                             */
/* goEntry                                                  */
/* goTimer                                                  */
/* goIdle                                                   */


#if defined(__APPLE__)
	#include <GLUT/glut.h>
#else
	#include <GL/glut.h>
#endif


/* in freeglut 3.0 GLdouble is replaced by double */

void glutSolidSphereCompatibilityWrapper( double radius, GLint slices, GLint stacks ) {
	glutSolidSphere( radius, slices, stacks );
}
void glutWireSphereCompatibilityWrapper( double radius, GLint slices, GLint stacks ) {
	glutWireSphere( radius, slices, stacks );
}
void glutSolidCubeCompatibilityWrapper( double size ) {
	glutSolidCube( size );
}
void glutWireCubeCompatibilityWrapper( double size ) {
	glutWireCube( size );
}
void glutSolidConeCompatibilityWrapper( double base, double height, GLint slices, GLint stacks ) {
	glutSolidCone( base, height, slices, stacks );
}
void glutWireConeCompatibilityWrapper( double base, double height, GLint slices, GLint stacks ) {
	glutWireCone( base, height, slices, stacks );
}
void glutSolidTorusCompatibilityWrapper( double innerRadius, double outerRadius, GLint nsides, GLint rings ) {
	glutSolidTorus( innerRadius, outerRadius, nsides, rings );
}
void glutWireTorusCompatibilityWrapper( double innerRadius, double outerRadius, GLint nsides, GLint rings ) {
	glutWireTorus( innerRadius, outerRadius, nsides, rings );
}
void glutSolidTeapotCompatibilityWrapper( double size ) {
	glutSolidTeapot( size );
}
void glutWireTeapotCompatibilityWrapper( double size ) {
	glutWireTeapot( size );
}


/* register callbacks */

void register_display() {
	glutDisplayFunc(&goDisplay);
}
void register_overlayDisplay() {
	glutOverlayDisplayFunc(&goOverlayDisplay);
}
int create_menu() {
	return glutCreateMenu(&goMenu);
}
int create_menuNullCallback() {
	return glutCreateMenu(NULL);
}
void register_reshape() {
	glutReshapeFunc(&goReshape);
}
void register_keyboard() {
	glutKeyboardFunc(&goKeyboard);
}
void register_mouse() {
	glutMouseFunc(&goMouse);
}
void register_motion() {
	glutMotionFunc(&goMotion);
}
void register_passiveMotion() {
	glutPassiveMotionFunc(&goPassiveMotion);
}
void register_visibility() {
	glutVisibilityFunc(&goVisibility);
}
void register_entry() {
	glutEntryFunc(&goEntry);
}
void register_special() {
	glutSpecialFunc(&goSpecial);
}
void register_spaceballMotion() {
	glutSpaceballMotionFunc(&goSpaceballMotion);
}
void register_spaceballRotate() {
	glutSpaceballRotateFunc(&goSpaceballRotate);
}
void register_spaceballButton() {
	glutSpaceballButtonFunc(&goSpaceballButton);
}
void register_buttonBox() {
	glutButtonBoxFunc(&goButtonBox);
}
void register_dials() {
	glutDialsFunc(&goDials);
}
void register_tabletMotion() {
	glutTabletMotionFunc(&goTabletMotion);
}
void register_tabletButton() {
	glutTabletButtonFunc(&goTabletButton);
}
void register_menuStatus() {
	glutMenuStatusFunc(&goMenuStatus);
}
void register_menuState() {
	glutMenuStateFunc(&goMenuState);
}
void register_idle() {
	glutIdleFunc(&goIdle);
}
void register_timer(unsigned int msecs, int timerId) {
	glutTimerFunc(msecs, &goTimer, timerId);
}
void register_windowStatus() {
	glutWindowStatusFunc(&goWindowStatus);
}
void register_keyboardUp() {
	glutKeyboardUpFunc(&goKeyboardUp);
}
void register_specialUp() {
	glutSpecialUpFunc(&goSpecialUp);
}
void register_joystick(int pollInterval) {
	glutJoystickFunc(&goJoystick, pollInterval);
}


/* unregister callbacks */

void unregister_display() {
	glutDisplayFunc(0);
}
void unregister_overlayDisplay() {
	glutOverlayDisplayFunc(0);
}
void unregister_reshape() {
	glutReshapeFunc(0);
}
void unregister_keyboard() {
	glutKeyboardFunc(0);
}
void unregister_mouse() {
	glutMouseFunc(0);
}
void unregister_motion() {
	glutMotionFunc(0);
}
void unregister_passiveMotion() {
	glutPassiveMotionFunc(0);
}
void unregister_visibility() {
	glutVisibilityFunc(0);
}
void unregister_entry() {
	glutEntryFunc(0);
}
void unregister_special() {
	glutSpecialFunc(0);
}
void unregister_spaceballMotion() {
	glutSpaceballMotionFunc(0);
}
void unregister_spaceballRotate() {
	glutSpaceballRotateFunc(0);
}
void unregister_spaceballButton() {
	glutSpaceballButtonFunc(0);
}
void unregister_buttonBox() {
	glutButtonBoxFunc(0);
}
void unregister_dials() {
	glutDialsFunc(0);
}
void unregister_tabletMotion() {
	glutTabletMotionFunc(0);
}
void unregister_tabletButton() {
	glutTabletButtonFunc(0);
}
void unregister_menuStatus() {
	glutMenuStatusFunc(0);
}
void unregister_menuState() {
	glutMenuStateFunc(0);
}
void unregister_idle() {
	glutIdleFunc(0);
}
void unregister_windowStatus() {
	glutWindowStatusFunc(0);
}
void unregister_keyboardUp() {
	glutKeyboardUpFunc(0);
}
void unregister_specialUp() {
	glutSpecialUpFunc(0);
}
void unregister_joystick(int pollInterval) {
	glutJoystickFunc(0, pollInterval);
}

/* Fonts */

void *const_GLUT_STROKE_ROMAN() {
	return GLUT_STROKE_ROMAN;
}
void *const_GLUT_STROKE_MONO_ROMAN() {
	return GLUT_STROKE_MONO_ROMAN;
}
void *const_GLUT_BITMAP_9_BY_15() {
	return GLUT_BITMAP_9_BY_15;
}
void *const_GLUT_BITMAP_8_BY_13() {
	return GLUT_BITMAP_8_BY_13;
}
void *const_GLUT_BITMAP_TIMES_ROMAN_10() {
	return GLUT_BITMAP_TIMES_ROMAN_10;
}
void *const_GLUT_BITMAP_TIMES_ROMAN_24() {
	return GLUT_BITMAP_TIMES_ROMAN_24;
}
void *const_GLUT_BITMAP_HELVETICA_10() {
	return GLUT_BITMAP_HELVETICA_10;
}
void *const_GLUT_BITMAP_HELVETICA_12() {
	return GLUT_BITMAP_HELVETICA_12;
}
void *const_GLUT_BITMAP_HELVETICA_18() {
	return GLUT_BITMAP_HELVETICA_18;
}

