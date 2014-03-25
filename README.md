sorep
--

Command-line app to count the StackOverflow reputation gain/loss for a user.
Since this data is not easily available in the SO API, it uses the graph data
from the user's profile page, so expect this to break as soon as they change
that UI.

Functionality is also available as a Go library if that is what you prefer.
[Docs are here](https://godoc.org/github.com/aliafshar/sorep)

    $ sorep --help
    Usage: sorep [options] username, [username, ...]
      -after="": Date after which to count. e.g. 2013-Sep-30
    exit status 2

    $ sorep 28380/ali-afshar
    aliafshar 22104

    $ sorep --after 2013-Sep-30 28380/ali-afshar
    aliafshar 1357
