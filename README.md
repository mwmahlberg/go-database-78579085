go-database-78579085
====================

This repo contains the code to [my answer][2] to the question
["How to connect to H2 database through TCP using h2go package?"][1] on StackOverflow.


Prerequisites
--------------

* A running docker instance.


Run
---

```shell
$ docker compose build
$ docker compose --env-file env up
```


[1]: https://stackoverflow.com/questions/78579085/how-to-connect-to-h2-database-through-tcp-using-h2go-package
[2]: https://stackoverflow.com/a/78598909/1296707