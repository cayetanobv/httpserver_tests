
import logging
import json
from flask import Flask, Response


app = Flask(__name__)

log = logging.getLogger('werkzeug')
log.setLevel(logging.ERROR)

@app.route("/")
def testHandler():
    data = {"Name":"HelloWorld","Types":["hello","world"]}
    return Response(response=data, mimetype="application/json")


if __name__ == "__main__":
    print("Listening at http://localhost:8090")
    app.run(port=8090)
