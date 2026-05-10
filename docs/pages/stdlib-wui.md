## Stdlib `wui`

The `wui` module provides functionalities for creating native Windows desktop applications with graphical user interface elements like windows, buttons, and input controls.

### Functions

#### Window & Dialog Creation

- **`new_window()`**
    Creates a new window object.
    ```go
    window := wui.new_window()
    ```

- **`new_file_open_dialog()`**
    Creates a new file open dialog object.
    ```go
    dlg := wui.new_file_open_dialog()
    ```

- **`new_file_save_dialog()`**
    Creates a new file save dialog object.
    ```go
    dlg := wui.new_file_save_dialog()
    ```

- **`new_folder_select_dialog()`**
    Creates a new folder selection dialog object.
    ```go
    dlg := wui.new_folder_select_dialog()
    ```

#### Control Creation

These functions create individual UI controls. They return a control object with its own set of methods.

- **`new_button()`**
    Creates a new button.
- **`new_checkbox()`**
    Creates a new checkbox.
- **`new_label()`**
    Creates a new label.
- **`new_editline()`**
    Creates a new single-line text input field.
- **`new_textedit()`**
    Creates a new multi-line text input field.
- **`new_combo_box()`**
    Creates a new combo box (dropdown list).
- **`new_stringlist()`**
    Creates a new list box for displaying strings.
- **`new_stringtable(header1, header2, ...)`**
    Creates a new table control with the specified column headers.
    - `header1`: The text for the first column header.
    - `header2, ...`: Additional column headers.
- **`new_slider()`**
    Creates a new slider control.
- **`new_progressbar()`**
    Creates a new progress bar.
- **`new_radiobutton()`**
    Creates a new radio button.
- **`new_intupdown()`**
    Creates a numeric up-down control for integers.
- **`new_floatupdown()`**
    Creates a numeric up-down control for floats.
- **`new_panel()`**
    Creates a new panel container.
- **`new_paintbox()`**
    Creates a new paint box for custom drawing.

#### Menu Creation

- **`new_menu(name)`**
    Creates a new menu (e.g., for a menu bar or a popup menu).
    - `name`: The text for the menu header.
- **`new_main_menu()`**
    Creates a new main menu bar for a window.
- **`new_menu_string(text)`**
    Creates a new clickable menu item with the specified text.
    - `text`: The text for the menu item.
- **`new_menu_separator()`**
    Creates a new menu separator line.

#### Resource Creation

- **`rgb_color(r, g, b)`**
    Creates a color object from red, green, and blue components (0-255).
    - `r`: Red component (0-255).
    - `g`: Green component (0-255).
    - `b`: Blue component (0-255).
- **`new_font(desc)`**
    Creates a new font object from a description map.
    - `desc`: A map containing font properties (e.g., `{name: "Arial", height: 12, bold: true}`).
- **`new_cursor_from_image(image_bytes, x, y)`**
    Creates a new cursor object from image data.
    - `image_bytes`: Byte array of the image data (PNG, JPEG, etc.).
    - `x`: The X-coordinate of the cursor's hotspot.
    - `y`: The Y-coordinate of the cursor's hotspot.
- **`new_icon_from_image(image_bytes)`**
    Creates a new icon object from image data.
    - `image_bytes`: Byte array of the image data (PNG, JPEG, etc.).

#### Message Boxes

- **`message_box(caption, text)`**
    Displays a simple message box.
    - `caption`: The title of the message box.
    - `text`: The message to display.

- **`message_box_error(caption, text)`**
    Displays an error message box.
    - `caption`: The title.
    - `text`: The message.

- **`message_box_info(caption, text)`**
    Displays an information message box.
    - `caption`: The title.
    - `text`: The message.

- **`message_box_warning(caption, text)`**
    Displays a warning message box.
    - `caption`: The title.
    - `text`: The message.

- **`message_box_question(caption, text)`**
    Displays a question message box.
    - `caption`: The title.
    - `text`: The question.

- **`message_box_ok_cancel(caption, text)`**
    Displays a message box with OK and Cancel buttons.
    - `caption`: The title.
    - `text`: The message.
    - **Returns**: `true` if OK was clicked, `false` if Cancel was clicked.

