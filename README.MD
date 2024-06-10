
# ASCII-ART-WEB

It's a web server for outputting the text in a graphic representation of ASCII.

Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of a program that outputs text in a graphical representation of ASCII

# Authors
zrdouane & oelhaouc

# Usage
1.To run program use the following command
```zsh
  go run .
```
2.Open the web browser and go to
```zsh
click on the link 'http://localhost:3000/'
```
#
```
# The algorithm works as follows:

It replaces all carriage return and tab characters in the input text with newline and four spaces respectively.
It splits the text into lines.
For each line, it converts each character into its ASCII art representation by looking it up in the ASCII art file specified by the style parameter.
It formats the ASCII art lines into a string and returns it.
The ASCII art files are loaded from the model directory using the LoadAsciiArtFromFile function in the same file. This function reads the ASCII art file line by line and stores the ASCII art representation of each character in a map. The map is then used to look up the ASCII art representation of each character when generating the ASCII art.

The server is implemented in Go using the net/http package. The server's entry point is the main function in the main.go file. The server listens on port 3000 and serves the ASCII art web GUI.