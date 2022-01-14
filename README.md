# Universal Enigma emulator

Configurable library for emulating the Enigma cyphering machine. Basic usage:

```go
package main

import (
	"fmt"
	"github.com/tomas-hanicinec/enigma/enigma"
)

func main() {
	e, err := enigma.NewEnigma(enigma.M3)

	text := enigma.Preprocess("Some text to encode") // handles case, spaces, ...
	fmt.Printf("Preprocessed text: %s\n", text)

	encoded, err := e.Encode(text)
	fmt.Printf("Encoded text: %s\n", encoded)

	e.RotorsReset() // must set rotors back to starting configuration before decoding
	decoded, err := e.Encode(encoded)
	decoded = enigma.Postprocess(decoded) // makes Enigma output more readable
	fmt.Printf("Decoded text: %s\n", decoded)
}
```

## Supported models and features

Supports all mainstream Enigma models, most notably German military models **I** and **M3**, four-rotor model **M4**, basic commercial Enigma (models **D** / **K**) and more. Also supports the **UKW-D** rewirable reflector used later in the war in models M3 and M4.

Full list of supported models along with their names, descriptions, design ect can be acquired as follows
```go
for _, model := range enigma.GetSupportedModels() {
	fmt.Printf("%s: %s\n", model.GetName(), model.GetDescription())
}
```

## Accepted inputs

Enigma machines can only encode **uppercase letters from the basic 26-letter alphabet**. This in practice led to various letter substitutions being used for common unsupported symbols like spaces and comas. One such substitution is provided by the `Preprocess()` function (and its complementary `Postprocess()`). It handles letter case, spaces and characters `.`, `,` and `-`.

**A few examples of unsupported texts**:
* `Friday` - lowercase letters, use `enigma.Preprocess("Friday")`
* `FRIDAY AFTERNOON.` - space and comma, use `enigma.Preprocess("FRIDAY AFTERNOON")`
* `FRIDAY (AFTERNOON)` - brackets not supported, remove or replace them
* `FRIDAY 24TH` - numbers not supported, remove or spell them out
* `PÈRE NOËL` - diacritics not supported, convert to plain letters

*Note on the substitutions used by `Preprocess()` and `Postprocess()` - they are optimized for common English language and might not work properly for some arbitrary character sequences or different languages, where the letter pairs used for substitutions are more common. For example hyphen is substituted by `YY`, so a "word" `BAYYES-NET` would become `BAYYESYYNET` after preprocess and `BA-ES-NET` after postprocess, which is not the original word. Collisions like this would still be very rare though.*

## Setup & configuration

Every Enigma model starts with a valid default configuration out of the box and is ready to start encoding. However, the configuration for all models still can (and should) be changed. This can be done on three different levels:

### "Local" configuration (changing just one thing)
When we already have a working Enigma model and want to tweak it (ie. when trying to break the code).
```go
e, err := enigma.NewEnigma(enigma.M4)

//select rotor types and put them into available slots
placedRotorTypes := map[enigma.RotorType]struct{}{} // cannot use the same rotor type twice
rotorTypes := map[enigma.RotorSlot]enigma.RotorType{}
for _, slot := range e.GetAvailableRotorSlots() {
	for _, availableType := range e.GetAvailableRotors(slot) {
		if _, ok := placedRotorTypes[availableType]; !ok {
			rotorTypes[slot] = availableType // use the first available unused
			placedRotorTypes[availableType] = struct{}{}
			break
		}
	}
}
err = e.RotorsSelect(rotorTypes) // rotor types can only be set all at once

// now adjust the middle rotor
err = e.RotorSetWheel(enigma.Middle, 'X') // allowed rotor wheel positions: A-Z
err = e.RotorSetRing(enigma.Middle, 5)    // allowed rotor ring positions: 1-26

// reflector changes
reflectors := e.GetAvailableReflectors()
err = e.ReflectorSelect(reflectors[0]) // select supported reflector
if e.GetReflectorType().IsMovable() {
    // this means the reflector can be rotated 
	err = e.ReflectorSetWheel(15) // allowed reflector wheel positions: 1-26
} else if e.GetReflectorType().IsRewirable() {
	// UKW-D rewirable reflector (about the configuration format see the note bellow)
	err = e.ReflectorRewire("AV BO CT DM EZ FN GX HQ IS KR LU PW") 
}

if e.HasPlugboard() {
    // this model also contains rewirable plugboard - let's set it up
    err = e.PlugboardSetup("AB CD EF GH")
}
```

### "Middle" configuration (fully configuring a single "subsystem")
When we want to completely reconfigure rotors, reflector, or plugboard and leave the rest as is.
```go
e, err := enigma.NewEnigma(enigma.Commercial)

// set up the rotors (we must already know the slots and available rotors for the current model)
err = e.RotorsSetup(map[enigma.RotorSlot]enigma.RotorConfig{
    enigma.Right: {
        RotorType:     enigma.Rotor_IK,
        WheelPosition: 'C',
        RingPosition:  12,
    },
    enigma.Middle: {enigma.Rotor_IIIK, 'Q', 10},  // one liner
    enigma.Left:   {RotorType: enigma.Rotor_IIK}, // just the type set, rest is on default
})

// set up the reflector
err = e.ReflectorSetup(enigma.ReflectorConfig{
    ReflectorType: enigma.UKW_B,
    WheelPosition: 15, // Commercial Enigma had a movable reflector, so ok to set position
    Wiring:        "",
})
```
### "Global" configuration (setting up the whole machine at once)
When we want to create new Enigma with full-custom configuration from scratch.
```go
e, err := enigma.NewEnigmaWithSetup(
    enigma.M4,
    map[enigma.RotorSlot]enigma.RotorConfig{
        enigma.Right:  {enigma.Rotor_I, 'W', 10},
        enigma.Middle: {enigma.Rotor_II, 'D', 5},
        enigma.Left:   {enigma.Rotor_III, 'A', 7},
        enigma.Fourth: {enigma.Rotor_gamma, 'X', 5},
    },
    enigma.ReflectorConfig{ReflectorType: enigma.UKW_BThin},
    "AB CD EF",
)
```

## Note on plugboard and UKW-D configuration

**Plugboards** are configured by a string containing pairs of uppercase letters, for example `AB CD EF GH`. Each pair represents one plug, so there is a maximum of 13 pairs in a valid configuration (no letter can be plugged twice and no letter can be plugged to itself). Partial configurations are allowed (not all plugs connected).

**UKW-D reflectors** are configured similarly by pairs of connected uppercase letters. There are a few differences though due to how these reflectors worked in real life:
* Letters Y and J were always hardwired together, so it is not possible to use these letters in any pair
* Partial configurations are not allowed, all letters (except Y and J) must be present in a valid configuration
* Valid UKW-D configuration is therefore always a string of 12 uppercase letter pairs with no letters repeating and Y and J missing