- **`message_box_yes_no(caption, text)`**
    Displays a message box with Yes and No buttons.
    - `caption`: The title.
    - `text`: The question.
    - **Returns**: `true` if Yes was clicked, `false` if No was clicked.

- **`message_box_custom(caption, text, flags)`**
    Displays a message box with custom buttons defined by flags.
    - `caption`: The title.
    - `text`: The message.
    - `flags`: An integer representing the combination of buttons (e.g., `0x00000004` for Yes/No).
    - **Returns**: An integer representing the ID of the button clicked (e.g., `6` for Yes, `7` for No).

#### Utility Functions

- **`enabled(control)`**
    Checks if a control is enabled.
    - `control`: The control object to check.
    - **Returns**: `true` if enabled, `false` otherwise.

- **`visible(control)`**
    Checks if a control is visible.
    - `control`: The control object to check.
    - **Returns**: `true` if visible, `false` otherwise.

### Object Methods

The following methods are available on the objects returned by the creation functions.

#### `window` Object

- **`show()`**
    Shows the window (non-modal).
- **`show_modal()`**
    Shows the window as a modal dialog.
- **`close()`**
    Closes the window.
- **`set_title(title)`**
    Sets the window's title.
    - `title`: The new title string.
- **`title()`**
    **Returns**: The current window title as a string.
- **`set_size(width, height)`**
    Sets the window's client area size.
    - `width`: The new width.
    - `height`: The new height.
- **`size()`**
    **Returns**: An array `[width, height]`.
- **`set_position(x, y)`**
    Sets the window's screen position.
    - `x`: The new X-coordinate.
    - `y`: The new Y-coordinate.
- **`position()`**
    **Returns**: An array `[x, y]`.
- **`add(control)`**
    Adds a control to the window.
    - `control`: The control object (e.g., from `new_button()`).
- **`remove(control)`**
    Removes a control from the window.
    - `control`: The control object to remove.

#### Common Control Methods

Many controls share a common set of methods.

- **`set_text(text)`**
    Sets the control's text.
    - `text`: The new text string.
- **`text()`**
    **Returns**: The control's current text.
- **`set_bounds(x, y, width, height)`**
    Sets the control's position and size relative to its parent.
    - `x`, `y`, `width`, `height`: Integer values.
- **`set_font(font)`**
    Sets the font for the control.
    - `font`: A font object from `new_font()`.
- **`font()`**
    **Returns**: The control's current font object.
- **`focus()`**
    Sets the keyboard focus to this control.
- **`has_focus()`**
    **Returns**: `true` if the control has focus, `false` otherwise.
- **`set_enabled(enabled)`**
    Enables or disables the control.
    - `enabled`: `true` to enable, `false` to disable.
- **`enabled()`**
    **Returns**: `true` if the control is enabled, `false` otherwise.
- **`set_visible(visible)`**
    Shows or hides the control.
    - `visible`: `true` to show, `false` to hide.
- **`visible()`**
    **Returns**: `true` if the control is visible, `false` otherwise.
- **`set_on_tab_focus(callback)`**
    Sets a function to be called when the control receives focus via the Tab key.
    - `callback`: A function with no arguments.
- **`set_bounds(x, y, width, height)`**
    Sets the control's position and size.

#### `button` Object

- **`set_onclick(callback)`**
    Sets a function to be called when the button is clicked.
    - `callback`: A function with no arguments.

#### `checkbox` Object

- **`set_checked(checked)`**
    Checks or unchecks the box.
    - `checked`: `true` to check, `false` to uncheck.
- **`checked()`**
    **Returns**: `true` if checked, `false` otherwise.
- **`set_on_change(callback)`**
    Sets a function to be called when the checked state changes.
    - `callback`: A function that receives one argument: `function(checked)`.

#### `combobox` Object

- **`set_items(items_array)`**
    Sets the list of items in the combo box.
    - `items_array`: An array of strings.
- **`add_item(item)`**
    Adds a single item to the list.
    - `item`: The string to add.
