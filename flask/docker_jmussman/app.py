from flask import Flask, render_template
import os
import random

app = Flask(__name__)

def page_not_found(e):
    return render_template('404a.html'), 404
app.register_error_handler(404, page_not_found)

@app.route("/")
def index():
    return render_template("index.html")
# To get picutes working had to put everything that is not a route in /static directory!

@app.route("/robots.txt")
def robot():
    return render_template("robots.txt")

@app.route("/ForensicsWebsite/html/<page>")
def forensite(page):
    if page == 'index.html':
        return render_template("ForensicsWebsite/html/index.html")
    if page == 'calendar.html':
        return render_template("ForensicsWebsite/html/calendar.html")
    if page == 'resources.html':
        return render_template("ForensicsWebsite/html/resources.html")
    if page == 'syllabus.html':
        return render_template("ForensicsWebsite/html/syllabus.html")

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=int(os.environ.get("PORT", 5000)))

