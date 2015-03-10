<img src="https://i.imgur.com/Fo1W2rL.png" width="269" height="62" alt="Switcher"/>

Switcher
========

Switcher is a proxy server which accepts connections and proxies based on which protocol is detected.

Currently implemented is:

  - SSH
  - IRC

The use case is running HTTP(S), SSH, and IRC on the same port.


Usage
-----

[Download release](https://github.com/jamescun/switcher/releases) or Build:

    make

To get help:

    $ ./switcher --help
    Switcher 1.0.0
    usage: switcher [options]

    Options:
      --listen   <:80>            Server Listen Address
      --ssh      <127.0.0.1:22>   SSH Server Address
      --irc      <127.0.0.1:6667> SSH Server Address
      --default  <127.0.0.1:8080> Default Server Address

    Examples:
      To serve SSH(127.0.0.1:22), IRC(127.0.0.1:6667), and HTTP(127.0.0.1:8080) on port 80
      $ switcher

      To serve SSH(127.0.0.1:2222), IRC(127.0.0.1:6668), and HTTPS(127.0.0.1:443) on port 443
      $ switcher --listen :443 --ssh 127.0.0.1:2222 --irc 127.0.0.1:6668 --default 127.0.0.1:443


Example
-------

Run switcher on HTTP port 80, proxy to SSH on 127.0.0.1:22, ZNC on 127.0.0.1:6667, and Nginx on 127.0.0.1:8080

    $ switcher --listen :80 --ssh 127.0.0.1:22  --irc 127.0.0.1:6667 --default 127.0.0.1:8080

To test HTTP:

    $ curl -I http://my-server.local
    HTTP/1.1 200 OK

To test SSH

    $ ssh james@my-server.local -p 80
    Password:


Why not sslh
------------

Switcher is heavily influenced by [sslh](https://github.com/yrutschle/sslh). It started out as a learning exercise to discover how sslh worked and attempt an implementation in Go.

The result is useful in its own right through use of Go's interfaces for protocol matching (making adding new protocols trivial), and lightweight goroutines (instead of forking, which is more CPU intensive under load).


License
-------

3-Clause "Modified" BSD Licence.

[License](LICENSE)
