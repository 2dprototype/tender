## Stdlib `audio`

The audio module provides functionalities for initializing an audio context and creating audio players for playing MP3 files.

### Functions

#### `init()`

Initializes the audio context with the specified options.

#### `player(data)`

Creates an audio player for playing MP3 audio data.

- `data`: Byte array containing MP3 audio data.

### Audio Context Functions

#### `suspend()`

Suspends the audio context.

#### `resume()`

Resumes the audio context.

### Audio Player Functions

#### `decoder`

Returns information about the audio decoder.

- `length()`: Returns the length of the audio in milliseconds.
- `sample_rate()`: Returns the sample rate of the audio.
- `seek(position)`: Seeks to the specified position in the audio.

#### `play()`

Plays the audio.

#### `pause()`

Pauses the audio playback.

#### `is_playing()`

Checks if the audio is currently playing.

#### `close()`

Closes the audio player.

#### `err()`

Returns any error encountered during audio playback.

#### `reset()`

Resets the audio player.

#### `buffered_size()`

Returns the size of the audio buffer.

#### `set_buffer_size(size)`

Sets the size of the audio buffer.

- `size`: Size of the audio buffer in bytes.

#### `set_volume(volume)`

Sets the volume level of the audio player.

- `volume`: Volume level (0.0 to 1.0).

#### `volume()`

Returns the current volume level of the audio player.

#### `seek(position)`

Seeks to the specified position in the audio.

- `position`: Position to seek to in milliseconds.

### Example Usage

```go
import "audio"

// Initialize the audio context
audio.init()

// Embed the audio data
var mp3_data = embed("myaudio.mp3")

// Create an audio player
var player = audio.player(mp3_data)

// Play the audio
player.play()

for player.is_playing() {}

println("Played!")
```
