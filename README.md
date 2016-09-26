# echoServer
It's an echo Server. You (and I) can test your simple HTTP client.

```
$ curl -v -X GET 'http://gae-echoservier.appspot.com/myapi'

{"header":{
    "Accept":"*/*",
    "Host":"gae-echoservier.appspot.com",
    "User-Agent":"curl/7.49.1",
    "X-Appengine-City":"minato",
    "X-Appengine-Citylatlong":"35.658068,139.751599",
    "X-Appengine-Country":"JP",
    "X-Appengine-Region":"13",
    "X-Cloud-Trace-Context":"1b773912f59f22601dbfe9f016598e45/3880758086095737070"
    },
 "method":"GET",
 "url":"/myapi"
}
```

# How to run this on Google App Engine

1 Create project on GCP console. (e.g. Project ID=my-echoserver)

2 change package name, remove `main()` function and add `init()` function

```
func init() {
	http.HandleFunc("/", handler)
}
```

3 Use appcfg.py to deploy it.

```
$ ls
LICENSE		main.go
$ appcfg.py -A my-echoserver -V v1 update .
```

4. Visit `https://my-echoserver.appspot.com`
