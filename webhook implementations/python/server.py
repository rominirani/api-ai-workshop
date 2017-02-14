from flask import Flask, request, jsonify
app = Flask(__name__)

@app.route("/")
def hello():
    return "Hello from APIAI Webhook Integration."

@app.route("/version")
def version():
    return "APIAI Webhook Integration. Version 1.0"

@app.route("/webhook", methods=['POST'])
def webhook():
    content = request.json
    //Extract out the parameters
    //Persist the record
    //Send email notification
    return jsonify({"speech":"Thank You for the feedback","displayText":"Thank You for the feedback","source":"Hotel Feedback System"})

if __name__ == "__main__":
    app.run()