- **`clear()`**
    Removes all items.
- **`set_selected_index(index)`**
    Selects the item at the given index (0-based).
    - `index`: The integer index.
- **`selected_index()`**
    **Returns**: The index of the currently selected item.
- **`items()`**
    **Returns**: An array of all item strings.
- **`set_on_change(callback)`**
    Sets a function to be called when the selection changes.
    - `callback`: A function that receives one argument: `function(new_index)`.

#### `editline` Object

- **`set_character_limit(limit)`**
    Sets the maximum number of characters the user can type.
    - `limit`: The integer limit.
- **`character_limit()`**
    **Returns**: The character limit.
- **`set_is_password(is_password)`**
    Sets whether the input field should mask its text (e.g., for passwords).
    - `is_password`: `true` to mask the text.
- **`is_password()`**
    **Returns**: `true` if it's a password field.
- **`set_read_only(read_only)`**
    Sets whether the text can be edited.
    - `read_only`: `true` to make it read-only.
- **`read_only()`**
    **Returns**: `true` if read-only.
- **`select_all()`**
    Selects all text in the field.
- **`set_selection(start, end)`**
    Selects a range of text.
    - `start`, `end`: Integer positions.
- **`cursor_position()`**
    **Returns**: An array `[start, end]` representing the current cursor selection.
- **`set_on_text_change(callback)`**
    Sets a function to be called whenever the text changes.
    - `callback`: A function with no arguments.

#### `intupdown` & `floatupdown` Objects

- **`set_min(min)`**
    Sets the minimum value.
- **`set_max(max)`**
    Sets the maximum value.
- **`set_min_max(min, max)`**
    Sets both min and max values.
- **`set_value(value)`**
    Sets the current value.
- **`value()`**
    **Returns**: The current value (`int` or `float`).
- **`min()`**, **`max()`**, **`min_max()`** (returns `[min, max]`)
- **`set_on_value_change(callback)`**
    Sets a function to be called when the value changes.
    - `callback`: A function that receives one argument: `function(new_value)`.

#### `label` Object

- **`set_alignment(alignment)`**
    Sets the text alignment. Use constants like `wui.align_left`, `wui.align_center`, `wui.align_right`.
    - `alignment`: The alignment constant (integer).
- **`alignment()`**
    **Returns**: The current alignment constant.

#### `paintbox` Object

- **`set_on_paint(callback)`**
    Sets the function that draws the control's content. This is the core method for custom drawing.
    - `callback`: A function that receives a `canvas` object: `function(canvas)`.
- **`set_on_mouse_move(callback)`**
    Sets a function called when the mouse moves over the control.
    - `callback`: A function that receives `x` and `y` coordinates: `function(x, y)`.
- **`paint()`**
    Manually triggers a repaint.

#### `canvas` Object (from `paintbox`)

- **`draw_rect(x, y, width, height, color)`**
    Draws the outline of a rectangle.
- **`fill_rect(x, y, width, height, color)`**
    Draws a filled rectangle.
- **`draw_ellipse(x, y, width, height, color)`**
    Draws the outline of an ellipse.
- **`fill_ellipse(x, y, width, height, color)`**
    Draws a filled ellipse.
- **`line(x1, y1, x2, y2, color)`**
    Draws a line.
- **`text_out(x, y, text, color)`**
    Draws text at a specific position.
- **`text_rect(x, y, w, h, text, color)`**
    Draws text within a bounding rectangle.
- **`set_font(font)`**
    Sets the font for subsequent text drawing.
- **`size()`**
    **Returns**: An array `[width, height]` of the canvas.
- **`width()`**, **`height()`**
- *Many other drawing primitives are available, including `arc`, `polygon`, `polyline`, and text measurement functions.*

#### `menu` Object

- **`add(menu_item)`**
    Adds a menu item (from `new_menu_string` or `new_menu_separator`) or a submenu (from `new_menu`) to this menu.

#### `menustring` Object

- **`set_text(text)`**
    Sets the text of the menu item.
- **`text()`**
    **Returns**: The menu item's text.
- *Note: The `set_on_click` method is currently not exposed.*