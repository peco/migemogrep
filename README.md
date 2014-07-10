migemogrep
==========

Search/grep against Japanese text using latin-1 text

Installation
------------

Please visit our releases page, and download the appropriate version/platform:

[https://github.com/peco/migemogrep/releases](https://github.com/peco/migemogrep/releases)

Or, if you are on OS X and are using homebrew

    $ brew tap peco/peco
    $ brew install migemogrep

And finally, if you want the latest bleeding edge version:

    $ go get github.com/peco/migemogrep

Usage
-----

```sh
$ migemogrep <pattern> <file>
$ cat file.txt | migemogrep <pattern>
```

![optimized](http://peco.github.io/images/migemogrep-demo.gif)
