
# gunicorn falcon_simple_http_server:app -b localhost:8090

import falcon
import json


class TestHandler:
    def on_get(self, req, resp):
        data = {"Name":"HelloWorld","Types":["hello","world"]}
        resp.status = falcon.HTTP_200
        resp.body = json.dumps(data)


app = falcon.API()
tests = TestHandler()
app.add_route('/', tests)
