gooey
=====

I'm learning Go and decided to make a YUI combo server. Probably not useful or correct.

Usage
-----

Provide the path to js files as an environment variable:

`$ GOOEY_JS_ROOT=/tmp/static/js/ go run combo.go`

or

`$ GOOEY_JS_ROOT=/tmp/static/js/ ./bin/combo`

In a browser:

`http://127.0.0.1:9000/combo/?baz.js&foo.js`

	function baz() {
		return 1;
	}
	var foo = {"foo": "bar"};
