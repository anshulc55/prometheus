from flask import Flask, request, Response
from prometheus_client import Counter, generate_latest

REQUEST_COUNT = Counter("app_request_counter", "Total Number of Requests made of Application")

app = Flask(__name__)

@app.route("/")
def hello_team():
    REQUEST_COUNT.inc()
    return "Hello Team, This is from Docker WSGI...."

@app.route("/metrics")
def metrics():
    return Response(generate_latest(), 200, mimetype="text/plain")

if __name__ == "__main__":
    app.run(hostname="0.0.0.0", port=5001)