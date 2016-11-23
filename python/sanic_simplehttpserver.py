
from sanic import Sanic
from sanic.response import json


app = Sanic(__name__)

@app.route("/")
async def testHandler(request):
    # print(request.headers)
    return json({"Name":"HelloWorld","Types":["hello","world"]})


if __name__ == "__main__":
    app.run(host="localhost", port=8090)
