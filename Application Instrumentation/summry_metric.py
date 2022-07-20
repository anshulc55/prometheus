import http.server
import time
from prometheus_client import start_http_server, Summary

REQUEST_LATENCY_TIME= Summary('request_latency_time','Response latency in seconds')

class HandleRequests(http.server.BaseHTTPRequestHandler):

    @REQUEST_LATENCY_TIME.time()
    def do_GET(self):
        #startTime = time.time()
        self.send_response(200)
        time.sleep(1)
        self.send_header("Content-type", "text/html")
        self.end_headers()
        self.wfile.write(bytes("<html><head><title>First python Application</title></head><body style='color: #333; margin-top: 30px;'><center><h2>Welcome to our first Python application.</center></h2></body></html>", "utf-8"))
        self.wfile.close
        #end_time = time.time() - startTime
        #REQUEST_LATENCY_TIME.observe(end_time) 

if __name__ == "__main__":
    start_http_server(5001)
    server = http.server.HTTPServer(('178.62.224.102', 5000), HandleRequests)
    server.serve_forever()