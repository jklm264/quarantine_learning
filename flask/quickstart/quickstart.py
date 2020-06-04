from flask import Flask
import flask
from markupsafe import escape

app = Flask(__name__)

@app.route('/')
def index():
    return flask.render_template("index.html")

@app.route('/page1')
def hello():
    return 'Hello world from Page 1!'

@app.route('/user/<username>')
def show_user_profile(username):
    # show the user profile for that user
    return 'User %s' % escape(username)

@app.route('/newfile')
def otherpage():
    return flask.render_template("filename.html")
